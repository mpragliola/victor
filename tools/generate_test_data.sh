#!/bin/bash

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

DEST="$DIR/../data/posts/testdata"

mkdir -p "$DEST"
rm -rf "$DEST/*"

for ((n=0;n<"$1";n++))
do
    FILE=$(cat /dev/urandom | tr -cd "a-z0-9" | fold -w 32 | head -n1)
    cp "$DIR/test.md" "$DEST/$FILE.md"
done
