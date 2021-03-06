// Code generated by entc, DO NOT EDIT.

package opt

import (
	"homework-backend/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Body applies equality check predicate on the "body" field. It's identical to BodyEQ.
func Body(v string) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBody), v))
	})
}

// Weight applies equality check predicate on the "weight" field. It's identical to WeightEQ.
func Weight(v float64) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWeight), v))
	})
}

// BodyEQ applies the EQ predicate on the "body" field.
func BodyEQ(v string) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBody), v))
	})
}

// BodyNEQ applies the NEQ predicate on the "body" field.
func BodyNEQ(v string) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBody), v))
	})
}

// BodyIn applies the In predicate on the "body" field.
func BodyIn(vs ...string) predicate.Opt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Opt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBody), v...))
	})
}

// BodyNotIn applies the NotIn predicate on the "body" field.
func BodyNotIn(vs ...string) predicate.Opt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Opt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBody), v...))
	})
}

// BodyGT applies the GT predicate on the "body" field.
func BodyGT(v string) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBody), v))
	})
}

// BodyGTE applies the GTE predicate on the "body" field.
func BodyGTE(v string) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBody), v))
	})
}

// BodyLT applies the LT predicate on the "body" field.
func BodyLT(v string) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBody), v))
	})
}

// BodyLTE applies the LTE predicate on the "body" field.
func BodyLTE(v string) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBody), v))
	})
}

// BodyContains applies the Contains predicate on the "body" field.
func BodyContains(v string) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBody), v))
	})
}

// BodyHasPrefix applies the HasPrefix predicate on the "body" field.
func BodyHasPrefix(v string) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBody), v))
	})
}

// BodyHasSuffix applies the HasSuffix predicate on the "body" field.
func BodyHasSuffix(v string) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBody), v))
	})
}

// BodyEqualFold applies the EqualFold predicate on the "body" field.
func BodyEqualFold(v string) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBody), v))
	})
}

// BodyContainsFold applies the ContainsFold predicate on the "body" field.
func BodyContainsFold(v string) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBody), v))
	})
}

// WeightEQ applies the EQ predicate on the "weight" field.
func WeightEQ(v float64) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWeight), v))
	})
}

// WeightNEQ applies the NEQ predicate on the "weight" field.
func WeightNEQ(v float64) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWeight), v))
	})
}

// WeightIn applies the In predicate on the "weight" field.
func WeightIn(vs ...float64) predicate.Opt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Opt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldWeight), v...))
	})
}

// WeightNotIn applies the NotIn predicate on the "weight" field.
func WeightNotIn(vs ...float64) predicate.Opt {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Opt(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldWeight), v...))
	})
}

// WeightGT applies the GT predicate on the "weight" field.
func WeightGT(v float64) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWeight), v))
	})
}

// WeightGTE applies the GTE predicate on the "weight" field.
func WeightGTE(v float64) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWeight), v))
	})
}

// WeightLT applies the LT predicate on the "weight" field.
func WeightLT(v float64) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWeight), v))
	})
}

// WeightLTE applies the LTE predicate on the "weight" field.
func WeightLTE(v float64) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWeight), v))
	})
}

// HasQuestion applies the HasEdge predicate on the "question" edge.
func HasQuestion() predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(QuestionTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, QuestionTable, QuestionPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasQuestionWith applies the HasEdge predicate on the "question" edge with a given conditions (other predicates).
func HasQuestionWith(preds ...predicate.Question) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(QuestionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, true, QuestionTable, QuestionPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Opt) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Opt) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
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
func Not(p predicate.Opt) predicate.Opt {
	return predicate.Opt(func(s *sql.Selector) {
		p(s.Not())
	})
}
