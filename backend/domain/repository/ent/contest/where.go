// Code generated by ent, DO NOT EDIT.

package contest

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Contest {
	return predicate.Contest(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Contest {
	return predicate.Contest(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Contest {
	return predicate.Contest(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Contest {
	return predicate.Contest(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Contest {
	return predicate.Contest(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Contest {
	return predicate.Contest(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Contest {
	return predicate.Contest(sql.FieldLTE(FieldID, id))
}

// Slug applies equality check predicate on the "slug" field. It's identical to SlugEQ.
func Slug(v string) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldSlug, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldDescription, v))
}

// TaskID applies equality check predicate on the "task_id" field. It's identical to TaskIDEQ.
func TaskID(v int) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldTaskID, v))
}

// ClarificationID applies equality check predicate on the "clarification_id" field. It's identical to ClarificationIDEQ.
func ClarificationID(v int) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldClarificationID, v))
}

// StartAt applies equality check predicate on the "start_at" field. It's identical to StartAtEQ.
func StartAt(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldStartAt, v))
}

// EndAt applies equality check predicate on the "end_at" field. It's identical to EndAtEQ.
func EndAt(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldEndAt, v))
}

// SlugEQ applies the EQ predicate on the "slug" field.
func SlugEQ(v string) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldSlug, v))
}

// SlugNEQ applies the NEQ predicate on the "slug" field.
func SlugNEQ(v string) predicate.Contest {
	return predicate.Contest(sql.FieldNEQ(FieldSlug, v))
}

// SlugIn applies the In predicate on the "slug" field.
func SlugIn(vs ...string) predicate.Contest {
	return predicate.Contest(sql.FieldIn(FieldSlug, vs...))
}

// SlugNotIn applies the NotIn predicate on the "slug" field.
func SlugNotIn(vs ...string) predicate.Contest {
	return predicate.Contest(sql.FieldNotIn(FieldSlug, vs...))
}

// SlugGT applies the GT predicate on the "slug" field.
func SlugGT(v string) predicate.Contest {
	return predicate.Contest(sql.FieldGT(FieldSlug, v))
}

// SlugGTE applies the GTE predicate on the "slug" field.
func SlugGTE(v string) predicate.Contest {
	return predicate.Contest(sql.FieldGTE(FieldSlug, v))
}

// SlugLT applies the LT predicate on the "slug" field.
func SlugLT(v string) predicate.Contest {
	return predicate.Contest(sql.FieldLT(FieldSlug, v))
}

// SlugLTE applies the LTE predicate on the "slug" field.
func SlugLTE(v string) predicate.Contest {
	return predicate.Contest(sql.FieldLTE(FieldSlug, v))
}

// SlugContains applies the Contains predicate on the "slug" field.
func SlugContains(v string) predicate.Contest {
	return predicate.Contest(sql.FieldContains(FieldSlug, v))
}

// SlugHasPrefix applies the HasPrefix predicate on the "slug" field.
func SlugHasPrefix(v string) predicate.Contest {
	return predicate.Contest(sql.FieldHasPrefix(FieldSlug, v))
}

// SlugHasSuffix applies the HasSuffix predicate on the "slug" field.
func SlugHasSuffix(v string) predicate.Contest {
	return predicate.Contest(sql.FieldHasSuffix(FieldSlug, v))
}

// SlugEqualFold applies the EqualFold predicate on the "slug" field.
func SlugEqualFold(v string) predicate.Contest {
	return predicate.Contest(sql.FieldEqualFold(FieldSlug, v))
}

