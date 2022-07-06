import { PluginRegistry } from "@hyperledger/cactus-core";

const {
  AccountId,
  PrivateKey,
  Client,
  AccountBalanceQuery,
} = require("@hashgraph/sdk");
import {
  IPluginLedgerConnector,
  IPluginWebService,
  ICactusPlugin,
  ICactusPluginOptions,
} from "@hyperledger/cactus-core-api";

import { Contract } from "web3-eth-contract";
import {
  Checks,
  Logger,
  LogLevelDesc,
  LoggerProvider,
} from "@hyperledger/cactus-common";

import {
  DeployContractSolidityBytecodeV1Request,
  DeployContractSolidityBytecodeV1Response,
  RunTransactionRequest,
  RunTransactionResponse,
} from "./generated/openapi/typescript-axios";
export interface IPluginLedgerConnectorHederaOptions
  extends ICactusPluginOptions {
  rpcApiHttpHost: string;
  rpcApiWsHost: string;
  pluginRegistry: PluginRegistry;
  logLevel?: LogLevelDesc;
  instanceId: string;
}
export interface GetBalanceRequest {
  operatorId: string;
  operatorKey: string;
}

export class PluginLedgerConnectorHedera
  implements
    IPluginLedgerConnector<
      DeployContractSolidityBytecodeV1Request,
      DeployContractSolidityBytecodeV1Response,
      RunTransactionRequest,
      RunTransactionResponse
    >,
    ICactusPlugin,
    IPluginWebService {
  private readonly instanceId: string;
  private readonly log: Logger;
  private readonly pluginRegistry: PluginRegistry;
  private contracts: {
    [name: string]: Contract;
  } = {};

  public static readonly CLASS_NAME = "PluginLedgerConnectorHedera";
  public get className(): string {
    return PluginLedgerConnectorHedera.CLASS_NAME;
  }

  constructor(public readonly options: IPluginLedgerConnectorHederaOptions) {
    const fnTag = `${this.className}#constructor()`;
    Checks.truthy(options, `${fnTag} arg options`);
    Checks.truthy(options.rpcApiHttpHost, `${fnTag} options.rpcApiHttpHost`);
    Checks.truthy(options.rpcApiWsHost, `${fnTag} options.rpcApiWsHost`);
    Checks.truthy(options.pluginRegistry, `${fnTag} options.pluginRegistry`);
    Checks.truthy(options.instanceId, `${fnTag} options.instanceId`);

    const level = this.options.logLevel || "INFO";
    const label = this.className;
    this.log = LoggerProvider.getOrCreate({ level, label });

    this.instanceId = options.instanceId;
  }

  public getInstanceId(): string {
    return this.instanceId;
  }

  public async shutdown(): Promise<void> {
    this.log.info(`Shutting down ${this.className}...`);
  }

  public async getBalance(req: GetBalanceRequest): Promise<string> {
    // Grab the OPERATOR_ID and OPERATOR_KEY from the .env file
    const operatorId = AccountId.fromString(req.operatorId);
    const operatorKey = PrivateKey.fromString(req.operatorKey);

    console.log(operatorId, operatorKey);

    // Build Hedera testnet and mirror node client
    const client = Client.forTestnet();

    // Set the operator account ID and operator private key
    client.setOperator(operatorId, operatorKey);

    // Create the account balance query
    const query = new AccountBalanceQuery().setAccountId(operatorId);
    //Submit the query to a Hedera network
    const accountBalance = await query.execute(client);
    //Print the balance of hbars
    console.log(
      "The hbar account balance for this account is " + accountBalance?.hbars,
    );
    // v2.0.7
    return accountBalance;
  }
  public getPackageName(): string {
    return `@hyperledger/cactus-plugin-ledger-connector-hedera`;
  }
}
