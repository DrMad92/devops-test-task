package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Entry struct {
	ID        int
	Name      string
	CreatedAt string
}

type Store interface {
	ListEntries() ([]Entry, error)
	AddEntry(name string) error
	DeleteEntry(id int) error
}

type dbStore struct {
	db *sql.DB
}

// OpenDB creates connection to db
func OpenDB(dbServer string) (*sql.DB, error) {
	checkDatabase(*dbName, dbServer)
	checkTable(dbServer)

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbServer, *port, *user, *password, *dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, err

}

//ListEntries lists rows from table entries
func (store *dbStore) ListEntries() ([]Entry, error) {
	var list []Entry
	sqlStatement := `SELECT * FROM entries;`
	rows, err := store.db.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		entry := Entry{}
		err = rows.Scan(&entry.ID, &entry.Name, &entry.CreatedAt)
		if err != nil {
			return nil, err
		}
		list = append(list, entry)
	}
	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return list, nil
}

// AddEntry adds new entry
func (store *dbStore) AddEntry(name string) error {
	sqlStatement := `
	INSERT INTO entries (name)
	VALUES ($1)`

	_, err := store.db.Exec(sqlStatement, name)

	return err
}

// DeleteEntry deletes entry
func (store *dbStore) DeleteEntry(id int) error {
	sqlStatement := `
	DELETE FROM entries
	WHERE id = $1;`

	res, err := store.db.Exec(sqlStatement, id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		fmt.Println("Nothing is deleted")
	} else {
		fmt.Println(id, "deleted")
	}
	return nil

}

// checkDatabase checks if database exists otherwise creates
func checkDatabase(dbname string, dbServer string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=postgres sslmode=disable",
		dbServer, *port, *user, *password)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	sqlStatement := `SELECT EXISTS(SELECT datname FROM pg_catalog.pg_database 
		WHERE datname=$1);`

	var exists bool
	err = db.QueryRow(sqlStatement, dbname).Scan(&exists)
	if err != nil {
		panic(err)
	}
	if !exists {
		fmt.Println("Creating database", dbname)
		sqlStatement = `CREATE DATABASE ` + dbname + ";"

		_, err = db.Exec(sqlStatement)
		if err != nil {
			panic(err)
		}
	}
}

//checkTable check if table exists otherwise creates
func checkTable(dbServer string) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbServer, *port, *user, *password, *dbName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `SELECT EXISTS(SELECT 1 FROM pg_tables
		WHERE schemaname = 'public'
		AND tablename = 'entries'
		);`

	var exists bool
	err = db.QueryRow(sqlStatement).Scan(&exists)
	if err != nil {
		panic(err)
	}

	if !exists {
		fmt.Println("Creating table entries")
		sqlStatement = `
		CREATE TABLE entries (
		id SERIAL PRIMARY KEY,
		name TEXT,
		created_at TIMESTAMP DEFAULT now()
		);`

		_, err = db.Exec(sqlStatement)
		if err != nil {
			panic(err)
		}
	}

}

var store Store

//InitStore initializes store
func InitStore(s Store) {
	store = s
}
