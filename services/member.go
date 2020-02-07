package services

import (
	"errors"

	"github.com/exercise/models"
	"github.com/exercise/repositories"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IMembersService interface {
	Create(comment *models.Member) (string, error)
	GetAllBy(org string) ([]*MemberView, error)
}

type MembersService struct {
	MembersRepository repositories.IMembersRepository
}

// Comment view
type MemberView struct {
	Org          string
	Login        string
	AvatarUrl    string
	Followers    int64
	Following    int64
	FollowersUrl string
	FollowingUrl string
}

func NewMemberService(memberRepository repositories.IMembersRepository) *MembersService {
	return &MembersService{
		MembersRepository: memberRepository,
	}
}

// Create member return doc id and error
func (o *MembersService) Create(member *models.Member) (string, error) {
	result, err := o.MembersRepository.Create(member)

	if err != nil {
		return "", err
	}

	if result != nil {
		if oid, ok := result.(primitive.ObjectID); ok {
			return oid.Hex(), nil
		}
		return "", errors.New("Unable to cast ObjectId")
	}

	return "", nil
}

// Get all by org return Array of member and error
func (o *MembersService) GetAllBy(org string) ([]*MemberView, error) {
	results, err := o.MembersRepository.GetAllby(org)

	if err != nil {
		log.Errorf("%v", err)
		return nil, err
	}

	// return empty array instead of null
	if len(results) == 0 {
		return make([]*MemberView, 0), nil
	}

	var members []*MemberView

	for _, member := range results {
		members = append(members, &MemberView{
			Org:       member.Org,
			Login:     member.Login,
			Followers: member.Followers,
			Following: member.Following,
		})
	}

	return members, nil
}
