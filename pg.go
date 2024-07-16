package main // TODO: make this its own module

import (
    "database/sql"
    "github.com/skye-lopez/go-query"
    _ "github.com/lib/pq"

    "fmt"
)

func getDB() goquery.GoQuery {
    // TODO: This is just a testing db, eventually will need a real one.
    connStr := "postgres://me:me@localhost/tft_tracker"
    db, connErr := sql.Open("postgres", connStr)
    if connErr != nil {
        fmt.Println("Error opening PG", connErr)
    }

    gq := goquery.NewGoQuery(db)
    gq.AddQueriesToMap("q")
    return gq
}
