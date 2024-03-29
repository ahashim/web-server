// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ahashim/web-server/ent/interaction"
	"github.com/ahashim/web-server/ent/pool"
	"github.com/ahashim/web-server/ent/predicate"
	"github.com/ahashim/web-server/ent/squeak"
	"github.com/ahashim/web-server/ent/user"
	"github.com/ahashim/web-server/types"
)

// SqueakUpdate is the builder for updating Squeak entities.
type SqueakUpdate struct {
	config
	hooks    []Hook
	mutation *SqueakMutation
}

// Where appends a list predicates to the SqueakUpdate builder.
func (su *SqueakUpdate) Where(ps ...predicate.Squeak) *SqueakUpdate {
	su.mutation.Where(ps...)
	return su
}

// SetBlockNumber sets the "block_number" field.
func (su *SqueakUpdate) SetBlockNumber(t *types.Uint256) *SqueakUpdate {
	su.mutation.SetBlockNumber(t)
	return su
}

// AddInteractionIDs adds the "interactions" edge to the Interaction entity by IDs.
func (su *SqueakUpdate) AddInteractionIDs(ids ...int) *SqueakUpdate {
	su.mutation.AddInteractionIDs(ids...)
	return su
}

// AddInteractions adds the "interactions" edges to the Interaction entity.
func (su *SqueakUpdate) AddInteractions(i ...*Interaction) *SqueakUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return su.AddInteractionIDs(ids...)
}

// SetPoolID sets the "pool" edge to the Pool entity by ID.
func (su *SqueakUpdate) SetPoolID(id int) *SqueakUpdate {
	su.mutation.SetPoolID(id)
	return su
}

// SetNillablePoolID sets the "pool" edge to the Pool entity by ID if the given value is not nil.
func (su *SqueakUpdate) SetNillablePoolID(id *int) *SqueakUpdate {
	if id != nil {
		su = su.SetPoolID(*id)
	}
	return su
}

// SetPool sets the "pool" edge to the Pool entity.
func (su *SqueakUpdate) SetPool(p *Pool) *SqueakUpdate {
	return su.SetPoolID(p.ID)
}

// SetCreatorID sets the "creator" edge to the User entity by ID.
func (su *SqueakUpdate) SetCreatorID(id int) *SqueakUpdate {
	su.mutation.SetCreatorID(id)
	return su
}

// SetNillableCreatorID sets the "creator" edge to the User entity by ID if the given value is not nil.
func (su *SqueakUpdate) SetNillableCreatorID(id *int) *SqueakUpdate {
	if id != nil {
		su = su.SetCreatorID(*id)
	}
	return su
}

