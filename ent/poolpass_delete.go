// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ahashim/web-server/ent/poolpass"
	"github.com/ahashim/web-server/ent/predicate"
)

// PoolPassDelete is the builder for deleting a PoolPass entity.
type PoolPassDelete struct {
	config
	hooks    []Hook
	mutation *PoolPassMutation
}

// Where appends a list predicates to the PoolPassDelete builder.
func (ppd *PoolPassDelete) Where(ps ...predicate.PoolPass) *PoolPassDelete {
	ppd.mutation.Where(ps...)
	return ppd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ppd *PoolPassDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ppd.hooks) == 0 {
		affected, err = ppd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PoolPassMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ppd.mutation = mutation
			affected, err = ppd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ppd.hooks) - 1; i >= 0; i-- {
			if ppd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ppd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ppd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppd *PoolPassDelete) ExecX(ctx context.Context) int {
	n, err := ppd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ppd *PoolPassDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: poolpass.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: poolpass.FieldID,
			},
		},
	}
	if ps := ppd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ppd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// PoolPassDeleteOne is the builder for deleting a single PoolPass entity.
type PoolPassDeleteOne struct {
	ppd *PoolPassDelete
}

// Exec executes the deletion query.
func (ppdo *PoolPassDeleteOne) Exec(ctx context.Context) error {
	n, err := ppdo.ppd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{poolpass.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ppdo *PoolPassDeleteOne) ExecX(ctx context.Context) {
	ppdo.ppd.ExecX(ctx)
}