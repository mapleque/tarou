.PHONY: clean run build hooks lint test tidy cover

run: simcity
	-mkdir bin
	./bin/simcity

simcity: src/engin/*.go src/client/*.go src/simcity/*.go src/simcity/main/main.go
	go mod tidy
	go build -ldflags="-s" -o ./bin/simcity src/simcity/main/main.go

hooks:
	ln -sf ../../.githooks/prepare-commit-msg .git/hooks/prepare-commit-msg
	ln -sf ../../.githooks/pre-commit .git/hooks/pre-commit

lint: src/engin/*.go src/client/*.go src/simcity/*.go src/simcity/main/main.go
	go mod tidy
	golangci-lint run -v

test: coverage.data

coverage.data: go.mod src/engin/*.go src/client/*.go src/simcity/*.go
	go mod tidy
	go test -v ./src/engin ./src/client ./src/simcity -coverprofile=coverage.data
	go tool cover -func=coverage.data

cover: coverage.data
	go tool cover -html=coverage.data

tidy:
	go mod tidy


clean:
	-rm -f service
	-rm -f coverage.data
