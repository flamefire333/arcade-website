FROM node:12.7-alpine AS buildng
WORKDIR /usr/src/app
COPY frontend/dist ./dist

FROM golang:alpine AS buildgo
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app
COPY backend/ .
COPY frontend/ .
RUN go get ./...
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/web-app ./*.go

FROM alpine:3.9
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=buildgo /go/src/app/bin /go/bin
COPY --from=buildgo /go/src/app/config config
COPY --from=buildng /usr/src/app/dist/games /frontend
EXPOSE 80
ENTRYPOINT /go/bin/web-app --port 80
