package service

import (
	"errors"

	"github.com/akshaym-3255/family-tree/internal/repositories"
)

type RelationshipService interface {
	AddRelationship(name string) error
	CheckRelationShipExists(name string) (bool, error)
}

type relationshipService struct {
	repo repositories.RelationshipRepository
}

func NewRelationshipService(repo repositories.RelationshipRepository) RelationshipService {
	return &relationshipService{repo: repo}
}

func (r *relationshipService) AddRelationship(name string) error {
	relationships, err := r.repo.GetRelationships()
	if err != nil {
		return err
	}

	for _, relationship := range relationships.Names {
		if relationship == name {
			return errors.New("relationship already exists")
		}
	}

	relationships.Names = append(relationships.Names, name)
	r.repo.UpdateRelationships(relationships)
	return nil
}

func (r relationshipService) CheckRelationShipExists(name string) (bool, error) {
	relationships, err := r.repo.GetRelationships()
	if err != nil {
		return false, err
	}

	for _, relationship := range relationships.Names {
		if relationship == name {
			return true, nil
		}
	}

	return false, nil
}
