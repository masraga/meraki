serve:
	@go run server/server.go

tests:
	@go test usecase/user-auth/register.go usecase/user-auth/register_test.go -v