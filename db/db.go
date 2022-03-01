package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	username = "root"
	password = "root"
	hostname = "localhost:33066"
	dbname   = "scrapper"
)

var db *sql.DB

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func execute(ctx context.Context, sql string) {

	res, err := db.ExecContext(ctx, sql)
	if err != nil {
		log.Fatal(err)
	}

	no, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("rows affected %d\n", no)

}

func openConnection() {

	d, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Fatal(err)
	}
	db = d
}

func ConfigureDatabase() {

	openConnection()
	defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	fmt.Printf("Creating database...")
	execute(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	fmt.Printf("Deleting songs...")
	execute(ctx, "delete from songs")
	fmt.Printf("Deleting artists...")
	execute(ctx, "delete from artists")
}

func InsertArtist(name string) (int64, error) {

	openConnection()
	defer db.Close()

	res, err := db.Exec("INSERT INTO artists (name) VALUES (?)",
		name)
	if err != nil {
		log.Fatal(err)
	}
	return res.LastInsertId()
}

func InsertSong(artistDbId int64, title string, lyrick string) (int64, error) {

	openConnection()
	defer db.Close()

	res, err := db.Exec("INSERT INTO songs (artists_id, title, lyrick) VALUES (?, ?, ?)",
		artistDbId, title, lyrick)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(".")
	return res.LastInsertId()
}
