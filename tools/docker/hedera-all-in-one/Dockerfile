ARG HEDERA_VERSION=1.5.1
ARG ORION_VERSION=1.6

FROM hyperledger/hedera:$HEDERA_VERSION AS hedera
FROM pegasyseng/orion:1.5.1 AS orion

COPY --from=hedera /opt/hedera/ /opt/hedera/

RUN mkdir /config/
RUN mkdir /config/orion/
RUN mkdir /config/hedera/
RUN mkdir /opt/hedera/public-keys/

ADD ./log-config.xml /config/log-config.xml
ADD ./orion.conf /config/orion/orion.conf
ADD ./nodeKey.pub /config/orion/nodeKey.pub
ADD ./nodeKey.key /config/orion/nodeKey.key
ADD ./bootnode_start.sh /opt/hedera/bootnode_start.sh
ADD ./key /opt/hedera/keys/key
ADD ./key.pub /opt/hedera/keys/key.pub

RUN apt-get update && apt-get install -y supervisor
RUN mkdir -p /var/log/supervisor
COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

EXPOSE 9001

ENTRYPOINT ["/usr/bin/supervisord"]
CMD ["--nodaemon"]
