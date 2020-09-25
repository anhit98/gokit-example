// main.go
package main

import (
	"database/sql"
	"flag"
	"fmt"
	"gokit-example/todo"
	"net/http"

	_ "github.com/lib/pq"
)

const dbsource = "postgresql://postgres:hihi@localhost:5432/postgres?sslmode=disable"

func main() {
	db, err := sql.Open("postgres", dbsource)

	if err != nil {
		panic(err)
	}

	repository := todo.NewRepo(db)

	service := todo.NewService(repository)

	endpoints := todo.MakeEndpoints(service)

	var httpAddr = flag.String("http", ":8080", "http listen address")
	fmt.Println("listening on port", *httpAddr)
	err = http.ListenAndServe(*httpAddr, todo.MakeHTTPHandler(endpoints))

	if err != nil {
		panic(err)
	}
}
