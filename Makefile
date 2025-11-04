
mockgen:
	mkdir -p bin
	GOBIN=$(PWD)/bin go install go.uber.org/mock/mockgen@latest

run:
	go mod tidy
	go run cmd/game/main.go
clean:
	rm -rf bin