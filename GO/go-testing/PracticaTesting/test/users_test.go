package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/backpack-bcgow6-mauricio-rosso/go-testing/cmd/server/handler"
	"github.com/backpack-bcgow6-mauricio-rosso/go-testing/internal/users"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer(mockStore mockStorage) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	repo := users.NewRepository(&mockStore)
	service := users.NewService(repo)
	handler := handler.NewUser(service)

	router := gin.Default()

	// // pr := r.Group("/products")
	// // pr.POST("/", p.Store())
	// // pr.GET("/", p.GetAll())

	router.PUT("/users/:id", handler.Update())
	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func TestPutOk(t *testing.T) {
	//arrange
	beforeUpdate := users.User{
		Id:           1,
		Name:         "Mauri",
		Surname:      "Rosso",
		Email:        "mauri@mercadolibre.com",
		Age:          1,
		Height:       1293,
		IsActive:     true,
		CreationDate: "01-01-2012",
	}
	afterUpdate := users.User{
		Id:           1,
		Name:         "Mauri",
		Surname:      "Rosso",
		Email:        "mauriupdated@mercadolibre.com",
		Age:          26,
		Height:       186,
		IsActive:     true,
		CreationDate: "01-01-2012",
	}
	mockStoreTest := mockStorage{
		DataMock: []users.User{
			beforeUpdate,
		},
		Error: "",
	}
	r := createServer()
	req, rr := createRequestTest(http.MethodPut, "/users/1", `{
		"name": "Mauri","surname": "Rosso","email": mauriupdated@mercadolibre.com,"age": 26,"height":186,"isactive":"true","creationdate":"01-01-2012"
		}`)
	//act
	var resp users.User
	r.ServeHTTP(rr, req)
	//assert
	assert.Equal(t, http.StatusOK, rr.Code)
	err := json.Unmarshal(rr.Body.Bytes(), &resp)
	assert.Nil(t, err)
	assert.Equal(t, afterUpdate, resp)
}
