PKGS = $$(go list ./... | grep -v /vendor/)

default:
	go build

test:
	go clean $(PKGS)
	go test $(PKGS) -check.v -coverprofile=coverage.txt -covermode=atomic

race:
	go clean $(PKGS)
	go test -race $(PKGS) -check.v -coverprofile=coverage.txt -covermode=atomic

profile:
	go clean $(PKGS)
	make
	
clean:
	rm -rf *.prof
	go clean $(PKGS)
	rm -rf frontend/node_modules
	rm -rf frontend/build

lint:
	go get -u github.com/golangci/golangci-lint/cmd/golangci-lint
	golangci-lint run

frontend/test:
	cd frontend && \
	npm install && \
	echo npm run test

frontend/build:
	cd frontend && \
	npm install && \
	npm run build

docker:
	docker build -t gfleury/squaas:latest .