// SlugContainsFold applies the ContainsFold predicate on the "slug" field.
func SlugContainsFold(v string) predicate.Contest {
	return predicate.Contest(sql.FieldContainsFold(FieldSlug, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Contest {
	return predicate.Contest(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Contest {
	return predicate.Contest(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Contest {
	return predicate.Contest(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Contest {
	return predicate.Contest(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Contest {
	return predicate.Contest(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Contest {
	return predicate.Contest(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Contest {
	return predicate.Contest(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Contest {
	return predicate.Contest(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Contest {
	return predicate.Contest(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Contest {
	return predicate.Contest(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Contest {
	return predicate.Contest(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Contest {
	return predicate.Contest(sql.FieldContainsFold(FieldDescription, v))
}

// TaskIDEQ applies the EQ predicate on the "task_id" field.
func TaskIDEQ(v int) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldTaskID, v))
}

// TaskIDNEQ applies the NEQ predicate on the "task_id" field.
func TaskIDNEQ(v int) predicate.Contest {
	return predicate.Contest(sql.FieldNEQ(FieldTaskID, v))
}

// TaskIDIn applies the In predicate on the "task_id" field.
func TaskIDIn(vs ...int) predicate.Contest {
	return predicate.Contest(sql.FieldIn(FieldTaskID, vs...))
}

// TaskIDNotIn applies the NotIn predicate on the "task_id" field.
func TaskIDNotIn(vs ...int) predicate.Contest {
	return predicate.Contest(sql.FieldNotIn(FieldTaskID, vs...))
}

// TaskIDGT applies the GT predicate on the "task_id" field.
func TaskIDGT(v int) predicate.Contest {
	return predicate.Contest(sql.FieldGT(FieldTaskID, v))
}

// TaskIDGTE applies the GTE predicate on the "task_id" field.
func TaskIDGTE(v int) predicate.Contest {
	return predicate.Contest(sql.FieldGTE(FieldTaskID, v))
}

// TaskIDLT applies the LT predicate on the "task_id" field.
func TaskIDLT(v int) predicate.Contest {
	return predicate.Contest(sql.FieldLT(FieldTaskID, v))
}

// TaskIDLTE applies the LTE predicate on the "task_id" field.
func TaskIDLTE(v int) predicate.Contest {
	return predicate.Contest(sql.FieldLTE(FieldTaskID, v))
}

// ClarificationIDEQ applies the EQ predicate on the "clarification_id" field.
func ClarificationIDEQ(v int) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldClarificationID, v))
}

// ClarificationIDNEQ applies the NEQ predicate on the "clarification_id" field.
func ClarificationIDNEQ(v int) predicate.Contest {
	return predicate.Contest(sql.FieldNEQ(FieldClarificationID, v))
}

// ClarificationIDIn applies the In predicate on the "clarification_id" field.
func ClarificationIDIn(vs ...int) predicate.Contest {
	return predicate.Contest(sql.FieldIn(FieldClarificationID, vs...))
}

// ClarificationIDNotIn applies the NotIn predicate on the "clarification_id" field.
func ClarificationIDNotIn(vs ...int) predicate.Contest {
	return predicate.Contest(sql.FieldNotIn(FieldClarificationID, vs...))
}

// ClarificationIDGT applies the GT predicate on the "clarification_id" field.
func ClarificationIDGT(v int) predicate.Contest {
	return predicate.Contest(sql.FieldGT(FieldClarificationID, v))
}

// ClarificationIDGTE applies the GTE predicate on the "clarification_id" field.
func ClarificationIDGTE(v int) predicate.Contest {
	return predicate.Contest(sql.FieldGTE(FieldClarificationID, v))
}

// ClarificationIDLT applies the LT predicate on the "clarification_id" field.
func ClarificationIDLT(v int) predicate.Contest {
	return predicate.Contest(sql.FieldLT(FieldClarificationID, v))
}

// ClarificationIDLTE applies the LTE predicate on the "clarification_id" field.
func ClarificationIDLTE(v int) predicate.Contest {
	return predicate.Contest(sql.FieldLTE(FieldClarificationID, v))
}

// StartAtEQ applies the EQ predicate on the "start_at" field.
func StartAtEQ(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldStartAt, v))
}

// StartAtNEQ applies the NEQ predicate on the "start_at" field.
func StartAtNEQ(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldNEQ(FieldStartAt, v))
}

// StartAtIn applies the In predicate on the "start_at" field.
func StartAtIn(vs ...time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldIn(FieldStartAt, vs...))
}

// StartAtNotIn applies the NotIn predicate on the "start_at" field.
func StartAtNotIn(vs ...time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldNotIn(FieldStartAt, vs...))
}

// StartAtGT applies the GT predicate on the "start_at" field.
func StartAtGT(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldGT(FieldStartAt, v))
}

// StartAtGTE applies the GTE predicate on the "start_at" field.
func StartAtGTE(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldGTE(FieldStartAt, v))
}

// StartAtLT applies the LT predicate on the "start_at" field.
func StartAtLT(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldLT(FieldStartAt, v))
}

// StartAtLTE applies the LTE predicate on the "start_at" field.
func StartAtLTE(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldLTE(FieldStartAt, v))
}

// EndAtEQ applies the EQ predicate on the "end_at" field.
func EndAtEQ(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldEQ(FieldEndAt, v))
}

// EndAtNEQ applies the NEQ predicate on the "end_at" field.
func EndAtNEQ(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldNEQ(FieldEndAt, v))
}

// EndAtIn applies the In predicate on the "end_at" field.
func EndAtIn(vs ...time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldIn(FieldEndAt, vs...))
}

// EndAtNotIn applies the NotIn predicate on the "end_at" field.
func EndAtNotIn(vs ...time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldNotIn(FieldEndAt, vs...))
}

// EndAtGT applies the GT predicate on the "end_at" field.
func EndAtGT(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldGT(FieldEndAt, v))
}

// EndAtGTE applies the GTE predicate on the "end_at" field.
func EndAtGTE(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldGTE(FieldEndAt, v))
}

// EndAtLT applies the LT predicate on the "end_at" field.
func EndAtLT(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldLT(FieldEndAt, v))
}

// EndAtLTE applies the LTE predicate on the "end_at" field.
func EndAtLTE(v time.Time) predicate.Contest {
	return predicate.Contest(sql.FieldLTE(FieldEndAt, v))
}

// HasTasks applies the HasEdge predicate on the "tasks" edge.
func HasTasks() predicate.Contest {
	return predicate.Contest(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, TasksTable, TasksPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasTasksWith applies the HasEdge predicate on the "tasks" edge with a given conditions (other predicates).
func HasTasksWith(preds ...predicate.Task) predicate.Contest {
	return predicate.Contest(func(s *sql.Selector) {
		step := newTasksStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasContestUsers applies the HasEdge predicate on the "contest_users" edge.
func HasContestUsers() predicate.Contest {
	return predicate.Contest(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, ContestUsersTable, ContestUsersPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasContestUsersWith applies the HasEdge predicate on the "contest_users" edge with a given conditions (other predicates).
func HasContestUsersWith(preds ...predicate.User) predicate.Contest {
	return predicate.Contest(func(s *sql.Selector) {
		step := newContestUsersStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Contest) predicate.Contest {
	return predicate.Contest(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Contest) predicate.Contest {
	return predicate.Contest(func(s *sql.Selector) {
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
func Not(p predicate.Contest) predicate.Contest {
	return predicate.Contest(func(s *sql.Selector) {
		p(s.Not())
	})
}
