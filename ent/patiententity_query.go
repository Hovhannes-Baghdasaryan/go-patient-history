// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"github.com/Hovhannes-Baghdasaryan/go-patient-history/ent/patiententity"
	"github.com/Hovhannes-Baghdasaryan/go-patient-history/ent/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// PatientEntityQuery is the builder for querying PatientEntity entities.
type PatientEntityQuery struct {
	config
	ctx        *QueryContext
	order      []patiententity.OrderOption
	inters     []Interceptor
	predicates []predicate.PatientEntity
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PatientEntityQuery builder.
func (peq *PatientEntityQuery) Where(ps ...predicate.PatientEntity) *PatientEntityQuery {
	peq.predicates = append(peq.predicates, ps...)
	return peq
}

// Limit the number of records to be returned by this query.
func (peq *PatientEntityQuery) Limit(limit int) *PatientEntityQuery {
	peq.ctx.Limit = &limit
	return peq
}

// Offset to start from.
func (peq *PatientEntityQuery) Offset(offset int) *PatientEntityQuery {
	peq.ctx.Offset = &offset
	return peq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (peq *PatientEntityQuery) Unique(unique bool) *PatientEntityQuery {
	peq.ctx.Unique = &unique
	return peq
}

// Order specifies how the records should be ordered.
func (peq *PatientEntityQuery) Order(o ...patiententity.OrderOption) *PatientEntityQuery {
	peq.order = append(peq.order, o...)
	return peq
}

// First returns the first PatientEntity entity from the query.
// Returns a *NotFoundError when no PatientEntity was found.
func (peq *PatientEntityQuery) First(ctx context.Context) (*PatientEntity, error) {
	nodes, err := peq.Limit(1).All(setContextOp(ctx, peq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{patiententity.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (peq *PatientEntityQuery) FirstX(ctx context.Context) *PatientEntity {
	node, err := peq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first PatientEntity ID from the query.
// Returns a *NotFoundError when no PatientEntity ID was found.
func (peq *PatientEntityQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = peq.Limit(1).IDs(setContextOp(ctx, peq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{patiententity.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (peq *PatientEntityQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := peq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single PatientEntity entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one PatientEntity entity is found.
// Returns a *NotFoundError when no PatientEntity entities are found.
func (peq *PatientEntityQuery) Only(ctx context.Context) (*PatientEntity, error) {
	nodes, err := peq.Limit(2).All(setContextOp(ctx, peq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{patiententity.Label}
	default:
		return nil, &NotSingularError{patiententity.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (peq *PatientEntityQuery) OnlyX(ctx context.Context) *PatientEntity {
	node, err := peq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only PatientEntity ID in the query.
// Returns a *NotSingularError when more than one PatientEntity ID is found.
// Returns a *NotFoundError when no entities are found.
func (peq *PatientEntityQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = peq.Limit(2).IDs(setContextOp(ctx, peq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{patiententity.Label}
	default:
		err = &NotSingularError{patiententity.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (peq *PatientEntityQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := peq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of PatientEntities.
func (peq *PatientEntityQuery) All(ctx context.Context) ([]*PatientEntity, error) {
	ctx = setContextOp(ctx, peq.ctx, "All")
	if err := peq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*PatientEntity, *PatientEntityQuery]()
	return withInterceptors[[]*PatientEntity](ctx, peq, qr, peq.inters)
}

// AllX is like All, but panics if an error occurs.
func (peq *PatientEntityQuery) AllX(ctx context.Context) []*PatientEntity {
	nodes, err := peq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of PatientEntity IDs.
func (peq *PatientEntityQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if peq.ctx.Unique == nil && peq.path != nil {
		peq.Unique(true)
	}
	ctx = setContextOp(ctx, peq.ctx, "IDs")
	if err = peq.Select(patiententity.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (peq *PatientEntityQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := peq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (peq *PatientEntityQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, peq.ctx, "Count")
	if err := peq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, peq, querierCount[*PatientEntityQuery](), peq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (peq *PatientEntityQuery) CountX(ctx context.Context) int {
	count, err := peq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (peq *PatientEntityQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, peq.ctx, "Exist")
	switch _, err := peq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (peq *PatientEntityQuery) ExistX(ctx context.Context) bool {
	exist, err := peq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PatientEntityQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (peq *PatientEntityQuery) Clone() *PatientEntityQuery {
	if peq == nil {
		return nil
	}
	return &PatientEntityQuery{
		config:     peq.config,
		ctx:        peq.ctx.Clone(),
		order:      append([]patiententity.OrderOption{}, peq.order...),
		inters:     append([]Interceptor{}, peq.inters...),
		predicates: append([]predicate.PatientEntity{}, peq.predicates...),
		// clone intermediate query.
		sql:  peq.sql.Clone(),
		path: peq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.PatientEntity.Query().
//		GroupBy(patiententity.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (peq *PatientEntityQuery) GroupBy(field string, fields ...string) *PatientEntityGroupBy {
	peq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PatientEntityGroupBy{build: peq}
	grbuild.flds = &peq.ctx.Fields
	grbuild.label = patiententity.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Name string `json:"name,omitempty"`
//	}
//
//	client.PatientEntity.Query().
//		Select(patiententity.FieldName).
//		Scan(ctx, &v)
func (peq *PatientEntityQuery) Select(fields ...string) *PatientEntitySelect {
	peq.ctx.Fields = append(peq.ctx.Fields, fields...)
	sbuild := &PatientEntitySelect{PatientEntityQuery: peq}
	sbuild.label = patiententity.Label
	sbuild.flds, sbuild.scan = &peq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PatientEntitySelect configured with the given aggregations.
func (peq *PatientEntityQuery) Aggregate(fns ...AggregateFunc) *PatientEntitySelect {
	return peq.Select().Aggregate(fns...)
}

func (peq *PatientEntityQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range peq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, peq); err != nil {
				return err
			}
		}
	}
	for _, f := range peq.ctx.Fields {
		if !patiententity.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if peq.path != nil {
		prev, err := peq.path(ctx)
		if err != nil {
			return err
		}
		peq.sql = prev
	}
	return nil
}

func (peq *PatientEntityQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*PatientEntity, error) {
	var (
		nodes = []*PatientEntity{}
		_spec = peq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*PatientEntity).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &PatientEntity{config: peq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, peq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (peq *PatientEntityQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := peq.querySpec()
	_spec.Node.Columns = peq.ctx.Fields
	if len(peq.ctx.Fields) > 0 {
		_spec.Unique = peq.ctx.Unique != nil && *peq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, peq.driver, _spec)
}

func (peq *PatientEntityQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(patiententity.Table, patiententity.Columns, sqlgraph.NewFieldSpec(patiententity.FieldID, field.TypeUUID))
	_spec.From = peq.sql
	if unique := peq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if peq.path != nil {
		_spec.Unique = true
	}
	if fields := peq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, patiententity.FieldID)
		for i := range fields {
			if fields[i] != patiententity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := peq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := peq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := peq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := peq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (peq *PatientEntityQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(peq.driver.Dialect())
	t1 := builder.Table(patiententity.Table)
	columns := peq.ctx.Fields
	if len(columns) == 0 {
		columns = patiententity.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if peq.sql != nil {
		selector = peq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if peq.ctx.Unique != nil && *peq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range peq.predicates {
		p(selector)
	}
	for _, p := range peq.order {
		p(selector)
	}
	if offset := peq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := peq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PatientEntityGroupBy is the group-by builder for PatientEntity entities.
type PatientEntityGroupBy struct {
	selector
	build *PatientEntityQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pegb *PatientEntityGroupBy) Aggregate(fns ...AggregateFunc) *PatientEntityGroupBy {
	pegb.fns = append(pegb.fns, fns...)
	return pegb
}

// Scan applies the selector query and scans the result into the given value.
func (pegb *PatientEntityGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pegb.build.ctx, "GroupBy")
	if err := pegb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PatientEntityQuery, *PatientEntityGroupBy](ctx, pegb.build, pegb, pegb.build.inters, v)
}

func (pegb *PatientEntityGroupBy) sqlScan(ctx context.Context, root *PatientEntityQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pegb.fns))
	for _, fn := range pegb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pegb.flds)+len(pegb.fns))
		for _, f := range *pegb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pegb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pegb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PatientEntitySelect is the builder for selecting fields of PatientEntity entities.
type PatientEntitySelect struct {
	*PatientEntityQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (pes *PatientEntitySelect) Aggregate(fns ...AggregateFunc) *PatientEntitySelect {
	pes.fns = append(pes.fns, fns...)
	return pes
}

// Scan applies the selector query and scans the result into the given value.
func (pes *PatientEntitySelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pes.ctx, "Select")
	if err := pes.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PatientEntityQuery, *PatientEntitySelect](ctx, pes.PatientEntityQuery, pes, pes.inters, v)
}

func (pes *PatientEntitySelect) sqlScan(ctx context.Context, root *PatientEntityQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(pes.fns))
	for _, fn := range pes.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*pes.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pes.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
