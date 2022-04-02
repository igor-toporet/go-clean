package gateway

import (
	"database/sql"

	"toporet/hop/goclean/entity"
)

type IGetAllBooks func() ([]entity.Book, error)

type BookStore struct {
	db *sql.DB
}

func NewBookStore(db *sql.DB) BookStore {
	return BookStore{db}
}

func (s BookStore) RetrieveAll() ([]entity.Book, error) {
	rows, err := s.db.Query("SELECT * FROM books")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var books []entity.Book

	for rows.Next() {
		var bk entity.Book

		err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
		if err != nil {
			return nil, err
		}

		books = append(books, bk)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}
