export PATH=$PATH:/usr/local/go/bin
go install github.com/gorilla/mux
go install github.com/rs/cors
go build server.go
