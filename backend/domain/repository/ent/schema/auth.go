package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type RefreshToken struct {
	ent.Schema
}

func (RefreshToken) Fields() []ent.Field {
	return []ent.Field{
		field.String("token").Unique(),
		field.Time("expires_at"),
		field.Bool("is_dead"),
	}
}

func (RefreshToken) Edges() []ent.Edge {
	return nil
}
