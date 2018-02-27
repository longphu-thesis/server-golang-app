#!/bin/bash

BASENAME=$(basename $PWD)
WORK_DIR="$HOME/go19_path/src/"

export GOROOT="$HOME/go19"
export GOPATH="$HOME/go19_path"

if [ ! -d "$WORK_DIR/$BASENAME" ]; then
    ln -s $PWD "$WORK_DIR"
fi

cd "$WORK_DIR/$BASENAME"