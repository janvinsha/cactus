import http from "http";
import { AddressInfo } from "net";
import test, { Test } from "tape-promise/tape";
import { v4 as uuidv4 } from "uuid";

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

import {
  PluginLedgerConnectorHedera,
  DefaultApi as HederaApi,
  PluginFactoryLedgerConnector,
} from "../../../../main/typescript/public-api";

import { Configuration } from "@hyperledger/cactus-core-api";

import {
  HederaQuery,
  RunTransactionRequest,
} from "../../../../main/typescript/generated/openapi/typescript-axios";

import OAS from "../../../../main/json/openapi.json";
import { installOpenapiValidationMiddleware } from "@hyperledger/cactus-core";

const testCase = "Hedera plugin openapi validation";
const logLevel: LogLevelDesc = "INFO";

test.onFailure(async () => {
  await Containers.logDiagnostics({ logLevel });
});

test("BEFORE " + testCase, async (t: Test) => {
  const pruning = pruneDockerAllIfGithubAction({ logLevel });
  await t.doesNotReject(pruning, "Pruning didn't throw OK");
  t.end();
});

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

  await installOpenapiValidationMiddleware({
    logLevel,
    app: expressApp,
    apiSpec: OAS,
  });

  await connector.getOrCreateWebServices();
  await connector.registerWebServices(expressApp);

  const adminID = `1234`;
  const adminPriv = "1234";

  const fRun = "runTransaction";
  const cOk = "without bad request error";
  const cWithoutParams = "not sending all required parameters";
  const cInvalidParams = "sending invalid parameters";

  test(`${testCase} - ${fRun} - ${cOk}`, async (t2: Test) => {
    const parameters = {
      commandName: HederaQuery.GetAccountInfo,
      baseConfig: {
        operatorId: adminID,
        operatoryKey: adminPriv,
      },
      params: [adminID],
    };
    const res = await apiClient.runTransaction(parameters);
    t2.ok(res);
    t2.equal(res.status, 200);

    t2.end();
  });

  test(`${testCase} - ${fRun} - ${cWithoutParams}`, async (t2: Test) => {
    try {
      const parameters = {
        commandName: HederaQuery.GetAccountInfo,
        baseConfig: {
          operatorId: adminID,
          operatorKey: adminPriv,
        },
        // params: [adminID],
      };
      await apiClient.runTransaction(
        (parameters as any) as RunTransactionRequest,
      );
    } catch (e) {
      t2.equal(
        e.response.status,
        400,
        `Endpoint ${fRun} without required params: response.status === 400 OK`,
      );
      const fields = e.response.data.map((param: any) =>
        param.path.replace(".body.", ""),
      );
      t2.ok(fields.includes("params"), "Rejected because params is required");
    }
    t2.end();
  });

  test(`${testCase} - ${fRun} - ${cInvalidParams}`, async (t2: Test) => {
    try {
      const parameters = {
        commandName: HederaQuery.GetAccountInfo,
        baseConfig: {
          operatorId: adminID,
          operatorKey: adminPriv,
        },
        params: [adminID],
        fake: 4,
      };
      await apiClient.runTransaction(
        (parameters as any) as RunTransactionRequest,
      );
    } catch (e) {
      t2.equal(
        e.response.status,
        400,
        `Endpoint ${fRun} with fake=4: response.status === 400 OK`,
      );
      const fields = e.response.data.map((param: any) =>
        param.path.replace(".body.", ""),
      );
      t2.ok(
        fields.includes("fake"),
        "Rejected because fake is not a valid parameter",
      );
    }
    t2.end();
  });

  t.end();
});

test("AFTER " + testCase, async (t: Test) => {
  const pruning = pruneDockerAllIfGithubAction({ logLevel });
  await t.doesNotReject(pruning, "Pruning didn't throw OK");
  t.end();
});
