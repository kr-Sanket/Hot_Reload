# Design Decisions

This document explains key design decisions made while building HotReload.

## Use of fsnotify

The `fsnotify` library was used as the file system event source because it provides efficient cross-platform file watching.

## Debounce Mechanism

Text editors often emit multiple file system events when saving files.

A debounce mechanism ensures only one rebuild occurs for rapid event sequences.

## Process Lifecycle Management

When restarting the server, the previous process must be fully terminated.

The process manager ensures that the old process is stopped before launching a new one.

## Log Streaming

Logs are streamed to the dashboard using **Server-Sent Events (SSE)** instead of WebSockets.

SSE was chosen because:

* the communication is one-way (server → browser)
* implementation is simpler
* it avoids additional dependencies.

## Directory Filtering

Common directories such as `.git`, `bin`, and `node_modules` are ignored to prevent unnecessary rebuilds and reduce file system event noise.

## Restart Cooldown

If the server crashes immediately after starting, a cooldown mechanism prevents rapid restart loops.
