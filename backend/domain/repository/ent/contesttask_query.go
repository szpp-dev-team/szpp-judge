// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contest"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contesttask"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/predicate"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/task"
)

// ContestTaskQuery is the builder for querying ContestTask entities.
type ContestTaskQuery struct {
	config
	ctx         *QueryContext
	order       []contesttask.OrderOption
	inters      []Interceptor
	predicates  []predicate.ContestTask
	withContest *ContestQuery
	withTask    *TaskQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ContestTaskQuery builder.
func (ctq *ContestTaskQuery) Where(ps ...predicate.ContestTask) *ContestTaskQuery {
	ctq.predicates = append(ctq.predicates, ps...)
	return ctq
}

// Limit the number of records to be returned by this query.
func (ctq *ContestTaskQuery) Limit(limit int) *ContestTaskQuery {
	ctq.ctx.Limit = &limit
	return ctq
}

// Offset to start from.
func (ctq *ContestTaskQuery) Offset(offset int) *ContestTaskQuery {
	ctq.ctx.Offset = &offset
	return ctq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ctq *ContestTaskQuery) Unique(unique bool) *ContestTaskQuery {
	ctq.ctx.Unique = &unique
	return ctq
}

// Order specifies how the records should be ordered.
func (ctq *ContestTaskQuery) Order(o ...contesttask.OrderOption) *ContestTaskQuery {
	ctq.order = append(ctq.order, o...)
	return ctq
}

// QueryContest chains the current query on the "contest" edge.
func (ctq *ContestTaskQuery) QueryContest() *ContestQuery {
	query := (&ContestClient{config: ctq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(contesttask.Table, contesttask.FieldID, selector),
			sqlgraph.To(contest.Table, contest.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, contesttask.ContestTable, contesttask.ContestColumn),
		)
		fromU = sqlgraph.SetNeighbors(ctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryTask chains the current query on the "task" edge.
func (ctq *ContestTaskQuery) QueryTask() *TaskQuery {
	query := (&TaskClient{config: ctq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ctq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ctq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(contesttask.Table, contesttask.FieldID, selector),
			sqlgraph.To(task.Table, task.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, contesttask.TaskTable, contesttask.TaskColumn),
		)
		fromU = sqlgraph.SetNeighbors(ctq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ContestTask entity from the query.
// Returns a *NotFoundError when no ContestTask was found.
func (ctq *ContestTaskQuery) First(ctx context.Context) (*ContestTask, error) {
	nodes, err := ctq.Limit(1).All(setContextOp(ctx, ctq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{contesttask.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ctq *ContestTaskQuery) FirstX(ctx context.Context) *ContestTask {
	node, err := ctq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ContestTask ID from the query.
// Returns a *NotFoundError when no ContestTask ID was found.
func (ctq *ContestTaskQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctq.Limit(1).IDs(setContextOp(ctx, ctq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{contesttask.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ctq *ContestTaskQuery) FirstIDX(ctx context.Context) int {
	id, err := ctq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ContestTask entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ContestTask entity is found.
// Returns a *NotFoundError when no ContestTask entities are found.
func (ctq *ContestTaskQuery) Only(ctx context.Context) (*ContestTask, error) {
	nodes, err := ctq.Limit(2).All(setContextOp(ctx, ctq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{contesttask.Label}
	default:
		return nil, &NotSingularError{contesttask.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ctq *ContestTaskQuery) OnlyX(ctx context.Context) *ContestTask {
	node, err := ctq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ContestTask ID in the query.
// Returns a *NotSingularError when more than one ContestTask ID is found.
// Returns a *NotFoundError when no entities are found.
func (ctq *ContestTaskQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ctq.Limit(2).IDs(setContextOp(ctx, ctq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{contesttask.Label}
	default:
		err = &NotSingularError{contesttask.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ctq *ContestTaskQuery) OnlyIDX(ctx context.Context) int {
	id, err := ctq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ContestTasks.
func (ctq *ContestTaskQuery) All(ctx context.Context) ([]*ContestTask, error) {
	ctx = setContextOp(ctx, ctq.ctx, "All")
	if err := ctq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ContestTask, *ContestTaskQuery]()
	return withInterceptors[[]*ContestTask](ctx, ctq, qr, ctq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ctq *ContestTaskQuery) AllX(ctx context.Context) []*ContestTask {
	nodes, err := ctq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ContestTask IDs.
func (ctq *ContestTaskQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ctq.ctx.Unique == nil && ctq.path != nil {
		ctq.Unique(true)
	}
	ctx = setContextOp(ctx, ctq.ctx, "IDs")
	if err = ctq.Select(contesttask.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ctq *ContestTaskQuery) IDsX(ctx context.Context) []int {
	ids, err := ctq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ctq *ContestTaskQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ctq.ctx, "Count")
	if err := ctq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ctq, querierCount[*ContestTaskQuery](), ctq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ctq *ContestTaskQuery) CountX(ctx context.Context) int {
	count, err := ctq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ctq *ContestTaskQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ctq.ctx, "Exist")
	switch _, err := ctq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ctq *ContestTaskQuery) ExistX(ctx context.Context) bool {
	exist, err := ctq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ContestTaskQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ctq *ContestTaskQuery) Clone() *ContestTaskQuery {
	if ctq == nil {
		return nil
	}
	return &ContestTaskQuery{
		config:      ctq.config,
		ctx:         ctq.ctx.Clone(),
		order:       append([]contesttask.OrderOption{}, ctq.order...),
		inters:      append([]Interceptor{}, ctq.inters...),
		predicates:  append([]predicate.ContestTask{}, ctq.predicates...),
		withContest: ctq.withContest.Clone(),
		withTask:    ctq.withTask.Clone(),
		// clone intermediate query.
		sql:  ctq.sql.Clone(),
		path: ctq.path,
	}
}

// WithContest tells the query-builder to eager-load the nodes that are connected to
// the "contest" edge. The optional arguments are used to configure the query builder of the edge.
func (ctq *ContestTaskQuery) WithContest(opts ...func(*ContestQuery)) *ContestTaskQuery {
	query := (&ContestClient{config: ctq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ctq.withContest = query
	return ctq
}

// WithTask tells the query-builder to eager-load the nodes that are connected to
// the "task" edge. The optional arguments are used to configure the query builder of the edge.
func (ctq *ContestTaskQuery) WithTask(opts ...func(*TaskQuery)) *ContestTaskQuery {
	query := (&TaskClient{config: ctq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ctq.withTask = query
	return ctq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Score int `json:"score,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ContestTask.Query().
//		GroupBy(contesttask.FieldScore).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ctq *ContestTaskQuery) GroupBy(field string, fields ...string) *ContestTaskGroupBy {
	ctq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ContestTaskGroupBy{build: ctq}
	grbuild.flds = &ctq.ctx.Fields
	grbuild.label = contesttask.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Score int `json:"score,omitempty"`
//	}
//
//	client.ContestTask.Query().
//		Select(contesttask.FieldScore).
//		Scan(ctx, &v)
func (ctq *ContestTaskQuery) Select(fields ...string) *ContestTaskSelect {
	ctq.ctx.Fields = append(ctq.ctx.Fields, fields...)
	sbuild := &ContestTaskSelect{ContestTaskQuery: ctq}
	sbuild.label = contesttask.Label
	sbuild.flds, sbuild.scan = &ctq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ContestTaskSelect configured with the given aggregations.
func (ctq *ContestTaskQuery) Aggregate(fns ...AggregateFunc) *ContestTaskSelect {
	return ctq.Select().Aggregate(fns...)
}

func (ctq *ContestTaskQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ctq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ctq); err != nil {
				return err
			}
		}
	}
	for _, f := range ctq.ctx.Fields {
		if !contesttask.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ctq.path != nil {
		prev, err := ctq.path(ctx)
		if err != nil {
			return err
		}
		ctq.sql = prev
	}
	return nil
}

func (ctq *ContestTaskQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ContestTask, error) {
	var (
		nodes       = []*ContestTask{}
		_spec       = ctq.querySpec()
		loadedTypes = [2]bool{
			ctq.withContest != nil,
			ctq.withTask != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ContestTask).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ContestTask{config: ctq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ctq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ctq.withContest; query != nil {
		if err := ctq.loadContest(ctx, query, nodes, nil,
			func(n *ContestTask, e *Contest) { n.Edges.Contest = e }); err != nil {
			return nil, err
		}
	}
	if query := ctq.withTask; query != nil {
		if err := ctq.loadTask(ctx, query, nodes, nil,
			func(n *ContestTask, e *Task) { n.Edges.Task = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ctq *ContestTaskQuery) loadContest(ctx context.Context, query *ContestQuery, nodes []*ContestTask, init func(*ContestTask), assign func(*ContestTask, *Contest)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ContestTask)
	for i := range nodes {
		fk := nodes[i].ContestID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(contest.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "contest_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ctq *ContestTaskQuery) loadTask(ctx context.Context, query *TaskQuery, nodes []*ContestTask, init func(*ContestTask), assign func(*ContestTask, *Task)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ContestTask)
	for i := range nodes {
		fk := nodes[i].TaskID
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
			return fmt.Errorf(`unexpected foreign-key "task_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ctq *ContestTaskQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ctq.querySpec()
	_spec.Node.Columns = ctq.ctx.Fields
	if len(ctq.ctx.Fields) > 0 {
		_spec.Unique = ctq.ctx.Unique != nil && *ctq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ctq.driver, _spec)
}

func (ctq *ContestTaskQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(contesttask.Table, contesttask.Columns, sqlgraph.NewFieldSpec(contesttask.FieldID, field.TypeInt))
	_spec.From = ctq.sql
	if unique := ctq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ctq.path != nil {
		_spec.Unique = true
	}
	if fields := ctq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contesttask.FieldID)
		for i := range fields {
			if fields[i] != contesttask.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ctq.withContest != nil {
			_spec.Node.AddColumnOnce(contesttask.FieldContestID)
		}
		if ctq.withTask != nil {
			_spec.Node.AddColumnOnce(contesttask.FieldTaskID)
		}
	}
	if ps := ctq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ctq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ctq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ctq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ctq *ContestTaskQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ctq.driver.Dialect())
	t1 := builder.Table(contesttask.Table)
	columns := ctq.ctx.Fields
	if len(columns) == 0 {
		columns = contesttask.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ctq.sql != nil {
		selector = ctq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ctq.ctx.Unique != nil && *ctq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ctq.predicates {
		p(selector)
	}
	for _, p := range ctq.order {
		p(selector)
	}
	if offset := ctq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ctq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ContestTaskGroupBy is the group-by builder for ContestTask entities.
type ContestTaskGroupBy struct {
	selector
	build *ContestTaskQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ctgb *ContestTaskGroupBy) Aggregate(fns ...AggregateFunc) *ContestTaskGroupBy {
	ctgb.fns = append(ctgb.fns, fns...)
	return ctgb
}

// Scan applies the selector query and scans the result into the given value.
func (ctgb *ContestTaskGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ctgb.build.ctx, "GroupBy")
	if err := ctgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ContestTaskQuery, *ContestTaskGroupBy](ctx, ctgb.build, ctgb, ctgb.build.inters, v)
}

func (ctgb *ContestTaskGroupBy) sqlScan(ctx context.Context, root *ContestTaskQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ctgb.fns))
	for _, fn := range ctgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ctgb.flds)+len(ctgb.fns))
		for _, f := range *ctgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ctgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ctgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ContestTaskSelect is the builder for selecting fields of ContestTask entities.
type ContestTaskSelect struct {
	*ContestTaskQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cts *ContestTaskSelect) Aggregate(fns ...AggregateFunc) *ContestTaskSelect {
	cts.fns = append(cts.fns, fns...)
	return cts
}

// Scan applies the selector query and scans the result into the given value.
func (cts *ContestTaskSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cts.ctx, "Select")
	if err := cts.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ContestTaskQuery, *ContestTaskSelect](ctx, cts.ContestTaskQuery, cts, cts.inters, v)
}

func (cts *ContestTaskSelect) sqlScan(ctx context.Context, root *ContestTaskQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cts.fns))
	for _, fn := range cts.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cts.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cts.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}