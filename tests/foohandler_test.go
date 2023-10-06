package tests

import (
	"automated_tests/handlers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHandlerGetFoo(t *testing.T) {
	// arrange
	server := httptest.NewServer(http.HandlerFunc(handlers.HandleGetFoo))

	// act
	res, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}

	// assert
	require.Equal(t, res.StatusCode, http.StatusOK)
}
