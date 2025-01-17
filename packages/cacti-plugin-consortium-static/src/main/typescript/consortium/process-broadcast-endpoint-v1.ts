import { Express, Request, Response } from "express";
import { stringify as safeStableStringify } from "safe-stable-stringify";
import {
  IWebServiceEndpoint,
  IExpressRequestHandler,
  IEndpointAuthzOptions,
} from "@hyperledger/cactus-core-api";

import { BroadcastResponseV1 } from "../generated/openapi/typescript-axios";

import {
  Logger,
  LogLevelDesc,
  LoggerProvider,
  Checks,
  IAsyncProvider,
} from "@hyperledger/cactus-common";

import {
  handleRestEndpointException,
  registerWebServiceEndpoint,
} from "@hyperledger/cactus-core";

import OAS from "../../json/openapi.json";
import { PluginConsortiumStatic } from "../plugin-consortium-static";
import { StaticConsortiumRepository } from "../repository/static-consortium-repository";
import { InternalServerError } from "http-errors-enhanced-cjs";

export interface IProcessBroadcastEndpointOptions {
  plugin: PluginConsortiumStatic;
  keyPairPem: string;
  consortiumRepo: StaticConsortiumRepository;
  logLevel?: LogLevelDesc;
}

export class ProcessBroadcastEndpoint implements IWebServiceEndpoint {
  public static readonly CLASS_NAME = "ProcessBroadcastEndpoint";

  private readonly log: Logger;
  private readonly plugin: PluginConsortiumStatic;

  constructor(public readonly options: IProcessBroadcastEndpointOptions) {
    const fnTag = `${this.className}#constructor()`;
    if (!options) {
      throw new Error(`${fnTag} options falsy.`);
    }
    if (!options.keyPairPem) {
      throw new Error(`${fnTag} options.keyPairPem falsy.`);
    }
    Checks.truthy(options.consortiumRepo, `${fnTag} options.consortiumRepo`);
    Checks.truthy(options.plugin, `${fnTag} options.plugin`);
    Checks.truthy(
      options.plugin instanceof PluginConsortiumStatic,
      `${fnTag} options.plugin instanceof PluginConsortiumStatic`,
    );
    this.plugin = options.plugin;

    const level = options.logLevel || "INFO";
    const label = "add-node-endpoint-v1";
    this.log = LoggerProvider.getOrCreate({ level, label });
  }

  public get className(): string {
    return ProcessBroadcastEndpoint.CLASS_NAME;
  }

  getAuthorizationOptionsProvider(): IAsyncProvider<IEndpointAuthzOptions> {
    // TODO: make this an injectable dependency in the constructor
    return {
      get: async () => ({
        isProtected: true,
        requiredRoles: [],
      }),
    };
  }

  public getExpressRequestHandler(): IExpressRequestHandler {
    return this.handleRequest.bind(this);
  }

  public get oasPath(): (typeof OAS.paths)["/api/v1/plugins/@hyperledger/cacti-plugin-consortium-static/receive-broadcast"] {
    return OAS.paths[
      "/api/v1/plugins/@hyperledger/cacti-plugin-consortium-static/receive-broadcast"
    ];
  }

  public getOperationId(): string {
    return this.oasPath.post.operationId;
  }

  public getPath(): string {
    return this.oasPath.post["x-hyperledger-cacti"].http.path;
  }

  public getVerbLowerCase(): string {
    return this.oasPath.post["x-hyperledger-cacti"].http.verbLowerCase;
  }

  public async registerExpress(
    expressApp: Express,
  ): Promise<IWebServiceEndpoint> {
    await registerWebServiceEndpoint(expressApp, this);
    return this;
  }

  async handleRequest(req: Request, res: Response): Promise<void> {
    const fnTag = `${this.className}#handleRequest()`;
    const reqTag = `${this.getVerbLowerCase()} - ${this.getPath()}`;

    this.log.debug(reqTag);
    try {
      const msgStringified = safeStableStringify(req.body.message);
      if (typeof msgStringified !== "string") {
        throw new InternalServerError(
          fnTag + "#safeStableStringify() returned with non-string value.",
        );
      }
      await this.plugin.verifySignature(
        req.body.signature,
        msgStringified,
        req.body.message.pubKey,
      );
      await this.options.plugin.processNewNodeRequest(req.body.message.message);
      const body: BroadcastResponseV1 = { result: true };
      res.status(200);
      res.json(body);
    } catch (ex: unknown) {
      const errorMsg = `${fnTag} request handler fn crashed for: ${reqTag}`;
      handleRestEndpointException({ errorMsg, log: this.log, error: ex, res });
    }
  }
}
