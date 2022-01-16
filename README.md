# go-online-store
<!-- To create dependencies file go.mod and go.sum -->
go mod init github.com/go-online-store
<!-- To add a new dependency -->
go get github.com/gorilla/mux@v1.6.2
<!-- To remove unused dependencies -->
go mod tidy
<!-- To build -->
go build .
<!-- To run -->
go run .
<!-- POSTMAN GET http://localhost:8080/api/v1/ -->
<!-- POSTMAN GET http://localhost:8080/api/v1/user/5/comment/3?location=Hogwarts%20Castle -->