include .env

$(eval export $(shell sed -ne 's/ *#.*$//; /./ s/=.*$$// p' .env))

build:
	gofumports -w -l .
	./scripts/build.sh
run:
	go build -o ./bin