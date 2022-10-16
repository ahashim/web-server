// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ahashim/web-server/ent/interaction"
	"github.com/ahashim/web-server/ent/predicate"
)

// InteractionDelete is the builder for deleting a Interaction entity.
type InteractionDelete struct {
	config
	hooks    []Hook
	mutation *InteractionMutation
}

// Where appends a list predicates to the InteractionDelete builder.
func (id *InteractionDelete) Where(ps ...predicate.Interaction) *InteractionDelete {
	id.mutation.Where(ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *InteractionDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(id.hooks) == 0 {
		affected, err = id.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*InteractionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			id.mutation = mutation
			affected, err = id.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(id.hooks) - 1; i >= 0; i-- {
			if id.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = id.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, id.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (id *InteractionDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (id *InteractionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: interaction.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: interaction.FieldID,
			},
		},
	}
	if ps := id.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, id.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// InteractionDeleteOne is the builder for deleting a single Interaction entity.
type InteractionDeleteOne struct {
	id *InteractionDelete
}

// Exec executes the deletion query.
func (ido *InteractionDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{interaction.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *InteractionDeleteOne) ExecX(ctx context.Context) {
	ido.id.ExecX(ctx)
}