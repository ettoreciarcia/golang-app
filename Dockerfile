FROM golang:1.19.1-alpine as build-step

WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /app/golang-app

FROM alpine:3.17.0
WORKDIR /app
COPY --from=build-step /app/golang-app .

ARG VERSION_NUMBER
ARG PORT_NUMBER

ENV PORT $PORT_NUMBER
ENV VERSION $VERSION_NUMBER
EXPOSE 8080

CMD [ "/app/golang-app" ]