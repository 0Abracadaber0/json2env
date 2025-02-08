#!/bin/bash

DIRECTORY="$HOME/bin"

if [ ! -d "$DIRECTORY" ]; then
  mkdir "$DIRECTORY"
fi

go build -o "$DIRECTORY"/json2env cmd/json2env/main.go

if [[ ":$PATH:" != *":$DIRECTORY:"* ]]; then
  export PATH="$DIRECTORY:$PATH"
fi