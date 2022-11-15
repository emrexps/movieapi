package service

import (
	"context"
	"fmt"
	"log"
	"movieapi/movie"
	"os"

	"github.com/jackc/pgx/v5"
)

const (
	connectionString = "postgres://postgres:changeme@localhost:5432/postgres"
)

func connectDb() (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), connectionString)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		return nil, err
	}
	return conn, nil
}

func FindAll() []movie.Movie {
	conn, err := connectDb()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	rows, err := conn.Query(context.Background(), "SELECT * FROM movie")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var rowSlice []movie.Movie
	for rows.Next() {
		var r movie.Movie
		err := rows.Scan(&r.Id, &r.MovieName, &r.ReleaseYear, &r.DirectedBy, &r.Genre)
		if err != nil {
			log.Fatal(err)
		}
		rowSlice = append(rowSlice, r)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return rowSlice
}

func FindById(id int) movie.Movie {
	conn, err := connectDb()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	var m movie.Movie
	err = conn.QueryRow(context.Background(), "SELECT ID,MOVIENAME,RELEASEYEAR,DIRECTEDBY,GENRE FROM Movie where id=$1 ", id).Scan(&m.Id, &m.MovieName, &m.ReleaseYear, &m.DirectedBy, &m.Genre)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}
	return m
}

func PostMovie(m movie.Movie) {
	conn, err := connectDb()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	sqlStatement := "INSERT INTO movie ( MOVIENAME,RELEASEYEAR,DIRECTEDBY,GENRE) VALUES ($1, $2, $3, $4)"
	_, err = conn.Exec(context.Background(), sqlStatement, m.MovieName, m.ReleaseYear, m.DirectedBy, m.Genre)
	if err != nil {
		panic(err)
	}
}
