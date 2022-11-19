// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ahashim/web-server/ent/pool"
	"github.com/ahashim/web-server/ent/poolpass"
	"github.com/ahashim/web-server/ent/user"
	"github.com/ahashim/web-server/types"
)

// PoolPassCreate is the builder for creating a PoolPass entity.
type PoolPassCreate struct {
	config
	mutation *PoolPassMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (ppc *PoolPassCreate) SetCreateTime(t time.Time) *PoolPassCreate {
	ppc.mutation.SetCreateTime(t)
	return ppc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ppc *PoolPassCreate) SetNillableCreateTime(t *time.Time) *PoolPassCreate {
	if t != nil {
		ppc.SetCreateTime(*t)
	}
	return ppc
}

// SetShares sets the "shares" field.
func (ppc *PoolPassCreate) SetShares(t *types.Uint256) *PoolPassCreate {
	ppc.mutation.SetShares(t)
	return ppc
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ppc *PoolPassCreate) SetUserID(id int) *PoolPassCreate {
	ppc.mutation.SetUserID(id)
	return ppc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (ppc *PoolPassCreate) SetNillableUserID(id *int) *PoolPassCreate {
	if id != nil {
		ppc = ppc.SetUserID(*id)
	}
	return ppc
}

// SetUser sets the "user" edge to the User entity.
func (ppc *PoolPassCreate) SetUser(u *User) *PoolPassCreate {
	return ppc.SetUserID(u.ID)
}

// SetPoolID sets the "pool" edge to the Pool entity by ID.
func (ppc *PoolPassCreate) SetPoolID(id int) *PoolPassCreate {
	ppc.mutation.SetPoolID(id)
	return ppc
}

// SetNillablePoolID sets the "pool" edge to the Pool entity by ID if the given value is not nil.
func (ppc *PoolPassCreate) SetNillablePoolID(id *int) *PoolPassCreate {
	if id != nil {
		ppc = ppc.SetPoolID(*id)
	}
	return ppc
}

// SetPool sets the "pool" edge to the Pool entity.
func (ppc *PoolPassCreate) SetPool(p *Pool) *PoolPassCreate {
	return ppc.SetPoolID(p.ID)
}

// Mutation returns the PoolPassMutation object of the builder.
func (ppc *PoolPassCreate) Mutation() *PoolPassMutation {
	return ppc.mutation
}

// Save creates the PoolPass in the database.
func (ppc *PoolPassCreate) Save(ctx context.Context) (*PoolPass, error) {
	var (
		err  error
		node *PoolPass
	)
	ppc.defaults()
	if len(ppc.hooks) == 0 {
		if err = ppc.check(); err != nil {
			return nil, err
		}
		node, err = ppc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PoolPassMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ppc.check(); err != nil {
				return nil, err
			}
			ppc.mutation = mutation
			if node, err = ppc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ppc.hooks) - 1; i >= 0; i-- {
			if ppc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ppc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, ppc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PoolPass)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PoolPassMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ppc *PoolPassCreate) SaveX(ctx context.Context) *PoolPass {
	v, err := ppc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ppc *PoolPassCreate) Exec(ctx context.Context) error {
	_, err := ppc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppc *PoolPassCreate) ExecX(ctx context.Context) {
	if err := ppc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ppc *PoolPassCreate) defaults() {
	if _, ok := ppc.mutation.CreateTime(); !ok {
		v := poolpass.DefaultCreateTime()
		ppc.mutation.SetCreateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ppc *PoolPassCreate) check() error {
	if _, ok := ppc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "PoolPass.create_time"`)}
	}
	if _, ok := ppc.mutation.Shares(); !ok {
		return &ValidationError{Name: "shares", err: errors.New(`ent: missing required field "PoolPass.shares"`)}
	}
	return nil
}

func (ppc *PoolPassCreate) sqlSave(ctx context.Context) (*PoolPass, error) {
	_node, _spec := ppc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ppc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ppc *PoolPassCreate) createSpec() (*PoolPass, *sqlgraph.CreateSpec) {
	var (
		_node = &PoolPass{config: ppc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: poolpass.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: poolpass.FieldID,
			},
		}
	)
	if value, ok := ppc.mutation.CreateTime(); ok {
		_spec.SetField(poolpass.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = value
	}
	if value, ok := ppc.mutation.Shares(); ok {
		_spec.SetField(poolpass.FieldShares, field.TypeInt, value)
		_node.Shares = value
	}
	if nodes := ppc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   poolpass.UserTable,
			Columns: []string{poolpass.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_pool_passes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := ppc.mutation.PoolIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   poolpass.PoolTable,
			Columns: []string{poolpass.PoolColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pool.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.pool_pool_passes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PoolPassCreateBulk is the builder for creating many PoolPass entities in bulk.
type PoolPassCreateBulk struct {
	config
	builders []*PoolPassCreate
}

// Save creates the PoolPass entities in the database.
func (ppcb *PoolPassCreateBulk) Save(ctx context.Context) ([]*PoolPass, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ppcb.builders))
	nodes := make([]*PoolPass, len(ppcb.builders))
	mutators := make([]Mutator, len(ppcb.builders))
	for i := range ppcb.builders {
		func(i int, root context.Context) {
			builder := ppcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PoolPassMutation)
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
					_, err = mutators[i+1].Mutate(root, ppcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ppcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ppcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ppcb *PoolPassCreateBulk) SaveX(ctx context.Context) []*PoolPass {
	v, err := ppcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ppcb *PoolPassCreateBulk) Exec(ctx context.Context) error {
	_, err := ppcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ppcb *PoolPassCreateBulk) ExecX(ctx context.Context) {
	if err := ppcb.Exec(ctx); err != nil {
		panic(err)
	}
}
