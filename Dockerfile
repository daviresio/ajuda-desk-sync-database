FROM golang:1.13.7 as build-env

RUN mkdir /app
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/sync-database

FROM golang:1.13.7-stretch
COPY --from=build-env /go/bin/sync-database /go/bin/sync-database
EXPOSE 9092
EXPOSE 9200
ENTRYPOINT ["/go/bin/sync-database"]