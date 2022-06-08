#!/bin/bash

set -e
set -o pipefail
set -a
set -x

$GO_BIN list -f '{{$dir := .Dir}}{{range $index, $fileName := .GoFiles}}{{$dir}}/{{$fileName}}
{{end}}' $PKG_PATTERN
