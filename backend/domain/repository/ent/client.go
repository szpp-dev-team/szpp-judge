// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/migrate"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/language"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/submit"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/testcaseresult"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
)

// Client is the client that holds all ent builders.
type Client struct {
	config
	// Schema is the client for creating, migrating and dropping schema.
	Schema *migrate.Schema
	// Language is the client for interacting with the Language builders.
	Language *LanguageClient
	// Submit is the client for interacting with the Submit builders.
	Submit *SubmitClient
	// TestcaseResult is the client for interacting with the TestcaseResult builders.
	TestcaseResult *TestcaseResultClient
	// User is the client for interacting with the User builders.
	User *UserClient
}

// NewClient creates a new client configured with the given options.
func NewClient(opts ...Option) *Client {
	cfg := config{log: log.Println, hooks: &hooks{}, inters: &inters{}}
	cfg.options(opts...)
	client := &Client{config: cfg}
	client.init()
	return client
}

func (c *Client) init() {
	c.Schema = migrate.NewSchema(c.driver)
	c.Language = NewLanguageClient(c.config)
	c.Submit = NewSubmitClient(c.config)
	c.TestcaseResult = NewTestcaseResultClient(c.config)
	c.User = NewUserClient(c.config)
}

type (
	// config is the configuration for the client and its builder.
	config struct {
		// driver used for executing database requests.
		driver dialect.Driver
		// debug enable a debug logging.
		debug bool
		// log used for logging on debug mode.
		log func(...any)
		// hooks to execute on mutations.
		hooks *hooks
		// interceptors to execute on queries.
		inters *inters
	}
	// Option function to configure the client.
	Option func(*config)
)

// options applies the options on the config object.
func (c *config) options(opts ...Option) {
	for _, opt := range opts {
		opt(c)
	}
	if c.debug {
		c.driver = dialect.Debug(c.driver, c.log)
	}
}

// Debug enables debug logging on the ent.Driver.
func Debug() Option {
	return func(c *config) {
		c.debug = true
	}
}

// Log sets the logging function for debug mode.
func Log(fn func(...any)) Option {
	return func(c *config) {
		c.log = fn
	}
}

// Driver configures the client driver.
func Driver(driver dialect.Driver) Option {
	return func(c *config) {
		c.driver = driver
	}
}

// Open opens a database/sql.DB specified by the driver name and
// the data source name, and returns a new client attached to it.
// Optional parameters can be added for configuring the client.
func Open(driverName, dataSourceName string, options ...Option) (*Client, error) {
	switch driverName {
	case dialect.MySQL, dialect.Postgres, dialect.SQLite:
		drv, err := sql.Open(driverName, dataSourceName)
		if err != nil {
			return nil, err
		}
		return NewClient(append(options, Driver(drv))...), nil
	default:
		return nil, fmt.Errorf("unsupported driver: %q", driverName)
	}
}

// Tx returns a new transactional client. The provided context
// is used until the transaction is committed or rolled back.
func (c *Client) Tx(ctx context.Context) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := newTx(ctx, c.driver)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = tx
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		Language:       NewLanguageClient(cfg),
		Submit:         NewSubmitClient(cfg),
		TestcaseResult: NewTestcaseResultClient(cfg),
		User:           NewUserClient(cfg),
	}, nil
}

// BeginTx returns a transactional client with specified options.
func (c *Client) BeginTx(ctx context.Context, opts *sql.TxOptions) (*Tx, error) {
	if _, ok := c.driver.(*txDriver); ok {
		return nil, errors.New("ent: cannot start a transaction within a transaction")
	}
	tx, err := c.driver.(interface {
		BeginTx(context.Context, *sql.TxOptions) (dialect.Tx, error)
	}).BeginTx(ctx, opts)
	if err != nil {
		return nil, fmt.Errorf("ent: starting a transaction: %w", err)
	}
	cfg := c.config
	cfg.driver = &txDriver{tx: tx, drv: c.driver}
	return &Tx{
		ctx:            ctx,
		config:         cfg,
		Language:       NewLanguageClient(cfg),
		Submit:         NewSubmitClient(cfg),
		TestcaseResult: NewTestcaseResultClient(cfg),
		User:           NewUserClient(cfg),
	}, nil
}

