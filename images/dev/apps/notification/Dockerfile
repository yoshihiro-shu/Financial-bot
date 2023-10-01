FROM golang:1.21-alpine3.18

ENV ROOT=/usr/src/app
WORKDIR ${ROOT}

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change

COPY go.mod ./go.mod
COPY go.sum ./go.sum

RUN go mod download

COPY apps/batch ./apps/notification
COPY internal ./internal

EXPOSE 8000

ENTRYPOINT [ "go", "run", "/usr/src/app/apps/notification/notification.go", "-f", "/usr/src/app/apps/notification/etc/notification.yaml" ]
