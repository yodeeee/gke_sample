package main

import (
	"context"
	"fmt"
	"io"
	"log"

	"cloud.google.com/go/spanner"
	database "cloud.google.com/go/spanner/admin/database/apiv1"
	"google.golang.org/api/iterator"
)

func main() {

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
			return
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

	// e := echo.New()

	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())


	// e.GET("/", func(c echo.Context) error {
	// 	return c.HTML(http.StatusOK, "Hello, Docker! <3")
	// })

	// e.GET("/ping", func(c echo.Context) error {
	// 	return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	// })

	// httpPort := os.Getenv("HTTP_PORT")
	// if httpPort == "" {
	// 	httpPort = "8080"
	// }

	// e.Logger.Fatal(e.Start(":" + httpPort))
    // getFromEnv()
    // http.HandleFunc("/", rootHandler)
    // http.HandleFunc("/getuser", getUserHandler)
    // http.HandleFunc("/adduser", addUserHandler)
    // http.ListenAndServe(":8080", nil)


    // defer adminClient.Close()
	// defer dataClient.Close()
	// if err := run(ctx, adminClient, dataClient, os.Stdout, cmd, db); err != nil {
	// 	os.Exit(1)
    // }
}

func read(w io.Writer, db string) error {
    ctx := context.Background()
    client, err := spanner.NewClient(ctx, db)
    if err != nil {
            return err
    }
    defer client.Close()

    iter := client.Single().Read(ctx, "Albums", spanner.AllKeys(),
            []string{"SingerId", "AlbumId", "AlbumTitle"})
    defer iter.Stop()
    for {
            row, err := iter.Next()
            if err == iterator.Done {
                    return nil
            }
            if err != nil {
                    return err
            }
            var singerID, albumID int64
            var albumTitle string
            if err := row.Columns(&singerID, &albumID, &albumTitle); err != nil {
                    return err
            }
            fmt.Fprintf(w, "%d %d %s\n", singerID, albumID, albumTitle)
    }
}



func createClients(ctx context.Context, db string) (*database.DatabaseAdminClient, *spanner.Client) {
	adminClient, err := database.NewDatabaseAdminClient(ctx)
	if err != nil {
		log.Fatal(err)
	}

	dataClient, err := spanner.NewClient(ctx, db)
	if err != nil {
		log.Fatal(err)
	}

	return adminClient, dataClient
}