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
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/predicate"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/task"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/testcase"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/testcaseset"
)

// TestcaseSetQuery is the builder for querying TestcaseSet entities.
type TestcaseSetQuery struct {
	config
	ctx           *QueryContext
	order         []testcaseset.OrderOption
	inters        []Interceptor
	predicates    []predicate.TestcaseSet
	withTasks     *TaskQuery
	withTestcases *TestcaseQuery
	withFKs       bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the TestcaseSetQuery builder.
func (tsq *TestcaseSetQuery) Where(ps ...predicate.TestcaseSet) *TestcaseSetQuery {
	tsq.predicates = append(tsq.predicates, ps...)
	return tsq
}

// Limit the number of records to be returned by this query.
func (tsq *TestcaseSetQuery) Limit(limit int) *TestcaseSetQuery {
	tsq.ctx.Limit = &limit
	return tsq
}

// Offset to start from.
func (tsq *TestcaseSetQuery) Offset(offset int) *TestcaseSetQuery {
	tsq.ctx.Offset = &offset
	return tsq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (tsq *TestcaseSetQuery) Unique(unique bool) *TestcaseSetQuery {
	tsq.ctx.Unique = &unique
	return tsq
}

// Order specifies how the records should be ordered.
func (tsq *TestcaseSetQuery) Order(o ...testcaseset.OrderOption) *TestcaseSetQuery {
	tsq.order = append(tsq.order, o...)
	return tsq
}

// QueryTasks chains the current query on the "tasks" edge.
func (tsq *TestcaseSetQuery) QueryTasks() *TaskQuery {
	query := (&TaskClient{config: tsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(testcaseset.Table, testcaseset.FieldID, selector),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, testcaseset.TasksTable, testcaseset.TasksColumn),
		)
		fromU = sqlgraph.SetNeighbors(tsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTestcases chains the current query on the "testcases" edge.
func (tsq *TestcaseSetQuery) QueryTestcases() *TestcaseQuery {
	query := (&TestcaseClient{config: tsq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := tsq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := tsq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(testcaseset.Table, testcaseset.FieldID, selector),
			sqlgraph.To(testcase.Table, testcase.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, testcaseset.TestcasesTable, testcaseset.TestcasesPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(tsq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first TestcaseSet entity from the query.
// Returns a *NotFoundError when no TestcaseSet was found.
func (tsq *TestcaseSetQuery) First(ctx context.Context) (*TestcaseSet, error) {
	nodes, err := tsq.Limit(1).All(setContextOp(ctx, tsq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{testcaseset.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (tsq *TestcaseSetQuery) FirstX(ctx context.Context) *TestcaseSet {
	node, err := tsq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first TestcaseSet ID from the query.
// Returns a *NotFoundError when no TestcaseSet ID was found.
func (tsq *TestcaseSetQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tsq.Limit(1).IDs(setContextOp(ctx, tsq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{testcaseset.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (tsq *TestcaseSetQuery) FirstIDX(ctx context.Context) int {
	id, err := tsq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single TestcaseSet entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one TestcaseSet entity is found.
// Returns a *NotFoundError when no TestcaseSet entities are found.
func (tsq *TestcaseSetQuery) Only(ctx context.Context) (*TestcaseSet, error) {
	nodes, err := tsq.Limit(2).All(setContextOp(ctx, tsq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{testcaseset.Label}
	default:
		return nil, &NotSingularError{testcaseset.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (tsq *TestcaseSetQuery) OnlyX(ctx context.Context) *TestcaseSet {
	node, err := tsq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only TestcaseSet ID in the query.
// Returns a *NotSingularError when more than one TestcaseSet ID is found.
// Returns a *NotFoundError when no entities are found.
func (tsq *TestcaseSetQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = tsq.Limit(2).IDs(setContextOp(ctx, tsq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{testcaseset.Label}
	default:
		err = &NotSingularError{testcaseset.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (tsq *TestcaseSetQuery) OnlyIDX(ctx context.Context) int {
	id, err := tsq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of TestcaseSets.
func (tsq *TestcaseSetQuery) All(ctx context.Context) ([]*TestcaseSet, error) {
	ctx = setContextOp(ctx, tsq.ctx, "All")
	if err := tsq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*TestcaseSet, *TestcaseSetQuery]()
	return withInterceptors[[]*TestcaseSet](ctx, tsq, qr, tsq.inters)
}

// AllX is like All, but panics if an error occurs.
func (tsq *TestcaseSetQuery) AllX(ctx context.Context) []*TestcaseSet {
	nodes, err := tsq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of TestcaseSet IDs.
func (tsq *TestcaseSetQuery) IDs(ctx context.Context) (ids []int, err error) {
	if tsq.ctx.Unique == nil && tsq.path != nil {
		tsq.Unique(true)
	}
	ctx = setContextOp(ctx, tsq.ctx, "IDs")
	if err = tsq.Select(testcaseset.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (tsq *TestcaseSetQuery) IDsX(ctx context.Context) []int {
	ids, err := tsq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (tsq *TestcaseSetQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, tsq.ctx, "Count")
	if err := tsq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, tsq, querierCount[*TestcaseSetQuery](), tsq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (tsq *TestcaseSetQuery) CountX(ctx context.Context) int {
	count, err := tsq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (tsq *TestcaseSetQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, tsq.ctx, "Exist")
	switch _, err := tsq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (tsq *TestcaseSetQuery) ExistX(ctx context.Context) bool {
	exist, err := tsq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the TestcaseSetQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (tsq *TestcaseSetQuery) Clone() *TestcaseSetQuery {
	if tsq == nil {
		return nil
	}
	return &TestcaseSetQuery{
		config:        tsq.config,
		ctx:           tsq.ctx.Clone(),
		order:         append([]testcaseset.OrderOption{}, tsq.order...),
		inters:        append([]Interceptor{}, tsq.inters...),
		predicates:    append([]predicate.TestcaseSet{}, tsq.predicates...),
		withTasks:     tsq.withTasks.Clone(),
		withTestcases: tsq.withTestcases.Clone(),
		// clone intermediate query.
		sql:  tsq.sql.Clone(),
		path: tsq.path,
	}
}

// WithTasks tells the query-builder to eager-load the nodes that are connected to
// the "tasks" edge. The optional arguments are used to configure the query builder of the edge.
func (tsq *TestcaseSetQuery) WithTasks(opts ...func(*TaskQuery)) *TestcaseSetQuery {
	query := (&TaskClient{config: tsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tsq.withTasks = query
	return tsq
}

// WithTestcases tells the query-builder to eager-load the nodes that are connected to
// the "testcases" edge. The optional arguments are used to configure the query builder of the edge.
func (tsq *TestcaseSetQuery) WithTestcases(opts ...func(*TestcaseQuery)) *TestcaseSetQuery {
	query := (&TestcaseClient{config: tsq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	tsq.withTestcases = query
	return tsq
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
//	client.TestcaseSet.Query().
//		GroupBy(testcaseset.FieldName).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (tsq *TestcaseSetQuery) GroupBy(field string, fields ...string) *TestcaseSetGroupBy {
	tsq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &TestcaseSetGroupBy{build: tsq}
	grbuild.flds = &tsq.ctx.Fields
	grbuild.label = testcaseset.Label
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
//	client.TestcaseSet.Query().
//		Select(testcaseset.FieldName).
//		Scan(ctx, &v)
func (tsq *TestcaseSetQuery) Select(fields ...string) *TestcaseSetSelect {
	tsq.ctx.Fields = append(tsq.ctx.Fields, fields...)
	sbuild := &TestcaseSetSelect{TestcaseSetQuery: tsq}
	sbuild.label = testcaseset.Label
	sbuild.flds, sbuild.scan = &tsq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a TestcaseSetSelect configured with the given aggregations.
func (tsq *TestcaseSetQuery) Aggregate(fns ...AggregateFunc) *TestcaseSetSelect {
	return tsq.Select().Aggregate(fns...)
}

func (tsq *TestcaseSetQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range tsq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, tsq); err != nil {
				return err
			}
		}
	}
	for _, f := range tsq.ctx.Fields {
		if !testcaseset.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if tsq.path != nil {
		prev, err := tsq.path(ctx)
		if err != nil {
			return err
		}
		tsq.sql = prev
	}
	return nil
}

func (tsq *TestcaseSetQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*TestcaseSet, error) {
	var (
		nodes       = []*TestcaseSet{}
		withFKs     = tsq.withFKs
		_spec       = tsq.querySpec()
		loadedTypes = [2]bool{
			tsq.withTasks != nil,
			tsq.withTestcases != nil,
		}
	)
	if tsq.withTasks != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, testcaseset.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*TestcaseSet).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &TestcaseSet{config: tsq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, tsq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := tsq.withTasks; query != nil {
		if err := tsq.loadTasks(ctx, query, nodes, nil,
			func(n *TestcaseSet, e *Task) { n.Edges.Tasks = e }); err != nil {
			return nil, err
		}
	}
	if query := tsq.withTestcases; query != nil {
		if err := tsq.loadTestcases(ctx, query, nodes,
			func(n *TestcaseSet) { n.Edges.Testcases = []*Testcase{} },
			func(n *TestcaseSet, e *Testcase) { n.Edges.Testcases = append(n.Edges.Testcases, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (tsq *TestcaseSetQuery) loadTasks(ctx context.Context, query *TaskQuery, nodes []*TestcaseSet, init func(*TestcaseSet), assign func(*TestcaseSet, *Task)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*TestcaseSet)
	for i := range nodes {
		if nodes[i].task_testcase_sets == nil {
			continue
		}
		fk := *nodes[i].task_testcase_sets
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(task.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "task_testcase_sets" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (tsq *TestcaseSetQuery) loadTestcases(ctx context.Context, query *TestcaseQuery, nodes []*TestcaseSet, init func(*TestcaseSet), assign func(*TestcaseSet, *Testcase)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*TestcaseSet)
	nids := make(map[int]map[*TestcaseSet]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(testcaseset.TestcasesTable)
		s.Join(joinT).On(s.C(testcase.FieldID), joinT.C(testcaseset.TestcasesPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(testcaseset.TestcasesPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(testcaseset.TestcasesPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*TestcaseSet]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Testcase](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "testcases" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}

func (tsq *TestcaseSetQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := tsq.querySpec()
	_spec.Node.Columns = tsq.ctx.Fields
	if len(tsq.ctx.Fields) > 0 {
		_spec.Unique = tsq.ctx.Unique != nil && *tsq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, tsq.driver, _spec)
}

func (tsq *TestcaseSetQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(testcaseset.Table, testcaseset.Columns, sqlgraph.NewFieldSpec(testcaseset.FieldID, field.TypeInt))
	_spec.From = tsq.sql
	if unique := tsq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if tsq.path != nil {
		_spec.Unique = true
	}
	if fields := tsq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, testcaseset.FieldID)
		for i := range fields {
			if fields[i] != testcaseset.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := tsq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := tsq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := tsq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := tsq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (tsq *TestcaseSetQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(tsq.driver.Dialect())
	t1 := builder.Table(testcaseset.Table)
	columns := tsq.ctx.Fields
	if len(columns) == 0 {
		columns = testcaseset.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if tsq.sql != nil {
		selector = tsq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if tsq.ctx.Unique != nil && *tsq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range tsq.predicates {
		p(selector)
	}
	for _, p := range tsq.order {
		p(selector)
	}
	if offset := tsq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := tsq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// TestcaseSetGroupBy is the group-by builder for TestcaseSet entities.
type TestcaseSetGroupBy struct {
	selector
	build *TestcaseSetQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (tsgb *TestcaseSetGroupBy) Aggregate(fns ...AggregateFunc) *TestcaseSetGroupBy {
	tsgb.fns = append(tsgb.fns, fns...)
	return tsgb
}

// Scan applies the selector query and scans the result into the given value.
func (tsgb *TestcaseSetGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tsgb.build.ctx, "GroupBy")
	if err := tsgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TestcaseSetQuery, *TestcaseSetGroupBy](ctx, tsgb.build, tsgb, tsgb.build.inters, v)
}

func (tsgb *TestcaseSetGroupBy) sqlScan(ctx context.Context, root *TestcaseSetQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(tsgb.fns))
	for _, fn := range tsgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*tsgb.flds)+len(tsgb.fns))
		for _, f := range *tsgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*tsgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tsgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// TestcaseSetSelect is the builder for selecting fields of TestcaseSet entities.
type TestcaseSetSelect struct {
	*TestcaseSetQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (tss *TestcaseSetSelect) Aggregate(fns ...AggregateFunc) *TestcaseSetSelect {
	tss.fns = append(tss.fns, fns...)
	return tss
}

// Scan applies the selector query and scans the result into the given value.
func (tss *TestcaseSetSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, tss.ctx, "Select")
	if err := tss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*TestcaseSetQuery, *TestcaseSetSelect](ctx, tss.TestcaseSetQuery, tss, tss.inters, v)
}

func (tss *TestcaseSetSelect) sqlScan(ctx context.Context, root *TestcaseSetQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(tss.fns))
	for _, fn := range tss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*tss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := tss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
