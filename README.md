You can run `./test.sh` or `make test`

The command in place is:

```shell
find ./pkg ./service -name go.mod -execdir go test ./... \;
```

which can be:
```shell
find . -name go.mod -execdir go test ./... \;
```

but this one will try to look for tests at root level and it'll throw the follow error for the empty go.mod (but it works ok):
```
go: warning: "./..." matched no packages
no packages to test
```
