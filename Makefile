build:
	go build -o bin/app cmd/main.go 
run:
	$(MAKE) build
	bin/app 
deps:
	go mod download