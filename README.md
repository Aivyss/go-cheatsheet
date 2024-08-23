# coverage test
```shell
go test -cover ./...
go test -cover ./... -coverprofile=cover.out
go tool cover -html=cover.out -o cover.html
open cover.html
```