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
	"github.com/BenMeredithConsult/locagri-apps/ent/nationality"
	"github.com/BenMeredithConsult/locagri-apps/ent/predicate"
)

// NationalityUpdate is the builder for updating Nationality entities.
type NationalityUpdate struct {
	config
	hooks    []Hook
	mutation *NationalityMutation
}

// Where appends a list predicates to the NationalityUpdate builder.
func (nu *NationalityUpdate) Where(ps ...predicate.Nationality) *NationalityUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetUpdatedAt sets the "updated_at" field.
func (nu *NationalityUpdate) SetUpdatedAt(t time.Time) *NationalityUpdate {
	nu.mutation.SetUpdatedAt(t)
	return nu
}

// SetName sets the "name" field.
func (nu *NationalityUpdate) SetName(s string) *NationalityUpdate {
	nu.mutation.SetName(s)
	return nu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (nu *NationalityUpdate) SetNillableName(s *string) *NationalityUpdate {
	if s != nil {
		nu.SetName(*s)
	}
	return nu
}

// SetCode sets the "code" field.
func (nu *NationalityUpdate) SetCode(s string) *NationalityUpdate {
	nu.mutation.SetCode(s)
	return nu
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (nu *NationalityUpdate) SetNillableCode(s *string) *NationalityUpdate {
	if s != nil {
		nu.SetCode(*s)
	}
	return nu
}

// Mutation returns the NationalityMutation object of the builder.
func (nu *NationalityUpdate) Mutation() *NationalityMutation {
	return nu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NationalityUpdate) Save(ctx context.Context) (int, error) {
	nu.defaults()
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NationalityUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NationalityUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NationalityUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nu *NationalityUpdate) defaults() {
	if _, ok := nu.mutation.UpdatedAt(); !ok {
		v := nationality.UpdateDefaultUpdatedAt()
		nu.mutation.SetUpdatedAt(v)
	}
}

func (nu *NationalityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(nationality.Table, nationality.Columns, sqlgraph.NewFieldSpec(nationality.FieldID, field.TypeInt))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.UpdatedAt(); ok {
		_spec.SetField(nationality.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := nu.mutation.Name(); ok {
		_spec.SetField(nationality.FieldName, field.TypeString, value)
	}
	if value, ok := nu.mutation.Code(); ok {
		_spec.SetField(nationality.FieldCode, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nationality.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NationalityUpdateOne is the builder for updating a single Nationality entity.
type NationalityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NationalityMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (nuo *NationalityUpdateOne) SetUpdatedAt(t time.Time) *NationalityUpdateOne {
	nuo.mutation.SetUpdatedAt(t)
	return nuo
}

// SetName sets the "name" field.
func (nuo *NationalityUpdateOne) SetName(s string) *NationalityUpdateOne {
	nuo.mutation.SetName(s)
	return nuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (nuo *NationalityUpdateOne) SetNillableName(s *string) *NationalityUpdateOne {
	if s != nil {
		nuo.SetName(*s)
	}
	return nuo
}

// SetCode sets the "code" field.
func (nuo *NationalityUpdateOne) SetCode(s string) *NationalityUpdateOne {
	nuo.mutation.SetCode(s)
	return nuo
}

// SetNillableCode sets the "code" field if the given value is not nil.
func (nuo *NationalityUpdateOne) SetNillableCode(s *string) *NationalityUpdateOne {
	if s != nil {
		nuo.SetCode(*s)
	}
	return nuo
}

// Mutation returns the NationalityMutation object of the builder.
func (nuo *NationalityUpdateOne) Mutation() *NationalityMutation {
	return nuo.mutation
}

// Where appends a list predicates to the NationalityUpdate builder.
func (nuo *NationalityUpdateOne) Where(ps ...predicate.Nationality) *NationalityUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NationalityUpdateOne) Select(field string, fields ...string) *NationalityUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated Nationality entity.
func (nuo *NationalityUpdateOne) Save(ctx context.Context) (*Nationality, error) {
	nuo.defaults()
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NationalityUpdateOne) SaveX(ctx context.Context) *Nationality {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NationalityUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NationalityUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (nuo *NationalityUpdateOne) defaults() {
	if _, ok := nuo.mutation.UpdatedAt(); !ok {
		v := nationality.UpdateDefaultUpdatedAt()
		nuo.mutation.SetUpdatedAt(v)
	}
}

func (nuo *NationalityUpdateOne) sqlSave(ctx context.Context) (_node *Nationality, err error) {
	_spec := sqlgraph.NewUpdateSpec(nationality.Table, nationality.Columns, sqlgraph.NewFieldSpec(nationality.FieldID, field.TypeInt))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Nationality.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, nationality.FieldID)
		for _, f := range fields {
			if !nationality.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != nationality.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.UpdatedAt(); ok {
		_spec.SetField(nationality.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := nuo.mutation.Name(); ok {
		_spec.SetField(nationality.FieldName, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Code(); ok {
		_spec.SetField(nationality.FieldCode, field.TypeString, value)
	}
	_node = &Nationality{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{nationality.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}
