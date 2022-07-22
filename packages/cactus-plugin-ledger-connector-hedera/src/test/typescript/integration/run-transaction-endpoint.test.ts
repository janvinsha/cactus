import http from "http";
import { AddressInfo } from "net";
import test, { Test } from "tape-promise/tape";
import { v4 as uuidv4 } from "uuid";
import { v4 as internalIpV4 } from "internal-ip";
import bodyParser from "body-parser";
import express from "express";

import {
  Containers,
  pruneDockerAllIfGithubAction,
} from "@hyperledger/cactus-test-tooling";
import { PluginRegistry } from "@hyperledger/cactus-core";
import { PluginImportType } from "@hyperledger/cactus-core-api";

import {
  IListenOptions,
  LogLevelDesc,
  Servers,
} from "@hyperledger/cactus-common";
import { RuntimeError } from "run-time-error";
import {
  PluginLedgerConnectorHedera,
  DefaultApi as HederaApi,
  PluginFactoryLedgerConnector,
} from "../../../main/typescript/public-api";

import { Configuration } from "@hyperledger/cactus-core-api";

import {
  HederaCommand,
  HederaQuery,
} from "../../../main/typescript/generated/openapi/typescript-axios";

const testCase = "runs tx on an hedera v1.2.0 ledger";
const logLevel: LogLevelDesc = "INFO";

test.onFailure(async () => {
  await Containers.logDiagnostics({ logLevel });
});

test("BEFORE " + testCase, async (t: Test) => {
  const hederaHost = await internalIpV4();
  if (!hederaHost) {
    throw new RuntimeError("Could not determine the internal IPV4 address.");
  }

  const pruning = pruneDockerAllIfGithubAction({ logLevel });
  await t.doesNotReject(pruning, "Pruning didn't throw OK");
  t.end();
});

