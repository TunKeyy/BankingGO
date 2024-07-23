# BankingGO

## Get library
```
go get github.com/<libname>
```
Next
```
go mod tidy
```

## Create Migration
```
migrate create -ext sql -dir db/migration -seq <migration_name>
```