// Code generated by ent, DO NOT EDIT.

package squeak

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/ahashim/web-server/ent/predicate"
	"github.com/ahashim/web-server/types"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		v := make([]any, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// BlockNumber applies equality check predicate on the "block_number" field. It's identical to BlockNumberEQ.
func BlockNumber(v *types.Uint256) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBlockNumber), v))
	})
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// BlockNumberEQ applies the EQ predicate on the "block_number" field.
func BlockNumberEQ(v *types.Uint256) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBlockNumber), v))
	})
}

// BlockNumberNEQ applies the NEQ predicate on the "block_number" field.
func BlockNumberNEQ(v *types.Uint256) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBlockNumber), v))
	})
}

// BlockNumberIn applies the In predicate on the "block_number" field.
func BlockNumberIn(vs ...*types.Uint256) predicate.Squeak {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldBlockNumber), v...))
	})
}

// BlockNumberNotIn applies the NotIn predicate on the "block_number" field.
func BlockNumberNotIn(vs ...*types.Uint256) predicate.Squeak {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldBlockNumber), v...))
	})
}

// BlockNumberGT applies the GT predicate on the "block_number" field.
func BlockNumberGT(v *types.Uint256) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBlockNumber), v))
	})
}

// BlockNumberGTE applies the GTE predicate on the "block_number" field.
func BlockNumberGTE(v *types.Uint256) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBlockNumber), v))
	})
}

// BlockNumberLT applies the LT predicate on the "block_number" field.
func BlockNumberLT(v *types.Uint256) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBlockNumber), v))
	})
}

// BlockNumberLTE applies the LTE predicate on the "block_number" field.
func BlockNumberLTE(v *types.Uint256) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBlockNumber), v))
	})
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldContent), v))
	})
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldContent), v))
	})
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Squeak {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldContent), v...))
	})
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Squeak {
	v := make([]any, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldContent), v...))
	})
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldContent), v))
	})
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldContent), v))
	})
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldContent), v))
	})
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldContent), v))
	})
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldContent), v))
	})
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldContent), v))
	})
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldContent), v))
	})
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldContent), v))
	})
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldContent), v))
	})
}

// HasInteractions applies the HasEdge predicate on the "interactions" edge.
func HasInteractions() predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InteractionsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InteractionsTable, InteractionsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInteractionsWith applies the HasEdge predicate on the "interactions" edge with a given conditions (other predicates).
func HasInteractionsWith(preds ...predicate.Interaction) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(InteractionsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, InteractionsTable, InteractionsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPool applies the HasEdge predicate on the "pool" edge.
func HasPool() predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PoolTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, PoolTable, PoolColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPoolWith applies the HasEdge predicate on the "pool" edge with a given conditions (other predicates).
func HasPoolWith(preds ...predicate.Pool) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PoolInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, PoolTable, PoolColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCreator applies the HasEdge predicate on the "creator" edge.
func HasCreator() predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CreatorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatorTable, CreatorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCreatorWith applies the HasEdge predicate on the "creator" edge with a given conditions (other predicates).
func HasCreatorWith(preds ...predicate.User) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(CreatorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CreatorTable, CreatorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(OwnerInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Squeak) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Squeak) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
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
func Not(p predicate.Squeak) predicate.Squeak {
	return predicate.Squeak(func(s *sql.Selector) {
		p(s.Not())
	})
}
