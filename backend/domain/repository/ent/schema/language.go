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
		field.String("name"), // display name
		field.String("slug").Unique(),
	}
}

func (Language) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("submits", Submit.Type),
	}
}
