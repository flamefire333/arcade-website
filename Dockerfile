FROM node:12.7-alpine AS build
WORKDIR /usr/src/app
COPY frontend-source/ .
RUN npm install
RUN npm run build

FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make git
WORKDIR /go/src/app
COPY backend/ .
COPY frontend/ .
RUN go get ./...
RUN GOOS=linux go build -ldflags="-s -w" -o ./bin/web-app ./*.go

FROM alpine:3.9
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/bin /go/bin
COPY --from=build /go/src/app /frontend
EXPOSE 80
ENTRYPOINT /go/bin/web-app --port 80
