package schema

import "entgo.io/ent"

// Contest holds the schema definition for the Contest entity.
type Contest struct {
	ent.Schema
}

// Fields of the Contest.
func (Contest) Fields() []ent.Field {
	return nil
}

// Edges of the Contest.
func (Contest) Edges() []ent.Edge {
	return nil
}
