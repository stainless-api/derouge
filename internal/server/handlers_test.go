package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stainless-api/derouge/internal/keystore"
	"github.com/stainless-api/derouge/internal/revocation"
	"github.com/stretchr/testify/assert"
)

func TestRevokeErrorContentType(t *testing.T) {
	ks, err := keystore.New("", true)
	if err != nil {
		t.Fatal(err)
	}
	dl := revocation.NewDenyList()
	h := NewHandlers(ks, dl)

	req := httptest.NewRequest("POST", "/revoke", strings.NewReader("not json"))
	w := httptest.NewRecorder()

	h.Revoke(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
}
