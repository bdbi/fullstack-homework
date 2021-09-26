package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Opt holds the schema definition for the QuestionOption entity.
type Opt struct {
	ent.Schema
}

// Fields of the Opt.
func (Opt) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).Default(uuid.New),
		field.String("body"),
		field.Float("weight"),
	}
}

// Edges of the Opt.
func (Opt) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("question", Question.Type).Ref("options"),
	}
}
