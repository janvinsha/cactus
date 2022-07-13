# Hedera All in One Image

An all in one hedera + Orion image that uses about 0.5 GB of RAM and one vCPU worth of hardware resources idling.
This makes it much easier to pull up for tests that are about functionality rather than running a production grade ledger.

Comes equipped with supervisord which provides access to tailing logs of both orion and hedera via a web based user interface that you can access via the port 9001 of the container (granted you exposed it on your host).

## Build an image locally

- To build the hedera v1.5.1 image locally use:

  ```sh
  docker build . -t hyperledger/cactus-hedera-all-in-one:latest
  ```

- To build the hedera v21.1.x image locally use:

  ```sh
  docker build ./tools/docker/hedera-all-in-one/ -f ./tools/docker/hedera-all-in-one/v21_1_x/Dockerfile -t baio21
  ```

## Shell into a running container:

```sh
docker run -it --entrypoint bash hyperledger/cactus-hedera-all-in-one
```

## Easiest way to run the image with defaults

```sh
docker-compose up
```

The following ports are open on the container:

```yaml
- 8545:8545/tcp # hedera RPC - HTTP
- 8546:8546/tcp # hedera RPC - WebSocket
- 8888:8888/tcp # orion - HTTP
- 9001:9001/tcp # supervisord - HTTP
- 9545:9545/tcp # hedera metrics
```

## Without docker-compose

```sh
docker run -p 0.0.0.0:8545:8545/tcp  -p 0.0.0.0:8546:8546/tcp  -p 0.0.0.0:8888:8888/tcp  -p 0.0.0.0:9001:9001/tcp  -p 0.0.0.0:9545:9545/tcp hyperledger/cactus-hedera-all-in-one:latest
```

## Logs of hedera and Orion via supervisord web UI:

Navigate your browser to http://localhost:9001
