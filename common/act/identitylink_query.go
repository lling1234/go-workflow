// Code generated by ent, DO NOT EDIT.

package act

import (
	"act/common/act/identitylink"
	"act/common/act/predicate"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// IdentityLinkQuery is the builder for querying IdentityLink entities.
type IdentityLinkQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.IdentityLink
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the IdentityLinkQuery builder.
func (ilq *IdentityLinkQuery) Where(ps ...predicate.IdentityLink) *IdentityLinkQuery {
	ilq.predicates = append(ilq.predicates, ps...)
	return ilq
}

// Limit adds a limit step to the query.
func (ilq *IdentityLinkQuery) Limit(limit int) *IdentityLinkQuery {
	ilq.limit = &limit
	return ilq
}

// Offset adds an offset step to the query.
func (ilq *IdentityLinkQuery) Offset(offset int) *IdentityLinkQuery {
	ilq.offset = &offset
	return ilq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ilq *IdentityLinkQuery) Unique(unique bool) *IdentityLinkQuery {
	ilq.unique = &unique
	return ilq
}

// Order adds an order step to the query.
func (ilq *IdentityLinkQuery) Order(o ...OrderFunc) *IdentityLinkQuery {
	ilq.order = append(ilq.order, o...)
	return ilq
}

// First returns the first IdentityLink entity from the query.
// Returns a *NotFoundError when no IdentityLink was found.
func (ilq *IdentityLinkQuery) First(ctx context.Context) (*IdentityLink, error) {
	nodes, err := ilq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{identitylink.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ilq *IdentityLinkQuery) FirstX(ctx context.Context) *IdentityLink {
	node, err := ilq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first IdentityLink ID from the query.
// Returns a *NotFoundError when no IdentityLink ID was found.
func (ilq *IdentityLinkQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ilq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{identitylink.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ilq *IdentityLinkQuery) FirstIDX(ctx context.Context) int {
	id, err := ilq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single IdentityLink entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one IdentityLink entity is found.
// Returns a *NotFoundError when no IdentityLink entities are found.
func (ilq *IdentityLinkQuery) Only(ctx context.Context) (*IdentityLink, error) {
	nodes, err := ilq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{identitylink.Label}
	default:
		return nil, &NotSingularError{identitylink.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ilq *IdentityLinkQuery) OnlyX(ctx context.Context) *IdentityLink {
	node, err := ilq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only IdentityLink ID in the query.
// Returns a *NotSingularError when more than one IdentityLink ID is found.
// Returns a *NotFoundError when no entities are found.
func (ilq *IdentityLinkQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ilq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{identitylink.Label}
	default:
		err = &NotSingularError{identitylink.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ilq *IdentityLinkQuery) OnlyIDX(ctx context.Context) int {
	id, err := ilq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of IdentityLinks.
func (ilq *IdentityLinkQuery) All(ctx context.Context) ([]*IdentityLink, error) {
	if err := ilq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return ilq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (ilq *IdentityLinkQuery) AllX(ctx context.Context) []*IdentityLink {
	nodes, err := ilq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of IdentityLink IDs.
func (ilq *IdentityLinkQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := ilq.Select(identitylink.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ilq *IdentityLinkQuery) IDsX(ctx context.Context) []int {
	ids, err := ilq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ilq *IdentityLinkQuery) Count(ctx context.Context) (int, error) {
	if err := ilq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return ilq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (ilq *IdentityLinkQuery) CountX(ctx context.Context) int {
	count, err := ilq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ilq *IdentityLinkQuery) Exist(ctx context.Context) (bool, error) {
	if err := ilq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return ilq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (ilq *IdentityLinkQuery) ExistX(ctx context.Context) bool {
	exist, err := ilq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the IdentityLinkQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ilq *IdentityLinkQuery) Clone() *IdentityLinkQuery {
	if ilq == nil {
		return nil
	}
	return &IdentityLinkQuery{
		config:     ilq.config,
		limit:      ilq.limit,
		offset:     ilq.offset,
		order:      append([]OrderFunc{}, ilq.order...),
		predicates: append([]predicate.IdentityLink{}, ilq.predicates...),
		// clone intermediate query.
		sql:    ilq.sql.Clone(),
		path:   ilq.path,
		unique: ilq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		UserID int64 `json:"user_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.IdentityLink.Query().
//		GroupBy(identitylink.FieldUserID).
//		Aggregate(act.Count()).
//		Scan(ctx, &v)
//
func (ilq *IdentityLinkQuery) GroupBy(field string, fields ...string) *IdentityLinkGroupBy {
	grbuild := &IdentityLinkGroupBy{config: ilq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := ilq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return ilq.sqlQuery(ctx), nil
	}
	grbuild.label = identitylink.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		UserID int64 `json:"user_id,omitempty"`
//	}
//
//	client.IdentityLink.Query().
//		Select(identitylink.FieldUserID).
//		Scan(ctx, &v)
//
func (ilq *IdentityLinkQuery) Select(fields ...string) *IdentityLinkSelect {
	ilq.fields = append(ilq.fields, fields...)
	selbuild := &IdentityLinkSelect{IdentityLinkQuery: ilq}
	selbuild.label = identitylink.Label
	selbuild.flds, selbuild.scan = &ilq.fields, selbuild.Scan
	return selbuild
}

func (ilq *IdentityLinkQuery) prepareQuery(ctx context.Context) error {
	for _, f := range ilq.fields {
		if !identitylink.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("act: invalid field %q for query", f)}
		}
	}
	if ilq.path != nil {
		prev, err := ilq.path(ctx)
		if err != nil {
			return err
		}
		ilq.sql = prev
	}
	return nil
}

func (ilq *IdentityLinkQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*IdentityLink, error) {
	var (
		nodes = []*IdentityLink{}
		_spec = ilq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*IdentityLink).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &IdentityLink{config: ilq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ilq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (ilq *IdentityLinkQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ilq.querySpec()
	_spec.Node.Columns = ilq.fields
	if len(ilq.fields) > 0 {
		_spec.Unique = ilq.unique != nil && *ilq.unique
	}
	return sqlgraph.CountNodes(ctx, ilq.driver, _spec)
}

func (ilq *IdentityLinkQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := ilq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("act: check existence: %w", err)
	}
	return n > 0, nil
}

func (ilq *IdentityLinkQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   identitylink.Table,
			Columns: identitylink.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: identitylink.FieldID,
			},
		},
		From:   ilq.sql,
		Unique: true,
	}
	if unique := ilq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := ilq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, identitylink.FieldID)
		for i := range fields {
			if fields[i] != identitylink.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ilq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ilq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ilq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ilq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ilq *IdentityLinkQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ilq.driver.Dialect())
	t1 := builder.Table(identitylink.Table)
	columns := ilq.fields
	if len(columns) == 0 {
		columns = identitylink.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ilq.sql != nil {
		selector = ilq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ilq.unique != nil && *ilq.unique {
		selector.Distinct()
	}
	for _, p := range ilq.predicates {
		p(selector)
	}
	for _, p := range ilq.order {
		p(selector)
	}
	if offset := ilq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ilq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// IdentityLinkGroupBy is the group-by builder for IdentityLink entities.
type IdentityLinkGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ilgb *IdentityLinkGroupBy) Aggregate(fns ...AggregateFunc) *IdentityLinkGroupBy {
	ilgb.fns = append(ilgb.fns, fns...)
	return ilgb
}

// Scan applies the group-by query and scans the result into the given value.
func (ilgb *IdentityLinkGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := ilgb.path(ctx)
	if err != nil {
		return err
	}
	ilgb.sql = query
	return ilgb.sqlScan(ctx, v)
}

func (ilgb *IdentityLinkGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range ilgb.fields {
		if !identitylink.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := ilgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ilgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (ilgb *IdentityLinkGroupBy) sqlQuery() *sql.Selector {
	selector := ilgb.sql.Select()
	aggregation := make([]string, 0, len(ilgb.fns))
	for _, fn := range ilgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(ilgb.fields)+len(ilgb.fns))
		for _, f := range ilgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(ilgb.fields...)...)
}

// IdentityLinkSelect is the builder for selecting fields of IdentityLink entities.
type IdentityLinkSelect struct {
	*IdentityLinkQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ils *IdentityLinkSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ils.prepareQuery(ctx); err != nil {
		return err
	}
	ils.sql = ils.IdentityLinkQuery.sqlQuery(ctx)
	return ils.sqlScan(ctx, v)
}

func (ils *IdentityLinkSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ils.sql.Query()
	if err := ils.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
