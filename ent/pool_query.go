// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ahashim/web-server/ent/pool"
	"github.com/ahashim/web-server/ent/poolpass"
	"github.com/ahashim/web-server/ent/predicate"
	"github.com/ahashim/web-server/ent/squeak"
)

// PoolQuery is the builder for querying Pool entities.
type PoolQuery struct {
	config
	limit          *int
	offset         *int
	unique         *bool
	order          []OrderFunc
	fields         []string
	predicates     []predicate.Pool
	withPoolPasses *PoolPassQuery
	withSqueak     *SqueakQuery
	withFKs        bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PoolQuery builder.
func (pq *PoolQuery) Where(ps ...predicate.Pool) *PoolQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit adds a limit step to the query.
func (pq *PoolQuery) Limit(limit int) *PoolQuery {
	pq.limit = &limit
	return pq
}

// Offset adds an offset step to the query.
func (pq *PoolQuery) Offset(offset int) *PoolQuery {
	pq.offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PoolQuery) Unique(unique bool) *PoolQuery {
	pq.unique = &unique
	return pq
}

// Order adds an order step to the query.
func (pq *PoolQuery) Order(o ...OrderFunc) *PoolQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryPoolPasses chains the current query on the "pool_passes" edge.
func (pq *PoolQuery) QueryPoolPasses() *PoolPassQuery {
	query := &PoolPassQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pool.Table, pool.FieldID, selector),
			sqlgraph.To(poolpass.Table, poolpass.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, pool.PoolPassesTable, pool.PoolPassesColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QuerySqueak chains the current query on the "squeak" edge.
func (pq *PoolQuery) QuerySqueak() *SqueakQuery {
	query := &SqueakQuery{config: pq.config}
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(pool.Table, pool.FieldID, selector),
			sqlgraph.To(squeak.Table, squeak.FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, pool.SqueakTable, pool.SqueakColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Pool entity from the query.
// Returns a *NotFoundError when no Pool was found.
func (pq *PoolQuery) First(ctx context.Context) (*Pool, error) {
	nodes, err := pq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{pool.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PoolQuery) FirstX(ctx context.Context) *Pool {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Pool ID from the query.
// Returns a *NotFoundError when no Pool ID was found.
func (pq *PoolQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{pool.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PoolQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Pool entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Pool entity is found.
// Returns a *NotFoundError when no Pool entities are found.
func (pq *PoolQuery) Only(ctx context.Context) (*Pool, error) {
	nodes, err := pq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{pool.Label}
	default:
		return nil, &NotSingularError{pool.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PoolQuery) OnlyX(ctx context.Context) *Pool {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Pool ID in the query.
// Returns a *NotSingularError when more than one Pool ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PoolQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{pool.Label}
	default:
		err = &NotSingularError{pool.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PoolQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Pools.
func (pq *PoolQuery) All(ctx context.Context) ([]*Pool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return pq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (pq *PoolQuery) AllX(ctx context.Context) []*Pool {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Pool IDs.
func (pq *PoolQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := pq.Select(pool.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PoolQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PoolQuery) Count(ctx context.Context) (int, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return pq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PoolQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PoolQuery) Exist(ctx context.Context) (bool, error) {
	if err := pq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return pq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PoolQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PoolQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PoolQuery) Clone() *PoolQuery {
	if pq == nil {
		return nil
	}
	return &PoolQuery{
		config:         pq.config,
		limit:          pq.limit,
		offset:         pq.offset,
		order:          append([]OrderFunc{}, pq.order...),
		predicates:     append([]predicate.Pool{}, pq.predicates...),
		withPoolPasses: pq.withPoolPasses.Clone(),
		withSqueak:     pq.withSqueak.Clone(),
		// clone intermediate query.
		sql:    pq.sql.Clone(),
		path:   pq.path,
		unique: pq.unique,
	}
}

// WithPoolPasses tells the query-builder to eager-load the nodes that are connected to
// the "pool_passes" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PoolQuery) WithPoolPasses(opts ...func(*PoolPassQuery)) *PoolQuery {
	query := &PoolPassQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withPoolPasses = query
	return pq
}

// WithSqueak tells the query-builder to eager-load the nodes that are connected to
// the "squeak" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PoolQuery) WithSqueak(opts ...func(*SqueakQuery)) *PoolQuery {
	query := &SqueakQuery{config: pq.config}
	for _, opt := range opts {
		opt(query)
	}
	pq.withSqueak = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Pool.Query().
//		GroupBy(pool.FieldCreateTime).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PoolQuery) GroupBy(field string, fields ...string) *PoolGroupBy {
	grbuild := &PoolGroupBy{config: pq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return pq.sqlQuery(ctx), nil
	}
	grbuild.label = pool.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.Pool.Query().
//		Select(pool.FieldCreateTime).
//		Scan(ctx, &v)
func (pq *PoolQuery) Select(fields ...string) *PoolSelect {
	pq.fields = append(pq.fields, fields...)
	selbuild := &PoolSelect{PoolQuery: pq}
	selbuild.label = pool.Label
	selbuild.flds, selbuild.scan = &pq.fields, selbuild.Scan
	return selbuild
}

// Aggregate returns a PoolSelect configured with the given aggregations.
func (pq *PoolQuery) Aggregate(fns ...AggregateFunc) *PoolSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PoolQuery) prepareQuery(ctx context.Context) error {
	for _, f := range pq.fields {
		if !pool.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PoolQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Pool, error) {
	var (
		nodes       = []*Pool{}
		withFKs     = pq.withFKs
		_spec       = pq.querySpec()
		loadedTypes = [2]bool{
			pq.withPoolPasses != nil,
			pq.withSqueak != nil,
		}
	)
	if pq.withSqueak != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, pool.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Pool).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Pool{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withPoolPasses; query != nil {
		if err := pq.loadPoolPasses(ctx, query, nodes,
			func(n *Pool) { n.Edges.PoolPasses = []*PoolPass{} },
			func(n *Pool, e *PoolPass) { n.Edges.PoolPasses = append(n.Edges.PoolPasses, e) }); err != nil {
			return nil, err
		}
	}
	if query := pq.withSqueak; query != nil {
		if err := pq.loadSqueak(ctx, query, nodes, nil,
			func(n *Pool, e *Squeak) { n.Edges.Squeak = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PoolQuery) loadPoolPasses(ctx context.Context, query *PoolPassQuery, nodes []*Pool, init func(*Pool), assign func(*Pool, *PoolPass)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Pool)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.PoolPass(func(s *sql.Selector) {
		s.Where(sql.InValues(pool.PoolPassesColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.pool_pool_passes
		if fk == nil {
			return fmt.Errorf(`foreign-key "pool_pool_passes" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "pool_pool_passes" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}
func (pq *PoolQuery) loadSqueak(ctx context.Context, query *SqueakQuery, nodes []*Pool, init func(*Pool), assign func(*Pool, *Squeak)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*Pool)
	for i := range nodes {
		if nodes[i].squeak_pool == nil {
			continue
		}
		fk := *nodes[i].squeak_pool
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	query.Where(squeak.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "squeak_pool" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (pq *PoolQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.fields
	if len(pq.fields) > 0 {
		_spec.Unique = pq.unique != nil && *pq.unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PoolQuery) sqlExist(ctx context.Context) (bool, error) {
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

func (pq *PoolQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   pool.Table,
			Columns: pool.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pool.FieldID,
			},
		},
		From:   pq.sql,
		Unique: true,
	}
	if unique := pq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := pq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, pool.FieldID)
		for i := range fields {
			if fields[i] != pool.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PoolQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(pool.Table)
	columns := pq.fields
	if len(columns) == 0 {
		columns = pool.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.unique != nil && *pq.unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PoolGroupBy is the group-by builder for Pool entities.
type PoolGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PoolGroupBy) Aggregate(fns ...AggregateFunc) *PoolGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the group-by query and scans the result into the given value.
func (pgb *PoolGroupBy) Scan(ctx context.Context, v any) error {
	query, err := pgb.path(ctx)
	if err != nil {
		return err
	}
	pgb.sql = query
	return pgb.sqlScan(ctx, v)
}

func (pgb *PoolGroupBy) sqlScan(ctx context.Context, v any) error {
	for _, f := range pgb.fields {
		if !pool.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pgb *PoolGroupBy) sqlQuery() *sql.Selector {
	selector := pgb.sql.Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pgb.fields)+len(pgb.fns))
		for _, f := range pgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pgb.fields...)...)
}

// PoolSelect is the builder for selecting fields of Pool entities.
type PoolSelect struct {
	*PoolQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PoolSelect) Aggregate(fns ...AggregateFunc) *PoolSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PoolSelect) Scan(ctx context.Context, v any) error {
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	ps.sql = ps.PoolQuery.sqlQuery(ctx)
	return ps.sqlScan(ctx, v)
}

func (ps *PoolSelect) sqlScan(ctx context.Context, v any) error {
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(ps.sql))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		ps.sql.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		ps.sql.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := ps.sql.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
