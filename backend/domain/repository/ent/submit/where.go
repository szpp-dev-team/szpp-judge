// Code generated by ent, DO NOT EDIT.

package submit

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Submit {
	return predicate.Submit(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Submit {
	return predicate.Submit(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Submit {
	return predicate.Submit(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Submit {
	return predicate.Submit(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Submit {
	return predicate.Submit(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Submit {
	return predicate.Submit(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Submit {
	return predicate.Submit(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Submit {
	return predicate.Submit(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Submit {
	return predicate.Submit(sql.FieldLTE(FieldID, id))
}

// HasUsers applies the HasEdge predicate on the "users" edge.
func HasUsers() predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, UsersTable, UsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUsersWith applies the HasEdge predicate on the "users" edge with a given conditions (other predicates).
func HasUsersWith(preds ...predicate.User) predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
		step := newUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTask applies the HasEdge predicate on the "task" edge.
func HasTask() predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, TaskTable, TaskColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTaskWith applies the HasEdge predicate on the "task" edge with a given conditions (other predicates).
func HasTaskWith(preds ...predicate.Task) predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
		step := newTaskStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasLanguage applies the HasEdge predicate on the "language" edge.
func HasLanguage() predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, LanguageTable, LanguageColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasLanguageWith applies the HasEdge predicate on the "language" edge with a given conditions (other predicates).
func HasLanguageWith(preds ...predicate.Language) predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
		step := newLanguageStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasTestcaseResults applies the HasEdge predicate on the "testcase_results" edge.
func HasTestcaseResults() predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, TestcaseResultsTable, TestcaseResultsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTestcaseResultsWith applies the HasEdge predicate on the "testcase_results" edge with a given conditions (other predicates).
func HasTestcaseResultsWith(preds ...predicate.TestcaseResult) predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
		step := newTestcaseResultsStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Submit) predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Submit) predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Submit) predicate.Submit {
	return predicate.Submit(func(s *sql.Selector) {
		p(s.Not())
	})
}
