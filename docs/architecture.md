# HotReload Architecture

## Overview

HotReload is a CLI tool that automatically rebuilds and restarts a Go server whenever source code changes.

The system is designed with a modular architecture where each component has a single responsibility.

## High Level Architecture

```
Developer edits code
        │
        ▼
 File Watcher (fsnotify)
        │
        ▼
 Debounce System
        │
        ▼
 Build Manager
        │
        ▼
 Process Manager
        │
        ▼
 Running Server
        │
        ▼
 LogHub → Dashboard
```

## Components

### CLI Interface

Responsible for parsing command line arguments:

```
hotreload --root <project-folder> --build "<build-command>" --exec "<run-command>"
```

Arguments:

| Flag      | Purpose                              |
| --------- | ------------------------------------ |
| `--root`  | Directory to watch                   |
| `--build` | Command used to build the project    |
| `--exec`  | Command used to run the built server |

---

### File Watcher

The watcher monitors the project directory for file changes using **fsnotify**.

Responsibilities:

* recursive directory watching
* ignore unnecessary folders (`.git`, `bin`, `node_modules`)
* detect newly created folders dynamically

---

### Debounce System

Many editors trigger multiple file system events during a single save.

Example:

```
WRITE main.go
WRITE main.go
WRITE main.go
```

Without debounce this would trigger multiple builds.

The debounce mechanism ensures only **one rebuild is triggered**.

---

### Build Manager

Responsible for executing the configured build command.

Example:

```
go build -o ./bin/server.exe ./testserver
```

If the build fails:

* error logs are shown
* server restart is skipped

---

### Process Manager

Handles server lifecycle management:

* starting the server
* stopping the existing server
* restarting after successful build

It also ensures the previous server process is **fully terminated before restart**.

---

### LogHub

LogHub is responsible for collecting and broadcasting logs.

Logs are sent to:

* terminal output
* browser dashboard

---

### Dashboard

The dashboard provides a web interface for monitoring logs.

```
http://localhost:8090
```

Logs are streamed using **Server-Sent Events (SSE)**.

---

## Design Principles

The system was designed using the following principles:

* modular architecture
* separation of concerns
* minimal dependencies
* real-time feedback for developers
