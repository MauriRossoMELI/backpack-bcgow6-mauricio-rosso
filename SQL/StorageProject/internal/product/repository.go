package product

import (
	"context"
	"database/sql"
	"errors"
	"storageproject/internal/domain"
)

type Repository interface {
	GetByName(ctx context.Context, name string) (domain.Product, error)
	GetById(ctx context.Context, id int) (domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	Save(ctx context.Context, b domain.Product) (int, error)
	Update(ctx context.Context, b domain.Product, id int) error
	Delete(ctx context.Context, id int) error
	Exists(ctx context.Context, id int) bool
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
	GET_ALL_PRODUCTS    = "SELECT id, name, type, count, price, id_warehouse FROM products"
	GET_PRODUCT_BY_NAME = "SELECT id, name, type, count, price, id_warehouse FROM products WHERE name=?;"
	GET_PRODUCT_BY_ID   = "SELECT id, name, type, count, price, id_warehouse FROM products WHERE id=?;"
	SAVE_PRODUCT        = "INSERT INTO products (name, type, count, price, id_warehouse) VALUES (?, ?, ?, ?, ?);"
	UPDATE_PRODUCT      = "UPDATE products SET name=?, type=?, count=?, price=?, id_warehouse=? WHERE id=?;"
	DELETE_PRODUCT      = "DELETE FROM products WHERE id=?"
	EXIST_PRODUCT       = "SELECT p.id FROM products p WHERE p.id=?"
)

func (r *repository) Exists(ctx context.Context, id int) bool {
	rows := r.db.QueryRow(EXIST_PRODUCT, id)
	err := rows.Scan(&id)
	return err == nil
}

func (r *repository) GetByName(ctx context.Context, name string) (domain.Product, error) {
	row := r.db.QueryRow(GET_PRODUCT_BY_NAME, name)
	var product domain.Product
	if err := row.Scan(&product.Id, &product.Name, &product.Type, &product.Count, &product.Price, &product.IdWarehouse); err != nil {
		return domain.Product{}, err
	}
	return product, nil
}

func (r *repository) GetById(ctx context.Context, id int) (domain.Product, error) {
	row := r.db.QueryRow(GET_PRODUCT_BY_ID, id)
	var product domain.Product
	if err := row.Scan(&product.Id, &product.Name, &product.Type, &product.Count, &product.Price, &product.IdWarehouse); err != nil {
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
		err := rows.Scan(&prod.Id, &prod.Name, &prod.Type, &prod.Count, &prod.Price, &prod.IdWarehouse)
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

func (r *repository) Save(ctx context.Context, p domain.Product) (int, error) {
	stm, err := r.db.Prepare(SAVE_PRODUCT) //preparamos la consulta
	if err != nil {
		return 0, err
	}

	//ejecutamos la consulta con aquellos valores a remplazar en la sentencia
	result, err := stm.Exec(p.Name, p.Type, p.Count, p.Price, p.IdWarehouse)
	if err != nil {
		return 0, err
	}

	//obtenemos el ultimo id
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

func (r *repository) Update(ctx context.Context, p domain.Product, id int) error {
	stm, err := r.db.Prepare(UPDATE_PRODUCT)
	if err != nil {
		return err
	}
	defer stm.Close() //cerramos para no perder memoria

	//ejecutamos la consulta con aquellos valores a remplazar en la sentencia
	result, err := stm.Exec(p.Name, p.Type, p.Count, p.Price, p.IdWarehouse, id)
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
