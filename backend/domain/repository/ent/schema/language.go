package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Language struct {
	ent.Schema
}

func (Language) Field() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name"),
		field.String("language_slug"),
	}
}

func (Language) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("submit", Submit.Type),
	}
}
