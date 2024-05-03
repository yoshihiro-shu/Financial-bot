FROM golang:1.22.2-alpine3.18 AS builder
LABEL org.opencontainers.image.source https://github.com/yoshihiro-shu/financial-bot

ENV CGO_ENABLED 0

RUN apk update --no-cache && apk add --no-cache tzdata

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download

COPY apps/notification ./apps/notification
COPY internal ./internal
COPY entity ./entity
COPY repository ./repository
COPY usecase ./usecase

RUN go build -ldflags="-s -w" -o /app/notification /build/apps/notification/notification.go


FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Tokyo /usr/share/zoneinfo/Asia/Tokyo
ENV TZ Asia/Tokyo

WORKDIR /app
COPY --from=builder /app/notification /app/notification
COPY --from=builder /build/apps/notification/etc/notification.yaml /app/etc/notification.yaml

CMD ["./notification"]
