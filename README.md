# flex_go
A simple application in Golang to find the number of licenses to request according to a user device list csv file.

### Prerequisites
golang version 1.16.4
a config.toml file is required in the project root, specifying relative `Filename` and `ApplicationId` for the application to filter for.
The data directory contains sample data to test with. The git repository also has GitHub Actions setup to build and test the code when a PR is opened.

# Build
```
$ go build
```
This will compile a binary `flex_go` in the project directory.

# Test
```
go test -v ./...
```

# Run
```
$ flex_go
2021/06/09 21:58:05 Parsing csv now...
2021/06/09 21:58:05 Parsing csv completed.
You need to order minimum 227 copies of application with ID 999
```
