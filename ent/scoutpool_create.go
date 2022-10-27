// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ahashim/web-server/ent/scoutpool"
	"github.com/ahashim/web-server/types"
)

// ScoutPoolCreate is the builder for creating a ScoutPool entity.
type ScoutPoolCreate struct {
	config
	mutation *ScoutPoolMutation
	hooks    []Hook
}

// SetAmount sets the "amount" field.
func (spc *ScoutPoolCreate) SetAmount(t *types.Uint256) *ScoutPoolCreate {
	spc.mutation.SetAmount(t)
	return spc
}

// Mutation returns the ScoutPoolMutation object of the builder.
func (spc *ScoutPoolCreate) Mutation() *ScoutPoolMutation {
	return spc.mutation
}

// Save creates the ScoutPool in the database.
func (spc *ScoutPoolCreate) Save(ctx context.Context) (*ScoutPool, error) {
	var (
		err  error
		node *ScoutPool
	)
	if len(spc.hooks) == 0 {
		if err = spc.check(); err != nil {
			return nil, err
		}
		node, err = spc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ScoutPoolMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = spc.check(); err != nil {
				return nil, err
			}
			spc.mutation = mutation
			if node, err = spc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(spc.hooks) - 1; i >= 0; i-- {
			if spc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = spc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, spc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*ScoutPool)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ScoutPoolMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (spc *ScoutPoolCreate) SaveX(ctx context.Context) *ScoutPool {
	v, err := spc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spc *ScoutPoolCreate) Exec(ctx context.Context) error {
	_, err := spc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spc *ScoutPoolCreate) ExecX(ctx context.Context) {
	if err := spc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (spc *ScoutPoolCreate) check() error {
	if _, ok := spc.mutation.Amount(); !ok {
		return &ValidationError{Name: "amount", err: errors.New(`ent: missing required field "ScoutPool.amount"`)}
	}
	return nil
}

func (spc *ScoutPoolCreate) sqlSave(ctx context.Context) (*ScoutPool, error) {
	_node, _spec := spc.createSpec()
	if err := sqlgraph.CreateNode(ctx, spc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (spc *ScoutPoolCreate) createSpec() (*ScoutPool, *sqlgraph.CreateSpec) {
	var (
		_node = &ScoutPool{config: spc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: scoutpool.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: scoutpool.FieldID,
			},
		}
	)
	if value, ok := spc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: scoutpool.FieldAmount,
		})
		_node.Amount = value
	}
	return _node, _spec
}

// ScoutPoolCreateBulk is the builder for creating many ScoutPool entities in bulk.
type ScoutPoolCreateBulk struct {
	config
	builders []*ScoutPoolCreate
}

// Save creates the ScoutPool entities in the database.
func (spcb *ScoutPoolCreateBulk) Save(ctx context.Context) ([]*ScoutPool, error) {
	specs := make([]*sqlgraph.CreateSpec, len(spcb.builders))
	nodes := make([]*ScoutPool, len(spcb.builders))
	mutators := make([]Mutator, len(spcb.builders))
	for i := range spcb.builders {
		func(i int, root context.Context) {
			builder := spcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ScoutPoolMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, spcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, spcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, spcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (spcb *ScoutPoolCreateBulk) SaveX(ctx context.Context) []*ScoutPool {
	v, err := spcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (spcb *ScoutPoolCreateBulk) Exec(ctx context.Context) error {
	_, err := spcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (spcb *ScoutPoolCreateBulk) ExecX(ctx context.Context) {
	if err := spcb.Exec(ctx); err != nil {
		panic(err)
	}
}
