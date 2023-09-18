package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ContestUsers struct {
	ent.Schema
}

func (ContestUsers) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("role"),
		field.Int("contest_id"),
		field.Int("user_id"),
	}
}

func (ContestUsers) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("contests", Contest.Type).Ref("contest_users").Field("contest_id").Unique().Required(),
		edge.To("users", User.Type).Field("user_id").Unique().Required(),
	}
}
