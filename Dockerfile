FROM golang:1.17.9-alpine AS builder
WORKDIR /app

ARG APP_NAME
ENV APP_NAME ${APP_NAME}
ARG APP_VERSION
ENV APP_VERSION ${APP_VERSION}

COPY . .

RUN GO111MODULE=on go mod download
RUN CGO_ENABLED=0 go build -ldflags "-X ${APP_PKG}/app.Name=${APP_NAME} -X ${APP_PKG}/app.Version=${APP_VERSION}" -o /app/httpecho main.go

FROM alpine
WORKDIR /app

ARG APP_NAME
ENV APP_NAME ${APP_NAME}
ARG APP_VERSION
ENV APP_VERSION ${APP_VERSION}

COPY --from=builder /app/httpecho /app/httpecho

CMD ["/app/httpecho"]