// Test fails because hedera is unable to connect to Postgres for some reason.
test(testCase, async (t: Test) => {
  const factory = new PluginFactoryLedgerConnector({
    pluginImportType: PluginImportType.Local,
  });

  const connector: PluginLedgerConnectorHedera = await factory.create({
    instanceId: uuidv4(),
    pluginRegistry: new PluginRegistry(),
  });

  const expressApp = express();
  expressApp.use(bodyParser.json({ limit: "250mb" }));
  const server = http.createServer(expressApp);
  const listenOptions: IListenOptions = {
    hostname: "localhost",
    port: 0,
    server,
  };

  const addressInfo = (await Servers.listen(listenOptions)) as AddressInfo;
  test.onFinish(async () => await Servers.shutdown(server));
  const { address, port } = addressInfo;
  const apiHost = `http://${address}:${port}`;
  const apiConfig = new Configuration({ basePath: apiHost });
  const apiClient = new HederaApi(apiConfig);

  await connector.getOrCreateWebServices();
  await connector.registerWebServices(expressApp);

  const adminID = `1234`;
  const adminPriv = "1234";

  {
    const req = {
      commandName: HederaCommand.CreateAccount,
      baseConfig: {
        operatorId: adminID,
        operatorKey: adminPriv,
      },
      params: [],
    };
    const res = await apiClient.runTransaction(req);
    t.ok(res);
    t.ok(res.data);
    t.equal(res.status, 200);
  }

  const transferAmount = 100;
  {
    const req = {
      commandName: HederaCommand.TransferHbars,
      baseConfig: {
        operatorId: adminID,
        operatorKey: adminPriv,
      },
      params: [adminID, adminPriv, adminID, transferAmount],
    };
    const res = await apiClient.runTransaction(req);
    t.ok(res);
    t.ok(res.data);
    t.equal(res.status, 200);
  }

  const tokenId = "0.0.46018668";
  {
    const req = {
      commandName: HederaCommand.TransferToken,
      baseConfig: {
        operatorId: adminID,
        operatorKey: adminPriv,
      },
      params: [adminID, adminPriv, adminID, transferAmount, tokenId],
    };
    const res = await apiClient.runTransaction(req);
    t.ok(res);
    t.ok(res.data);
    t.equal(res.status, 200);
  }

  const topicMemo = "Lorem Ipsum";
  {
    const req = {
      commandName: HederaCommand.CreateTopic,
      baseConfig: {
        operatorId: adminID,
        operatorKey: adminPriv,
      },
      params: [topicMemo],
    };
    const res = await apiClient.runTransaction(req);
    t.ok(res);
    t.ok(res.data);
    t.equal(res.status, 200);
  }

  const topicId = "0.0.46018668";
  const message = "Lorem Ipsum";
  {
    const req = {
      commandName: HederaCommand.SubmitMessage,
      baseConfig: {
        operatorId: adminID,
        operatorKey: adminPriv,
      },
      params: [adminID, adminPriv, topicId, message],
    };
    const res = await apiClient.runTransaction(req);
    t.ok(res);
    t.ok(res.data);
    t.equal(res.status, 200);
  }

  const adminPublicKey = "1234";
  {
    const req = {
      commandName: HederaCommand.CreateToken,
      baseConfig: {
        operatorId: adminID,
        operatorKey: adminPriv,
      },
      params: [
        adminID,
        adminPriv,
        adminPublicKey,
        adminID,
        adminPriv,
        topicId,
        message,
      ],
    };
    const res = await apiClient.runTransaction(req);
    t.ok(res);
    t.ok(res.data);
    t.equal(res.status, 200);
  }

  const maxSupply = 100;
  {
    const req = {
      commandName: HederaCommand.CreateNft,
      baseConfig: {
        operatorId: adminID,
        operatorKey: adminPriv,
      },
      params: [
        adminID,
        adminPriv,
        adminPublicKey,
        adminID,
        adminPriv,
        topicId,
        message,
        maxSupply,
      ],
    };
    const res = await apiClient.runTransaction(req);
    t.ok(res);
    t.ok(res.data);
    t.equal(res.status, 200);
  }

  const filePublicKey = "1234";
  const filePrivateKey = "1234";
  const fileContent = "11";
  {
    const req = {
      commandName: HederaCommand.CreateFile,
      baseConfig: {
        operatorId: adminID,
        operatorKey: adminPriv,
      },
      params: [filePublicKey, filePrivateKey, fileContent],
    };
    const res = await apiClient.runTransaction(req);
    t.ok(res);
    t.ok(res.data);
    t.equal(res.status, 200);
  }

  {
    const req = {
      commandName: HederaQuery.GetAccountBalance,
      baseConfig: {
        operatorId: adminID,
        operatorKey: adminPriv,
      },
      params: [adminID],
    };
    const res = await apiClient.runTransaction(req);
    t.ok(res);
    t.ok(res.data);
    t.equal(res.status, 200);
  }

  {
    const req = {
      commandName: HederaQuery.GetAccountInfo,
      baseConfig: {
        operatorId: adminID,
        operatorKey: adminPriv,
      },
      params: [adminID],
    };
    const res = await apiClient.runTransaction(req);
    t.ok(res);
    t.ok(res.data);
    t.equal(res.status, 200);
  }

  {
    const req = {
      commandName: HederaQuery.GetTopicInfo,
      baseConfig: {
        operatorId: adminID,
        operatorKey: adminPriv,
      },
      params: [topicId],
    };
    const res = await apiClient.runTransaction(req);
    t.ok(res);
    t.ok(res.data);
    t.equal(res.status, 200);
  }

  {
    const req = {
      commandName: HederaQuery.GetAccountTokenBalance,
      baseConfig: {
        operatorId: adminID,
        operatorKey: adminPriv,
      },
      params: [adminID],
    };
    const res = await apiClient.runTransaction(req);
    t.ok(res);
    t.ok(res.data);
    t.equal(res.status, 200);
  }

  {
    const req = {
      commandName: HederaQuery.GetTokenInfo,
      baseConfig: {
        operatorId: adminID,
        operatorKey: adminPriv,
      },
      params: [tokenId],
    };
    const res = await apiClient.runTransaction(req);
    t.ok(res);
    t.ok(res.data);
    t.equal(res.status, 200);
  }
  const fileId = "1234";
  {
    const req = {
      commandName: HederaQuery.GetFileContents,
      baseConfig: {
        operatorId: adminID,
        operatorKey: adminPriv,
      },
      params: [fileId],
    };
    const res = await apiClient.runTransaction(req);
    t.ok(res);
    t.ok(res.data);
    t.equal(res.status, 200);
  }

  {
    const req = {
      commandName: HederaQuery.GetFileInfo,
      baseConfig: {
        operatorId: adminID,
        operatorKey: adminPriv,
      },
      params: [fileId],
    };
    const res = await apiClient.runTransaction(req);
    console.log(res.data.transactionReceipt);
    t.ok(res);
    t.ok(res.data);
    t.equal(res.status, 200);
  }
  const contractId = "1234";
  {
    const req = {
      commandName: HederaQuery.GetSmartContractBytecode,
      baseConfig: {
        operatorId: adminID,
        operatorKey: adminPriv,
      },
      params: [contractId],
    };
    const res = await apiClient.runTransaction(req);
    console.log(res.data.transactionReceipt);
    t.ok(res);
    t.ok(res.data);
    t.equal(res.status, 200);
  }
  const scheduleId = "1234";
  {
    const req = {
      commandName: HederaQuery.GetScheduleInfo,
      baseConfig: {
        operatorId: adminID,
        operatorKey: adminPriv,
      },

      params: [scheduleId],
    };
    const res = await apiClient.runTransaction(req);
    t.ok(res);
    t.ok(res.data);
    t.equal(res.status, 200);
  }

  t.end();
});

test("AFTER " + testCase, async (t: Test) => {
  const pruning = pruneDockerAllIfGithubAction({ logLevel });
  await t.doesNotReject(pruning, "Pruning didn't throw OK");
  t.end();
});
