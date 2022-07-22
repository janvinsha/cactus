export {
  IPluginLedgerConnectorHederaOptions,
  PluginLedgerConnectorHedera,
} from "./plugin-ledger-connector-hedera";
export { PluginFactoryLedgerConnector } from "./plugin-factory-ledger-connector";

import { IPluginFactoryOptions } from "@hyperledger/cactus-core-api";
import { PluginFactoryLedgerConnector } from "./plugin-factory-ledger-connector";

export {
  HederaApiClient,
  HederaApiClientOptions,
} from "./api-client/hedera-api-client";

export * from "./generated/openapi/typescript-axios/api";

export async function createPluginFactory(
  pluginFactoryOptions: IPluginFactoryOptions,
): Promise<PluginFactoryLedgerConnector> {
  return new PluginFactoryLedgerConnector(pluginFactoryOptions);
}
