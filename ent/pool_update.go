// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ahashim/web-server/ent/pool"
	"github.com/ahashim/web-server/ent/predicate"
	"github.com/ahashim/web-server/types"
)

// PoolUpdate is the builder for updating Pool entities.
type PoolUpdate struct {
	config
	hooks    []Hook
	mutation *PoolMutation
}

// Where appends a list predicates to the PoolUpdate builder.
func (pu *PoolUpdate) Where(ps ...predicate.Pool) *PoolUpdate {
	pu.mutation.Where(ps...)
	return pu
}

// SetAmount sets the "amount" field.
func (pu *PoolUpdate) SetAmount(t *types.Uint256) *PoolUpdate {
	pu.mutation.SetAmount(t)
	return pu
}

// SetShares sets the "shares" field.
func (pu *PoolUpdate) SetShares(t *types.Uint256) *PoolUpdate {
	pu.mutation.SetShares(t)
	return pu
}

// SetBlockNumber sets the "block_number" field.
func (pu *PoolUpdate) SetBlockNumber(t *types.Uint256) *PoolUpdate {
	pu.mutation.SetBlockNumber(t)
	return pu
}

// Mutation returns the PoolMutation object of the builder.
func (pu *PoolUpdate) Mutation() *PoolMutation {
	return pu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pu *PoolUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PoolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			if pu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PoolUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PoolUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PoolUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PoolUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pool.Table,
			Columns: pool.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pool.FieldID,
			},
		},
	}
	if ps := pu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Amount(); ok {
		_spec.SetField(pool.FieldAmount, field.TypeInt, value)
	}
	if value, ok := pu.mutation.Shares(); ok {
		_spec.SetField(pool.FieldShares, field.TypeInt, value)
	}
	if value, ok := pu.mutation.BlockNumber(); ok {
		_spec.SetField(pool.FieldBlockNumber, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pool.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PoolUpdateOne is the builder for updating a single Pool entity.
type PoolUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PoolMutation
}

// SetAmount sets the "amount" field.
func (puo *PoolUpdateOne) SetAmount(t *types.Uint256) *PoolUpdateOne {
	puo.mutation.SetAmount(t)
	return puo
}

// SetShares sets the "shares" field.
func (puo *PoolUpdateOne) SetShares(t *types.Uint256) *PoolUpdateOne {
	puo.mutation.SetShares(t)
	return puo
}

// SetBlockNumber sets the "block_number" field.
func (puo *PoolUpdateOne) SetBlockNumber(t *types.Uint256) *PoolUpdateOne {
	puo.mutation.SetBlockNumber(t)
	return puo
}

// Mutation returns the PoolMutation object of the builder.
func (puo *PoolUpdateOne) Mutation() *PoolMutation {
	return puo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (puo *PoolUpdateOne) Select(field string, fields ...string) *PoolUpdateOne {
	puo.fields = append([]string{field}, fields...)
	return puo
}

// Save executes the query and returns the updated Pool entity.
func (puo *PoolUpdateOne) Save(ctx context.Context) (*Pool, error) {
	var (
		err  error
		node *Pool
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PoolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			if puo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = puo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, puo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Pool)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PoolMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PoolUpdateOne) SaveX(ctx context.Context) *Pool {
	node, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (puo *PoolUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PoolUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PoolUpdateOne) sqlSave(ctx context.Context) (_node *Pool, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pool.Table,
			Columns: pool.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pool.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Pool.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := puo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pool.FieldID)
		for _, f := range fields {
			if !pool.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != pool.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := puo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := puo.mutation.Amount(); ok {
		_spec.SetField(pool.FieldAmount, field.TypeInt, value)
	}
	if value, ok := puo.mutation.Shares(); ok {
		_spec.SetField(pool.FieldShares, field.TypeInt, value)
	}
	if value, ok := puo.mutation.BlockNumber(); ok {
		_spec.SetField(pool.FieldBlockNumber, field.TypeInt, value)
	}
	_node = &Pool{config: puo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{pool.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
