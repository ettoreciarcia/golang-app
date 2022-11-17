FROM golang:1.19.1-alpine

WORKDIR /app
COPY go.mod ./
RUN go mod download

COPY *.go ./

ARG VERSION_NUMBER
ARG PORT_NUMBER

ENV PORT $PORT_NUMBER
ENV VERSION $VERSION_NUMBER

RUN go build -o /golang-app
EXPOSE 8080

CMD [ "/golang-app" ]