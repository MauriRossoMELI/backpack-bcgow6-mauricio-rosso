package products

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	repo := NewRepository()
	service := NewService(repo)
	handler := NewHandler(service)
	router := gin.Default()

	router.GET("/api/v1/products", handler.GetProducts)
	return router
}

func createFailServer() *gin.Engine {
	repo := NewMockedRepository() // this will return err
	svc := NewService(repo)
	handler := NewHandler(svc)
	router := gin.Default()

	router.GET("/api/v1/products", handler.GetProducts)
	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	//req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestGetAllStatusOk(t *testing.T) {
	server := createServer()
	req, response := createRequestTest(http.MethodGet, "/api/v1/products?seller_id=asd", "")
	// act
	server.ServeHTTP(response, req)
	// assert
	assert.Equal(t, 200, response.Code)
}

func TestGetAllStatusBadRequest(t *testing.T) {
	server := createServer()
	req, response := createRequestTest(http.MethodGet, "/api/v1/products?", "")
	// act
	server.ServeHTTP(response, req)
	// assert
	assert.Equal(t, 400, response.Code)
}

func TestStatusInternalServerError(t *testing.T) {
	server := createFailServer()
	req, response := createRequestTest(http.MethodGet, "/api/v1/products?seller_id=asd", "")
	// act
	server.ServeHTTP(response, req)
	// assert
	assert.Equal(t, response.Code, 500)
}
