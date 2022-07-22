import { PluginRegistry } from "@hyperledger/cactus-core";

const {
  AccountId,
  AccountInfoQuery,
  TopicInfoQuery,
  TokenInfoQuery,
  Client,
  AccountBalanceQuery,
  AccountCreateTransaction,
  Hbar,
  PrivateKey,
  FileContentsQuery,
  FileInfoQuery,
  TransferTransaction,
  ContractByteCodeQuery,
  ScheduleInfoQuery,
  TopicCreateTransaction,
  TopicMessageSubmitTransaction,
  TokenCreateTransaction,
  TokenType,
  PublicKey,
  FileCreateTransaction,
} = require("@hashgraph/sdk");

import type { Express } from "express";
import { RuntimeError } from "run-time-error";
import {
  IPluginLedgerConnector,
  IWebServiceEndpoint,
  IPluginWebService,
  ICactusPlugin,
  ICactusPluginOptions,
  ConsensusAlgorithmFamily,
} from "@hyperledger/cactus-core-api";

import { RunTransactionEndpoint } from "./web-services/run-transaction-endpoint";

import {
  Checks,
  Logger,
  LogLevelDesc,
  LoggerProvider,
} from "@hyperledger/cactus-common";

import { PrometheusExporter } from "./prometheus-exporter/prometheus-exporter";
import {
  GetPrometheusExporterMetricsEndpointV1,
  IGetPrometheusExporterMetricsEndpointV1Options,
} from "./web-services/get-prometheus-exporter-metrics-endpoint-v1";

import {
  HederaCommand,
  HederaQuery,
  RunTransactionRequest,
  RunTransactionResponse,
} from "./generated/openapi/typescript-axios";

export interface IPluginLedgerConnectorHederaOptions
  extends ICactusPluginOptions {
  pluginRegistry: PluginRegistry;
  prometheusExporter?: PrometheusExporter;
  logLevel?: LogLevelDesc;
  instanceId: string;
}
export interface GetBalanceRequest {
  operatorId: string;
}
export interface TransferHbarsRequest {
  myAccountId: string;
  myPrivateKey: string;
  toAccountId: string;
  amount: number;
}