// SetCreator sets the "creator" edge to the User entity.
func (su *SqueakUpdate) SetCreator(u *User) *SqueakUpdate {
	return su.SetCreatorID(u.ID)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (su *SqueakUpdate) SetOwnerID(id int) *SqueakUpdate {
	su.mutation.SetOwnerID(id)
	return su
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (su *SqueakUpdate) SetNillableOwnerID(id *int) *SqueakUpdate {
	if id != nil {
		su = su.SetOwnerID(*id)
	}
	return su
}

// SetOwner sets the "owner" edge to the User entity.
func (su *SqueakUpdate) SetOwner(u *User) *SqueakUpdate {
	return su.SetOwnerID(u.ID)
}

// Mutation returns the SqueakMutation object of the builder.
func (su *SqueakUpdate) Mutation() *SqueakMutation {
	return su.mutation
}

// ClearInteractions clears all "interactions" edges to the Interaction entity.
func (su *SqueakUpdate) ClearInteractions() *SqueakUpdate {
	su.mutation.ClearInteractions()
	return su
}

// RemoveInteractionIDs removes the "interactions" edge to Interaction entities by IDs.
func (su *SqueakUpdate) RemoveInteractionIDs(ids ...int) *SqueakUpdate {
	su.mutation.RemoveInteractionIDs(ids...)
	return su
}

// RemoveInteractions removes "interactions" edges to Interaction entities.
func (su *SqueakUpdate) RemoveInteractions(i ...*Interaction) *SqueakUpdate {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return su.RemoveInteractionIDs(ids...)
}

// ClearPool clears the "pool" edge to the Pool entity.
func (su *SqueakUpdate) ClearPool() *SqueakUpdate {
	su.mutation.ClearPool()
	return su
}

// ClearCreator clears the "creator" edge to the User entity.
func (su *SqueakUpdate) ClearCreator() *SqueakUpdate {
	su.mutation.ClearCreator()
	return su
}

// ClearOwner clears the "owner" edge to the User entity.
func (su *SqueakUpdate) ClearOwner() *SqueakUpdate {
	su.mutation.ClearOwner()
	return su
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (su *SqueakUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(su.hooks) == 0 {
		affected, err = su.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SqueakMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			su.mutation = mutation
			affected, err = su.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(su.hooks) - 1; i >= 0; i-- {
			if su.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = su.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, su.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (su *SqueakUpdate) SaveX(ctx context.Context) int {
	affected, err := su.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (su *SqueakUpdate) Exec(ctx context.Context) error {
	_, err := su.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (su *SqueakUpdate) ExecX(ctx context.Context) {
	if err := su.Exec(ctx); err != nil {
		panic(err)
	}
}

func (su *SqueakUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   squeak.Table,
			Columns: squeak.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: squeak.FieldID,
			},
		},
	}
	if ps := su.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := su.mutation.BlockNumber(); ok {
		_spec.SetField(squeak.FieldBlockNumber, field.TypeInt, value)
	}
	if su.mutation.InteractionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.RemovedInteractionsIDs(); len(nodes) > 0 && !su.mutation.InteractionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.InteractionsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.PoolCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.PoolIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.CreatorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.CreatorIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if su.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := su.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, su.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{squeak.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// SqueakUpdateOne is the builder for updating a single Squeak entity.
type SqueakUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *SqueakMutation
}

// SetBlockNumber sets the "block_number" field.
func (suo *SqueakUpdateOne) SetBlockNumber(t *types.Uint256) *SqueakUpdateOne {
	suo.mutation.SetBlockNumber(t)
	return suo
}

// AddInteractionIDs adds the "interactions" edge to the Interaction entity by IDs.
func (suo *SqueakUpdateOne) AddInteractionIDs(ids ...int) *SqueakUpdateOne {
	suo.mutation.AddInteractionIDs(ids...)
	return suo
}

// AddInteractions adds the "interactions" edges to the Interaction entity.
func (suo *SqueakUpdateOne) AddInteractions(i ...*Interaction) *SqueakUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return suo.AddInteractionIDs(ids...)
}

// SetPoolID sets the "pool" edge to the Pool entity by ID.
func (suo *SqueakUpdateOne) SetPoolID(id int) *SqueakUpdateOne {
	suo.mutation.SetPoolID(id)
	return suo
}

// SetNillablePoolID sets the "pool" edge to the Pool entity by ID if the given value is not nil.
func (suo *SqueakUpdateOne) SetNillablePoolID(id *int) *SqueakUpdateOne {
	if id != nil {
		suo = suo.SetPoolID(*id)
	}
	return suo
}

// SetPool sets the "pool" edge to the Pool entity.
func (suo *SqueakUpdateOne) SetPool(p *Pool) *SqueakUpdateOne {
	return suo.SetPoolID(p.ID)
}

// SetCreatorID sets the "creator" edge to the User entity by ID.
func (suo *SqueakUpdateOne) SetCreatorID(id int) *SqueakUpdateOne {
	suo.mutation.SetCreatorID(id)
	return suo
}

// SetNillableCreatorID sets the "creator" edge to the User entity by ID if the given value is not nil.
func (suo *SqueakUpdateOne) SetNillableCreatorID(id *int) *SqueakUpdateOne {
	if id != nil {
		suo = suo.SetCreatorID(*id)
	}
	return suo
}

// SetCreator sets the "creator" edge to the User entity.
func (suo *SqueakUpdateOne) SetCreator(u *User) *SqueakUpdateOne {
	return suo.SetCreatorID(u.ID)
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (suo *SqueakUpdateOne) SetOwnerID(id int) *SqueakUpdateOne {
	suo.mutation.SetOwnerID(id)
	return suo
}

// SetNillableOwnerID sets the "owner" edge to the User entity by ID if the given value is not nil.
func (suo *SqueakUpdateOne) SetNillableOwnerID(id *int) *SqueakUpdateOne {
	if id != nil {
		suo = suo.SetOwnerID(*id)
	}
	return suo
}

// SetOwner sets the "owner" edge to the User entity.
func (suo *SqueakUpdateOne) SetOwner(u *User) *SqueakUpdateOne {
	return suo.SetOwnerID(u.ID)
}

// Mutation returns the SqueakMutation object of the builder.
func (suo *SqueakUpdateOne) Mutation() *SqueakMutation {
	return suo.mutation
}

// ClearInteractions clears all "interactions" edges to the Interaction entity.
func (suo *SqueakUpdateOne) ClearInteractions() *SqueakUpdateOne {
	suo.mutation.ClearInteractions()
	return suo
}

// RemoveInteractionIDs removes the "interactions" edge to Interaction entities by IDs.
func (suo *SqueakUpdateOne) RemoveInteractionIDs(ids ...int) *SqueakUpdateOne {
	suo.mutation.RemoveInteractionIDs(ids...)
	return suo
}

// RemoveInteractions removes "interactions" edges to Interaction entities.
func (suo *SqueakUpdateOne) RemoveInteractions(i ...*Interaction) *SqueakUpdateOne {
	ids := make([]int, len(i))
	for j := range i {
		ids[j] = i[j].ID
	}
	return suo.RemoveInteractionIDs(ids...)
}

// ClearPool clears the "pool" edge to the Pool entity.
func (suo *SqueakUpdateOne) ClearPool() *SqueakUpdateOne {
	suo.mutation.ClearPool()
	return suo
}

// ClearCreator clears the "creator" edge to the User entity.
func (suo *SqueakUpdateOne) ClearCreator() *SqueakUpdateOne {
	suo.mutation.ClearCreator()
	return suo
}

// ClearOwner clears the "owner" edge to the User entity.
func (suo *SqueakUpdateOne) ClearOwner() *SqueakUpdateOne {
	suo.mutation.ClearOwner()
	return suo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (suo *SqueakUpdateOne) Select(field string, fields ...string) *SqueakUpdateOne {
	suo.fields = append([]string{field}, fields...)
	return suo
}

// Save executes the query and returns the updated Squeak entity.
func (suo *SqueakUpdateOne) Save(ctx context.Context) (*Squeak, error) {
	var (
		err  error
		node *Squeak
	)
	if len(suo.hooks) == 0 {
		node, err = suo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*SqueakMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			suo.mutation = mutation
			node, err = suo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(suo.hooks) - 1; i >= 0; i-- {
			if suo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = suo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, suo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (suo *SqueakUpdateOne) SaveX(ctx context.Context) *Squeak {
	node, err := suo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (suo *SqueakUpdateOne) Exec(ctx context.Context) error {
	_, err := suo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (suo *SqueakUpdateOne) ExecX(ctx context.Context) {
	if err := suo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (suo *SqueakUpdateOne) sqlSave(ctx context.Context) (_node *Squeak, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   squeak.Table,
			Columns: squeak.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: squeak.FieldID,
			},
		},
	}
	id, ok := suo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Squeak.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := suo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, squeak.FieldID)
		for _, f := range fields {
			if !squeak.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != squeak.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := suo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := suo.mutation.BlockNumber(); ok {
		_spec.SetField(squeak.FieldBlockNumber, field.TypeInt, value)
	}
	if suo.mutation.InteractionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.RemovedInteractionsIDs(); len(nodes) > 0 && !suo.mutation.InteractionsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.InteractionsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.PoolCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.PoolIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.CreatorCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.CreatorIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if suo.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := suo.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Squeak{config: suo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, suo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{squeak.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
