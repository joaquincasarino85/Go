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

func OpenConnection() {

	d, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Fatal(err)
	}
	db = d
	defer db.Close()
}

func CreateDatabase() {

	d, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Fatal(err)
	}
	db = d

	defer db.Close()

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()

	res, err := db.ExecContext(ctx, "CREATE DATABASE IF NOT EXISTS "+dbname)
	if err != nil {
		log.Fatal(err)
	}

	no, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("rows affected %d\n", no)
}

func InsertArtist(name string) (int64, error) {

	d, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Fatal(err)
	}
	db = d
	defer db.Close()

	res, err := db.Exec("INSERT INTO artists (name) VALUES (?)",
		name)
	if err != nil {
		log.Fatal(err)
	}
	return res.LastInsertId()
}

func InsertSong(artistDbId int64, title string, lyrick string) (int64, error) {

	d, err := sql.Open("mysql", dsn())
	if err != nil {
		log.Fatal(err)
	}
	db = d
	defer db.Close()

	res, err := db.Exec("INSERT INTO songs (artists_id, title, lyrick) VALUES (?, ?, ?)",
		artistDbId, title, lyrick)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(".")
	return res.LastInsertId()
}
