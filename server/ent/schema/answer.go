package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Answer holds the schema definition for the Answer entity.
type Answer struct {
	ent.Schema
}

// Fields of the Answer.
func (Answer) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("question_id"),
		field.String("body").Optional().Nillable(),
		field.String("option_id").Optional().Nillable(),
	}
}

// Edges of the Answer.
func (Answer) Edges() []ent.Edge {
	return nil
}
