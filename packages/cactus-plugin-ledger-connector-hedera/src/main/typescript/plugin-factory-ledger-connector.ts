import {
  IPluginFactoryOptions,
  PluginFactory,
} from "@hyperledger/cactus-core-api";
import {
  IPluginLedgerConnectorHederaOptions,
  PluginLedgerConnectorHedera,
} from "./plugin-ledger-connector-hedera";

export class PluginFactoryLedgerConnector extends PluginFactory<
  PluginLedgerConnectorHedera,
  IPluginLedgerConnectorHederaOptions,
  IPluginFactoryOptions
> {
  async create(
    pluginOptions: IPluginLedgerConnectorHederaOptions,
  ): Promise<PluginLedgerConnectorHedera> {
    return new PluginLedgerConnectorHedera(pluginOptions);
  }
}
