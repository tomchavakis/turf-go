Testing using VS Code

```
go test -v conversions_test.go conversions.go
go test -v measurement_test.go measurement.go
go test -v measurement_test.go measurement.go conversions.go constants.go point.go
```