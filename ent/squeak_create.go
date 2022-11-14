// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ahashim/web-server/ent/interaction"
	"github.com/ahashim/web-server/ent/pool"
	"github.com/ahashim/web-server/ent/squeak"
	"github.com/ahashim/web-server/ent/user"
	"github.com/ahashim/web-server/types"
)

// SqueakCreate is the builder for creating a Squeak entity.
type SqueakCreate struct {
	config
	mutation *SqueakMutation
	hooks    []Hook
}

// SetBlockNumber sets the "block_number" field.
func (sc *SqueakCreate) SetBlockNumber(t *types.Uint256) *SqueakCreate {
	sc.mutation.SetBlockNumber(t)
	return sc
}

// SetContent sets the "content" field.
func (sc *SqueakCreate) SetContent(s string) *SqueakCreate {
	sc.mutation.SetContent(s)
	return sc
}

// AddInteractionIDs adds the "interactions" edge to the Interaction entity by IDs.
func (sc *SqueakCreate) AddInteractionIDs(ids ...int) *SqueakCreate {
	sc.mutation.AddInteractionIDs(ids...)
	return sc
}

// AddInteractions adds the "interactions" edges to the Interaction entity.
func (sc *SqueakCreate) AddInteractions(i ...*Interaction) *SqueakCreate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return sc.AddInteractionIDs(ids...)
}

// SetPoolID sets the "pool" edge to the Pool entity by ID.
func (sc *SqueakCreate) SetPoolID(id int) *SqueakCreate {
	sc.mutation.SetPoolID(id)
	return sc
}

// SetNillablePoolID sets the "pool" edge to the Pool entity by ID if the given value is not nil.
func (sc *SqueakCreate) SetNillablePoolID(id *int) *SqueakCreate {
	if id != nil {
		sc = sc.SetPoolID(*id)
	}
	return sc
}

// SetPool sets the "pool" edge to the Pool entity.
func (sc *SqueakCreate) SetPool(p *Pool) *SqueakCreate {
	return sc.SetPoolID(p.ID)
}

// SetCreatorID sets the "creator" edge to the User entity by ID.
func (sc *SqueakCreate) SetCreatorID(id int) *SqueakCreate {
	sc.mutation.SetCreatorID(id)
	return sc
}

// SetNillableCreatorID sets the "creator" edge to the User entity by ID if the given value is not nil.
func (sc *SqueakCreate) SetNillableCreatorID(id *int) *SqueakCreate {
	if id != nil {
		sc = sc.SetCreatorID(*id)
	}
	return sc
}

// SetCreator sets the "creator" edge to the User entity.
func (sc *SqueakCreate) SetCreator(u *User) *SqueakCreate {
	return sc.SetCreatorID(u.ID)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (sc *SqueakCreate) SetOwnerID(id int) *SqueakCreate {
	sc.mutation.SetOwnerID(id)
	return sc
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (sc *SqueakCreate) SetNillableOwnerID(id *int) *SqueakCreate {
	if id != nil {
		sc = sc.SetOwnerID(*id)
	}
	return sc
}

// SetOwner sets the "owner" edge to the User entity.
func (sc *SqueakCreate) SetOwner(u *User) *SqueakCreate {
	return sc.SetOwnerID(u.ID)
}

// Mutation returns the SqueakMutation object of the builder.
func (sc *SqueakCreate) Mutation() *SqueakMutation {
	return sc.mutation
}

// Save creates the Squeak in the database.
func (sc *SqueakCreate) Save(ctx context.Context) (*Squeak, error) {
	var (
		err  error
		node *Squeak
	)
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SqueakMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, sc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Squeak)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from SqueakMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SqueakCreate) SaveX(ctx context.Context) *Squeak {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SqueakCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SqueakCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SqueakCreate) check() error {
	if _, ok := sc.mutation.BlockNumber(); !ok {
		return &ValidationError{Name: "block_number", err: errors.New(`ent: missing required field "Squeak.block_number"`)}
	}
	if _, ok := sc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Squeak.content"`)}
	}
	if v, ok := sc.mutation.Content(); ok {
		if err := squeak.ContentValidator(v); err != nil {
			return &ValidationError{Name: "content", err: fmt.Errorf(`ent: validator failed for field "Squeak.content": %w`, err)}
		}
	}
	return nil
}

func (sc *SqueakCreate) sqlSave(ctx context.Context) (*Squeak, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *SqueakCreate) createSpec() (*Squeak, *sqlgraph.CreateSpec) {
	var (
		_node = &Squeak{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: squeak.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: squeak.FieldID,
			},
		}
	)
	if value, ok := sc.mutation.BlockNumber(); ok {
		_spec.SetField(squeak.FieldBlockNumber, field.TypeInt, value)
		_node.BlockNumber = value
	}
	if value, ok := sc.mutation.Content(); ok {
		_spec.SetField(squeak.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if nodes := sc.mutation.InteractionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   squeak.InteractionsTable,
			Columns: []string{squeak.InteractionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: interaction.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.PoolIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   squeak.PoolTable,
			Columns: []string{squeak.PoolColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.CreatorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   squeak.CreatorTable,
			Columns: []string{squeak.CreatorColumn},
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
		_node.user_created = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   squeak.OwnerTable,
			Columns: []string{squeak.OwnerColumn},
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
		_node.user_owned = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// SqueakCreateBulk is the builder for creating many Squeak entities in bulk.
type SqueakCreateBulk struct {
	config
	builders []*SqueakCreate
}

// Save creates the Squeak entities in the database.
func (scb *SqueakCreateBulk) Save(ctx context.Context) ([]*Squeak, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Squeak, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SqueakMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SqueakCreateBulk) SaveX(ctx context.Context) []*Squeak {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SqueakCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SqueakCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
