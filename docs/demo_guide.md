# Demo Guide

This guide explains how to run and demonstrate the HotReload tool.

## Start the Tool

Run:

```
go run ./cmd/hotreload --root ./testserver --build "go build -o ./bin/server.exe ./testserver" --exec ".\\bin\\server.exe"
```

## Open the Test Server

```
http://localhost:8080
```

## Open Log Dashboard

```
http://localhost:8090
```

## Trigger Hot Reload

1. Edit `testserver/main.go`
2. Save the file

Expected logs:

```
File change detected
Build successful
Server restarted
```

Refresh the browser to see updated server output.

## Dashboard Logs

The dashboard displays logs in real time showing system activity.
