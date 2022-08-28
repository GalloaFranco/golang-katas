mkdir coverage
go test ./... -v -cover -coverprofile=coverage/coverage.out
go tool cover -html=coverage/coverage.out -o coverage/coverage.html
open -a "Google Chrome" coverage/coverage.html
sleep 2
rm -r coverage
