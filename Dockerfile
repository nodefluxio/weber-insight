FROM node:17-alpine AS build-views
WORKDIR /src/app/views
COPY /views/package.json /views/yarn.lock ./
RUN apk --no-cache --virtual build-dependencies add \
    python3 \
    make \
    g++ \
    git \
    && yarn install --production \
    && apk del build-dependencies

FROM golang:1.17-alpine3.14 AS build-golang
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
COPY --from=build-views /src/app/views/node_modules ./views/node_modules
RUN go install

FROM alpine:3.14 AS runtime
WORKDIR /app
COPY --from=build-golang /go/bin/weber-insight /usr/bin/
COPY --from=build-golang /go/src/app/views /app/views
EXPOSE 8080
CMD ["weber-insight"]
