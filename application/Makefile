build:
	GOOS=linux GOARCH=amd64 go build -o ../ansible/roles/app/files/app-linux-amd64 -v main.go
	GOOS=darwin GOARCH=amd64 go build -o app-macos-amd64 -v main.go

test:
	go test -v 

run:
	go run main.go

deps:
	go get github.com/jackc/pgx/v4