build:
	dep ensure
	env GOOS=linux go build -ldflags="-s -w" -o bin/read read/main.go
	env GOOS=linux go build -ldflags="-s -w" -o bin/create create/main.go