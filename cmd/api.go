package main

import (
	"fmt"
	"github.com/mukeshpilaniya/key-value/db"
	"log"
	"net/http"
	"os"
	"time"
)

type Application struct {
	db          *db.Store
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func (app *Application) Serve() error {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", 8080),
		Handler:      app.routes(),
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 5 * time.Second,
	}
	app.infoLogger.Println("starting backend server at port 8080")
	return srv.ListenAndServe()
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Llongfile)
	var inMemory db.Store
	dbConnection := inMemory.NewStore()
	app := &Application{
		db:          dbConnection,
		errorLogger: errorLog,
		infoLogger:  infoLog,
	}
	err := app.Serve()
	if err != nil {
		panic(err)
	}
}
