build-ui:
	cd ./ui; npm run build

build: build-ui package
	go build -o ./dist

package:
	go install github.com/markbates/pkger/cmd/pkger
	pkger

run: build-ui package
	go run ./main.go

run-dev:
	go run -devMode