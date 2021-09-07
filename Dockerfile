FROM golang

WORKDIR /app-watcher

COPY app-watcher .

RUN go mod download && \
    go install .

EXPOSE 8080

ENTRYPOINT ["app-watcher"]

