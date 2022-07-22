import { Logger, Checks } from "@hyperledger/cactus-common";
import { LogLevelDesc, LoggerProvider } from "@hyperledger/cactus-common";
import { Constants } from "@hyperledger/cactus-core-api";
import { DefaultApi } from "../generated/openapi/typescript-axios";
import { Configuration } from "../generated/openapi/typescript-axios/configuration";

export class HederaApiClientOptions extends Configuration {
  readonly logLevel?: LogLevelDesc;
  readonly wsApiHost?: string;
  readonly wsApiPath?: string;
}

export class HederaApiClient extends DefaultApi {
  public static readonly CLASS_NAME = "HederaApiClient";

  private readonly log: Logger;
  private readonly wsApiHost: string;
  private readonly wsApiPath: string;

  public get className(): string {
    return HederaApiClient.CLASS_NAME;
  }

  constructor(public readonly options: HederaApiClientOptions) {
    super(options);
    const fnTag = `${this.className}#constructor()`;
    Checks.truthy(options, `${fnTag} arg options`);

    const level = this.options.logLevel || "INFO";
    const label = this.className;
    this.log = LoggerProvider.getOrCreate({ level, label });

    this.wsApiHost = options.wsApiHost || options.basePath || location.host;
    this.wsApiPath = options.wsApiPath || Constants.SocketIoConnectionPathV1;
    this.log.debug(`Created ${this.className} OK.`);
    this.log.debug(`wsApiHost=${this.wsApiHost}`);
    this.log.debug(`wsApiPath=${this.wsApiPath}`);
    this.log.debug(`basePath=${this.options.basePath}`);
  }
}
