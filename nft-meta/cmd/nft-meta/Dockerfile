FROM debian:stable-slim

COPY nft-meta /usr/local/bin/nft-meta
COPY NFTMeta.viper.yaml /etc/NFTMeta/NFTMeta.viper.yaml
ENV ENV_DATA_DIR=/data \
    ENV_LOG_DIR=/var/log \
    ENV_LOG_LEVEL=debug \
    ENV_ETH_ENDPOINT="" \
    ENV_INVERSTOR_KEY=""
RUN mkdir $ENV_DATA_DIR
RUN apt-get  update \
    && apt-get install -y --no-install-recommends ca-certificates curl

WORKDIR /usr/local/bin/

CMD ./nft-meta run