include .env

$(eval export $(shell sed -ne 's/ *#.*$//; /./ s/=.*$$// p' .env))

all:
	gofumports -w -l .
build-client:
	go build -o bin/client cmd/client/main.go
build-server:
	go build -o bin/server cmd/server/main.go
run-client:
	go run cmd/client/main.go
run-server:
	go run cmd/server/main.go
git:
	git add --all
	git commit -m "WIP"
	git push