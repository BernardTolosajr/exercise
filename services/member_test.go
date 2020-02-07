package services

import (
	"errors"
	"testing"

	"github.com/exercise/models"
	"github.com/exercise/repositories"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestMemberServiceCreateNewCommentWhenSuccess(t *testing.T) {
	mock := repositories.MemberRepositoryMock{}

	member := &models.Member{
		Org:   "org",
		Login: "login",
	}

	mock.On("Create", member).Return(primitive.NewObjectID(), nil)

	service := NewMemberService(&mock)

	id, _ := service.Create(member)

	mock.AssertCalled(t, "Create", member)

	assert.Equal(t, 24, len(id))
}

func TestMemberServiceCreateNewCommentWhenResultIsNil(t *testing.T) {
	mock := repositories.MemberRepositoryMock{}

	member := &models.Member{
		Org:   "org",
		Login: "login",
	}

	mock.On("Create", member).Return(nil, nil)

	service := NewMemberService(&mock)

	id, _ := service.Create(member)
	assert.Equal(t, "", id)
}

func TestMemberServiceCreateNewCommentWhenWrongObjectId(t *testing.T) {
	mock := repositories.MemberRepositoryMock{}

	member := &models.Member{
		Org:   "org",
		Login: "login",
	}

	mock.On("Create", member).Return("1", nil)

	service := NewMemberService(&mock)

	id, _ := service.Create(member)
	assert.Equal(t, "", id)
}

func TestMemberServiceCreateNewCommentWhenFailed(t *testing.T) {
	mock := repositories.MemberRepositoryMock{}

	member := &models.Member{
		Org:   "org",
		Login: "login",
	}

	mock.On("Create", member).Return("", errors.New("ops"))

	service := NewMemberService(&mock)

	_, err := service.Create(member)

	assert.Equal(t, "ops", err.Error())
}

func TestMemberServiceGetAllWhenSuccessReturnArrayOfMember(t *testing.T) {
	org := "foo"

	mock := repositories.MemberRepositoryMock{}

	comments := []*models.Member{&models.Member{Login: "foo"}}

	mock.On("GetAllby", org).Return(comments, nil)

	service := NewMemberService(&mock)

	results, _ := service.GetAllBy(org)

	assert.Equal(t, 1, len(results))
}

func TestMemberServiceGetAllWhenSuccessReturnError(t *testing.T) {
	org := "foo"

	mock := repositories.MemberRepositoryMock{}

	mock.On("GetAllby", org).Return([]*models.Member{}, errors.New("ops"))

	service := NewMemberService(&mock)

	_, err := service.GetAllBy(org)

	assert.Equal(t, "ops", err.Error())
}

func TestMemberServiceGetAllWhenSuccessReturnEmptyArray(t *testing.T) {
	org := "foo"

	mock := repositories.MemberRepositoryMock{}

	mock.On("GetAllby", org).Return([]*models.Member{}, nil)

	service := NewMemberService(&mock)

	results, _ := service.GetAllBy(org)

	assert.Equal(t, 0, len(results))
}
