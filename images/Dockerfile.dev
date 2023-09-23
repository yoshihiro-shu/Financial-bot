FROM golang:1.21

ENV ROOT=/usr/src/app
WORKDIR ${ROOT}

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./ ./
RUN go mod tidy
EXPOSE 8080

CMD ["go", "run", "./main.go"]