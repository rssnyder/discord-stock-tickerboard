build-linux:
	env GOOS=linux GOARCH=amd64 go build -o ./bin/board

build-osx:
	env GOOS=darwin GOARCH=amd64 go build -o ./bin/board
	
build-openbsd:
	env GOOS=openbsd GOARCH=amd64 go build -o ./bin/board

run:
	./bin/board -logLevel=0

run-dev:
	./bin/board -logLevel=1
