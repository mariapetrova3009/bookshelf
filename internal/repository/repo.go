package repository

import (
	"bookshelf/internal/models"
	"context"
	"database/sql"
	"errors"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

// CreateBook insert new book in db
func (r *Repo) CreateBook(ctx context.Context, book models.BookRequest) (models.Book, error) {
	q := `INSERT INTO books (title, author, year, isbn, out_of_stock, read, rating)
	      VALUES ($1, $2, $3, $4, $5, $6, $7)
	      RETURNING id, title, author, year, isbn, out_of_stock, read, rating, created_at, updated_at`

	var created models.Book

	err := r.db.QueryRowContext(ctx, q,
		book.Title,
		book.Author,
		book.Year,
		book.ISBN,
		book.OutOfStock,
		book.Read,
		book.Rating,
	).Scan(
		&created.Id,
		&created.Title,
		&created.Author,
		&created.Year,
		&created.ISBN,
		&created.OutOfStock,
		&created.Read,
		&created.Rating,
		&created.CreatedAt,
		&created.UpdatedAt,
	)

	if err != nil {
		return models.Book{}, err
	}

	return created, nil
}

// GetBooks returns books from bd
func (r *Repo) GetBooks(ctx context.Context, limit, offset int) ([]models.Book, error) {
	q := `SELECT * FROM books LIMIT $1 OFFSET $2`
	rows, err := r.db.QueryContext(ctx, q, limit, offset)
	if err != nil {
		return []models.Book{}, err
	}
	books := make([]models.Book, 0, limit)
	defer rows.Close()
	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Author,
			&book.Year,
			&book.ISBN,
			&book.OutOfStock,
			&book.Read,
			&book.Rating,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return books, nil
}

func (r *Repo) UpdateBook(ctx context.Context, id uint64, book models.BookRequest) (models.Book, error) {
	q := `UPDATE books SET
		title = $1,
		author = $2,
		year = $3,
		isbn = $4,
		out_of_stock = $5,
		read = $6,
		rating = $7,
		updated_at = NOW()
	WHERE id = $8
	RETURNING id, title, author, year, isbn, out_of_stock, read, rating, created_at, updated_at`

	var updated models.Book

	err := r.db.QueryRowContext(ctx, q,
		book.Title,
		book.Author,
		book.Year,
		book.ISBN,
		book.OutOfStock,
		book.Read,
		book.Rating,
		id,
	).Scan(
		&updated.Id,
		&updated.Title,
		&updated.Author,
		&updated.Year,
		&updated.ISBN,
		&updated.OutOfStock,
		&updated.Read,
		&updated.Rating,
		&updated.CreatedAt,
		&updated.UpdatedAt,
	)

	if err != nil {
		if err == sql.ErrNoRows {
			return models.Book{}, ErrNotFound
		}
		return models.Book{}, err
	}
	return updated, nil
}

func (r *Repo) GetDetails(ctx context.Context, id uint64) (models.Book, error) {
	q := `SELECT * FROM books WHERE id = $1`
	book := models.Book{}
	err := r.db.QueryRowContext(ctx, q, id).Scan(
		&book.Id,
		&book.Title,
		&book.Author,
		&book.Year,
		&book.ISBN,
		&book.OutOfStock,
		&book.Read,
		&book.Rating,
		&book.CreatedAt,
		&book.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Book{}, ErrNotFound
		}
		return models.Book{}, err
	}
	return book, nil
}

func (r *Repo) DeleteBook(ctx context.Context, id uint64) error {
	q := `DELETE FROM books WHERE id = $1`

	res, err := r.db.ExecContext(ctx, q, id)
	if err != nil {
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return ErrNotFound
	}
	return nil
}

func (r *Repo) RecommendBooks(ctx context.Context, limit int) ([]models.Book, error) {
	q := `SELECT * FROM books 
	WHERE read=TRUE ORDER BY rating DESC, year DESC LIMIT $1`
	rows, err := r.db.QueryContext(ctx, q, limit)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	books := make([]models.Book, 0, limit)
	for rows.Next() {
		var book models.Book
		err := rows.Scan(
			&book.Id,
			&book.Title,
			&book.Author,
			&book.Year,
			&book.ISBN,
			&book.OutOfStock,
			&book.Read,
			&book.Rating,
			&book.CreatedAt,
			&book.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		books = append(books, book)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return books, nil
}

func (r *Repo) MarkOutOfStock(ctx context.Context, id uint64) (models.Book, error) {
	q := `UPDATE books SET out_of_stock=TRUE,
	updated_at = NOW()
	WHERE id=$1 
	RETURNING id, title, author, year, isbn, out_of_stock, read, rating, created_at, updated_at`
	var book models.Book
	err := r.db.QueryRowContext(ctx, q, id).Scan(
		&book.Id,
		&book.Title,
		&book.Author,
		&book.Year,
		&book.ISBN,
		&book.OutOfStock,
		&book.Read,
		&book.Rating,
		&book.CreatedAt,
		&book.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Book{}, ErrNotFound
		}
		return models.Book{}, err
	}
	return book, nil
}
