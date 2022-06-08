package main

import (
	"gogrpc/api"
	"gogrpc/server"
	"log"
	"net"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

func main() {

	db, err := sqlx.Open("mysql", "root:secret@tcp(localhost:3306)/gogrpc")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	srv := server.NewGRPCServer(db)

	api.RegisterFindBooksAuthorsServer(s, srv)

	listener, err := net.Listen("tcp", ":8080")

	if err != nil {
		log.Fatal(err)
	}

	if err := s.Serve(listener); err != nil {
		log.Fatal(err)
	}

}
