package product

import (
	"context"
	"database/sql"
	"errors"
	"storageproject/internal/domain"
)

type Repository interface {
	GetByName(ctx context.Context, name string) (domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	Delete(ctx context.Context, id int) error
	// Save(ctx context.Context, b domain.Movie) (int64, error)
	// Exists(ctx context.Context, id int) bool
	// Update(ctx context.Context, b domain.Movie, id int) error
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{
		db: db,
	}
}

const (
	// SAVE_MOVIE     = "INSERT INTO movies (title, rating, awards, length, genre_id) VALUES (?, ?, ?, ?, ?);"
	// UPDATE_MOVIE   = "UPDATE movies SET title=?, rating=?, awards=?, length=?, genre_id=? WHERE id=?;"

	GET_ALL_PRODUCTS = "SELECT id, name, type, count, price FROM products"
	GET_PRODUCT      = "SELECT id, name, type, count, price FROM products WHERE name=?;"
	EXIST_PRODUCT    = "SELECT p.id FROM products p WHERE p.id=?"
	DELETE_PRODUCT   = "DELETE FROM products WHERE id=?"
)

func (r *repository) Exists(ctx context.Context, id int) bool {
	rows := r.db.QueryRow(EXIST_PRODUCT, id)
	err := rows.Scan(&id)
	return err == nil
}

func (r *repository) GetByName(ctx context.Context, name string) (domain.Product, error) {
	row := r.db.QueryRow(GET_PRODUCT, name)
	var product domain.Product
	if err := row.Scan(&product.Id, &product.Name, &product.Type, &product.Count, &product.Price); err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error) {
	var products []domain.Product
	rows, err := r.db.Query(GET_ALL_PRODUCTS)
	if err != nil {
		return []domain.Product{}, err
	}

	for rows.Next() {
		var prod domain.Product
		err := rows.Scan(&prod.Id, &prod.Name, &prod.Type, &prod.Count, &prod.Price)
		if err != nil {
			return []domain.Product{}, err
		}
		products = append(products, prod)
	}

	return products, nil
}

func (r *repository) Delete(ctx context.Context, id int) error {
	stm, err := r.db.Prepare(DELETE_PRODUCT)
	if err != nil {
		return err
	}
	defer stm.Close() //cerramos para no perder memoria

	result, err := stm.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("error: no affected rows")
	}

	return nil
}

/*
func (r *repository) Save(ctx context.Context, m domain.Movie) (int64, error) {
	stm, err := r.db.Prepare(SAVE_MOVIE) //preparamos la consulta
	if err != nil {
		return 0, err
	}

	//ejecutamos la consulta con aquellos valores a remplazar en la sentencia
	result, err := stm.Exec(m.Title, m.Rating, m.Awards, m.Length, m.Genre_id)
	if err != nil {
		return 0, err
	}

	//obtenemos el ultimo id
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (r *repository) Update(ctx context.Context, m domain.Movie, id int) error {
	stm, err := r.db.Prepare(UPDATE_MOVIE)
	if err != nil {
		return err
	}
	defer stm.Close() //cerramos para no perder memoria

	//ejecutamos la consulta con aquellos valores a remplazar en la sentencia
	result, err := stm.Exec(m.Title, m.Rating, m.Awards, m.Length, m.Genre_id, id)
	if err != nil {
		return err
	}

	affected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if affected < 1 {
		return errors.New("error: no affected rows")
	}
	return nil
}
*/
