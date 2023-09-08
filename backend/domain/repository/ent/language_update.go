// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/language"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/predicate"
	"github.com/szpp-dev-team/szpp-judge/backend/domain/repository/ent/submit"
)

// LanguageUpdate is the builder for updating Language entities.
type LanguageUpdate struct {
	config
	hooks    []Hook
	mutation *LanguageMutation
}

// Where appends a list predicates to the LanguageUpdate builder.
func (lu *LanguageUpdate) Where(ps ...predicate.Language) *LanguageUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetSubmitID sets the "submit" edge to the Submit entity by ID.
func (lu *LanguageUpdate) SetSubmitID(id int) *LanguageUpdate {
	lu.mutation.SetSubmitID(id)
	return lu
}

// SetNillableSubmitID sets the "submit" edge to the Submit entity by ID if the given value is not nil.
func (lu *LanguageUpdate) SetNillableSubmitID(id *int) *LanguageUpdate {
	if id != nil {
		lu = lu.SetSubmitID(*id)
	}
	return lu
}

// SetSubmit sets the "submit" edge to the Submit entity.
func (lu *LanguageUpdate) SetSubmit(s *Submit) *LanguageUpdate {
	return lu.SetSubmitID(s.ID)
}

// Mutation returns the LanguageMutation object of the builder.
func (lu *LanguageUpdate) Mutation() *LanguageMutation {
	return lu.mutation
}

// ClearSubmit clears the "submit" edge to the Submit entity.
func (lu *LanguageUpdate) ClearSubmit() *LanguageUpdate {
	lu.mutation.ClearSubmit()
	return lu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LanguageUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LanguageUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LanguageUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LanguageUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LanguageUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(language.Table, language.Columns, sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if lu.mutation.SubmitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   language.SubmitTable,
			Columns: []string{language.SubmitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.SubmitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   language.SubmitTable,
			Columns: []string{language.SubmitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{language.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LanguageUpdateOne is the builder for updating a single Language entity.
type LanguageUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LanguageMutation
}

// SetSubmitID sets the "submit" edge to the Submit entity by ID.
func (luo *LanguageUpdateOne) SetSubmitID(id int) *LanguageUpdateOne {
	luo.mutation.SetSubmitID(id)
	return luo
}

// SetNillableSubmitID sets the "submit" edge to the Submit entity by ID if the given value is not nil.
func (luo *LanguageUpdateOne) SetNillableSubmitID(id *int) *LanguageUpdateOne {
	if id != nil {
		luo = luo.SetSubmitID(*id)
	}
	return luo
}

// SetSubmit sets the "submit" edge to the Submit entity.
func (luo *LanguageUpdateOne) SetSubmit(s *Submit) *LanguageUpdateOne {
	return luo.SetSubmitID(s.ID)
}

// Mutation returns the LanguageMutation object of the builder.
func (luo *LanguageUpdateOne) Mutation() *LanguageMutation {
	return luo.mutation
}

// ClearSubmit clears the "submit" edge to the Submit entity.
func (luo *LanguageUpdateOne) ClearSubmit() *LanguageUpdateOne {
	luo.mutation.ClearSubmit()
	return luo
}

// Where appends a list predicates to the LanguageUpdate builder.
func (luo *LanguageUpdateOne) Where(ps ...predicate.Language) *LanguageUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LanguageUpdateOne) Select(field string, fields ...string) *LanguageUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Language entity.
func (luo *LanguageUpdateOne) Save(ctx context.Context) (*Language, error) {
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LanguageUpdateOne) SaveX(ctx context.Context) *Language {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LanguageUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LanguageUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LanguageUpdateOne) sqlSave(ctx context.Context) (_node *Language, err error) {
	_spec := sqlgraph.NewUpdateSpec(language.Table, language.Columns, sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Language.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, language.FieldID)
		for _, f := range fields {
			if !language.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != language.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if luo.mutation.SubmitCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   language.SubmitTable,
			Columns: []string{language.SubmitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.SubmitIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   language.SubmitTable,
			Columns: []string{language.SubmitColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(submit.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Language{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{language.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
