package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/exercise/models"
	"github.com/exercise/services"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateNewMemberHandlerWhenSuccess(t *testing.T) {
	member := &models.Member{
		Org: "foo",
	}

	payload := []byte(`{"comment":"bar"}`)
	req, err := http.NewRequest("POST", "/orgs/foo/members", bytes.NewBuffer(payload))

	if err != nil {
		t.Fatal(err)
	}

	//Hack to fake gorilla/mux vars
	vars := map[string]string{
		"name": "foo",
	}

	req = mux.SetURLVars(req, vars)

	//setup mock service
	mock := &services.MemberServiceMock{}

	mock.On("Create", member).Return("1", nil)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(PostMemberHandler(mock))

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	// Check the response body is what we expect.
	assert.Equal(t, "{\"id\":\"1\"}\n", rr.Body.String())
	// make sure our depedency is called with correct parameter
	mock.AssertCalled(t, "Create", member)
}

func TestGetMemberHandlerWhenSuccessReturnArrayOfComment(t *testing.T) {
	org := "foo"
	req, err := http.NewRequest("GET", "/orgs/foo/members", nil)

	if err != nil {
		t.Fatal(err)
	}

	//Hack to fake gorilla/mux vars
	vars := map[string]string{
		"name": org,
	}

	req = mux.SetURLVars(req, vars)

	//setup mock service
	mock := &services.MemberServiceMock{}

	members := []*services.MemberView{&services.MemberView{
		Org: "theorg",
	}}

	mock.On("GetAllBy", org).Return(members, nil)

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(GetMembersHandler(mock))

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)

	// FIXME to long, maybe we can sanitize this.
	expected := "{\"members\":[{\"Org\":\"theorg\",\"Login\":\"\",\"AvatarUrl\":\"\",\"Followers\":0,\"Following\":0,\"FollowersUrl\":\"\",\"FollowingUrl\":\"\"}]}\n"

	// Check the response body is what we expect.
	assert.Equal(t, expected, rr.Body.String())
	// make sure our depedency is called with correct parameter
	mock.AssertCalled(t, "GetAllBy", org)
}
