package service

import (
	"errors"
	"fmt"

	"github.com/akshaym-3255/family-tree/internal/models"
	"github.com/akshaym-3255/family-tree/internal/repositories"
)

type MemberService interface {
	AddMember(name string) error
	AddConnection(first_member string, second_member string, relation string) error
	GetCountOfRelationShip(name string, relation string) (int, error)
	GetFather(name string) (string, error)
}

type memberService struct {
	repo                repositories.MemberRepository
	relationshipService RelationshipService
}

func NewMemberService(repo repositories.MemberRepository, relationshipService RelationshipService) MemberService {
	return &memberService{repo: repo, relationshipService: relationshipService}
}

func (m *memberService) AddMember(name string) error {
	exists, err := m.checkMemberExits(name)
	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("person already exists %s", name)
	}
	member := models.Member{}
	member.Name = name
	members, err := m.repo.GetMembers()
	if err != nil {
		return err
	}
	members = append(members, member)
	m.repo.UpdateMembers(members)
	return nil

}

func (m *memberService) AddConnection(first_member string, second_member string, relation string) error {
	members, err := m.repo.GetMembers()
	if err != nil {
		return err
	}
	// check relation exits in array
	exists, err := m.relationshipService.CheckRelationShipExists(relation)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("relation not found %s", relation)
	}
	// check members exits
	exists, err = m.checkMemberExits(first_member)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("person not found %s", first_member)
	}

	exists, err = m.checkMemberExits(second_member)
	if err != nil {
		return err
	}

	if !exists {
		return fmt.Errorf("person not found %s", second_member)
	}

	for index, member := range members {
		if member.Name == second_member {
			if member.Connections != nil {
				relationFound := false
				for conn_index, connection := range member.Connections {
					if connection.Relation == relation {
						for _, name := range members[index].Connections[conn_index].Ids {
							if name == first_member {
								return fmt.Errorf("relation already exits")
							}
						}
						members[index].Connections[conn_index].Ids = append(members[index].Connections[conn_index].Ids, first_member)
						relationFound = true
						break
					}
				}
				if !relationFound {
					connection := models.Connections{
						Relation: relation,
					}
					// connection.Ids = make([]string, 1)
					connection.Ids = append(connection.Ids, first_member)
					members[index].Connections = append(members[index].Connections, connection)
					break
				}
			} else {
				members[index].Connections = make([]models.Connections, 1)
				members[index].Connections[0].Relation = relation
				members[index].Connections[0].Ids = make([]string, 1)
				members[index].Connections[0].Ids[0] = first_member
				break
			}

		}
	}

	m.repo.UpdateMembers(members)
	return nil
}

func (m *memberService) GetFather(name string) (string, error) {
	members, err := m.repo.GetMembers()
	if err != nil {
		return "", err
	}
	for _, member := range members {
		if member.Name == name {
			for _, connection := range member.Connections {
				if connection.Relation == "father" {
					return connection.Ids[0], nil
				}
			}
		}
	}
	return "", errors.New("no relation found")
}

func (m *memberService) GetCountOfRelationShip(name string, relation string) (int, error) {
	singularRelation := m.getPluralToSingularRelation(relation)
	if singularRelation == "" {
		return 0, fmt.Errorf("check the relationship")
	}
	members, err := m.repo.GetMembers()
	if err != nil {
		return 0, err
	}
	for _, member := range members {
		if member.Name == name {
			for _, connection := range member.Connections {
				if connection.Relation == singularRelation {
					return len(connection.Ids), nil
				}
			}
		}
	}
	return 0, nil
}

func (m *memberService) checkMemberExits(name string) (bool, error) {
	members, err := m.repo.GetMembers()

	if err != nil {
		return false, err
	}
	for _, member := range members {
		if member.Name == name {
			return true, nil
		}
	}
	return false, nil
}

func (m *memberService) getPluralToSingularRelation(relation string) string {
	switch relation {
	case "sons":
		return "son"
	case "wives":
		return "wife"
	case "daughters":
		return "daughter"
	}
	return ""
}
