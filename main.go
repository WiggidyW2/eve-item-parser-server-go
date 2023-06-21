package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/WiggidyW/eve-item-parser-server-go/server"

	pb "github.com/WiggidyW/eve-item-parser-server-go/proto"
)

func main() {
	serve_addr := os.Getenv("SERVE_ADDRESS")
	if serve_addr == "" {
		panic("SERVE_ADDRESS environment variable is unset")
	}
	db_url := os.Getenv("DB_URL")
	if db_url == "" {
		panic("DB_URL environment variable is unset")
	}
	db_max_readers_str := os.Getenv("DB_MAX_READERS")
	if db_max_readers_str == "" {
		panic("DB_MAX_READERS environment variable is unset")
	}
	db_max_readers, err := strconv.Atoi(db_max_readers_str)
	if err != nil {
		panic(fmt.Sprintf(
			"DB_MAX_READERS environment variable is invalid: %e",
			err,
		))
	}
	service, err := server.NewService(db_url, db_max_readers)
	if err != nil {
		panic(fmt.Sprintf("Problem initializing service: %e", err))
	}
	server := pb.NewItemParserServer(service)
	http.ListenAndServe(serve_addr, server)
}
