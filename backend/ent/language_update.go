// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/BenMeredithConsult/locagri-apps/ent/language"
	"github.com/BenMeredithConsult/locagri-apps/ent/predicate"
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

// SetUpdatedAt sets the "updated_at" field.
func (lu *LanguageUpdate) SetUpdatedAt(t time.Time) *LanguageUpdate {
	lu.mutation.SetUpdatedAt(t)
	return lu
}

// SetName sets the "name" field.
func (lu *LanguageUpdate) SetName(s string) *LanguageUpdate {
	lu.mutation.SetName(s)
	return lu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (lu *LanguageUpdate) SetNillableName(s *string) *LanguageUpdate {
	if s != nil {
		lu.SetName(*s)
	}
	return lu
}

// SetCode sets the "code" field.
func (lu *LanguageUpdate) SetCode(s string) *LanguageUpdate {
	lu.mutation.SetCode(s)
	return lu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (lu *LanguageUpdate) SetNillableCode(s *string) *LanguageUpdate {
	if s != nil {
		lu.SetCode(*s)
	}
	return lu
}

// Mutation returns the LanguageMutation object of the builder.
func (lu *LanguageUpdate) Mutation() *LanguageMutation {
	return lu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LanguageUpdate) Save(ctx context.Context) (int, error) {
	lu.defaults()
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

// defaults sets the default values of the builder before save.
func (lu *LanguageUpdate) defaults() {
	if _, ok := lu.mutation.UpdatedAt(); !ok {
		v := language.UpdateDefaultUpdatedAt()
		lu.mutation.SetUpdatedAt(v)
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
	if value, ok := lu.mutation.UpdatedAt(); ok {
		_spec.SetField(language.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := lu.mutation.Name(); ok {
		_spec.SetField(language.FieldName, field.TypeString, value)
	}
	if value, ok := lu.mutation.Code(); ok {
		_spec.SetField(language.FieldCode, field.TypeString, value)
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

// SetUpdatedAt sets the "updated_at" field.
func (luo *LanguageUpdateOne) SetUpdatedAt(t time.Time) *LanguageUpdateOne {
	luo.mutation.SetUpdatedAt(t)
	return luo
}

// SetName sets the "name" field.
func (luo *LanguageUpdateOne) SetName(s string) *LanguageUpdateOne {
	luo.mutation.SetName(s)
	return luo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (luo *LanguageUpdateOne) SetNillableName(s *string) *LanguageUpdateOne {
	if s != nil {
		luo.SetName(*s)
	}
	return luo
}

// SetCode sets the "code" field.
func (luo *LanguageUpdateOne) SetCode(s string) *LanguageUpdateOne {
	luo.mutation.SetCode(s)
	return luo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (luo *LanguageUpdateOne) SetNillableCode(s *string) *LanguageUpdateOne {
	if s != nil {
		luo.SetCode(*s)
	}
	return luo
}

// Mutation returns the LanguageMutation object of the builder.
func (luo *LanguageUpdateOne) Mutation() *LanguageMutation {
	return luo.mutation
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
	luo.defaults()
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

// defaults sets the default values of the builder before save.
func (luo *LanguageUpdateOne) defaults() {
	if _, ok := luo.mutation.UpdatedAt(); !ok {
		v := language.UpdateDefaultUpdatedAt()
		luo.mutation.SetUpdatedAt(v)
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
	if value, ok := luo.mutation.UpdatedAt(); ok {
		_spec.SetField(language.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := luo.mutation.Name(); ok {
		_spec.SetField(language.FieldName, field.TypeString, value)
	}
	if value, ok := luo.mutation.Code(); ok {
		_spec.SetField(language.FieldCode, field.TypeString, value)
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