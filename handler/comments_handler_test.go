package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/exercise/models"
	"github.com/exercise/services"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewCommentsHandlerWhenSuccess(t *testing.T) {
	comment := &models.Comment{
		Org:     "foo",
		Comment: "bar",
	}

	payload := []byte(`{"comment":"bar"}`)
	req, err := http.NewRequest("POST", "/orgs/foo/comments", bytes.NewBuffer(payload))

	if err != nil {
		t.Fatal(err)
	}

	//Hack to fake gorilla/mux vars
	vars := map[string]string{
		"name": "foo",
	}

	req = mux.SetURLVars(req, vars)

	//setup mock service
	mock := &services.CommentServiceMock{}

	mock.On("Create", comment).Return("1", nil)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostCommentsHandler(mock))

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	// Check the response body is what we expect.
	assert.Equal(t, "{\"id\":\"1\"}\n", rr.Body.String())
	// make sure our depedency is called with correct parameter
	mock.AssertCalled(t, "Create", comment)
}

func TestCreateNewCommentsHandlerWhenFailed(t *testing.T) {
	comment := &models.Comment{
		Org:     "foo",
		Comment: "bar",
	}

	payload := []byte(`{"comment":"bar"}`)
	req, err := http.NewRequest("POST", "/orgs/foo/comments", bytes.NewBuffer(payload))

	if err != nil {
		t.Fatal(err)
	}

	//Hack to fake gorilla/mux vars
	vars := map[string]string{
		"name": "foo",
	}

	req = mux.SetURLVars(req, vars)

	//setup mock service
	mock := &services.CommentServiceMock{}

	mock.On("Create", comment).Return("", errors.New("something happend"))

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostCommentsHandler(mock))

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	assert.Equal(t, 422, rr.Code)
	// Check the response body is what we expect.
	assert.Equal(t, "{\"message\":\"something happend\"}\n", rr.Body.String())
}

func TestGetCommentsHandlerWhenSuccessReturnArrayOfComment(t *testing.T) {
	org := "foo"
	req, err := http.NewRequest("GET", "/orgs/foo/comments", nil)

	if err != nil {
		t.Fatal(err)
	}

	//Hack to fake gorilla/mux vars
	vars := map[string]string{
		"name": org,
	}

	req = mux.SetURLVars(req, vars)

	//setup mock service
	mock := &services.CommentServiceMock{}

	comments := []*services.Comment{&services.Comment{Comment: "awesome"}}

	mock.On("GetAllBy", org).Return(comments, nil)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetCommentsHandler(mock))

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	// Check the response body is what we expect.
	assert.Equal(t, "{\"comments\":[{\"Comment\":\"awesome\"}]}\n", rr.Body.String())
	// make sure our depedency is called with correct parameter
	mock.AssertCalled(t, "GetAllBy", org)
}
