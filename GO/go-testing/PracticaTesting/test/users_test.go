package test

import (
	"bytes"
	"net/http"
	"net/http/httptest"

	"github.com/backpack-bcgow6-mauricio-rosso/go-testing/internal/users"
	"github.com/gin-gonic/gin"
)

func createServer() *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	myStubStore := users.StubStore{}
	repo := users.NewRepository(&myStubStore)
	service := users.NewService(repo)

	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())

	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}
