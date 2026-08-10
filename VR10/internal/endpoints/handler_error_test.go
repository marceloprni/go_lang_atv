package endpoints

import (
	internalerrors "emailn/internal/internal-errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_HandlerError_when_endpoint_returns_error(t *testing.T) {
	// Create a mock handler that returns an error
	assert := assert.New(t)
	endpoint := func(w http.ResponseWriter, r *http.Request) (interface{}, int, error) {
		return nil, 0, internalerrors.ErrInternal
	}

	handlerFunc := HandlerError(endpoint)

	// Create a test request
	req, _ := http.NewRequest("GET", "/", nil)
	res := httptest.NewRecorder()

	// Call the handler
	handlerFunc.ServeHTTP(res, req)
	assert.Equal(http.StatusInternalServerError, res.Code)

	assert.Contains(res.Body.String(), internalerrors.ErrInternal.Error())

}
