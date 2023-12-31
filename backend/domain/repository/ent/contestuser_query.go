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
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contestuser"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/predicate"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
)

// ContestUserQuery is the builder for querying ContestUser entities.
type ContestUserQuery struct {
	config
	ctx         *QueryContext
	order       []contestuser.OrderOption
	inters      []Interceptor
	predicates  []predicate.ContestUser
	withContest *ContestQuery
	withUser    *UserQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ContestUserQuery builder.
func (cuq *ContestUserQuery) Where(ps ...predicate.ContestUser) *ContestUserQuery {
	cuq.predicates = append(cuq.predicates, ps...)
	return cuq
}

// Limit the number of records to be returned by this query.
func (cuq *ContestUserQuery) Limit(limit int) *ContestUserQuery {
	cuq.ctx.Limit = &limit
	return cuq
}

// Offset to start from.
func (cuq *ContestUserQuery) Offset(offset int) *ContestUserQuery {
	cuq.ctx.Offset = &offset
	return cuq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (cuq *ContestUserQuery) Unique(unique bool) *ContestUserQuery {
	cuq.ctx.Unique = &unique
	return cuq
}

// Order specifies how the records should be ordered.
func (cuq *ContestUserQuery) Order(o ...contestuser.OrderOption) *ContestUserQuery {
	cuq.order = append(cuq.order, o...)
	return cuq
}

// QueryContest chains the current query on the "contest" edge.
func (cuq *ContestUserQuery) QueryContest() *ContestQuery {
	query := (&ContestClient{config: cuq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cuq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(contestuser.Table, contestuser.FieldID, selector),
			sqlgraph.To(contest.Table, contest.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, contestuser.ContestTable, contestuser.ContestColumn),
		)
		fromU = sqlgraph.SetNeighbors(cuq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUser chains the current query on the "user" edge.
func (cuq *ContestUserQuery) QueryUser() *UserQuery {
	query := (&UserClient{config: cuq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := cuq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := cuq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(contestuser.Table, contestuser.FieldID, selector),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, contestuser.UserTable, contestuser.UserColumn),
		)
		fromU = sqlgraph.SetNeighbors(cuq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first ContestUser entity from the query.
// Returns a *NotFoundError when no ContestUser was found.
func (cuq *ContestUserQuery) First(ctx context.Context) (*ContestUser, error) {
	nodes, err := cuq.Limit(1).All(setContextOp(ctx, cuq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{contestuser.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (cuq *ContestUserQuery) FirstX(ctx context.Context) *ContestUser {
	node, err := cuq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ContestUser ID from the query.
// Returns a *NotFoundError when no ContestUser ID was found.
func (cuq *ContestUserQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cuq.Limit(1).IDs(setContextOp(ctx, cuq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{contestuser.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (cuq *ContestUserQuery) FirstIDX(ctx context.Context) int {
	id, err := cuq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ContestUser entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ContestUser entity is found.
// Returns a *NotFoundError when no ContestUser entities are found.
func (cuq *ContestUserQuery) Only(ctx context.Context) (*ContestUser, error) {
	nodes, err := cuq.Limit(2).All(setContextOp(ctx, cuq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{contestuser.Label}
	default:
		return nil, &NotSingularError{contestuser.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (cuq *ContestUserQuery) OnlyX(ctx context.Context) *ContestUser {
	node, err := cuq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ContestUser ID in the query.
// Returns a *NotSingularError when more than one ContestUser ID is found.
// Returns a *NotFoundError when no entities are found.
func (cuq *ContestUserQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = cuq.Limit(2).IDs(setContextOp(ctx, cuq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{contestuser.Label}
	default:
		err = &NotSingularError{contestuser.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (cuq *ContestUserQuery) OnlyIDX(ctx context.Context) int {
	id, err := cuq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ContestUsers.
func (cuq *ContestUserQuery) All(ctx context.Context) ([]*ContestUser, error) {
	ctx = setContextOp(ctx, cuq.ctx, "All")
	if err := cuq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ContestUser, *ContestUserQuery]()
	return withInterceptors[[]*ContestUser](ctx, cuq, qr, cuq.inters)
}

// AllX is like All, but panics if an error occurs.
func (cuq *ContestUserQuery) AllX(ctx context.Context) []*ContestUser {
	nodes, err := cuq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ContestUser IDs.
func (cuq *ContestUserQuery) IDs(ctx context.Context) (ids []int, err error) {
	if cuq.ctx.Unique == nil && cuq.path != nil {
		cuq.Unique(true)
	}
	ctx = setContextOp(ctx, cuq.ctx, "IDs")
	if err = cuq.Select(contestuser.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (cuq *ContestUserQuery) IDsX(ctx context.Context) []int {
	ids, err := cuq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (cuq *ContestUserQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, cuq.ctx, "Count")
	if err := cuq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, cuq, querierCount[*ContestUserQuery](), cuq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (cuq *ContestUserQuery) CountX(ctx context.Context) int {
	count, err := cuq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (cuq *ContestUserQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, cuq.ctx, "Exist")
	switch _, err := cuq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (cuq *ContestUserQuery) ExistX(ctx context.Context) bool {
	exist, err := cuq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ContestUserQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (cuq *ContestUserQuery) Clone() *ContestUserQuery {
	if cuq == nil {
		return nil
	}
	return &ContestUserQuery{
		config:      cuq.config,
		ctx:         cuq.ctx.Clone(),
		order:       append([]contestuser.OrderOption{}, cuq.order...),
		inters:      append([]Interceptor{}, cuq.inters...),
		predicates:  append([]predicate.ContestUser{}, cuq.predicates...),
		withContest: cuq.withContest.Clone(),
		withUser:    cuq.withUser.Clone(),
		// clone intermediate query.
		sql:  cuq.sql.Clone(),
		path: cuq.path,
	}
}

// WithContest tells the query-builder to eager-load the nodes that are connected to
// the "contest" edge. The optional arguments are used to configure the query builder of the edge.
func (cuq *ContestUserQuery) WithContest(opts ...func(*ContestQuery)) *ContestUserQuery {
	query := (&ContestClient{config: cuq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cuq.withContest = query
	return cuq
}

// WithUser tells the query-builder to eager-load the nodes that are connected to
// the "user" edge. The optional arguments are used to configure the query builder of the edge.
func (cuq *ContestUserQuery) WithUser(opts ...func(*UserQuery)) *ContestUserQuery {
	query := (&UserClient{config: cuq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	cuq.withUser = query
	return cuq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Role string `json:"role,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ContestUser.Query().
//		GroupBy(contestuser.FieldRole).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (cuq *ContestUserQuery) GroupBy(field string, fields ...string) *ContestUserGroupBy {
	cuq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ContestUserGroupBy{build: cuq}
	grbuild.flds = &cuq.ctx.Fields
	grbuild.label = contestuser.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Role string `json:"role,omitempty"`
//	}
//
//	client.ContestUser.Query().
//		Select(contestuser.FieldRole).
//		Scan(ctx, &v)
func (cuq *ContestUserQuery) Select(fields ...string) *ContestUserSelect {
	cuq.ctx.Fields = append(cuq.ctx.Fields, fields...)
	sbuild := &ContestUserSelect{ContestUserQuery: cuq}
	sbuild.label = contestuser.Label
	sbuild.flds, sbuild.scan = &cuq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ContestUserSelect configured with the given aggregations.
func (cuq *ContestUserQuery) Aggregate(fns ...AggregateFunc) *ContestUserSelect {
	return cuq.Select().Aggregate(fns...)
}

func (cuq *ContestUserQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range cuq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, cuq); err != nil {
				return err
			}
		}
	}
	for _, f := range cuq.ctx.Fields {
		if !contestuser.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if cuq.path != nil {
		prev, err := cuq.path(ctx)
		if err != nil {
			return err
		}
		cuq.sql = prev
	}
	return nil
}

func (cuq *ContestUserQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ContestUser, error) {
	var (
		nodes       = []*ContestUser{}
		_spec       = cuq.querySpec()
		loadedTypes = [2]bool{
			cuq.withContest != nil,
			cuq.withUser != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ContestUser).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ContestUser{config: cuq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, cuq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := cuq.withContest; query != nil {
		if err := cuq.loadContest(ctx, query, nodes, nil,
			func(n *ContestUser, e *Contest) { n.Edges.Contest = e }); err != nil {
			return nil, err
		}
	}
	if query := cuq.withUser; query != nil {
		if err := cuq.loadUser(ctx, query, nodes, nil,
			func(n *ContestUser, e *User) { n.Edges.User = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (cuq *ContestUserQuery) loadContest(ctx context.Context, query *ContestQuery, nodes []*ContestUser, init func(*ContestUser), assign func(*ContestUser, *Contest)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ContestUser)
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
func (cuq *ContestUserQuery) loadUser(ctx context.Context, query *UserQuery, nodes []*ContestUser, init func(*ContestUser), assign func(*ContestUser, *User)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*ContestUser)
	for i := range nodes {
		fk := nodes[i].UserID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(user.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "user_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (cuq *ContestUserQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := cuq.querySpec()
	_spec.Node.Columns = cuq.ctx.Fields
	if len(cuq.ctx.Fields) > 0 {
		_spec.Unique = cuq.ctx.Unique != nil && *cuq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, cuq.driver, _spec)
}

func (cuq *ContestUserQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(contestuser.Table, contestuser.Columns, sqlgraph.NewFieldSpec(contestuser.FieldID, field.TypeInt))
	_spec.From = cuq.sql
	if unique := cuq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if cuq.path != nil {
		_spec.Unique = true
	}
	if fields := cuq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, contestuser.FieldID)
		for i := range fields {
			if fields[i] != contestuser.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if cuq.withContest != nil {
			_spec.Node.AddColumnOnce(contestuser.FieldContestID)
		}
		if cuq.withUser != nil {
			_spec.Node.AddColumnOnce(contestuser.FieldUserID)
		}
	}
	if ps := cuq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := cuq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := cuq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := cuq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (cuq *ContestUserQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(cuq.driver.Dialect())
	t1 := builder.Table(contestuser.Table)
	columns := cuq.ctx.Fields
	if len(columns) == 0 {
		columns = contestuser.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if cuq.sql != nil {
		selector = cuq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if cuq.ctx.Unique != nil && *cuq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range cuq.predicates {
		p(selector)
	}
	for _, p := range cuq.order {
		p(selector)
	}
	if offset := cuq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := cuq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ContestUserGroupBy is the group-by builder for ContestUser entities.
type ContestUserGroupBy struct {
	selector
	build *ContestUserQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (cugb *ContestUserGroupBy) Aggregate(fns ...AggregateFunc) *ContestUserGroupBy {
	cugb.fns = append(cugb.fns, fns...)
	return cugb
}

// Scan applies the selector query and scans the result into the given value.
func (cugb *ContestUserGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cugb.build.ctx, "GroupBy")
	if err := cugb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ContestUserQuery, *ContestUserGroupBy](ctx, cugb.build, cugb, cugb.build.inters, v)
}

func (cugb *ContestUserGroupBy) sqlScan(ctx context.Context, root *ContestUserQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(cugb.fns))
	for _, fn := range cugb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*cugb.flds)+len(cugb.fns))
		for _, f := range *cugb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*cugb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cugb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ContestUserSelect is the builder for selecting fields of ContestUser entities.
type ContestUserSelect struct {
	*ContestUserQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (cus *ContestUserSelect) Aggregate(fns ...AggregateFunc) *ContestUserSelect {
	cus.fns = append(cus.fns, fns...)
	return cus
}

// Scan applies the selector query and scans the result into the given value.
func (cus *ContestUserSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, cus.ctx, "Select")
	if err := cus.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ContestUserQuery, *ContestUserSelect](ctx, cus.ContestUserQuery, cus, cus.inters, v)
}

func (cus *ContestUserSelect) sqlScan(ctx context.Context, root *ContestUserQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(cus.fns))
	for _, fn := range cus.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*cus.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := cus.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
