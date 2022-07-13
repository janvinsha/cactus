# `@hyperledger/cactus-plugin-ledger-connector-hedera`

This project will provide integration between Hyperledger `Cactus` and a non-hyperledger project Hedera Hashgraph, outcome of the project:

- Documented, ready-to-use integration of Hedera and Cactus.
- Documented example of integration using Hedera's Testnet and Hedera's Previewnet.
- Documented example of integration between another blockchain and Hedera using Cactus.

## Summary

- [Getting Started](#getting-started)
- [Architecture](#architecture)
- [Usage](#usage)
- [Runing the tests](#running-the-tests)
- [Built With](#built-with)
- [Contributing](#contributing)
- [License](#license)
- [Acknowledgments](#acknowledgments)

## Getting Started

Clone the git repository on your local machine. Follow these instructions that will get you a copy of the project up and running on
your local machine for development and testing purposes.

### Prerequisites

In the root of the project to install the dependencies execute the command:

```sh
npm run configure
```

### Compiling

In the project root folder, run this command to compile the plugin and create the dist directory:

```sh
npm run tsc
```

### Architecture

The sequence diagrams for various endpoints are mentioned below

#### run-transaction-endpoint

### Usage

> Extensive documentation and examples in the [readthedocs](https://readthedocs.org/projects/hyperledger-cactus/) (WIP)

### Building/running the container image locally

In the Cactus project root say:

```sh
DOCKER_BUILDKIT=1 docker build -f ./packages/cactus-plugin-ledger-connector-hedera/Dockerfile . -t cactus-plugin-ledger-connector-hedera
```

Build with a specific version of the npm package:

```sh
DOCKER_BUILDKIT=1 docker build --build-arg NPM_PKG_VERSION=latest -f ./packages/cactus-plugin-ledger-connector-hedera/Dockerfile . -t cactus-plugin-ledger-connector-hedera
```

#### Running the container

Launch container with plugin configuration as an **environment variable**:

```sh
docker run \
  --rm \
  --publish 3000:3000 \
  --publish 4000:4000 \
  --publish 9000:9000 \
  --env PLUGINS='[{"packageName": "@hyperledger/cactus-plugin-ledger-connector-hedera", "type": "org.hyperledger.cactus.plugin_import_type.LOCAL", "action": "org.hyperledger.cactus.plugin_import_action.INSTALL",  "options": {"rpcApiHttpHost": "http://localhost:8545", "instanceId": "some-unique-hedera-connector-instance-id"}}]' \
  cactus-plugin-ledger-connector-hedera
```

Launch container with plugin configuration as a **CLI argument**:

```sh
docker run \
  --rm \
  --publish 3000:3000 \
  --publish 4000:4000 \
  --publish 9000:9000 \
  cactus-plugin-ledger-connector-hedera \
    ./node_modules/.bin/cactusapi \
    --plugins='[{"packageName": "@hyperledger/cactus-plugin-ledger-connector-hedera", "type": "org.hyperledger.cactus.plugin_import_type.LOCAL", "action": "org.hyperledger.cactus.plugin_import_action.INSTALL",  "options": {"rpcApiHttpHost": "http://localhost:8545", "instanceId": "some-unique-hedera-connector-instance-id"}}]'
```

Launch container with **configuration file** mounted from host machine:

```sh

echo '[{"packageName": "@hyperledger/cactus-plugin-ledger-connector-hedera", "type": "org.hyperledger.cactus.plugin_import_type.LOCAL", "action": "org.hyperledger.cactus.plugin_import_action.INSTALL", "options": {"rpcApiHttpHost": "http://localhost:8545", "instanceId": "some-unique-hedera-connector-instance-id"}}]' > cactus.json

docker run \
  --rm \
  --publish 3000:3000 \
  --publish 4000:4000 \
  --publish 9000:9000 \
  --mount type=bind,source="$(pwd)"/cactus.json,target=/cactus.json \
  cactus-plugin-ledger-connector-hedera \
    ./node_modules/.bin/cactusapi \
    --config-file=/cactus.json
```

#### Testing API calls with the container

**Terminal Window 1 (Cactus API Server)**

```sh
docker run \
  --network host \
  --rm \
  --publish 3000:3000 \
  --publish 4000:4000 \
  --publish 9000:9000 \
  --env AUTHORIZATION_CONFIG_JSON="{}"\
  --env AUTHORIZATION_PROTOCOL=NONE \
  --env PLUGINS='[{"packageName": "@hyperledger/cactus-plugin-ledger-connector-hedera", "type": "org.hyperledger.cactus.plugin_import_type.LOCAL", "action": "org.hyperledger.cactus.plugin_import_action.INSTALL", "options": {"rpcApiHttpHost": "http://localhost:8545", "instanceId": "some-unique-hedera-connector-instance-id"}}]' \
  cactus-plugin-ledger-connector-hedera
```

**Terminal Window 2 (curl - replace transaction request as needed)**

```sh
curl --location --request POST 'http://127.0.0.1:4000/api/v1/plugins/@hyperledger/cactus-plugin-ledger-connector-hedera/run-transaction' \
--header 'Content-Type: application/json' \
--data-raw '{
  commandName: 'createAccount',
  baseConfig: {
    operatorId: 'admin@test',
    operatorKey:
      'f101537e319568c765b2cc89698325604991dca57b9716b58016b253506cab70',
    timeoutLimit: 5000
  },
  params: []
}'
```

The above should produce a response that looks similar to this:

```json
{
  "success": true,
  "data": {
    "transactionReceipt": {
      "txHash": "c3ffd772f26950243aa357ab4f21b9703d5172490b66ddc285355230d6df60b8",
      "status": "COMMITTED"
    }
  }
}
```

## Running the tests

To check that all has been installed correctly and that the plugin has no errors run the tests:

- Run this command at the project's root:

```sh
yarn run test:plugin-ledger-connector-hedera
```

## Contributing

We welcome contributions to Hyperledger Cactus in many forms, and there’s always plenty to do!

Please review [CONTIRBUTING.md](../../CONTRIBUTING.md) to get started.

## License

This distribution is published under the Apache License Version 2.0 found in the [LICENSE](../../LICENSE) file.

## Acknowledgments
