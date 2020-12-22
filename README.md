# Go UI POC

## Goals
This POC has these goals:
----
- Serve a React app in dev mode
- Serve a React app production build
- Bundle the production build files with the binary
- Serve a REST api and UI off different endpoints

## Requirements

To run this project, you will need:
- Go 1.15+
- Node
- NPM

## How to run

The `Makefile` gives a pretty good outline of the process.

If `-devMode` is `true`, then the app does the following:
- Run `npm run start` in separate process for the UI dev server
- Proxy the UI dev server at `/ui`
- Start the API on `/api`

To run in non-dev-mode, you need to build the production UI bundle and embed them using `pkger`:
- `npm run build` in the `/ui` directory
- `pkger`
- `go run ./main.go`

You can also use the `Makefile` to automate this:
- `make run-dev` to run in dev-mode.
- `make build` to build the file to `dist`. 
- `make run` to run the production version of the app in place.