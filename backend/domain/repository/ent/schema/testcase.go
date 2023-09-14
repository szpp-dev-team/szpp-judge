package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type Testcase struct {
	ent.Schema
}

func (Testcase) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name"),
		field.String("description").Optional().Nillable(),
		field.Time("created_at"),
		field.Time("updated_at").Optional().Nillable(),
	}
}

func (Testcase) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("testcase_sets", TestcaseSet.Type).Ref("testcases"),
		edge.From("task", Task.Type).Ref("testcases").Unique().Required(),
	}
}

func (Testcase) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Edges("task").Unique(),
	}
}
