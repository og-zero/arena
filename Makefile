include .env

$(eval export $(shell sed -ne 's/ *#.*$//; /./ s/=.*$$// p' .env))

test:
	go test ./...
build:
	go build -o bin/client cmd/client/main.go
	go build -o bin/server cmd/server/main.go
run:
	go run cmd/server/main.go
git:
	gofumports -w -l .
	git commit -am "WIP"
	git push
	git fetch --all