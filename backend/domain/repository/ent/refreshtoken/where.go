// Code generated by ent, DO NOT EDIT.

package refreshtoken

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLTE(FieldID, id))
}

// Token applies equality check predicate on the "token" field. It's identical to TokenEQ.
func Token(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldToken, v))
}

// ExpiresAt applies equality check predicate on the "expires_at" field. It's identical to ExpiresAtEQ.
func ExpiresAt(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldExpiresAt, v))
}

// IsDead applies equality check predicate on the "is_dead" field. It's identical to IsDeadEQ.
func IsDead(v bool) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldIsDead, v))
}

// TokenEQ applies the EQ predicate on the "token" field.
func TokenEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldToken, v))
}

// TokenNEQ applies the NEQ predicate on the "token" field.
func TokenNEQ(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldToken, v))
}

// TokenIn applies the In predicate on the "token" field.
func TokenIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIn(FieldToken, vs...))
}

// TokenNotIn applies the NotIn predicate on the "token" field.
func TokenNotIn(vs ...string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotIn(FieldToken, vs...))
}

// TokenGT applies the GT predicate on the "token" field.
func TokenGT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGT(FieldToken, v))
}

// TokenGTE applies the GTE predicate on the "token" field.
func TokenGTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGTE(FieldToken, v))
}

// TokenLT applies the LT predicate on the "token" field.
func TokenLT(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLT(FieldToken, v))
}

// TokenLTE applies the LTE predicate on the "token" field.
func TokenLTE(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLTE(FieldToken, v))
}

// TokenContains applies the Contains predicate on the "token" field.
func TokenContains(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContains(FieldToken, v))
}

// TokenHasPrefix applies the HasPrefix predicate on the "token" field.
func TokenHasPrefix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasPrefix(FieldToken, v))
}

// TokenHasSuffix applies the HasSuffix predicate on the "token" field.
func TokenHasSuffix(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldHasSuffix(FieldToken, v))
}

// TokenEqualFold applies the EqualFold predicate on the "token" field.
func TokenEqualFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEqualFold(FieldToken, v))
}

// TokenContainsFold applies the ContainsFold predicate on the "token" field.
func TokenContainsFold(v string) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldContainsFold(FieldToken, v))
}

// ExpiresAtEQ applies the EQ predicate on the "expires_at" field.
func ExpiresAtEQ(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldExpiresAt, v))
}

// ExpiresAtNEQ applies the NEQ predicate on the "expires_at" field.
func ExpiresAtNEQ(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldExpiresAt, v))
}

// ExpiresAtIn applies the In predicate on the "expires_at" field.
func ExpiresAtIn(vs ...time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldIn(FieldExpiresAt, vs...))
}

// ExpiresAtNotIn applies the NotIn predicate on the "expires_at" field.
func ExpiresAtNotIn(vs ...time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNotIn(FieldExpiresAt, vs...))
}

// ExpiresAtGT applies the GT predicate on the "expires_at" field.
func ExpiresAtGT(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGT(FieldExpiresAt, v))
}

// ExpiresAtGTE applies the GTE predicate on the "expires_at" field.
func ExpiresAtGTE(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldGTE(FieldExpiresAt, v))
}

// ExpiresAtLT applies the LT predicate on the "expires_at" field.
func ExpiresAtLT(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLT(FieldExpiresAt, v))
}

// ExpiresAtLTE applies the LTE predicate on the "expires_at" field.
func ExpiresAtLTE(v time.Time) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldLTE(FieldExpiresAt, v))
}

// IsDeadEQ applies the EQ predicate on the "is_dead" field.
func IsDeadEQ(v bool) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldEQ(FieldIsDead, v))
}

// IsDeadNEQ applies the NEQ predicate on the "is_dead" field.
func IsDeadNEQ(v bool) predicate.RefreshToken {
	return predicate.RefreshToken(sql.FieldNEQ(FieldIsDead, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.RefreshToken) predicate.RefreshToken {
	return predicate.RefreshToken(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.RefreshToken) predicate.RefreshToken {
	return predicate.RefreshToken(func(s *sql.Selector) {
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
func Not(p predicate.RefreshToken) predicate.RefreshToken {
	return predicate.RefreshToken(func(s *sql.Selector) {
		p(s.Not())
	})
}
