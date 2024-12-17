FROM --platform=linux/amd64 golang:1.17

ARG SERVICE_NAME
ENV ENV_SERVICE_NAME=$SERVICE_NAME

RUN GO111MODULE=on go get github.com/cespare/reflex@latest

RUN apt-get update && \
    apt-get install -y ffmpeg && \
    rm -rf /var/lib/apt/lists/*

# change this if you have place the code dir in other place
WORKDIR /go/src/personal/audio_retrieval/audio_retrieval
RUN echo "-r '(\.go$|go\.mod)' -s go run ./cmd/$ENV_SERVICE_NAME/" > /reflex.conf
ENTRYPOINT ["reflex", "-c", "/reflex.conf"]