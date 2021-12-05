package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"cloud.google.com/go/spanner"
	"github.com/labstack/echo"
	"google.golang.org/api/iterator"
)

func main() {

    e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})
    e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})
    e.GET("/access_db", access_db)
    e.Logger.Fatal(e.Start(":3000"))
}

func access_db(c echo.Context) error {
	ctx := context.Background()

	// This database must exist.
	databaseName := "projects/test-project/instances/test-instance/databases/test-database"

	client, err := spanner.NewClient(ctx, databaseName)
 
	if err != nil {
		log.Fatalf("Failed to create client %v", err)
	}
	defer client.Close()

	stmt := spanner.Statement{SQL: "SELECT 1"}
	iter := client.Single().Query(ctx, stmt)
	defer iter.Stop()

	for {
		row, err := iter.Next()
		if err == iterator.Done {
			fmt.Println("Done")
            return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
		}
		if err != nil {
			log.Fatalf("Query failed with %v", err)
		}

		var i int64
		if row.Columns(&i) != nil {
			log.Fatalf("Failed to parse row %v", err)
		}
		fmt.Printf("Got value %v\n", i)
	}
}