export PATH=$PATH:/usr/local/go/bin
go get -u github.com/gorilla/mux
go get -u github.com/rs/cors
go build server.go
./server
