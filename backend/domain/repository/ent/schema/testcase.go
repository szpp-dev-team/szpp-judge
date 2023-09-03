package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Testcase struct {
	ent.Schema
}

func (Testcase) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name").Unique(),
		field.String("description"),
		field.Time("created_at"),
		field.Time("updated_at").Optional().Nillable(),
	}
}

func (Testcase) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("testcase_sets", TestcaseSet.Type).Ref("testcases"),
	}
}
