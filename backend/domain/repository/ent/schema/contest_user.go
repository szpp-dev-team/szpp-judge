package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ContestUser struct {
	ent.Schema
}

func (ContestUser) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("role"),
		field.Int("contest_id"),
		field.Int("user_id"),
	}
}

func (ContestUser) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("contest", Contest.Type).
			Field("contest_id").
			Unique().
			Required(),
		edge.To("user", User.Type).
			Field("user_id").
			Unique().
			Required(),
	}
}
