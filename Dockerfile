FROM golang:1.20 AS build
WORKDIR /go/src
COPY server/go ./go
COPY server/main.go .
COPY server/go.mod .
COPY server/go.sum .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o swagger .

FROM scratch AS runtime
COPY --from=build /go/src/swagger ./
EXPOSE 8080/tcp
ENTRYPOINT ["./swagger"]