// Debug returns a new debug-client. It's used to get verbose logging on specific operations.
//
//	client.Debug().
//		Language.
//		Query().
//		Count(ctx)
func (c *Client) Debug() *Client {
	if c.debug {
		return c
	}
	cfg := c.config
	cfg.driver = dialect.Debug(c.driver, c.log)
	client := &Client{config: cfg}
	client.init()
	return client
}

// Close closes the database connection and prevents new queries from starting.
func (c *Client) Close() error {
	return c.driver.Close()
}

// Use adds the mutation hooks to all the entity clients.
// In order to add hooks to a specific client, call: `client.Node.Use(...)`.
func (c *Client) Use(hooks ...Hook) {
	c.Language.Use(hooks...)
	c.Submit.Use(hooks...)
	c.TestcaseResult.Use(hooks...)
	c.User.Use(hooks...)
}

// Intercept adds the query interceptors to all the entity clients.
// In order to add interceptors to a specific client, call: `client.Node.Intercept(...)`.
func (c *Client) Intercept(interceptors ...Interceptor) {
	c.Language.Intercept(interceptors...)
	c.Submit.Intercept(interceptors...)
	c.TestcaseResult.Intercept(interceptors...)
	c.User.Intercept(interceptors...)
}

// Mutate implements the ent.Mutator interface.
func (c *Client) Mutate(ctx context.Context, m Mutation) (Value, error) {
	switch m := m.(type) {
	case *LanguageMutation:
		return c.Language.mutate(ctx, m)
	case *SubmitMutation:
		return c.Submit.mutate(ctx, m)
	case *TestcaseResultMutation:
		return c.TestcaseResult.mutate(ctx, m)
	case *UserMutation:
		return c.User.mutate(ctx, m)
	default:
		return nil, fmt.Errorf("ent: unknown mutation type %T", m)
	}
}

// LanguageClient is a client for the Language schema.
type LanguageClient struct {
	config
}

