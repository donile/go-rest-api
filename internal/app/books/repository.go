package books

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) Add(book *Book) error {
	insert := "INSERT INTO book (author, title) VALUES ($1, $2)"
	_, err := r.db.Exec(insert, book.Author, book.Title)
	return err
}
