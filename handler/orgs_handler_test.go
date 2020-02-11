package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/exercise/models"
	"github.com/exercise/utils"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewOrganizationSuccessHandler(t *testing.T) {
	org := &models.Organization{
		Login:       "foo",
		ProfileName: "bar",
		Admin:       "admin",
	}

	payload := []byte(`{"login":"foo","profile_name":"bar","admin":"admin"}`)
	req, err := http.NewRequest("POST", "/orgs", bytes.NewBuffer(payload))

	if err != nil {
		t.Fatal(err)
	}

	//setup mock service
	mock := &utils.OrganizationServiceMock{}

	mock.On("Create", org).Return("1", nil)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostOrganizationHandler(mock))

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	// Check the response body is what we expect.
	assert.Equal(t, "{\"id\":\"1\"}\n", rr.Body.String())
	mock.AssertCalled(t, "Create", org)
}

func TestCreateNewOrganizationFailedHandler(t *testing.T) {
	payload := []byte(`{"login":"foo","profile_name":"bar","admin":"admin"}`)
	req, err := http.NewRequest("POST", "/organizations", bytes.NewBuffer(payload))

	if err != nil {
		t.Fatal(err)
	}

	//setup mock service
	mock := &utils.OrganizationServiceMock{}

	mock.On("Create", &models.Organization{
		Login:       "foo",
		ProfileName: "bar",
		Admin:       "admin",
	}).Return("", errors.New("something happend"))

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostOrganizationHandler(mock))

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	assert.Equal(t, 422, rr.Code)
	// Check the response body is what we expect.
	assert.Equal(t, "{\"message\":\"something happend\"}\n", rr.Body.String())
}
