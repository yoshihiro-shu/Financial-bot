FROM golang:1.21-alpine3.18

ENV ROOT=/usr/src/app
WORKDIR ${ROOT}

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change

COPY . .
RUN go mod tidy

EXPOSE 8000

ENTRYPOINT [ "go", "run", "/usr/src/app/apps/batch/batch.go", "-f", "/usr/src/app/apps/batch/etc/batch.yaml" ]
