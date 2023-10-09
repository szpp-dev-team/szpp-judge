package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Task struct {
	ent.Schema
}

func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.Int("id"),
		field.String("title"),
		field.String("statement"),
		field.String("difficulty"),
		field.Uint("exec_time_limit"),   // ms
		field.Uint("exec_memory_limit"), // MiB

		field.Time("created_at"),
		field.Time("updated_at").Optional().Nillable(),
	}
}

func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("testcase_sets", TestcaseSet.Type),
		edge.To("testcases", Testcase.Type),
		edge.To("submits", Submit.Type),
		edge.From("user", User.Type).Ref("tasks").Unique().Required(),
		edge.From("contests", Contest.Type).Ref("tasks").Through("contest_task", ContestTask.Type),
	}
}
