


OpenAPI specification for Data Connector Agent
https://github.com/hasura/graphql-engine/blob/master/dc-agents/dc-api-types/src/agent.openapi.json

Generate go-server on https://editor.swagger.io/

Run server
---
akashanand@Akashs-MacBook-Pro server % go run main.go
main.go:19:2: "./go" is relative, but relative import paths are not supported in module mode
---

Fix server
---
go mod init example.com

In main.go
replace

sw "./go"

with

sw "example.com/go"

go mod tidy
---

Run server
---
akashanand@Akashs-MacBook-Pro server % go run main.go
# example.com/go
go/model_scalar_value_comparison.go:12:8: undefined: ModelMap
go/model_apply_binary_array_comparison_operator.go:18:11: undefined: ModelMap
go/model_scalar_argument_value.go:12:8: undefined: ModelMap
go/model_capabilities.go:14:12: undefined: Object
go/model_capabilities.go:16:11: undefined: Object
go/model_capabilities.go:18:13: undefined: Object
go/model_capabilities.go:20:11: undefined: Object
go/model_mutation_capabilities.go:12:10: undefined: Object
go/model_mutation_capabilities.go:16:13: undefined: Object
go/model_mutation_capabilities.go:18:10: undefined: Object
go/model_mutation_capabilities.go:18:10: too many errors
---

Fix server
---
In server/go/model_scalar_value_comparison.go
add

type ModelMap interface{}

type Object interface{}

type ScalarType interface{}

type GraphQlType string
---

Build docker image
---
akashanand@Akashs-MacBook-Pro server % docker build -t migrate-data-agent:dev .
[+] Building 3.7s (9/11)                                                                                                                                                                                              
 => [internal] load .dockerignore                                                                                                                                                                                0.0s
 => => transferring context: 2B                                                                                                                                                                                  0.0s
 => [internal] load build definition from Dockerfile                                                                                                                                                             0.0s
 => => transferring dockerfile: 340B                                                                                                                                                                             0.0s
 => [internal] load metadata for docker.io/library/golang:1.10                                                                                                                                                   1.1s
 => [build 1/6] FROM docker.io/library/golang:1.10@sha256:6d5e79878a3e4f1b30b7aa4d24fb6ee6184e905a9b172fc72593935633be4c46                                                                                       0.0s
 => [internal] load build context                                                                                                                                                                                0.0s
 => => transferring context: 11.61kB                                                                                                                                                                             0.0s
 => CACHED [build 2/6] WORKDIR /go/src                                                                                                                                                                           0.0s
 => CACHED [build 3/6] COPY go ./go                                                                                                                                                                              0.0s
 => CACHED [build 4/6] COPY main.go .                                                                                                                                                                            0.0s
 => ERROR [build 5/6] RUN go get -d -v ./...                                                                                                                                                                     2.5s
------                                                                                                                                                                                                                
 > [build 5/6] RUN go get -d -v ./...:                                                                                                                                                                                
#0 0.168 Fetching https://example.com/go?go-get=1                                                                                                                                                                     
#0 1.375 Parsing meta tags from https://example.com/go?go-get=1 (status code 404)                                                                                                                                     
#0 1.375 package example.com/go: unrecognized import path "example.com/go" (parse https://example.com/go?go-get=1: no go-import meta tags ())
#0 1.385 github.com/gorilla/mux (download)
------
Dockerfile:7
--------------------
   5 |     
   6 |     ENV CGO_ENABLED=0
   7 | >>> RUN go get -d -v ./...
   8 |     
   9 |     RUN go build -a -installsuffix cgo -o swagger .
--------------------
ERROR: failed to solve: process "/bin/sh -c go get -d -v ./..." did not complete successfully: exit code: 1
---

Fix docker image
---
Replace
sw "example.com/go"
in main.go
with
sw "./go"
---