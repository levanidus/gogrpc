package repository

import (
	"github.com/jmoiron/sqlx"
)

type Book struct {
	name string `db:"name,prefix=books."`
}

type Author struct {
	name string `db:"name"`
}

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) FindBooksByAuthor(req string) ([]string, error) {
	books := []string{}

	rows, err := r.db.Query(`select books.name from books join authors_books  
									on books.id = authors_books.book_id
									join authors
									on authors_books.author_id = authors.id
									where authors.name=?`, req)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book Book
		if err := rows.Scan(&book.name); err != nil {
			return nil, err
		}
		books = append(books, book.name)
	}

	return books, nil
}
