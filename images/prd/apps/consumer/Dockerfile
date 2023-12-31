FROM golang:1.22rc1-alpine3.18 AS builder
LABEL org.opencontainers.image.source https://github.com/yoshihiro-shu/financial-bot

ENV CGO_ENABLED 0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY apps/consumer ./apps/consumer
COPY internal ./internal
COPY entity ./entity
COPY repository ./repository
COPY usecase ./usecase

RUN go build -ldflags="-s -w" -o /app/consumer /build/apps/consumer/*.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Tokyo /usr/share/zoneinfo/Asia/Tokyo
ENV TZ Asia/Tokyo

WORKDIR /app
COPY --from=builder /app/consumer /app/consumer

CMD ["./consumer"]
