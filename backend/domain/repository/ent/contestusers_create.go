// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contest"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contestusers"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
)

// ContestUsersCreate is the builder for creating a ContestUsers entity.
type ContestUsersCreate struct {
	config
	mutation *ContestUsersMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetRole sets the "role" field.
func (cuc *ContestUsersCreate) SetRole(s string) *ContestUsersCreate {
	cuc.mutation.SetRole(s)
	return cuc
}

// SetContestID sets the "contest_id" field.
func (cuc *ContestUsersCreate) SetContestID(i int) *ContestUsersCreate {
	cuc.mutation.SetContestID(i)
	return cuc
}

// SetUserID sets the "user_id" field.
func (cuc *ContestUsersCreate) SetUserID(i int) *ContestUsersCreate {
	cuc.mutation.SetUserID(i)
	return cuc
}

// SetID sets the "id" field.
func (cuc *ContestUsersCreate) SetID(i int) *ContestUsersCreate {
	cuc.mutation.SetID(i)
	return cuc
}

// SetContestsID sets the "contests" edge to the Contest entity by ID.
func (cuc *ContestUsersCreate) SetContestsID(id int) *ContestUsersCreate {
	cuc.mutation.SetContestsID(id)
	return cuc
}

// SetContests sets the "contests" edge to the Contest entity.
func (cuc *ContestUsersCreate) SetContests(c *Contest) *ContestUsersCreate {
	return cuc.SetContestsID(c.ID)
}

// SetUsersID sets the "users" edge to the User entity by ID.
func (cuc *ContestUsersCreate) SetUsersID(id int) *ContestUsersCreate {
	cuc.mutation.SetUsersID(id)
	return cuc
}

// SetUsers sets the "users" edge to the User entity.
func (cuc *ContestUsersCreate) SetUsers(u *User) *ContestUsersCreate {
	return cuc.SetUsersID(u.ID)
}

// Mutation returns the ContestUsersMutation object of the builder.
func (cuc *ContestUsersCreate) Mutation() *ContestUsersMutation {
	return cuc.mutation
}

// Save creates the ContestUsers in the database.
func (cuc *ContestUsersCreate) Save(ctx context.Context) (*ContestUsers, error) {
	return withHooks(ctx, cuc.sqlSave, cuc.mutation, cuc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cuc *ContestUsersCreate) SaveX(ctx context.Context) *ContestUsers {
	v, err := cuc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cuc *ContestUsersCreate) Exec(ctx context.Context) error {
	_, err := cuc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuc *ContestUsersCreate) ExecX(ctx context.Context) {
	if err := cuc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuc *ContestUsersCreate) check() error {
	if _, ok := cuc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "ContestUsers.role"`)}
	}
	if _, ok := cuc.mutation.ContestID(); !ok {
		return &ValidationError{Name: "contest_id", err: errors.New(`ent: missing required field "ContestUsers.contest_id"`)}
	}
	if _, ok := cuc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "ContestUsers.user_id"`)}
	}
	if _, ok := cuc.mutation.ContestsID(); !ok {
		return &ValidationError{Name: "contests", err: errors.New(`ent: missing required edge "ContestUsers.contests"`)}
	}
	if _, ok := cuc.mutation.UsersID(); !ok {
		return &ValidationError{Name: "users", err: errors.New(`ent: missing required edge "ContestUsers.users"`)}
	}
	return nil
}

func (cuc *ContestUsersCreate) sqlSave(ctx context.Context) (*ContestUsers, error) {
	if err := cuc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cuc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cuc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	cuc.mutation.id = &_node.ID
	cuc.mutation.done = true
	return _node, nil
}

