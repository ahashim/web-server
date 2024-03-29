// Code generated by ent, DO NOT EDIT.

package poolpass

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ahashim/web-server/ent/predicate"
	"github.com/ahashim/web-server/types"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// Shares applies equality check predicate on the "shares" field. It's identical to SharesEQ.
func Shares(v *types.Uint256) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShares), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.PoolPass {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.PoolPass {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// SharesEQ applies the EQ predicate on the "shares" field.
func SharesEQ(v *types.Uint256) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShares), v))
	})
}

// SharesNEQ applies the NEQ predicate on the "shares" field.
func SharesNEQ(v *types.Uint256) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldShares), v))
	})
}

// SharesIn applies the In predicate on the "shares" field.
func SharesIn(vs ...*types.Uint256) predicate.PoolPass {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldShares), v...))
	})
}

// SharesNotIn applies the NotIn predicate on the "shares" field.
func SharesNotIn(vs ...*types.Uint256) predicate.PoolPass {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldShares), v...))
	})
}

// SharesGT applies the GT predicate on the "shares" field.
func SharesGT(v *types.Uint256) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldShares), v))
	})
}

// SharesGTE applies the GTE predicate on the "shares" field.
func SharesGTE(v *types.Uint256) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldShares), v))
	})
}

// SharesLT applies the LT predicate on the "shares" field.
func SharesLT(v *types.Uint256) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldShares), v))
	})
}

// SharesLTE applies the LTE predicate on the "shares" field.
func SharesLTE(v *types.Uint256) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldShares), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPool applies the HasEdge predicate on the "pool" edge.
func HasPool() predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PoolTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PoolTable, PoolColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPoolWith applies the HasEdge predicate on the "pool" edge with a given conditions (other predicates).
func HasPoolWith(preds ...predicate.Pool) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PoolInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PoolTable, PoolColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.PoolPass) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.PoolPass) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
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
func Not(p predicate.PoolPass) predicate.PoolPass {
	return predicate.PoolPass(func(s *sql.Selector) {
		p(s.Not())
	})
}
