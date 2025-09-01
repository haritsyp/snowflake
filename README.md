# Haritsyp Snowflake for Golang

Snowflake ID generator for Golang, supports int64 numeric IDs, thread-safe, and parsing into components.

## Features

- Generate unique 64-bit integer IDs similar to Kra8 Laravel Snowflake
- Thread-safe (`sync.Mutex`) for concurrent use
- Parse generated IDs to get timestamp, datacenter, node, and sequence
- Configurable datacenter and node IDs
- Ready for integration in any Golang project

## Install

Using Go modules:

```bash
go get github.com/haritsyp/go-snowflake
