package product

import (
	"context"
	"errors"
	"regexp"
	"storageproject/internal/domain"
	"storageproject/tests"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var errorForzado = errors.New("error forzado")

var product_test = domain.Product{
	Id:          1,
	Name:        "TestName 1",
	Type:        "TestType 1",
	Count:       1,
	Price:       1,
	IdWarehouse: 1,
}

/////////// ----> TESTS OK & FAIL <---- ///////////

func TestExists_OK(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"id"}
	rows := sqlmock.NewRows(columns)

	rows.AddRow(product_test.Id)
	mock.ExpectQuery(regexp.QuoteMeta(EXIST_PRODUCT)).WithArgs(1).WillReturnRows(rows)

	repo := NewRepository(db)

	resp := repo.Exists(context.TODO(), 1)

	assert.True(t, resp)
}

func TestExists_Fail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"id"}
	rows := sqlmock.NewRows(columns)

	mock.ExpectQuery(regexp.QuoteMeta(EXIST_PRODUCT)).WithArgs(999).WillReturnRows(rows)

	repo := NewRepository(db)

	resp := repo.Exists(context.TODO(), 999)

	assert.False(t, resp)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAll_Ok(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"id", "name", "type", "count", "price", "id_warehouse"}
	rows := sqlmock.NewRows(columns)
	products := []domain.Product{{Id: 1, Name: "nameTest1", Type: "typeTest1", Count: 1, Price: 1, IdWarehouse: 1}, {Id: 2, Name: "nameTest2", Type: "typeTest2", Count: 2, Price: 2, IdWarehouse: 2}}

	for _, product := range products {
		rows.AddRow(product.Id, product.Name, product.Type, product.Count, product.Price, product.IdWarehouse)
	}

	mock.ExpectQuery(regexp.QuoteMeta(GET_ALL_PRODUCTS)).WillReturnRows(rows)

	repo := NewRepository(db)
	resultProducts, err := repo.GetAll(context.TODO())

	// recordar tener el producto con id 1 y tener un warehouse asociado a este producto
	assert.NoError(t, err)
	assert.Equal(t, products, resultProducts)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAll_Fail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"id", "name", "type", "count", "price", "id_warehouse"}
	rows := sqlmock.NewRows(columns)
	products := []domain.Product{{Id: 1, Name: "nameTest1", Type: "typeTest1", Count: 1, Price: 1, IdWarehouse: 1}, {Id: 2, Name: "nameTest2", Type: "typeTest2", Count: 2, Price: 2, IdWarehouse: 2}}

	for _, product := range products {
		rows.AddRow(product.Id, product.Name, product.Type, product.Count, product.Price, product.IdWarehouse)
	}

	mock.ExpectQuery(regexp.QuoteMeta(GET_ALL_PRODUCTS)).WillReturnError(errorForzado)

	repo := NewRepository(db)
	resultProducts, err := repo.GetAll(context.TODO())

	// recordar tener el producto con id 1 y tener un warehouse asociado a este producto
	assert.EqualError(t, err, errorForzado.Error())
	assert.Empty(t, resultProducts)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetById_Ok(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"id", "name", "type", "count", "price", "id_warehouse"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(product_test.Id, product_test.Name, product_test.Type, product_test.Count, product_test.Price, product_test.IdWarehouse)

	// acá como en el metodo se utiliza QueryRow seria ExpectQuery, en el caso de que en el metodo se utilize Exec ahí si utilizariamos ExpectExec
	mock.ExpectPrepare(regexp.QuoteMeta(GET_PRODUCT_BY_ID)).ExpectQuery().WithArgs(product_test.Id).WillReturnRows(rows)
	//USE THIS IF WE DONT EXPECT A QUERY PREPARED
	// mock.ExpectQuery(regexp.QuoteMeta(GET_PRODUCT_BY_ID)).WithArgs(product_test.Id).WillReturnRows(rows)

	repo := NewRepository(db)

	//SI NECESITAMOS UN TIMEOUT
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	ctx := context.Background()

	productResult, err := repo.GetById(ctx, product_test.Id)
	assert.NoError(t, err)
	assert.Equal(t, product_test.Name, productResult.Name)
	assert.Equal(t, product_test.Id, productResult.Id)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestGetById_Fail(t *testing.T) {
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	columns := []string{"id", "name", "type", "count", "price", "id_warehouse"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(product_test.Id, product_test.Name, product_test.Type, product_test.Count, product_test.Price, product_test.IdWarehouse)

	mock.ExpectPrepare(regexp.QuoteMeta(GET_PRODUCT_BY_ID)).ExpectQuery().WithArgs(product_test.Id).WillReturnError(errorForzado)

	repo := NewRepository(db)
	resultProducts, err := repo.GetById(context.TODO(), product_test.Id)

	assert.EqualError(t, err, errorForzado.Error())
	assert.Empty(t, resultProducts)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestSave_Ok(t *testing.T) {
	db, mock, err := tests.NewDBMock(t)
	assert.NoError(t, err)
	defer db.Close()

	ctx := context.TODO()
	repo := NewRepository(db)

	mock.ExpectPrepare(regexp.QuoteMeta(SAVE_PRODUCT)).ExpectExec().WithArgs(product_test.Name, product_test.Type, product_test.Count, product_test.Price, product_test.IdWarehouse).WillReturnResult(sqlmock.NewResult(1, 1))

	// Save
	id, err := repo.Save(ctx, domain.Product{Id: product_test.Id, Name: product_test.Name, Type: product_test.Type, Count: product_test.Count, Price: product_test.Price, IdWarehouse: product_test.IdWarehouse})
	assert.NoError(t, err)
	assert.Equal(t, product_test.Id, id)

	columns := []string{"id", "name", "type", "count", "price", "id_warehouse"}
	rows := sqlmock.NewRows(columns)
	rows.AddRow(product_test.Id, product_test.Name, product_test.Type, product_test.Count, product_test.Price, product_test.IdWarehouse)
	mock.ExpectPrepare(regexp.QuoteMeta(GET_PRODUCT_BY_ID)).ExpectQuery().WithArgs(product_test.Id).WillReturnRows(rows)

	// GetById
	product, err := repo.GetById(ctx, product_test.Id)
	assert.NoError(t, err)
	assert.Equal(t, product_test, product)
}

func TestSave_Fail(t *testing.T) {
	db, mock, err := tests.NewDBMock(t)
	assert.NoError(t, err)
	defer db.Close()

	p := &domain.Product{}

	ctx := context.TODO()
	repo := NewRepository(db)

	mock.ExpectPrepare(regexp.QuoteMeta(SAVE_PRODUCT)).
		ExpectExec().
		WithArgs(p.Name, p.Type, p.Count, p.Price, p.IdWarehouse).
		WillReturnError(errors.New("you have not provided the necessary fields to insert"))

	// Store failed
	id, err := repo.Save(ctx, *p)
	assert.Equal(t, 0, id)
	assert.NotNil(t, err)
	assert.Error(t, err)
}
func TestUpdate_Ok(t *testing.T) { //PROBAR ROWS AFFECTED PARA AUMENTAR EL COVERAGE ;)
	db, mock, err := tests.NewDBMock(t)
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta(UPDATE_PRODUCT)).ExpectExec().WithArgs(product_test.Name, product_test.Type, product_test.Count, product_test.Price, product_test.IdWarehouse, product_test.Id).WillReturnResult(sqlmock.NewResult(0, 1)) //Es 0 porque NO te retorna un id la func originial.

	// Update
	ctx := context.TODO()
	repo := NewRepository(db)
	errUpd := repo.Update(ctx, domain.Product{Name: product_test.Name, Type: product_test.Type, Count: product_test.Count, Price: product_test.Price, IdWarehouse: product_test.IdWarehouse}, product_test.Id)

	assert.NoError(t, errUpd)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestUpdate_Fail(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta(UPDATE_PRODUCT)).ExpectExec().WithArgs(product_test.Name, product_test.Type, product_test.Count, product_test.Price, product_test.IdWarehouse, product_test.Id).WillReturnError(errorForzado)

	ctx := context.TODO()
	repo := NewRepository(db)

	// act
	errUpd := repo.Update(ctx, domain.Product{Name: product_test.Name, Type: product_test.Type, Count: product_test.Count, Price: product_test.Price, IdWarehouse: product_test.IdWarehouse}, product_test.Id)

	// assert
	assert.EqualError(t, errUpd, errorForzado.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete_Ok(t *testing.T) { //PROBAR ROWS AFFECTED PARA AUMENTAR EL COVERAGE ;)
	db, mock, err := tests.NewDBMock(t)
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta(DELETE_PRODUCT)).ExpectExec().WithArgs(product_test.Id).WillReturnResult(sqlmock.NewResult(0, 1)) //Es 0 porque NO te retorna un id la func original.

	// Update
	ctx := context.TODO()
	repo := NewRepository(db)
	errDelete := repo.Delete(ctx, product_test.Id)

	assert.NoError(t, errDelete)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestDelete_Fail(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta(DELETE_PRODUCT)).ExpectExec().WithArgs(product_test.Id).WillReturnError(errorForzado)

	repo := NewRepository(db)

	// act
	err = repo.Delete(context.TODO(), product_test.Id)

	// assert
	assert.EqualError(t, err, errorForzado.Error())
	assert.NoError(t, mock.ExpectationsWereMet())
}

/////////// ----> TESTS EXTRAS <---- ///////////

func TestDelete_FailRowsAffected(t *testing.T) {
	// arrange
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	mock.ExpectPrepare(regexp.QuoteMeta(DELETE_PRODUCT)).ExpectExec().WithArgs(product_test.Id).WillReturnResult(sqlmock.NewResult(1, 2))

	repo := NewRepository(db)

	// act
	errDel := repo.Delete(context.TODO(), product_test.Id)

	// assert
	assert.EqualError(t, errDel, "error: no affected rows")
}
