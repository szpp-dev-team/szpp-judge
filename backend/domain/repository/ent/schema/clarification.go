package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Clarification struct {
	ent.Schema
}

// Fields of the ContestClarification.
func (Clarification) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("content"),
		field.Bool("is_public"),
		field.Time("created_at"),
		field.Time("updated_at"),
		field.String("answer_content").
			Optional().
			Nillable(),
		field.Time("answer_created_at").
			Optional().
			Nillable(),
		field.Time("answer_updated_at").
			Optional().
			Nillable(),
	}
}

func (Clarification) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("contest", Contest.Type).
			Ref("clarifications").
			Unique().
			Required(),
		edge.From("task", Task.Type).
			Ref("clarifications").
			Required(),
		edge.From("user", User.Type).
			Ref("clarifications").
			Required(),
		edge.From("answer_user", User.Type).
			Ref("answered_clarifications").
			Required(),
	}
}