export class PluginLedgerConnectorHedera
  implements
    IPluginLedgerConnector<
      never,
      never,
      RunTransactionRequest,
      RunTransactionResponse
    >,
    ICactusPlugin,
    IPluginWebService {
  private readonly instanceId: string;
  public prometheusExporter: PrometheusExporter;
  private readonly log: Logger;
  private endpoints: IWebServiceEndpoint[] | undefined;

  public static readonly CLASS_NAME = "PluginLedgerConnectorHedera";

  public get className(): string {
    return PluginLedgerConnectorHedera.CLASS_NAME;
  }

  constructor(public readonly options: IPluginLedgerConnectorHederaOptions) {
    const fnTag = `${this.className}#constructor()`;
    Checks.truthy(options, `${fnTag} arg options`);
    Checks.truthy(options.pluginRegistry, `${fnTag} options.pluginRegistry`);
    Checks.truthy(options.instanceId, `${fnTag} options.instanceId`);

    const level = this.options.logLevel || "INFO";
    const label = this.className;
    this.log = LoggerProvider.getOrCreate({ level, label });

    this.instanceId = options.instanceId;

    this.prometheusExporter =
      options.prometheusExporter ||
      new PrometheusExporter({ pollingIntervalInMin: 1 });
    Checks.truthy(
      this.prometheusExporter,
      `${fnTag} options.prometheusExporter`,
    );

    this.prometheusExporter.startMetricsCollection();
  }
  getOpenApiSpec(): unknown {
    throw new Error("Method not implemented.");
  }
  deployContract(): Promise<never> {
    throw new RuntimeError("Method not implemented.");
  }

  getConsensusAlgorithmFamily(): Promise<ConsensusAlgorithmFamily> {
    throw new Error("Method not implemented.");
  }
  hasTransactionFinality(): Promise<boolean> {
    throw new Error("Method not implemented.");
  }

  public getInstanceId(): string {
    return this.instanceId;
  }
  public async onPluginInit(): Promise<unknown> {
    return;
  }
  public async shutdown(): Promise<void> {
    this.log.info(`Shutting down ${this.className}...`);
  }

  async registerWebServices(app: Express): Promise<IWebServiceEndpoint[]> {
    const webServices = await this.getOrCreateWebServices();
    await Promise.all(webServices.map((ws) => ws.registerExpress(app)));
    return webServices;
  }
  public getPrometheusExporter(): PrometheusExporter {
    return this.prometheusExporter;
  }

  public async getPrometheusExporterMetrics(): Promise<string> {
    const res: string = await this.prometheusExporter.getPrometheusMetrics();
    this.log.debug(`getPrometheusExporterMetrics() response: %o`, res);
    return res;
  }

  public async getOrCreateWebServices(): Promise<IWebServiceEndpoint[]> {
    if (Array.isArray(this.endpoints)) {
      return this.endpoints;
    }
    const endpoints: IWebServiceEndpoint[] = [];
    {
      const endpoint = new RunTransactionEndpoint({
        connector: this,
        logLevel: this.options.logLevel,
      });
      endpoints.push(endpoint);
    }
    {
      const opts: IGetPrometheusExporterMetricsEndpointV1Options = {
        connector: this,
        logLevel: this.options.logLevel,
      };
      const endpoint = new GetPrometheusExporterMetricsEndpointV1(opts);
      endpoints.push(endpoint);
    }
    this.endpoints = endpoints;
    return endpoints;
  }

  public async transact(
    req: RunTransactionRequest,
  ): Promise<RunTransactionResponse> {
    const { baseConfig } = req;

    if (!baseConfig || !baseConfig.operatorKey || !baseConfig.operatorId) {
      this.log.debug(
        "Certain field within Hedera basic configuration is missing!",
      );
      throw new RuntimeError("Some fields in baseConfig is undefined");
    }

    const operatorId = AccountId.fromString(baseConfig.operatorId);
    const operatorKey = PrivateKey.fromString(baseConfig.operatorKey);
    const client = Client.forTestnet();
    switch (req.commandName) {
      case HederaQuery.GetAccountBalance: {
        try {
          //expects a req.params id of account to check balance for
          const accountId = AccountId.fromString(req.params[0]);
          client.setOperator(operatorId, operatorKey);
          const query = new AccountBalanceQuery().setAccountId(accountId);
          const accountBalance = await query.execute(client);
          console.log(
            "The hbar account balance for this account is " +
              accountBalance?.hbars,
          );
          return { transactionReceipt: accountBalance };
        } catch (err) {
          throw new RuntimeError(err);
        }
      }

      case HederaQuery.GetAccountInfo: {
        try {
          //expects a req.params id of account to check balance for
          const accountId = AccountId.fromString(req.params[0]);
          client.setOperator(operatorId, operatorKey);

          //Create the account info query
          const query = new AccountInfoQuery().setAccountId(accountId);

          //Sign with client operator private key and submit the query to a Hedera network
          const accountInfo = await query.execute(client);

          //Print the account info to the console
          console.log(accountInfo);

          //v2.0.0
          return { transactionReceipt: accountInfo };
        } catch (err) {
          throw new RuntimeError(err);
        }
      }

      case HederaQuery.GetTopicInfo: {
        try {
          //expects a topicId
          client.setOperator(operatorId, operatorKey);
          //Create the account info query
          const query = new TopicInfoQuery().setTopicId(req.params[0]);

          //Submit the query to a Hedera network
          const info = await query.execute(client);

          //Print the account key to the console
          console.log(info);
          return { transactionReceipt: info };
        } catch (err) {
          throw new RuntimeError(err);
        }
      }

      case HederaQuery.GetAccountTokenBalance: {
        try {
          const accountId = AccountId.fromString(req.params[0]);
          client.setOperator(operatorId, operatorKey);

          const query = new AccountBalanceQuery().setAccountId(accountId);

          const tokenBalance = await query.execute(client);

          console.log(
            "The token balance(s) for this account: " +
              tokenBalance.tokens.toString(),
          );

          return { transactionReceipt: tokenBalance };
        } catch (err) {
          throw new RuntimeError(err);
        }
      }

      case HederaQuery.GetTokenInfo: {
        try {
          //expects a tokenId
          client.setOperator(operatorId, operatorKey);
          //Create the query
          const query = new TokenInfoQuery().setTokenId(req.params[0]);

          //Sign with the client operator private key, submit the query to the network and get the token supply
          const tokenSupply = (await query.execute(client)).totalSupply;

          console.log("The total supply of this token is " + tokenSupply);

          //v2.0.7
          return { transactionReceipt: tokenSupply };
        } catch (err) {
          throw new RuntimeError(err);
        }
      }

      case HederaQuery.GetFileContents: {
        try {
          //expects a fileId
          client.setOperator(operatorId, operatorKey);

          const query = new FileContentsQuery().setFileId(req.params[0]);

          //Sign with client operator private key and submit the query to a Hedera network
          const contents = await query.execute(client);

          console.log(contents.toString());
          return { transactionReceipt: contents };
        } catch (err) {
          throw new RuntimeError(err);
        }
      }

      case HederaQuery.GetFileInfo: {
        try {
          client.setOperator(operatorId, operatorKey);
          const query = new FileInfoQuery().setFileId(req.params[0]);
          //Sign the query with the client operator private key and submit to a Hedera network
          const getInfo = await query.execute(client);

          console.log("File info response: " + getInfo);
          return { transactionReceipt: getInfo };
        } catch (err) {
          throw new RuntimeError(err);
        }
      }

      case HederaQuery.GetSmartContractBytecode: {
        try {
          //expects a contractId
          client.setOperator(operatorId, operatorKey);
          //Create the query
          const query = new ContractByteCodeQuery().setContractId(
            req.params[0],
          );

          //Sign with the client operator private key and submit to a Hedera network
          const bytecode = await query.execute(client);
          return { transactionReceipt: bytecode };
        } catch (err) {
          throw new RuntimeError(err);
        }
      }

      case HederaQuery.GetScheduleInfo: {
        try {
          //expects a schedule id
          client.setOperator(operatorId, operatorKey);
          //Create the query
          const query = new ScheduleInfoQuery().setScheduleId(req.params[0]);

          //Sign with the client operator private key and submit the query request to a node in a Hedera network
          const info = await query.execute(client);
          return { transactionReceipt: info };
        } catch (err) {
          throw new RuntimeError(err);
        }
      }

      case HederaCommand.CreateAccount: {
        //this method creates a new account

        try {
          client.setOperator(operatorId, operatorKey);

          const privateKey = await PrivateKey.generateED25519();
          const transaction = new AccountCreateTransaction()
            .setKey(privateKey.publicKey)
            .setInitialBalance(new Hbar(1000));
          const txResponse = await transaction.execute(client);
          const receipt = await txResponse.getReceipt(client);
          const newAccountId = receipt.accountId;

          console.log("The new account ID is " + newAccountId);
          const accountBalance = await new AccountBalanceQuery()
            .setAccountId(newAccountId)
            .execute(client);

          console.log(
            "The new account balance is: " +
              accountBalance.hbars.toTinybars() +
              " tinybar.",
          );
          return { transactionReceipt: newAccountId };
        } catch (err) {
          throw new RuntimeError(err);
        }
      }

      case HederaCommand.TransferHbars: {
        //this method transfers hbars from one account to another
        //re.params contain myAccountId, myPrivateKey, toAccountId, amount
        try {
          const myAccountId = AccountId.fromString(req.params[0]);
          const myPrivateKey = PrivateKey.fromString(req.params[1]);
          const toAccountId = AccountId.fromString(req.params[2]);
          client.setOperator(myAccountId, myPrivateKey);
          const sendHbar = await new TransferTransaction()
            .addHbarTransfer(myAccountId, Hbar.fromTinybars(-req.params[3]))
            .addHbarTransfer(toAccountId, Hbar.fromTinybars(req.params[3]))
            .execute(client);
          const transactionReceipt = await sendHbar.getReceipt(client);
          console.log(
            "The transfer transaction from my account to the another account was: " +
              transactionReceipt.status.toString(),
          );

          const queryCost = await new AccountBalanceQuery()
            .setAccountId(req.params[2])
            .getCost(client);

          console.log("The cost of query is: " + queryCost);

          const getNewBalance = await new AccountBalanceQuery()
            .setAccountId(req.params[2])
            .execute(client);

          console.log(
            "The account balance after the transfer is: " +
              getNewBalance.hbars.toTinybars() +
              " tinybar.",
          );
          return { transactionReceipt: transactionReceipt };
        } catch (err) {
          throw new RuntimeError(err);
        }
      }

      case HederaCommand.TransferToken: {
        //this method transfers hbars from one account to another
        //re.params contain myAccountId, myPrivateKey, toAccountId, amount, tokenId
        try {
          const myAccountId = AccountId.fromString(req.params[0]);
          const myPrivateKey = PrivateKey.fromString(req.params[1]);
          const toAccountId = AccountId.fromString(req.params[2]);
          client.setOperator(myAccountId, myPrivateKey);
          const sendHbar = await new TransferTransaction()
            .addTokenTransfer(req.params[4], myAccountId, -req.params[3])
            .addTokenTransfer(req.params[4], toAccountId, req.params[3])
            .execute(client);
          const receipt = await sendHbar.getReceipt(client);
          const transactionStatus = receipt.status;

          console.log(
            "The transaction consensus status " + transactionStatus.toString(),
          );
          return { transactionReceipt: transactionStatus };
        } catch (err) {
          throw new RuntimeError(err);
        }
      }

      case HederaCommand.CreateTopic: {
        //this method creates a topic
        //it expects a topic memo fro the req.params
        try {
          client.setOperator(operatorId, operatorKey);
          //Create the transaction
          const transaction = new TopicCreateTransaction().setTopicMemo(
            req.params[0],
          );

          //Sign with the client operator private key and submit the transaction to a Hedera network
          const txResponse = await transaction.execute(client);

          //Request the receipt of the transaction
          const receipt = await txResponse.getReceipt(client);

          //Get the topic ID
          const newTopicId = receipt.topicId;

          console.log("The new topic ID is " + newTopicId);

          return { transactionReceipt: newTopicId };
        } catch (err) {
          throw new RuntimeError(err);
        }
      }

      case HederaCommand.SubmitMessage: {
        //this method submits a message to the method
        //it expects an accountId, accountKey, topic Id, Message
        const myAccountId = AccountId.fromString(req.params[0]);
        const myPrivateKey = PrivateKey.fromString(req.params[1]);
        try {
          client.setOperator(myAccountId, myPrivateKey);
          //Create the transaction
          const transaction = await new TopicMessageSubmitTransaction({
            topicId: req.params[2],
            message: req.params[3],
          }).execute(client);

          return { transactionReceipt: transaction };
        } catch (err) {
          throw new RuntimeError(err);
        }
      }

      case HederaCommand.CreateToken: {
        const adminAccountId = AccountId.fromString(req.params[0]);
        const adminPrivateKey = PrivateKey.fromString(req.params[1]);
        const adminPublicKey = PublicKey.fromString(req.params[2]);
        const treasuryAccountId = AccountId.fromString(req.params[3]);
        const treasuryKey = PrivateKey.fromString(req.params[4]);

        //this method submits a message to the method
        //it expects an adminAccountId, adminPrivateKey, adminPublicKey,treasuryAccountId,treasuryKey, topic Id, Message

        try {
          client.setOperator(adminAccountId, adminPrivateKey);
          //Create the transaction and freeze for manual signing
          //This method expects a token name, symbol
          const transaction = await new TokenCreateTransaction()
            .setTokenName(req.params[5])
            .setTokenSymbol(req.params[6])
            .setTreasuryAccountId(treasuryAccountId)
            .setInitialSupply(0)
            .setAdminKey(adminPublicKey)
            .setMaxTransactionFee(new Hbar(30)) //Change the default max transaction fee
            .freezeWith(client);

          //Sign the transaction with the token adminKey and the token treasury account private key
          const signTx = await (await transaction.sign(adminPrivateKey)).sign(
            treasuryKey,
          );

          //Sign the transaction with the client operator private key and submit to a Hedera network
          const txResponse = await signTx.execute(client);

          //Get the receipt of the transaction
          const receipt = await txResponse.getReceipt(client);

          //Get the token ID from the receipt
          const tokenId = receipt.tokenId;

          console.log("The new token ID is " + tokenId);

          //v2.0.5
          return { transactionReceipt: tokenId };
        } catch (err) {
          throw new RuntimeError(err);
        }
      }

      //////////CREATE NFT NO CUSTOM FEE SCHEDULE FOR THIS, WILL BE IMPLEMENTED LATER
      case HederaCommand.CreateNft: {
        const adminAccountId = AccountId.fromString(req.params[0]);
        const adminPrivateKey = PrivateKey.fromString(req.params[1]);
        const adminPublicKey = PublicKey.fromString(req.params[2]);
        const treasuryAccountId = AccountId.fromString(req.params[3]);
        const treasuryKey = PrivateKey.fromString(req.params[4]);

        //this method submits a message to the method
        //it expects an adminAccountId, adminPrivateKey, adminPublicKey,treasuryAccountId,treasuryKey, topic Id, Message,max supply

        try {
          client.setOperator(adminAccountId, adminPrivateKey);
          //Create the transaction and freeze for manual signing
          //This method expects a token name, symbol
          const transaction = await new TokenCreateTransaction()
            .setTokenName(req.params[5])
            .setTokenSymbol(req.params[6])
            .setTokenType(TokenType.NonFungibleUnique)
            .setTreasuryAccountId(treasuryAccountId)
            .setDecimals(0)
            .setInitialSupply(0)
            .setAdminKey(adminPublicKey)
            .setMaxSupply(req.params[7])
            .setMaxTransactionFee(new Hbar(100)) //Change the default max transaction fee
            .freezeWith(client);

          //Sign the transaction with the token adminKey and the token treasury account private key
          const signTx = await (await transaction.sign(adminPrivateKey)).sign(
            treasuryKey,
          );

          //Sign the transaction with the client operator private key and submit to a Hedera network
          const txResponse = await signTx.execute(client);

          //Get the receipt of the transaction
          const receipt = await txResponse.getReceipt(client);

          //Get the token ID from the receipt
          const tokenId = receipt.tokenId;

          console.log("The new token ID is " + tokenId);

          //v2.0.5
          return { transactionReceipt: tokenId };
        } catch (err) {
          throw new RuntimeError(err);
        }
      }

      case HederaCommand.CreateFile: {
        //this method creates a file
        //expects an filePublicKey, file privateKey, file content,
        try {
          const filePublicKey = PublicKey.fromString(req.params[0]);
          const fileKey = PrivateKey.fromString(req.params[1]);
          client.setOperator(operatorId, operatorKey);
          //Create the transaction
          const transaction = await new FileCreateTransaction()
            .setKeys([filePublicKey]) //A different key then the client operator key
            .setContents("the file contents")
            .setMaxTransactionFee(new Hbar(2))
            .freezeWith(client);

          //Sign with the file private key
          const signTx = await transaction.sign(fileKey);

          //Sign with the client operator private key and submit to a Hedera network
          const submitTx = await signTx.execute(client);

          //Request the receipt
          const receipt = await submitTx.getReceipt(client);

          //Get the file ID
          const newFileId = receipt.fileId;

          console.log("The new file ID is: " + newFileId);
          return { transactionReceipt: newFileId };
        } catch (err) {
          throw new RuntimeError(err);
        }
      }

      default: {
        throw new RuntimeError(
          "command or query does not exist, or is not supported in current version",
        );
      }
    }
  }

  public getPackageName(): string {
    return `@hyperledger/cactus-plugin-ledger-connector-hedera`;
  }
}