// NewLanguageClient returns a client for the Language from the given config.
func NewLanguageClient(c config) *LanguageClient {
	return &LanguageClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `language.Hooks(f(g(h())))`.
func (c *LanguageClient) Use(hooks ...Hook) {
	c.hooks.Language = append(c.hooks.Language, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `language.Intercept(f(g(h())))`.
func (c *LanguageClient) Intercept(interceptors ...Interceptor) {
	c.inters.Language = append(c.inters.Language, interceptors...)
}

// Create returns a builder for creating a Language entity.
func (c *LanguageClient) Create() *LanguageCreate {
	mutation := newLanguageMutation(c.config, OpCreate)
	return &LanguageCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Language entities.
func (c *LanguageClient) CreateBulk(builders ...*LanguageCreate) *LanguageCreateBulk {
	return &LanguageCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Language.
func (c *LanguageClient) Update() *LanguageUpdate {
	mutation := newLanguageMutation(c.config, OpUpdate)
	return &LanguageUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *LanguageClient) UpdateOne(l *Language) *LanguageUpdateOne {
	mutation := newLanguageMutation(c.config, OpUpdateOne, withLanguage(l))
	return &LanguageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *LanguageClient) UpdateOneID(id int) *LanguageUpdateOne {
	mutation := newLanguageMutation(c.config, OpUpdateOne, withLanguageID(id))
	return &LanguageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Language.
func (c *LanguageClient) Delete() *LanguageDelete {
	mutation := newLanguageMutation(c.config, OpDelete)
	return &LanguageDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *LanguageClient) DeleteOne(l *Language) *LanguageDeleteOne {
	return c.DeleteOneID(l.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *LanguageClient) DeleteOneID(id int) *LanguageDeleteOne {
	builder := c.Delete().Where(language.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &LanguageDeleteOne{builder}
}

// Query returns a query builder for Language.
func (c *LanguageClient) Query() *LanguageQuery {
	return &LanguageQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeLanguage},
		inters: c.Interceptors(),
	}
}

// Get returns a Language entity by its id.
func (c *LanguageClient) Get(ctx context.Context, id int) (*Language, error) {
	return c.Query().Where(language.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *LanguageClient) GetX(ctx context.Context, id int) *Language {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QuerySubmit queries the submit edge of a Language.
func (c *LanguageClient) QuerySubmit(l *Language) *SubmitQuery {
	query := (&SubmitClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := l.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(language.Table, language.FieldID, id),
			sqlgraph.To(submit.Table, submit.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, language.SubmitTable, language.SubmitColumn),
		)
		fromV = sqlgraph.Neighbors(l.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *LanguageClient) Hooks() []Hook {
	return c.hooks.Language
}

// Interceptors returns the client interceptors.
func (c *LanguageClient) Interceptors() []Interceptor {
	return c.inters.Language
}

func (c *LanguageClient) mutate(ctx context.Context, m *LanguageMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&LanguageCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&LanguageUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&LanguageUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&LanguageDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Language mutation op: %q", m.Op())
	}
}

// SubmitClient is a client for the Submit schema.
type SubmitClient struct {
	config
}

// NewSubmitClient returns a client for the Submit from the given config.
func NewSubmitClient(c config) *SubmitClient {
	return &SubmitClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `submit.Hooks(f(g(h())))`.
func (c *SubmitClient) Use(hooks ...Hook) {
	c.hooks.Submit = append(c.hooks.Submit, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `submit.Intercept(f(g(h())))`.
func (c *SubmitClient) Intercept(interceptors ...Interceptor) {
	c.inters.Submit = append(c.inters.Submit, interceptors...)
}

// Create returns a builder for creating a Submit entity.
func (c *SubmitClient) Create() *SubmitCreate {
	mutation := newSubmitMutation(c.config, OpCreate)
	return &SubmitCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of Submit entities.
func (c *SubmitClient) CreateBulk(builders ...*SubmitCreate) *SubmitCreateBulk {
	return &SubmitCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for Submit.
func (c *SubmitClient) Update() *SubmitUpdate {
	mutation := newSubmitMutation(c.config, OpUpdate)
	return &SubmitUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *SubmitClient) UpdateOne(s *Submit) *SubmitUpdateOne {
	mutation := newSubmitMutation(c.config, OpUpdateOne, withSubmit(s))
	return &SubmitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *SubmitClient) UpdateOneID(id int) *SubmitUpdateOne {
	mutation := newSubmitMutation(c.config, OpUpdateOne, withSubmitID(id))
	return &SubmitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for Submit.
func (c *SubmitClient) Delete() *SubmitDelete {
	mutation := newSubmitMutation(c.config, OpDelete)
	return &SubmitDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *SubmitClient) DeleteOne(s *Submit) *SubmitDeleteOne {
	return c.DeleteOneID(s.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *SubmitClient) DeleteOneID(id int) *SubmitDeleteOne {
	builder := c.Delete().Where(submit.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &SubmitDeleteOne{builder}
}

// Query returns a query builder for Submit.
func (c *SubmitClient) Query() *SubmitQuery {
	return &SubmitQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeSubmit},
		inters: c.Interceptors(),
	}
}

// Get returns a Submit entity by its id.
func (c *SubmitClient) Get(ctx context.Context, id int) (*Submit, error) {
	return c.Query().Where(submit.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *SubmitClient) GetX(ctx context.Context, id int) *Submit {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// QueryUser queries the user edge of a Submit.
func (c *SubmitClient) QueryUser(s *Submit) *UserQuery {
	query := (&UserClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(submit.Table, submit.FieldID, id),
			sqlgraph.To(user.Table, user.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, submit.UserTable, submit.UserColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryLanguage queries the language edge of a Submit.
func (c *SubmitClient) QueryLanguage(s *Submit) *LanguageQuery {
	query := (&LanguageClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(submit.Table, submit.FieldID, id),
			sqlgraph.To(language.Table, language.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, submit.LanguageTable, submit.LanguageColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// QueryTestcaseResult queries the testcase_result edge of a Submit.
func (c *SubmitClient) QueryTestcaseResult(s *Submit) *TestcaseResultQuery {
	query := (&TestcaseResultClient{config: c.config}).Query()
	query.path = func(context.Context) (fromV *sql.Selector, _ error) {
		id := s.ID
		step := sqlgraph.NewStep(
			sqlgraph.From(submit.Table, submit.FieldID, id),
			sqlgraph.To(testcaseresult.Table, testcaseresult.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, submit.TestcaseResultTable, submit.TestcaseResultColumn),
		)
		fromV = sqlgraph.Neighbors(s.driver.Dialect(), step)
		return fromV, nil
	}
	return query
}

// Hooks returns the client hooks.
func (c *SubmitClient) Hooks() []Hook {
	return c.hooks.Submit
}

// Interceptors returns the client interceptors.
func (c *SubmitClient) Interceptors() []Interceptor {
	return c.inters.Submit
}

func (c *SubmitClient) mutate(ctx context.Context, m *SubmitMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&SubmitCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&SubmitUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&SubmitUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&SubmitDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown Submit mutation op: %q", m.Op())
	}
}

// TestcaseResultClient is a client for the TestcaseResult schema.
type TestcaseResultClient struct {
	config
}

// NewTestcaseResultClient returns a client for the TestcaseResult from the given config.
func NewTestcaseResultClient(c config) *TestcaseResultClient {
	return &TestcaseResultClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `testcaseresult.Hooks(f(g(h())))`.
func (c *TestcaseResultClient) Use(hooks ...Hook) {
	c.hooks.TestcaseResult = append(c.hooks.TestcaseResult, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `testcaseresult.Intercept(f(g(h())))`.
func (c *TestcaseResultClient) Intercept(interceptors ...Interceptor) {
	c.inters.TestcaseResult = append(c.inters.TestcaseResult, interceptors...)
}

// Create returns a builder for creating a TestcaseResult entity.
func (c *TestcaseResultClient) Create() *TestcaseResultCreate {
	mutation := newTestcaseResultMutation(c.config, OpCreate)
	return &TestcaseResultCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of TestcaseResult entities.
func (c *TestcaseResultClient) CreateBulk(builders ...*TestcaseResultCreate) *TestcaseResultCreateBulk {
	return &TestcaseResultCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for TestcaseResult.
func (c *TestcaseResultClient) Update() *TestcaseResultUpdate {
	mutation := newTestcaseResultMutation(c.config, OpUpdate)
	return &TestcaseResultUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *TestcaseResultClient) UpdateOne(tr *TestcaseResult) *TestcaseResultUpdateOne {
	mutation := newTestcaseResultMutation(c.config, OpUpdateOne, withTestcaseResult(tr))
	return &TestcaseResultUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *TestcaseResultClient) UpdateOneID(id int) *TestcaseResultUpdateOne {
	mutation := newTestcaseResultMutation(c.config, OpUpdateOne, withTestcaseResultID(id))
	return &TestcaseResultUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for TestcaseResult.
func (c *TestcaseResultClient) Delete() *TestcaseResultDelete {
	mutation := newTestcaseResultMutation(c.config, OpDelete)
	return &TestcaseResultDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *TestcaseResultClient) DeleteOne(tr *TestcaseResult) *TestcaseResultDeleteOne {
	return c.DeleteOneID(tr.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *TestcaseResultClient) DeleteOneID(id int) *TestcaseResultDeleteOne {
	builder := c.Delete().Where(testcaseresult.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &TestcaseResultDeleteOne{builder}
}

// Query returns a query builder for TestcaseResult.
func (c *TestcaseResultClient) Query() *TestcaseResultQuery {
	return &TestcaseResultQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeTestcaseResult},
		inters: c.Interceptors(),
	}
}

// Get returns a TestcaseResult entity by its id.
func (c *TestcaseResultClient) Get(ctx context.Context, id int) (*TestcaseResult, error) {
	return c.Query().Where(testcaseresult.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *TestcaseResultClient) GetX(ctx context.Context, id int) *TestcaseResult {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *TestcaseResultClient) Hooks() []Hook {
	return c.hooks.TestcaseResult
}

// Interceptors returns the client interceptors.
func (c *TestcaseResultClient) Interceptors() []Interceptor {
	return c.inters.TestcaseResult
}

func (c *TestcaseResultClient) mutate(ctx context.Context, m *TestcaseResultMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&TestcaseResultCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&TestcaseResultUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&TestcaseResultUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&TestcaseResultDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown TestcaseResult mutation op: %q", m.Op())
	}
}

// UserClient is a client for the User schema.
type UserClient struct {
	config
}

// NewUserClient returns a client for the User from the given config.
func NewUserClient(c config) *UserClient {
	return &UserClient{config: c}
}

// Use adds a list of mutation hooks to the hooks stack.
// A call to `Use(f, g, h)` equals to `user.Hooks(f(g(h())))`.
func (c *UserClient) Use(hooks ...Hook) {
	c.hooks.User = append(c.hooks.User, hooks...)
}

// Intercept adds a list of query interceptors to the interceptors stack.
// A call to `Intercept(f, g, h)` equals to `user.Intercept(f(g(h())))`.
func (c *UserClient) Intercept(interceptors ...Interceptor) {
	c.inters.User = append(c.inters.User, interceptors...)
}

// Create returns a builder for creating a User entity.
func (c *UserClient) Create() *UserCreate {
	mutation := newUserMutation(c.config, OpCreate)
	return &UserCreate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// CreateBulk returns a builder for creating a bulk of User entities.
func (c *UserClient) CreateBulk(builders ...*UserCreate) *UserCreateBulk {
	return &UserCreateBulk{config: c.config, builders: builders}
}

// Update returns an update builder for User.
func (c *UserClient) Update() *UserUpdate {
	mutation := newUserMutation(c.config, OpUpdate)
	return &UserUpdate{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOne returns an update builder for the given entity.
func (c *UserClient) UpdateOne(u *User) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUser(u))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// UpdateOneID returns an update builder for the given id.
func (c *UserClient) UpdateOneID(id int) *UserUpdateOne {
	mutation := newUserMutation(c.config, OpUpdateOne, withUserID(id))
	return &UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// Delete returns a delete builder for User.
func (c *UserClient) Delete() *UserDelete {
	mutation := newUserMutation(c.config, OpDelete)
	return &UserDelete{config: c.config, hooks: c.Hooks(), mutation: mutation}
}

// DeleteOne returns a builder for deleting the given entity.
func (c *UserClient) DeleteOne(u *User) *UserDeleteOne {
	return c.DeleteOneID(u.ID)
}

// DeleteOneID returns a builder for deleting the given entity by its id.
func (c *UserClient) DeleteOneID(id int) *UserDeleteOne {
	builder := c.Delete().Where(user.ID(id))
	builder.mutation.id = &id
	builder.mutation.op = OpDeleteOne
	return &UserDeleteOne{builder}
}

// Query returns a query builder for User.
func (c *UserClient) Query() *UserQuery {
	return &UserQuery{
		config: c.config,
		ctx:    &QueryContext{Type: TypeUser},
		inters: c.Interceptors(),
	}
}

// Get returns a User entity by its id.
func (c *UserClient) Get(ctx context.Context, id int) (*User, error) {
	return c.Query().Where(user.ID(id)).Only(ctx)
}

// GetX is like Get, but panics if an error occurs.
func (c *UserClient) GetX(ctx context.Context, id int) *User {
	obj, err := c.Get(ctx, id)
	if err != nil {
		panic(err)
	}
	return obj
}

// Hooks returns the client hooks.
func (c *UserClient) Hooks() []Hook {
	return c.hooks.User
}

// Interceptors returns the client interceptors.
func (c *UserClient) Interceptors() []Interceptor {
	return c.inters.User
}

func (c *UserClient) mutate(ctx context.Context, m *UserMutation) (Value, error) {
	switch m.Op() {
	case OpCreate:
		return (&UserCreate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdate:
		return (&UserUpdate{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpUpdateOne:
		return (&UserUpdateOne{config: c.config, hooks: c.Hooks(), mutation: m}).Save(ctx)
	case OpDelete, OpDeleteOne:
		return (&UserDelete{config: c.config, hooks: c.Hooks(), mutation: m}).Exec(ctx)
	default:
		return nil, fmt.Errorf("ent: unknown User mutation op: %q", m.Op())
	}
}

// hooks and interceptors per client, for fast access.
type (
	hooks struct {
		Language, Submit, TestcaseResult, User []ent.Hook
	}
	inters struct {
		Language, Submit, TestcaseResult, User []ent.Interceptor
	}
)
