#!/bin/zsh
echo "\n👋 Basically me run server with postgres on Mac"
DATABASE_URL="postgres://localhost:5432/agrobank2" go run server.go
