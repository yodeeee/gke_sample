package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// var (
//     dbuser   string
//     dbpass   string
//     dbhost   string
//     dbport   string
//     dbname   string
//     hostname string
// )

// func rootHandler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "/ is requested")
// }

// func getUserHandler(w http.ResponseWriter, r *http.Request) {
//     id, name := getUser()
//     fmt.Fprintf(w, fmt.Sprintf("%s: name is [%d %s]\n", hostname, id, name))
//     fmt.Printf(fmt.Sprintf("%s: name is [%d %s]\n", hostname, id, name))

// }

// func addUserHandler(w http.ResponseWriter, r *http.Request) {
//     name := addUser()
//     fmt.Fprintf(w, fmt.Sprintf("%s: added user [%s]\n", hostname, name))
//     fmt.Printf(fmt.Sprintf("%s: added user [%s]\n", hostname, name))

// }

func main() {
    fmt.Println("Hello, world.")
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Hello, Docker! <3")
	})

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8080"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
    // getFromEnv()
    // http.HandleFunc("/", rootHandler)
    // http.HandleFunc("/getuser", getUserHandler)
    // http.HandleFunc("/adduser", addUserHandler)
    // http.ListenAndServe(":8080", nil)
}

// func getFromEnv() {
//     dbuser = os.Getenv("DBUSER")
//     dbpass = os.Getenv("DBPASS")
//     dbhost = os.Getenv("DBHOST")
//     dbport = os.Getenv("DBPORT")
//     dbname = os.Getenv("DBNAME")
//     hostname, _ = os.Hostname()
// }

// func getUser() (int, string) {
//     db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpass, dbhost, dbport, dbname))
//     if err != nil {
//         panic(err.Error())
//     }
//     defer db.Close()

//     stmtOut, err := db.Prepare("SELECT id,name FROM users ORDER BY RAND() LIMIT 1;")
//     if err != nil {
//         panic(err.Error())
//     }
//     defer stmtOut.Close()

//     var name string
//     var id int

//     err = stmtOut.QueryRow().Scan(&id, &name)
//     if err != nil {
//         panic(err.Error())
//     }

//     return id, name
// }

// func addUser() string {
//     db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbuser, dbpass, dbhost, dbport, dbname))
//     if err != nil {
//         return err.Error()
//     }
//     defer db.Close()
    

//     stmtIns, err := db.Prepare("INSERT INTO users(name) VALUES(?)!")
//     if err != nil {
//         return err.Error()
//     }
//     defer stmtIns.Close()

//     name := randomdata.SillyName()
//     _, err = stmtIns.Exec(name)
//     if err != nil {
//         return err.Error()
//     }

//     return name
// }