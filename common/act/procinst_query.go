// Code generated by ent, DO NOT EDIT.

package act

import (
	"act/common/act/predicate"
	"act/common/act/procinst"
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ProcInstQuery is the builder for querying ProcInst entities.
type ProcInstQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.ProcInst
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ProcInstQuery builder.
func (piq *ProcInstQuery) Where(ps ...predicate.ProcInst) *ProcInstQuery {
	piq.predicates = append(piq.predicates, ps...)
	return piq
}

// Limit adds a limit step to the query.
func (piq *ProcInstQuery) Limit(limit int) *ProcInstQuery {
	piq.limit = &limit
	return piq
}

// Offset adds an offset step to the query.
func (piq *ProcInstQuery) Offset(offset int) *ProcInstQuery {
	piq.offset = &offset
	return piq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (piq *ProcInstQuery) Unique(unique bool) *ProcInstQuery {
	piq.unique = &unique
	return piq
}

// Order adds an order step to the query.
func (piq *ProcInstQuery) Order(o ...OrderFunc) *ProcInstQuery {
	piq.order = append(piq.order, o...)
	return piq
}

// First returns the first ProcInst entity from the query.
// Returns a *NotFoundError when no ProcInst was found.
func (piq *ProcInstQuery) First(ctx context.Context) (*ProcInst, error) {
	nodes, err := piq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{procinst.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (piq *ProcInstQuery) FirstX(ctx context.Context) *ProcInst {
	node, err := piq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ProcInst ID from the query.
// Returns a *NotFoundError when no ProcInst ID was found.
func (piq *ProcInstQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = piq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{procinst.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (piq *ProcInstQuery) FirstIDX(ctx context.Context) int {
	id, err := piq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ProcInst entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ProcInst entity is found.
// Returns a *NotFoundError when no ProcInst entities are found.
func (piq *ProcInstQuery) Only(ctx context.Context) (*ProcInst, error) {
	nodes, err := piq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{procinst.Label}
	default:
		return nil, &NotSingularError{procinst.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (piq *ProcInstQuery) OnlyX(ctx context.Context) *ProcInst {
	node, err := piq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ProcInst ID in the query.
// Returns a *NotSingularError when more than one ProcInst ID is found.
// Returns a *NotFoundError when no entities are found.
func (piq *ProcInstQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = piq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{procinst.Label}
	default:
		err = &NotSingularError{procinst.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (piq *ProcInstQuery) OnlyIDX(ctx context.Context) int {
	id, err := piq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ProcInsts.
func (piq *ProcInstQuery) All(ctx context.Context) ([]*ProcInst, error) {
	if err := piq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return piq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (piq *ProcInstQuery) AllX(ctx context.Context) []*ProcInst {
	nodes, err := piq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ProcInst IDs.
func (piq *ProcInstQuery) IDs(ctx context.Context) ([]int, error) {
	var ids []int
	if err := piq.Select(procinst.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (piq *ProcInstQuery) IDsX(ctx context.Context) []int {
	ids, err := piq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (piq *ProcInstQuery) Count(ctx context.Context) (int, error) {
	if err := piq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return piq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (piq *ProcInstQuery) CountX(ctx context.Context) int {
	count, err := piq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (piq *ProcInstQuery) Exist(ctx context.Context) (bool, error) {
	if err := piq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return piq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (piq *ProcInstQuery) ExistX(ctx context.Context) bool {
	exist, err := piq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ProcInstQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (piq *ProcInstQuery) Clone() *ProcInstQuery {
	if piq == nil {
		return nil
	}
	return &ProcInstQuery{
		config:     piq.config,
		limit:      piq.limit,
		offset:     piq.offset,
		order:      append([]OrderFunc{}, piq.order...),
		predicates: append([]predicate.ProcInst{}, piq.predicates...),
		// clone intermediate query.
		sql:    piq.sql.Clone(),
		path:   piq.path,
		unique: piq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		ProcDefID int64 `json:"proc_def_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ProcInst.Query().
//		GroupBy(procinst.FieldProcDefID).
//		Aggregate(act.Count()).
//		Scan(ctx, &v)
//
func (piq *ProcInstQuery) GroupBy(field string, fields ...string) *ProcInstGroupBy {
	grbuild := &ProcInstGroupBy{config: piq.config}
	grbuild.fields = append([]string{field}, fields...)
	grbuild.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := piq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return piq.sqlQuery(ctx), nil
	}
	grbuild.label = procinst.Label
	grbuild.flds, grbuild.scan = &grbuild.fields, grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		ProcDefID int64 `json:"proc_def_id,omitempty"`
//	}
//
//	client.ProcInst.Query().
//		Select(procinst.FieldProcDefID).
//		Scan(ctx, &v)
//
func (piq *ProcInstQuery) Select(fields ...string) *ProcInstSelect {
	piq.fields = append(piq.fields, fields...)
	selbuild := &ProcInstSelect{ProcInstQuery: piq}
	selbuild.label = procinst.Label
	selbuild.flds, selbuild.scan = &piq.fields, selbuild.Scan
	return selbuild
}

func (piq *ProcInstQuery) prepareQuery(ctx context.Context) error {
	for _, f := range piq.fields {
		if !procinst.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("act: invalid field %q for query", f)}
		}
	}
	if piq.path != nil {
		prev, err := piq.path(ctx)
		if err != nil {
			return err
		}
		piq.sql = prev
	}
	return nil
}

func (piq *ProcInstQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ProcInst, error) {
	var (
		nodes = []*ProcInst{}
		_spec = piq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		return (*ProcInst).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		node := &ProcInst{config: piq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, piq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (piq *ProcInstQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := piq.querySpec()
	_spec.Node.Columns = piq.fields
	if len(piq.fields) > 0 {
		_spec.Unique = piq.unique != nil && *piq.unique
	}
	return sqlgraph.CountNodes(ctx, piq.driver, _spec)
}

func (piq *ProcInstQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := piq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("act: check existence: %w", err)
	}
	return n > 0, nil
}

func (piq *ProcInstQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   procinst.Table,
			Columns: procinst.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: procinst.FieldID,
			},
		},
		From:   piq.sql,
		Unique: true,
	}
	if unique := piq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := piq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, procinst.FieldID)
		for i := range fields {
			if fields[i] != procinst.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := piq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := piq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := piq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := piq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (piq *ProcInstQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(piq.driver.Dialect())
	t1 := builder.Table(procinst.Table)
	columns := piq.fields
	if len(columns) == 0 {
		columns = procinst.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if piq.sql != nil {
		selector = piq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if piq.unique != nil && *piq.unique {
		selector.Distinct()
	}
	for _, p := range piq.predicates {
		p(selector)
	}
	for _, p := range piq.order {
		p(selector)
	}
	if offset := piq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := piq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ProcInstGroupBy is the group-by builder for ProcInst entities.
type ProcInstGroupBy struct {
	config
	selector
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pigb *ProcInstGroupBy) Aggregate(fns ...AggregateFunc) *ProcInstGroupBy {
	pigb.fns = append(pigb.fns, fns...)
	return pigb
}

// Scan applies the group-by query and scans the result into the given value.
func (pigb *ProcInstGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := pigb.path(ctx)
	if err != nil {
		return err
	}
	pigb.sql = query
	return pigb.sqlScan(ctx, v)
}

func (pigb *ProcInstGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range pigb.fields {
		if !procinst.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := pigb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pigb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (pigb *ProcInstGroupBy) sqlQuery() *sql.Selector {
	selector := pigb.sql.Select()
	aggregation := make([]string, 0, len(pigb.fns))
	for _, fn := range pigb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(pigb.fields)+len(pigb.fns))
		for _, f := range pigb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(pigb.fields...)...)
}

// ProcInstSelect is the builder for selecting fields of ProcInst entities.
type ProcInstSelect struct {
	*ProcInstQuery
	selector
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (pis *ProcInstSelect) Scan(ctx context.Context, v interface{}) error {
	if err := pis.prepareQuery(ctx); err != nil {
		return err
	}
	pis.sql = pis.ProcInstQuery.sqlQuery(ctx)
	return pis.sqlScan(ctx, v)
}

func (pis *ProcInstSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := pis.sql.Query()
	if err := pis.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
