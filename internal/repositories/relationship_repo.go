package repositories

import (
	"embed"
	"encoding/json"
	"os"

	"github.com/akshaym-3255/family-tree/internal/models"
)

var RelationFile embed.FS

type RelationshipRepository interface {
	GetRelationships() (models.Relationship, error)
	UpdateRelationships(models.Relationship) error
}

type relationshipRepository struct {
}

func NewRelationshipRepository() RelationshipRepository {
	return &relationshipRepository{}
}

func (r *relationshipRepository) GetRelationships() (models.Relationship, error) {
	var relationships models.Relationship

	data, err := RelationFile.ReadFile("internal/database/relations.json")
	if err != nil {
		return models.Relationship{}, err
	}
	json.Unmarshal(data, &relationships)
	return relationships, nil
}

func (r *relationshipRepository) UpdateRelationships(relationships models.Relationship) error {
	bytes, err := json.Marshal(relationships)
	if err != nil {
		return err
	}
	os.WriteFile("internal/database/relations.json", bytes, 0644)

	return nil
}
