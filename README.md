# TChat

## Overview

A terminal-chat developed in Golang.

## Project Structure

```
tchat
├── cmd
│   └── tchat
│       └── main.go
├── go.mod
├── internal
│   ├── client
│   │   └── client.go
│   ├── connection
│   │   └── connection.go
│   ├── consts
│   │   └── consts.go
│   ├── server
│   │   └── server.go
│   └── types
│       └── types.go
└── README.md
```

## Usage Instructions

User 1
```
go run . --listen localhost:3000
```

User 2
```
go run . --connect localhost:3000
```