package integration

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bmizerany/assert"
)

func TestGetComments(t *testing.T) {
	defer DropCommentCollection(mongo)

	req, err := http.NewRequest(http.MethodGet, "/orgs/foo/comments", nil)
	if err != nil {
		t.Errorf("error creating request: %v", err)
	}

	rr := httptest.NewRecorder()

	r.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	// Check the response body is what we expect.
	assert.Equal(t, "{\"comments\":[{\"Comment\":\"test\"}]}\n", rr.Body.String())
}
