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

		field.Enum("judge_type").Values("normal", "eps", "interactive", "custom"),
		field.Bool("case_insensitive").Optional().Nillable(),  // normal
		field.Uint("ndigits").Optional().Nillable(),           // eps
		field.String("judge_code_path").Optional().Nillable(), // interactive, custom

		field.Time("created_at"),
		field.Time("updated_at").Optional().Nillable(),
	}
}

func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("testcase_sets", TestcaseSet.Type),
		edge.To("testcases", Testcase.Type),
		edge.From("user", User.Type).Ref("tasks").Unique().Required(),
	}
}
