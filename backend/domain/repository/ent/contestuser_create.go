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
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/contestuser"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/user"
)

// ContestUserCreate is the builder for creating a ContestUser entity.
type ContestUserCreate struct {
	config
	mutation *ContestUserMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetRole sets the "role" field.
func (cuc *ContestUserCreate) SetRole(s string) *ContestUserCreate {
	cuc.mutation.SetRole(s)
	return cuc
}

// SetContestID sets the "contest_id" field.
func (cuc *ContestUserCreate) SetContestID(i int) *ContestUserCreate {
	cuc.mutation.SetContestID(i)
	return cuc
}

// SetUserID sets the "user_id" field.
func (cuc *ContestUserCreate) SetUserID(i int) *ContestUserCreate {
	cuc.mutation.SetUserID(i)
	return cuc
}

// SetID sets the "id" field.
func (cuc *ContestUserCreate) SetID(i int) *ContestUserCreate {
	cuc.mutation.SetID(i)
	return cuc
}

// SetContest sets the "contest" edge to the Contest entity.
func (cuc *ContestUserCreate) SetContest(c *Contest) *ContestUserCreate {
	return cuc.SetContestID(c.ID)
}

// SetUser sets the "user" edge to the User entity.
func (cuc *ContestUserCreate) SetUser(u *User) *ContestUserCreate {
	return cuc.SetUserID(u.ID)
}

// Mutation returns the ContestUserMutation object of the builder.
func (cuc *ContestUserCreate) Mutation() *ContestUserMutation {
	return cuc.mutation
}

// Save creates the ContestUser in the database.
func (cuc *ContestUserCreate) Save(ctx context.Context) (*ContestUser, error) {
	return withHooks(ctx, cuc.sqlSave, cuc.mutation, cuc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cuc *ContestUserCreate) SaveX(ctx context.Context) *ContestUser {
	v, err := cuc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cuc *ContestUserCreate) Exec(ctx context.Context) error {
	_, err := cuc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuc *ContestUserCreate) ExecX(ctx context.Context) {
	if err := cuc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuc *ContestUserCreate) check() error {
	if _, ok := cuc.mutation.Role(); !ok {
		return &ValidationError{Name: "role", err: errors.New(`ent: missing required field "ContestUser.role"`)}
	}
	if _, ok := cuc.mutation.ContestID(); !ok {
		return &ValidationError{Name: "contest_id", err: errors.New(`ent: missing required field "ContestUser.contest_id"`)}
	}
	if _, ok := cuc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "ContestUser.user_id"`)}
	}
	if _, ok := cuc.mutation.ContestID(); !ok {
		return &ValidationError{Name: "contest", err: errors.New(`ent: missing required edge "ContestUser.contest"`)}
	}
	if _, ok := cuc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "ContestUser.user"`)}
	}
	return nil
}

