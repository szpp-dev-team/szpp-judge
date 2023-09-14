package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type TestcaseSet struct {
	ent.Schema
}

func (TestcaseSet) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("name"),
		field.Int("score"),
		field.Bool("is_sample"),
		field.Time("created_at"),
		field.Time("updated_at").Optional().Nillable(),
	}
}

func (TestcaseSet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("task", Task.Type).Ref("testcase_sets").Unique().Required(),
		edge.To("testcases", Testcase.Type),
	}
}

func (TestcaseSet) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").Edges("task").Unique(),
	}
}
