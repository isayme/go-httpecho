FROM golang:1.13.4-alpine AS builder

ARG APP_PKG
WORKDIR /go/src/${APP_PKG}

COPY . .

ARG APP_NAME
ARG APP_VERSION
RUN CGO_ENABLED=0 go build -ldflags "-X ${APP_PKG}/app.Name=${APP_NAME} -X ${APP_PKG}/app.Version=${APP_VERSION}" -o /app/httpecho main.go

FROM alpine
WORKDIR /app

COPY --from=builder /app/httpecho /app/httpecho

CMD ["/app/httpecho"]
