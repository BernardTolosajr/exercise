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
