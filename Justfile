alias t := test
alias tf := test-force
alias tc := test-cover
alias b := build

test:
  go test -v ./...

test-force:
  go test --count=1 -v ./...

test-cover TYPE:
  go test -coverprofile cover.out ./... && go tool cover -{{TYPE}} cover.out

build NAME:
  go build -o bin/{{NAME}} ./...
