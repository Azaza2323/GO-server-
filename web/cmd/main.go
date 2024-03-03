package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"server/mysql"
)

type application struct {
	debug    bool
	errorLog *log.Logger
	infoLog  *log.Logger
	users    *mysql.UserModel
	books    *mysql.BookModel
}

func main() {
	addr := flag.String("addr", ":1111", "HTTP network address")
	debug := flag.Bool("debug", false, "Enable debug mode")
	dsn := flag.String("dsn", "root:123123@/bookstore?parseTime=true", "Main Data")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()
	app := &application{
		debug:    *debug,
		errorLog: errorLog,
		infoLog:  infoLog,
		users:    &mysql.UserModel{DB: db},
		books:    &mysql.BookModel{DB: db},
	}

	handler := app.routes()

	handler = corsHandler(handler)

	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  handler,
	}
	infoLog.Printf("Starting server on %s", *addr)
	err = srv.ListenAndServe()
	if err != nil {
		errorLog.Fatal(err)
	}
}
func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}
func corsHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow requests from any origin
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}