func (cuc *ContestUsersCreate) createSpec() (*ContestUsers, *sqlgraph.CreateSpec) {
	var (
		_node = &ContestUsers{config: cuc.config}
		_spec = sqlgraph.NewCreateSpec(contestusers.Table, sqlgraph.NewFieldSpec(contestusers.FieldID, field.TypeInt))
	)
	_spec.OnConflict = cuc.conflict
	if id, ok := cuc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cuc.mutation.Role(); ok {
		_spec.SetField(contestusers.FieldRole, field.TypeString, value)
		_node.Role = value
	}
	if nodes := cuc.mutation.ContestsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   contestusers.ContestsTable,
			Columns: []string{contestusers.ContestsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(contest.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ContestID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cuc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   contestusers.UsersTable,
			Columns: []string{contestusers.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ContestUsers.Create().
//		SetRole(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ContestUsersUpsert) {
//			SetRole(v+v).
//		}).
//		Exec(ctx)
func (cuc *ContestUsersCreate) OnConflict(opts ...sql.ConflictOption) *ContestUsersUpsertOne {
	cuc.conflict = opts
	return &ContestUsersUpsertOne{
		create: cuc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ContestUsers.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cuc *ContestUsersCreate) OnConflictColumns(columns ...string) *ContestUsersUpsertOne {
	cuc.conflict = append(cuc.conflict, sql.ConflictColumns(columns...))
	return &ContestUsersUpsertOne{
		create: cuc,
	}
}

type (
	// ContestUsersUpsertOne is the builder for "upsert"-ing
	//  one ContestUsers node.
	ContestUsersUpsertOne struct {
		create *ContestUsersCreate
	}

	// ContestUsersUpsert is the "OnConflict" setter.
	ContestUsersUpsert struct {
		*sql.UpdateSet
	}
)

// SetRole sets the "role" field.
func (u *ContestUsersUpsert) SetRole(v string) *ContestUsersUpsert {
	u.Set(contestusers.FieldRole, v)
	return u
}

// UpdateRole sets the "role" field to the value that was provided on create.
func (u *ContestUsersUpsert) UpdateRole() *ContestUsersUpsert {
	u.SetExcluded(contestusers.FieldRole)
	return u
}

// SetContestID sets the "contest_id" field.
func (u *ContestUsersUpsert) SetContestID(v int) *ContestUsersUpsert {
	u.Set(contestusers.FieldContestID, v)
	return u
}

// UpdateContestID sets the "contest_id" field to the value that was provided on create.
func (u *ContestUsersUpsert) UpdateContestID() *ContestUsersUpsert {
	u.SetExcluded(contestusers.FieldContestID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *ContestUsersUpsert) SetUserID(v int) *ContestUsersUpsert {
	u.Set(contestusers.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ContestUsersUpsert) UpdateUserID() *ContestUsersUpsert {
	u.SetExcluded(contestusers.FieldUserID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ContestUsers.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(contestusers.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ContestUsersUpsertOne) UpdateNewValues() *ContestUsersUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(contestusers.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ContestUsers.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ContestUsersUpsertOne) Ignore() *ContestUsersUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ContestUsersUpsertOne) DoNothing() *ContestUsersUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ContestUsersCreate.OnConflict
// documentation for more info.
func (u *ContestUsersUpsertOne) Update(set func(*ContestUsersUpsert)) *ContestUsersUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ContestUsersUpsert{UpdateSet: update})
	}))
	return u
}

// SetRole sets the "role" field.
func (u *ContestUsersUpsertOne) SetRole(v string) *ContestUsersUpsertOne {
	return u.Update(func(s *ContestUsersUpsert) {
		s.SetRole(v)
	})
}

// UpdateRole sets the "role" field to the value that was provided on create.
func (u *ContestUsersUpsertOne) UpdateRole() *ContestUsersUpsertOne {
	return u.Update(func(s *ContestUsersUpsert) {
		s.UpdateRole()
	})
}

// SetContestID sets the "contest_id" field.
func (u *ContestUsersUpsertOne) SetContestID(v int) *ContestUsersUpsertOne {
	return u.Update(func(s *ContestUsersUpsert) {
		s.SetContestID(v)
	})
}

// UpdateContestID sets the "contest_id" field to the value that was provided on create.
func (u *ContestUsersUpsertOne) UpdateContestID() *ContestUsersUpsertOne {
	return u.Update(func(s *ContestUsersUpsert) {
		s.UpdateContestID()
	})
}

// SetUserID sets the "user_id" field.
func (u *ContestUsersUpsertOne) SetUserID(v int) *ContestUsersUpsertOne {
	return u.Update(func(s *ContestUsersUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ContestUsersUpsertOne) UpdateUserID() *ContestUsersUpsertOne {
	return u.Update(func(s *ContestUsersUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *ContestUsersUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ContestUsersCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ContestUsersUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ContestUsersUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ContestUsersUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ContestUsersCreateBulk is the builder for creating many ContestUsers entities in bulk.
type ContestUsersCreateBulk struct {
	config
	builders []*ContestUsersCreate
	conflict []sql.ConflictOption
}

// Save creates the ContestUsers entities in the database.
func (cucb *ContestUsersCreateBulk) Save(ctx context.Context) ([]*ContestUsers, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cucb.builders))
	nodes := make([]*ContestUsers, len(cucb.builders))
	mutators := make([]Mutator, len(cucb.builders))
	for i := range cucb.builders {
		func(i int, root context.Context) {
			builder := cucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContestUsersMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, cucb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = cucb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, cucb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil && nodes[i].ID == 0 {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, cucb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (cucb *ContestUsersCreateBulk) SaveX(ctx context.Context) []*ContestUsers {
	v, err := cucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cucb *ContestUsersCreateBulk) Exec(ctx context.Context) error {
	_, err := cucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cucb *ContestUsersCreateBulk) ExecX(ctx context.Context) {
	if err := cucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ContestUsers.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ContestUsersUpsert) {
//			SetRole(v+v).
//		}).
//		Exec(ctx)
func (cucb *ContestUsersCreateBulk) OnConflict(opts ...sql.ConflictOption) *ContestUsersUpsertBulk {
	cucb.conflict = opts
	return &ContestUsersUpsertBulk{
		create: cucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ContestUsers.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cucb *ContestUsersCreateBulk) OnConflictColumns(columns ...string) *ContestUsersUpsertBulk {
	cucb.conflict = append(cucb.conflict, sql.ConflictColumns(columns...))
	return &ContestUsersUpsertBulk{
		create: cucb,
	}
}

// ContestUsersUpsertBulk is the builder for "upsert"-ing
// a bulk of ContestUsers nodes.
type ContestUsersUpsertBulk struct {
	create *ContestUsersCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ContestUsers.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(contestusers.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ContestUsersUpsertBulk) UpdateNewValues() *ContestUsersUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(contestusers.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ContestUsers.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ContestUsersUpsertBulk) Ignore() *ContestUsersUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ContestUsersUpsertBulk) DoNothing() *ContestUsersUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ContestUsersCreateBulk.OnConflict
// documentation for more info.
func (u *ContestUsersUpsertBulk) Update(set func(*ContestUsersUpsert)) *ContestUsersUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ContestUsersUpsert{UpdateSet: update})
	}))
	return u
}

// SetRole sets the "role" field.
func (u *ContestUsersUpsertBulk) SetRole(v string) *ContestUsersUpsertBulk {
	return u.Update(func(s *ContestUsersUpsert) {
		s.SetRole(v)
	})
}

// UpdateRole sets the "role" field to the value that was provided on create.
func (u *ContestUsersUpsertBulk) UpdateRole() *ContestUsersUpsertBulk {
	return u.Update(func(s *ContestUsersUpsert) {
		s.UpdateRole()
	})
}

// SetContestID sets the "contest_id" field.
func (u *ContestUsersUpsertBulk) SetContestID(v int) *ContestUsersUpsertBulk {
	return u.Update(func(s *ContestUsersUpsert) {
		s.SetContestID(v)
	})
}

// UpdateContestID sets the "contest_id" field to the value that was provided on create.
func (u *ContestUsersUpsertBulk) UpdateContestID() *ContestUsersUpsertBulk {
	return u.Update(func(s *ContestUsersUpsert) {
		s.UpdateContestID()
	})
}

// SetUserID sets the "user_id" field.
func (u *ContestUsersUpsertBulk) SetUserID(v int) *ContestUsersUpsertBulk {
	return u.Update(func(s *ContestUsersUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ContestUsersUpsertBulk) UpdateUserID() *ContestUsersUpsertBulk {
	return u.Update(func(s *ContestUsersUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *ContestUsersUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ContestUsersCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ContestUsersCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ContestUsersUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
