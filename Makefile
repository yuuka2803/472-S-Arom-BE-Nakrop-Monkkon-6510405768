test-log:
@ go test ./domain/usecases/ -coverprofile=coverage.out 
@ go tool cover -html=coverage.out -o coverage.html