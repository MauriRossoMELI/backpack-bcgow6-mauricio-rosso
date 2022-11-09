package tests

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"storageproject/cmd/server/handler"
	"storageproject/internal/product"
	"storageproject/pkg/db"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var s = createServer()

func createServer() *gin.Engine {
	// os.Setenv("DBUSER", "root")
	// os.Setenv("DBPASS", "")
	// os.Setenv("DBNAME", "storage")

	_, db := db.ConnectDatabase()
	repo := product.NewRepository(db)
	serv := product.NewService(repo)

	p := handler.NewProduct(serv)

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	pr := r.Group("/api/v1/products")
	pr.GET("/", p.GetByName())
	// pr.POST("/", p.Store())

	return r
}

func createRequest(method, url, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Set("Content-Type", "application/json")

	return req, httptest.NewRecorder()
}

func TestGetByNameProduct_Ok(t *testing.T) {
	// req, rr := createRequest(http.MethodGet, "/api/v1/products/", `{"name":"ProductExample1"}`)
	req, rr := createRequest(http.MethodGet, "/api/v1/products/ProductExample1", "")

	s.ServeHTTP(rr, req)

	// assert code
	assert.Equal(t, http.StatusOK, rr.Code)
}

// func TestStoreProduct_Ok(t *testing.T) {
// 	new := domain.Product{
// 		Name:  "producto nuevo",
// 		Type:  "producto tipo",
// 		Count: 3,
// 		Price: 84.4,
// 	}

// 	product, err := json.Marshal(new)
// 	assert.Nil(t, err)

// 	req, rr := createRequest(http.MethodPost, "/api/v1/products/", string(product))
// 	s.ServeHTTP(rr, req)

// 	// assert code
// 	assert.Equal(t, http.StatusCreated, rr.Code)

// 	// struct for assertion
// 	p := struct{ Data domain.Product }{}
// 	err = json.Unmarshal(rr.Body.Bytes(), &p)
// 	assert.Nil(t, err)

// 	new.Id = p.Data.Id
// 	assert.Equal(t, new, p.Data)
// }
