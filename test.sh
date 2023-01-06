#!/bin/bash

find ./pkg ./service -name go.mod -execdir go test ./... \;