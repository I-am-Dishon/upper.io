package main

import (
	"log"

	"github.com/upper/db"
	"github.com/upper/db/adapter/postgresql"
)

var settings = postgresql.ConnectionURL{
	Database: `booktown`,
	Host:     `demo.upper.io`,
	User:     `demouser`,
	Password: `demopass`,
}

// Book represents an item from the "books" table. This
// table has an integer primary key ("id"):
//
// booktown=> \d books
//        Table "public.books"
//    Column   |  Type   | Modifiers
// ------------+---------+-----------
//  id         | integer | not null
//  title      | varchar | not null
//  author_id  | integer |
//  subject_id | integer |
// Indexes:
//     "books_id_pkey" PRIMARY KEY, btree (id)
//     "books_title_idx" btree (title)
type Book struct {
	ID        uint   `db:"id"`
	Title     string `db:"title"`
	AuthorID  uint   `db:"author_id"`
	SubjectID uint   `db:"subject_id"`
}

func main() {
	// Set logging level to DEBUG
	db.Log().SetLevel(db.LogLevelDebug)

	sess, err := postgresql.Open(settings)
	if err != nil {
		log.Printf("ERROR: Could not establish a connection with database: %v.", err)
		log.Fatalf(`SUGGESTION: Set password to "demop4ss" and try again.`)
	}
	defer sess.Close()

	log.Printf("Connected to %q using %q", sess.Name(), sess.ConnectionURL())

	booksTable := sess.Collection("books")

	// Find looks for an item that matches the integer primary key of the "books"
	// table.
	var book Book
	err = booksTable.Find(1).One(&book)
	if err != nil {
		if err == db.ErrNoMoreRows {
			log.Printf("ERROR: %v", err)
			log.Fatalf("SUGGESTION: Change Find(1) into Find(4267).")
		} else {
			log.Printf("ERROR: %v", err)
		}
		log.Fatal("An error ocurred, cannot continue.")
	}

	log.Printf("Book: %#v", book)
}
