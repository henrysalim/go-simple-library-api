package repository

import (
	"context"
	"database/sql"
	"errors"
	"simple-library-api/internal/model"
)

type BookRepository struct {
	DB *sql.DB
}

func NewBookRepository(db *sql.DB) *BookRepository {
	return &BookRepository{DB: db}
}

var ErrBookNotFound = errors.New("book not found")

func (r *BookRepository) GetBooks(ctx context.Context) ([]model.Book, error) {
	rows, err := r.DB.QueryContext(ctx, "SELECT id, title, author, year FROM books")
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	var books []model.Book
	for rows.Next() {
		var b model.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.Author, &b.Year); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}

func (r *BookRepository) CreateBook(ctx context.Context, b *model.Book) error {
	res, err := r.DB.ExecContext(ctx, "INSERT INTO books (title, author, year) VALUES (?, ?, ?)", b.Title, b.Author, b.Year)
	if err != nil {
		return err
	}
	id, _ := res.LastInsertId()
	b.ID = int(id)
	return nil
}

func (r *BookRepository) UpdateBook(ctx context.Context, b *model.Book) error {
	query := "UPDATE books SET title = ?, author = ?, year = ? WHERE id = ?"

	//	ExecContext returns a result and an error
	res, err := r.DB.ExecContext(ctx, query, b.Title, b.Author, b.Year, b.ID)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrBookNotFound
	}

	return nil
}

func (r *BookRepository) DeleteBook(ctx context.Context, id int) error {
	query := "DELETE FROM books WHERE id = ?"

	res, err := r.DB.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return ErrBookNotFound
	}

	return nil
}
