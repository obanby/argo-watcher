FROM golang

WORKDIR /app-watcher

COPY argo-watcher .

RUN go mod download && \
    go install .

EXPOSE 8080

ENTRYPOINT ["argo-watcher"]
