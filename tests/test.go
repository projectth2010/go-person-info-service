package tests

import (
	"go-person-info-service/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRegisterUser(t *testing.T) {
	router := gin.Default()
	router.POST("/register", controllers.RegisterUser)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/register", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
}
