// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ahashim/web-server/ent/passwordtoken"
	"github.com/ahashim/web-server/ent/predicate"
	"github.com/ahashim/web-server/ent/user"
)

// PasswordTokenUpdate is the builder for updating PasswordToken entities.
type PasswordTokenUpdate struct {
	config
	hooks    []Hook
	mutation *PasswordTokenMutation
}

// Where appends a list predicates to the PasswordTokenUpdate builder.
func (ptu *PasswordTokenUpdate) Where(ps ...predicate.PasswordToken) *PasswordTokenUpdate {
	ptu.mutation.Where(ps...)
	return ptu
}

// SetHash sets the "hash" field.
func (ptu *PasswordTokenUpdate) SetHash(s string) *PasswordTokenUpdate {
	ptu.mutation.SetHash(s)
	return ptu
}

// SetCreatedAt sets the "created_at" field.
func (ptu *PasswordTokenUpdate) SetCreatedAt(t time.Time) *PasswordTokenUpdate {
	ptu.mutation.SetCreatedAt(t)
	return ptu
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ptu *PasswordTokenUpdate) SetNillableCreatedAt(t *time.Time) *PasswordTokenUpdate {
	if t != nil {
		ptu.SetCreatedAt(*t)
	}
	return ptu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ptu *PasswordTokenUpdate) SetUserID(id int) *PasswordTokenUpdate {
	ptu.mutation.SetUserID(id)
	return ptu
}

// SetUser sets the "user" edge to the User entity.
func (ptu *PasswordTokenUpdate) SetUser(u *User) *PasswordTokenUpdate {
	return ptu.SetUserID(u.ID)
}

// Mutation returns the PasswordTokenMutation object of the builder.
func (ptu *PasswordTokenUpdate) Mutation() *PasswordTokenMutation {
	return ptu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ptu *PasswordTokenUpdate) ClearUser() *PasswordTokenUpdate {
	ptu.mutation.ClearUser()
	return ptu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ptu *PasswordTokenUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(ptu.hooks) == 0 {
		if err = ptu.check(); err != nil {
			return 0, err
		}
		affected, err = ptu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PasswordTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptu.check(); err != nil {
				return 0, err
			}
			ptu.mutation = mutation
			affected, err = ptu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ptu.hooks) - 1; i >= 0; i-- {
			if ptu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ptu *PasswordTokenUpdate) SaveX(ctx context.Context) int {
	affected, err := ptu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ptu *PasswordTokenUpdate) Exec(ctx context.Context) error {
	_, err := ptu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptu *PasswordTokenUpdate) ExecX(ctx context.Context) {
	if err := ptu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptu *PasswordTokenUpdate) check() error {
	if v, ok := ptu.mutation.Hash(); ok {
		if err := passwordtoken.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "PasswordToken.hash": %w`, err)}
		}
	}
	if _, ok := ptu.mutation.UserID(); ptu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PasswordToken.user"`)
	}
	return nil
}

func (ptu *PasswordTokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   passwordtoken.Table,
			Columns: passwordtoken.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: passwordtoken.FieldID,
			},
		},
	}
	if ps := ptu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptu.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: passwordtoken.FieldHash,
		})
	}
	if value, ok := ptu.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: passwordtoken.FieldCreatedAt,
		})
	}
	if ptu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   passwordtoken.UserTable,
			Columns: []string{passwordtoken.UserColumn},
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
	if nodes := ptu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   passwordtoken.UserTable,
			Columns: []string{passwordtoken.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, ptu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{passwordtoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// PasswordTokenUpdateOne is the builder for updating a single PasswordToken entity.
type PasswordTokenUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PasswordTokenMutation
}

// SetHash sets the "hash" field.
func (ptuo *PasswordTokenUpdateOne) SetHash(s string) *PasswordTokenUpdateOne {
	ptuo.mutation.SetHash(s)
	return ptuo
}

// SetCreatedAt sets the "created_at" field.
func (ptuo *PasswordTokenUpdateOne) SetCreatedAt(t time.Time) *PasswordTokenUpdateOne {
	ptuo.mutation.SetCreatedAt(t)
	return ptuo
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (ptuo *PasswordTokenUpdateOne) SetNillableCreatedAt(t *time.Time) *PasswordTokenUpdateOne {
	if t != nil {
		ptuo.SetCreatedAt(*t)
	}
	return ptuo
}

// SetUserID sets the "user" edge to the User entity by ID.
func (ptuo *PasswordTokenUpdateOne) SetUserID(id int) *PasswordTokenUpdateOne {
	ptuo.mutation.SetUserID(id)
	return ptuo
}

// SetUser sets the "user" edge to the User entity.
func (ptuo *PasswordTokenUpdateOne) SetUser(u *User) *PasswordTokenUpdateOne {
	return ptuo.SetUserID(u.ID)
}

// Mutation returns the PasswordTokenMutation object of the builder.
func (ptuo *PasswordTokenUpdateOne) Mutation() *PasswordTokenMutation {
	return ptuo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (ptuo *PasswordTokenUpdateOne) ClearUser() *PasswordTokenUpdateOne {
	ptuo.mutation.ClearUser()
	return ptuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ptuo *PasswordTokenUpdateOne) Select(field string, fields ...string) *PasswordTokenUpdateOne {
	ptuo.fields = append([]string{field}, fields...)
	return ptuo
}

// Save executes the query and returns the updated PasswordToken entity.
func (ptuo *PasswordTokenUpdateOne) Save(ctx context.Context) (*PasswordToken, error) {
	var (
		err  error
		node *PasswordToken
	)
	if len(ptuo.hooks) == 0 {
		if err = ptuo.check(); err != nil {
			return nil, err
		}
		node, err = ptuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PasswordTokenMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ptuo.check(); err != nil {
				return nil, err
			}
			ptuo.mutation = mutation
			node, err = ptuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ptuo.hooks) - 1; i >= 0; i-- {
			if ptuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ptuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ptuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ptuo *PasswordTokenUpdateOne) SaveX(ctx context.Context) *PasswordToken {
	node, err := ptuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ptuo *PasswordTokenUpdateOne) Exec(ctx context.Context) error {
	_, err := ptuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ptuo *PasswordTokenUpdateOne) ExecX(ctx context.Context) {
	if err := ptuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ptuo *PasswordTokenUpdateOne) check() error {
	if v, ok := ptuo.mutation.Hash(); ok {
		if err := passwordtoken.HashValidator(v); err != nil {
			return &ValidationError{Name: "hash", err: fmt.Errorf(`ent: validator failed for field "PasswordToken.hash": %w`, err)}
		}
	}
	if _, ok := ptuo.mutation.UserID(); ptuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PasswordToken.user"`)
	}
	return nil
}

func (ptuo *PasswordTokenUpdateOne) sqlSave(ctx context.Context) (_node *PasswordToken, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   passwordtoken.Table,
			Columns: passwordtoken.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: passwordtoken.FieldID,
			},
		},
	}
	id, ok := ptuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PasswordToken.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ptuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, passwordtoken.FieldID)
		for _, f := range fields {
			if !passwordtoken.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != passwordtoken.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ptuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ptuo.mutation.Hash(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: passwordtoken.FieldHash,
		})
	}
	if value, ok := ptuo.mutation.CreatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: passwordtoken.FieldCreatedAt,
		})
	}
	if ptuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   passwordtoken.UserTable,
			Columns: []string{passwordtoken.UserColumn},
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
	if nodes := ptuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   passwordtoken.UserTable,
			Columns: []string{passwordtoken.UserColumn},
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
	_node = &PasswordToken{config: ptuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ptuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{passwordtoken.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