func (cuc *ContestUserCreate) sqlSave(ctx context.Context) (*ContestUser, error) {
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

func (cuc *ContestUserCreate) createSpec() (*ContestUser, *sqlgraph.CreateSpec) {
	var (
		_node = &ContestUser{config: cuc.config}
		_spec = sqlgraph.NewCreateSpec(contestuser.Table, sqlgraph.NewFieldSpec(contestuser.FieldID, field.TypeInt))
	)
	_spec.OnConflict = cuc.conflict
	if id, ok := cuc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cuc.mutation.Role(); ok {
		_spec.SetField(contestuser.FieldRole, field.TypeString, value)
		_node.Role = value
	}
	if nodes := cuc.mutation.ContestIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   contestuser.ContestTable,
			Columns: []string{contestuser.ContestColumn},
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
	if nodes := cuc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   contestuser.UserTable,
			Columns: []string{contestuser.UserColumn},
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
//	client.ContestUser.Create().
//		SetRole(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ContestUserUpsert) {
//			SetRole(v+v).
//		}).
//		Exec(ctx)
func (cuc *ContestUserCreate) OnConflict(opts ...sql.ConflictOption) *ContestUserUpsertOne {
	cuc.conflict = opts
	return &ContestUserUpsertOne{
		create: cuc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ContestUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cuc *ContestUserCreate) OnConflictColumns(columns ...string) *ContestUserUpsertOne {
	cuc.conflict = append(cuc.conflict, sql.ConflictColumns(columns...))
	return &ContestUserUpsertOne{
		create: cuc,
	}
}

type (
	// ContestUserUpsertOne is the builder for "upsert"-ing
	//  one ContestUser node.
	ContestUserUpsertOne struct {
		create *ContestUserCreate
	}

	// ContestUserUpsert is the "OnConflict" setter.
	ContestUserUpsert struct {
		*sql.UpdateSet
	}
)

// SetRole sets the "role" field.
func (u *ContestUserUpsert) SetRole(v string) *ContestUserUpsert {
	u.Set(contestuser.FieldRole, v)
	return u
}

// UpdateRole sets the "role" field to the value that was provided on create.
func (u *ContestUserUpsert) UpdateRole() *ContestUserUpsert {
	u.SetExcluded(contestuser.FieldRole)
	return u
}

// SetContestID sets the "contest_id" field.
func (u *ContestUserUpsert) SetContestID(v int) *ContestUserUpsert {
	u.Set(contestuser.FieldContestID, v)
	return u
}

// UpdateContestID sets the "contest_id" field to the value that was provided on create.
func (u *ContestUserUpsert) UpdateContestID() *ContestUserUpsert {
	u.SetExcluded(contestuser.FieldContestID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *ContestUserUpsert) SetUserID(v int) *ContestUserUpsert {
	u.Set(contestuser.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ContestUserUpsert) UpdateUserID() *ContestUserUpsert {
	u.SetExcluded(contestuser.FieldUserID)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ContestUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(contestuser.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ContestUserUpsertOne) UpdateNewValues() *ContestUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(contestuser.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ContestUser.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ContestUserUpsertOne) Ignore() *ContestUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ContestUserUpsertOne) DoNothing() *ContestUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ContestUserCreate.OnConflict
// documentation for more info.
func (u *ContestUserUpsertOne) Update(set func(*ContestUserUpsert)) *ContestUserUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ContestUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetRole sets the "role" field.
func (u *ContestUserUpsertOne) SetRole(v string) *ContestUserUpsertOne {
	return u.Update(func(s *ContestUserUpsert) {
		s.SetRole(v)
	})
}

// UpdateRole sets the "role" field to the value that was provided on create.
func (u *ContestUserUpsertOne) UpdateRole() *ContestUserUpsertOne {
	return u.Update(func(s *ContestUserUpsert) {
		s.UpdateRole()
	})
}

// SetContestID sets the "contest_id" field.
func (u *ContestUserUpsertOne) SetContestID(v int) *ContestUserUpsertOne {
	return u.Update(func(s *ContestUserUpsert) {
		s.SetContestID(v)
	})
}

// UpdateContestID sets the "contest_id" field to the value that was provided on create.
func (u *ContestUserUpsertOne) UpdateContestID() *ContestUserUpsertOne {
	return u.Update(func(s *ContestUserUpsert) {
		s.UpdateContestID()
	})
}

// SetUserID sets the "user_id" field.
func (u *ContestUserUpsertOne) SetUserID(v int) *ContestUserUpsertOne {
	return u.Update(func(s *ContestUserUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ContestUserUpsertOne) UpdateUserID() *ContestUserUpsertOne {
	return u.Update(func(s *ContestUserUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *ContestUserUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ContestUserCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ContestUserUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ContestUserUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ContestUserUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ContestUserCreateBulk is the builder for creating many ContestUser entities in bulk.
type ContestUserCreateBulk struct {
	config
	builders []*ContestUserCreate
	conflict []sql.ConflictOption
}

// Save creates the ContestUser entities in the database.
func (cucb *ContestUserCreateBulk) Save(ctx context.Context) ([]*ContestUser, error) {
	specs := make([]*sqlgraph.CreateSpec, len(cucb.builders))
	nodes := make([]*ContestUser, len(cucb.builders))
	mutators := make([]Mutator, len(cucb.builders))
	for i := range cucb.builders {
		func(i int, root context.Context) {
			builder := cucb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ContestUserMutation)
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
func (cucb *ContestUserCreateBulk) SaveX(ctx context.Context) []*ContestUser {
	v, err := cucb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cucb *ContestUserCreateBulk) Exec(ctx context.Context) error {
	_, err := cucb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cucb *ContestUserCreateBulk) ExecX(ctx context.Context) {
	if err := cucb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ContestUser.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ContestUserUpsert) {
//			SetRole(v+v).
//		}).
//		Exec(ctx)
func (cucb *ContestUserCreateBulk) OnConflict(opts ...sql.ConflictOption) *ContestUserUpsertBulk {
	cucb.conflict = opts
	return &ContestUserUpsertBulk{
		create: cucb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ContestUser.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (cucb *ContestUserCreateBulk) OnConflictColumns(columns ...string) *ContestUserUpsertBulk {
	cucb.conflict = append(cucb.conflict, sql.ConflictColumns(columns...))
	return &ContestUserUpsertBulk{
		create: cucb,
	}
}

// ContestUserUpsertBulk is the builder for "upsert"-ing
// a bulk of ContestUser nodes.
type ContestUserUpsertBulk struct {
	create *ContestUserCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ContestUser.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(contestuser.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ContestUserUpsertBulk) UpdateNewValues() *ContestUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(contestuser.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ContestUser.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ContestUserUpsertBulk) Ignore() *ContestUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ContestUserUpsertBulk) DoNothing() *ContestUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ContestUserCreateBulk.OnConflict
// documentation for more info.
func (u *ContestUserUpsertBulk) Update(set func(*ContestUserUpsert)) *ContestUserUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ContestUserUpsert{UpdateSet: update})
	}))
	return u
}

// SetRole sets the "role" field.
func (u *ContestUserUpsertBulk) SetRole(v string) *ContestUserUpsertBulk {
	return u.Update(func(s *ContestUserUpsert) {
		s.SetRole(v)
	})
}

// UpdateRole sets the "role" field to the value that was provided on create.
func (u *ContestUserUpsertBulk) UpdateRole() *ContestUserUpsertBulk {
	return u.Update(func(s *ContestUserUpsert) {
		s.UpdateRole()
	})
}

// SetContestID sets the "contest_id" field.
func (u *ContestUserUpsertBulk) SetContestID(v int) *ContestUserUpsertBulk {
	return u.Update(func(s *ContestUserUpsert) {
		s.SetContestID(v)
	})
}

// UpdateContestID sets the "contest_id" field to the value that was provided on create.
func (u *ContestUserUpsertBulk) UpdateContestID() *ContestUserUpsertBulk {
	return u.Update(func(s *ContestUserUpsert) {
		s.UpdateContestID()
	})
}

// SetUserID sets the "user_id" field.
func (u *ContestUserUpsertBulk) SetUserID(v int) *ContestUserUpsertBulk {
	return u.Update(func(s *ContestUserUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *ContestUserUpsertBulk) UpdateUserID() *ContestUserUpsertBulk {
	return u.Update(func(s *ContestUserUpsert) {
		s.UpdateUserID()
	})
}

// Exec executes the query.
func (u *ContestUserUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ContestUserCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ContestUserCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ContestUserUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
