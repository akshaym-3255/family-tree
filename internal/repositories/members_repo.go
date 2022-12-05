package repositories

import (
	"encoding/json"
	"os"

	"github.com/akshaym-3255/family-tree/internal/models"
)

type MemberRepository interface {
	GetMembers() ([]models.Member, error)
	UpdateMembers([]models.Member) error
}

type memberRepository struct {
}

func NewMemberRepository() MemberRepository {
	return &memberRepository{}
}

func (p *memberRepository) GetMembers() ([]models.Member, error) {
	var members []models.Member

	data, err := os.ReadFile("internal/database/members.json")
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &members)
	return members, nil
}

func (p *memberRepository) UpdateMembers(members []models.Member) error {

	bytes, err := json.Marshal(members)
	if err != nil {
		return err
	}
	os.WriteFile("internal/database/members.json", bytes, 0777)

	return nil
}
