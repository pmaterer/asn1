test:
	go test -v ./...

lint:
	staticcheck ./...

fmt:
	go fmt ./...