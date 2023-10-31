package test

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"testing"
)

func TestGetUserId(t *testing.T) {
	recorder := httptest.NewRecorder()
	context := GetTestGinContext(recorder)

	params := []gin.Param{
		{
			Key:   "id",
			Value: "1",
		},
	}

	u := url.Values{}
	u.Set("foo", "bar")

	MakeGet(context, params, u)
	GetUserId(context)

	assert.EqualValues(t, http.StatusOK, recorder.Code)
	got, _ := strconv.Atoi(recorder.Body.String())
	assert.EqualValues(t, 1, got)

}
func GetTestGinContext(recorder *httptest.ResponseRecorder) *gin.Context {
	gin.SetMode(gin.TestMode)

	ctx, _ := gin.CreateTestContext(recorder)
	ctx.Request = &http.Request{
		Header: make(http.Header),
		URL:    &url.URL{},
	}
	return ctx
}

func MakeGet(c *gin.Context, param gin.Params, u url.Values) {
	c.Request.Method = "GET"
	c.Request.Header.Set("Content-Type", "application/json")
	c.Params = param
	c.Request.URL.RawQuery = u.Encode()
}
