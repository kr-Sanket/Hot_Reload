# HotReload

HotReload is a CLI tool that automatically rebuilds and restarts a Go server whenever source files change.

It improves developer productivity by eliminating the need to manually stop, rebuild, and restart the server after every code change.

---

## Problem

When developing backend services in Go, even a small change requires:

1. Stopping the server
2. Rebuilding the project
3. Starting the server again

This process is repetitive and slows down development.

HotReload automates this workflow.

---

## Features

* Watches project folders recursively for file changes
* Automatically rebuilds the project when source files change
* Automatically restarts the running server
* Streams server logs in real time
* Debounces rapid file events to avoid unnecessary rebuilds
* Automatically watches newly created directories
* Ignores unnecessary directories such as `.git`, `bin`, and `node_modules`
* Prevents rapid restart loops if the server crashes

---

## Architecture Overview

The tool is built as a modular system with the following components:

### CLI Interface

Parses command line arguments:

```
hotreload --root <project-folder> --build "<build-command>" --exec "<run-command>"
```

### File Watcher

Uses `fsnotify` to monitor file system events such as:

* file modifications
* file creations
* file deletions

The watcher recursively monitors directories and dynamically adds new folders created during runtime.

### Debounce System

Editors often trigger multiple file events for a single save operation.
The debounce mechanism waits briefly before triggering a rebuild to avoid redundant builds.

### Build Manager

Executes the provided build command using the Go `exec` package.
If the build fails, the running server continues operating.

### Process Manager

Manages the server lifecycle:

* stops the currently running server
* starts the newly built server
* streams logs to the terminal

A cooldown mechanism prevents rapid restart loops if the server crashes immediately after starting.

---

## Project Structure

```
hotreload
│
├── cmd
│   └── hotreload
│       └── main.go
│
├── internal
│   ├── builder
│   ├── config
│   ├── debounce
│   ├── process
│   └── watcher
│
├── testserver
│   └── main.go
│
├── go.mod
└── README.md
```

---

## Example Usage

Run HotReload with a Go project:

```
go run ./cmd/hotreload \
  --root ./testserver \
  --build "go build -o ./bin/server.exe ./testserver" \
  --exec ".\\bin\\server.exe"
```

Now whenever a file in `testserver` changes:

```
File change detected
Starting build...
Build successful
Stopping server...
Starting server...
```

The server restarts automatically.

---

## Demo Server

A simple HTTP server is included in the `testserver` directory to demonstrate hot reload functionality.

```
http://localhost:8080
```

Edit the server code and save the file to see the automatic rebuild and restart.

---

## Design Goals

The tool was designed with the following goals:

* fast rebuild cycle
* minimal dependencies
* simple architecture
* reliability for long-running development sessions

---

## Future Improvements

Possible enhancements include:

* file filtering based on extensions
* cross-platform process management improvements
* configurable debounce delay
* better logging options

---

## Author

Sanket Kumar
