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
	"github.com/BenMeredithConsult/locagri-apps/ent/country"
	"github.com/BenMeredithConsult/locagri-apps/ent/predicate"
	"github.com/BenMeredithConsult/locagri-apps/ent/user"
)

// UserUpdate is the builder for updating User entities.
type UserUpdate struct {
	config
	hooks    []Hook
	mutation *UserMutation
}

// Where appends a list predicates to the UserUpdate builder.
func (uu *UserUpdate) Where(ps ...predicate.User) *UserUpdate {
	uu.mutation.Where(ps...)
	return uu
}

// SetUpdatedAt sets the "updated_at" field.
func (uu *UserUpdate) SetUpdatedAt(t time.Time) *UserUpdate {
	uu.mutation.SetUpdatedAt(t)
	return uu
}

// SetFirstName sets the "first_name" field.
func (uu *UserUpdate) SetFirstName(s string) *UserUpdate {
	uu.mutation.SetFirstName(s)
	return uu
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableFirstName(s *string) *UserUpdate {
	if s != nil {
		uu.SetFirstName(*s)
	}
	return uu
}

// SetLastName sets the "last_name" field.
func (uu *UserUpdate) SetLastName(s string) *UserUpdate {
	uu.mutation.SetLastName(s)
	return uu
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLastName(s *string) *UserUpdate {
	if s != nil {
		uu.SetLastName(*s)
	}
	return uu
}

// SetPhone sets the "phone" field.
func (uu *UserUpdate) SetPhone(s string) *UserUpdate {
	uu.mutation.SetPhone(s)
	return uu
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uu *UserUpdate) SetNillablePhone(s *string) *UserUpdate {
	if s != nil {
		uu.SetPhone(*s)
	}
	return uu
}

// SetRole sets the "role" field.
func (uu *UserUpdate) SetRole(u user.Role) *UserUpdate {
	uu.mutation.SetRole(u)
	return uu
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uu *UserUpdate) SetNillableRole(u *user.Role) *UserUpdate {
	if u != nil {
		uu.SetRole(*u)
	}
	return uu
}

// SetIDType sets the "id_type" field.
func (uu *UserUpdate) SetIDType(s string) *UserUpdate {
	uu.mutation.SetIDType(s)
	return uu
}

// SetNillableIDType sets the "id_type" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIDType(s *string) *UserUpdate {
	if s != nil {
		uu.SetIDType(*s)
	}
	return uu
}

// ClearIDType clears the value of the "id_type" field.
func (uu *UserUpdate) ClearIDType() *UserUpdate {
	uu.mutation.ClearIDType()
	return uu
}

// SetIDNumber sets the "id_number" field.
func (uu *UserUpdate) SetIDNumber(s string) *UserUpdate {
	uu.mutation.SetIDNumber(s)
	return uu
}

// SetNillableIDNumber sets the "id_number" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIDNumber(s *string) *UserUpdate {
	if s != nil {
		uu.SetIDNumber(*s)
	}
	return uu
}

// ClearIDNumber clears the value of the "id_number" field.
func (uu *UserUpdate) ClearIDNumber() *UserUpdate {
	uu.mutation.ClearIDNumber()
	return uu
}

// SetIDPhoto sets the "id_photo" field.
func (uu *UserUpdate) SetIDPhoto(s string) *UserUpdate {
	uu.mutation.SetIDPhoto(s)
	return uu
}

// SetNillableIDPhoto sets the "id_photo" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIDPhoto(s *string) *UserUpdate {
	if s != nil {
		uu.SetIDPhoto(*s)
	}
	return uu
}

// ClearIDPhoto clears the value of the "id_photo" field.
func (uu *UserUpdate) ClearIDPhoto() *UserUpdate {
	uu.mutation.ClearIDPhoto()
	return uu
}

// SetProfilePhoto sets the "profile_photo" field.
func (uu *UserUpdate) SetProfilePhoto(s string) *UserUpdate {
	uu.mutation.SetProfilePhoto(s)
	return uu
}

// SetNillableProfilePhoto sets the "profile_photo" field if the given value is not nil.
func (uu *UserUpdate) SetNillableProfilePhoto(s *string) *UserUpdate {
	if s != nil {
		uu.SetProfilePhoto(*s)
	}
	return uu
}

// ClearProfilePhoto clears the value of the "profile_photo" field.
func (uu *UserUpdate) ClearProfilePhoto() *UserUpdate {
	uu.mutation.ClearProfilePhoto()
	return uu
}

// SetAddress sets the "address" field.
func (uu *UserUpdate) SetAddress(s string) *UserUpdate {
	uu.mutation.SetAddress(s)
	return uu
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (uu *UserUpdate) SetNillableAddress(s *string) *UserUpdate {
	if s != nil {
		uu.SetAddress(*s)
	}
	return uu
}

// ClearAddress clears the value of the "address" field.
func (uu *UserUpdate) ClearAddress() *UserUpdate {
	uu.mutation.ClearAddress()
	return uu
}

// SetCity sets the "city" field.
func (uu *UserUpdate) SetCity(s string) *UserUpdate {
	uu.mutation.SetCity(s)
	return uu
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCity(s *string) *UserUpdate {
	if s != nil {
		uu.SetCity(*s)
	}
	return uu
}

// ClearCity clears the value of the "city" field.
func (uu *UserUpdate) ClearCity() *UserUpdate {
	uu.mutation.ClearCity()
	return uu
}

// SetNationality sets the "nationality" field.
func (uu *UserUpdate) SetNationality(s string) *UserUpdate {
	uu.mutation.SetNationality(s)
	return uu
}

// SetNillableNationality sets the "nationality" field if the given value is not nil.
func (uu *UserUpdate) SetNillableNationality(s *string) *UserUpdate {
	if s != nil {
		uu.SetNationality(*s)
	}
	return uu
}

// ClearNationality clears the value of the "nationality" field.
func (uu *UserUpdate) ClearNationality() *UserUpdate {
	uu.mutation.ClearNationality()
	return uu
}

// SetLanguage sets the "language" field.
func (uu *UserUpdate) SetLanguage(u user.Language) *UserUpdate {
	uu.mutation.SetLanguage(u)
	return uu
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (uu *UserUpdate) SetNillableLanguage(u *user.Language) *UserUpdate {
	if u != nil {
		uu.SetLanguage(*u)
	}
	return uu
}

// SetCountryID sets the "country_id" field.
func (uu *UserUpdate) SetCountryID(i int) *UserUpdate {
	uu.mutation.SetCountryID(i)
	return uu
}

// SetNillableCountryID sets the "country_id" field if the given value is not nil.
func (uu *UserUpdate) SetNillableCountryID(i *int) *UserUpdate {
	if i != nil {
		uu.SetCountryID(*i)
	}
	return uu
}

// SetIsWorker sets the "is_worker" field.
func (uu *UserUpdate) SetIsWorker(b bool) *UserUpdate {
	uu.mutation.SetIsWorker(b)
	return uu
}

// SetNillableIsWorker sets the "is_worker" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsWorker(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsWorker(*b)
	}
	return uu
}

// SetIsVerified sets the "is_verified" field.
func (uu *UserUpdate) SetIsVerified(b bool) *UserUpdate {
	uu.mutation.SetIsVerified(b)
	return uu
}

// SetNillableIsVerified sets the "is_verified" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsVerified(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsVerified(*b)
	}
	return uu
}

// SetIsBlocked sets the "is_blocked" field.
func (uu *UserUpdate) SetIsBlocked(b bool) *UserUpdate {
	uu.mutation.SetIsBlocked(b)
	return uu
}

// SetNillableIsBlocked sets the "is_blocked" field if the given value is not nil.
func (uu *UserUpdate) SetNillableIsBlocked(b *bool) *UserUpdate {
	if b != nil {
		uu.SetIsBlocked(*b)
	}
	return uu
}

// SetReason sets the "reason" field.
func (uu *UserUpdate) SetReason(s string) *UserUpdate {
	uu.mutation.SetReason(s)
	return uu
}

// SetNillableReason sets the "reason" field if the given value is not nil.
func (uu *UserUpdate) SetNillableReason(s *string) *UserUpdate {
	if s != nil {
		uu.SetReason(*s)
	}
	return uu
}

// ClearReason clears the value of the "reason" field.
func (uu *UserUpdate) ClearReason() *UserUpdate {
	uu.mutation.ClearReason()
	return uu
}

// SetCountry sets the "country" edge to the Country entity.
func (uu *UserUpdate) SetCountry(c *Country) *UserUpdate {
	return uu.SetCountryID(c.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uu *UserUpdate) Mutation() *UserMutation {
	return uu.mutation
}

// ClearCountry clears the "country" edge to the Country entity.
func (uu *UserUpdate) ClearCountry() *UserUpdate {
	uu.mutation.ClearCountry()
	return uu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (uu *UserUpdate) Save(ctx context.Context) (int, error) {
	uu.defaults()
	return withHooks(ctx, uu.sqlSave, uu.mutation, uu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uu *UserUpdate) SaveX(ctx context.Context) int {
	affected, err := uu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uu *UserUpdate) Exec(ctx context.Context) error {
	_, err := uu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uu *UserUpdate) ExecX(ctx context.Context) {
	if err := uu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uu *UserUpdate) defaults() {
	if _, ok := uu.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uu.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uu *UserUpdate) check() error {
	if v, ok := uu.mutation.FirstName(); ok {
		if err := user.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "User.first_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.LastName(); ok {
		if err := user.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "User.last_name": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Phone(); ok {
		if err := user.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "User.phone": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Role(); ok {
		if err := user.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "User.role": %w`, err)}
		}
	}
	if v, ok := uu.mutation.Language(); ok {
		if err := user.LanguageValidator(v); err != nil {
			return &ValidationError{Name: "language", err: fmt.Errorf(`ent: validator failed for field "User.language": %w`, err)}
		}
	}
	if uu.mutation.CountryCleared() && len(uu.mutation.CountryIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "User.country"`)
	}
	return nil
}

func (uu *UserUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := uu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	if ps := uu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uu.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uu.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uu.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if value, ok := uu.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if value, ok := uu.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeEnum, value)
	}
	if value, ok := uu.mutation.IDType(); ok {
		_spec.SetField(user.FieldIDType, field.TypeString, value)
	}
	if uu.mutation.IDTypeCleared() {
		_spec.ClearField(user.FieldIDType, field.TypeString)
	}
	if value, ok := uu.mutation.IDNumber(); ok {
		_spec.SetField(user.FieldIDNumber, field.TypeString, value)
	}
	if uu.mutation.IDNumberCleared() {
		_spec.ClearField(user.FieldIDNumber, field.TypeString)
	}
	if value, ok := uu.mutation.IDPhoto(); ok {
		_spec.SetField(user.FieldIDPhoto, field.TypeString, value)
	}
	if uu.mutation.IDPhotoCleared() {
		_spec.ClearField(user.FieldIDPhoto, field.TypeString)
	}
	if value, ok := uu.mutation.ProfilePhoto(); ok {
		_spec.SetField(user.FieldProfilePhoto, field.TypeString, value)
	}
	if uu.mutation.ProfilePhotoCleared() {
		_spec.ClearField(user.FieldProfilePhoto, field.TypeString)
	}
	if value, ok := uu.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
	}
	if uu.mutation.AddressCleared() {
		_spec.ClearField(user.FieldAddress, field.TypeString)
	}
	if value, ok := uu.mutation.City(); ok {
		_spec.SetField(user.FieldCity, field.TypeString, value)
	}
	if uu.mutation.CityCleared() {
		_spec.ClearField(user.FieldCity, field.TypeString)
	}
	if value, ok := uu.mutation.Nationality(); ok {
		_spec.SetField(user.FieldNationality, field.TypeString, value)
	}
	if uu.mutation.NationalityCleared() {
		_spec.ClearField(user.FieldNationality, field.TypeString)
	}
	if value, ok := uu.mutation.Language(); ok {
		_spec.SetField(user.FieldLanguage, field.TypeEnum, value)
	}
	if value, ok := uu.mutation.IsWorker(); ok {
		_spec.SetField(user.FieldIsWorker, field.TypeBool, value)
	}
	if value, ok := uu.mutation.IsVerified(); ok {
		_spec.SetField(user.FieldIsVerified, field.TypeBool, value)
	}
	if value, ok := uu.mutation.IsBlocked(); ok {
		_spec.SetField(user.FieldIsBlocked, field.TypeBool, value)
	}
	if value, ok := uu.mutation.Reason(); ok {
		_spec.SetField(user.FieldReason, field.TypeString, value)
	}
	if uu.mutation.ReasonCleared() {
		_spec.ClearField(user.FieldReason, field.TypeString)
	}
	if uu.mutation.CountryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.CountryTable,
			Columns: []string{user.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(country.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uu.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.CountryTable,
			Columns: []string{user.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(country.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, uu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	uu.mutation.done = true
	return n, nil
}

// UserUpdateOne is the builder for updating a single User entity.
type UserUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (uuo *UserUpdateOne) SetUpdatedAt(t time.Time) *UserUpdateOne {
	uuo.mutation.SetUpdatedAt(t)
	return uuo
}

// SetFirstName sets the "first_name" field.
func (uuo *UserUpdateOne) SetFirstName(s string) *UserUpdateOne {
	uuo.mutation.SetFirstName(s)
	return uuo
}

// SetNillableFirstName sets the "first_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableFirstName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetFirstName(*s)
	}
	return uuo
}

// SetLastName sets the "last_name" field.
func (uuo *UserUpdateOne) SetLastName(s string) *UserUpdateOne {
	uuo.mutation.SetLastName(s)
	return uuo
}

// SetNillableLastName sets the "last_name" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLastName(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetLastName(*s)
	}
	return uuo
}

// SetPhone sets the "phone" field.
func (uuo *UserUpdateOne) SetPhone(s string) *UserUpdateOne {
	uuo.mutation.SetPhone(s)
	return uuo
}

// SetNillablePhone sets the "phone" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillablePhone(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetPhone(*s)
	}
	return uuo
}

// SetRole sets the "role" field.
func (uuo *UserUpdateOne) SetRole(u user.Role) *UserUpdateOne {
	uuo.mutation.SetRole(u)
	return uuo
}

// SetNillableRole sets the "role" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableRole(u *user.Role) *UserUpdateOne {
	if u != nil {
		uuo.SetRole(*u)
	}
	return uuo
}

// SetIDType sets the "id_type" field.
func (uuo *UserUpdateOne) SetIDType(s string) *UserUpdateOne {
	uuo.mutation.SetIDType(s)
	return uuo
}

// SetNillableIDType sets the "id_type" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIDType(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetIDType(*s)
	}
	return uuo
}

// ClearIDType clears the value of the "id_type" field.
func (uuo *UserUpdateOne) ClearIDType() *UserUpdateOne {
	uuo.mutation.ClearIDType()
	return uuo
}

// SetIDNumber sets the "id_number" field.
func (uuo *UserUpdateOne) SetIDNumber(s string) *UserUpdateOne {
	uuo.mutation.SetIDNumber(s)
	return uuo
}

// SetNillableIDNumber sets the "id_number" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIDNumber(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetIDNumber(*s)
	}
	return uuo
}

// ClearIDNumber clears the value of the "id_number" field.
func (uuo *UserUpdateOne) ClearIDNumber() *UserUpdateOne {
	uuo.mutation.ClearIDNumber()
	return uuo
}

// SetIDPhoto sets the "id_photo" field.
func (uuo *UserUpdateOne) SetIDPhoto(s string) *UserUpdateOne {
	uuo.mutation.SetIDPhoto(s)
	return uuo
}

// SetNillableIDPhoto sets the "id_photo" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIDPhoto(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetIDPhoto(*s)
	}
	return uuo
}

// ClearIDPhoto clears the value of the "id_photo" field.
func (uuo *UserUpdateOne) ClearIDPhoto() *UserUpdateOne {
	uuo.mutation.ClearIDPhoto()
	return uuo
}

// SetProfilePhoto sets the "profile_photo" field.
func (uuo *UserUpdateOne) SetProfilePhoto(s string) *UserUpdateOne {
	uuo.mutation.SetProfilePhoto(s)
	return uuo
}

// SetNillableProfilePhoto sets the "profile_photo" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableProfilePhoto(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetProfilePhoto(*s)
	}
	return uuo
}

// ClearProfilePhoto clears the value of the "profile_photo" field.
func (uuo *UserUpdateOne) ClearProfilePhoto() *UserUpdateOne {
	uuo.mutation.ClearProfilePhoto()
	return uuo
}

// SetAddress sets the "address" field.
func (uuo *UserUpdateOne) SetAddress(s string) *UserUpdateOne {
	uuo.mutation.SetAddress(s)
	return uuo
}

// SetNillableAddress sets the "address" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableAddress(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetAddress(*s)
	}
	return uuo
}

// ClearAddress clears the value of the "address" field.
func (uuo *UserUpdateOne) ClearAddress() *UserUpdateOne {
	uuo.mutation.ClearAddress()
	return uuo
}

// SetCity sets the "city" field.
func (uuo *UserUpdateOne) SetCity(s string) *UserUpdateOne {
	uuo.mutation.SetCity(s)
	return uuo
}

// SetNillableCity sets the "city" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCity(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetCity(*s)
	}
	return uuo
}

// ClearCity clears the value of the "city" field.
func (uuo *UserUpdateOne) ClearCity() *UserUpdateOne {
	uuo.mutation.ClearCity()
	return uuo
}

// SetNationality sets the "nationality" field.
func (uuo *UserUpdateOne) SetNationality(s string) *UserUpdateOne {
	uuo.mutation.SetNationality(s)
	return uuo
}

// SetNillableNationality sets the "nationality" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableNationality(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetNationality(*s)
	}
	return uuo
}

// ClearNationality clears the value of the "nationality" field.
func (uuo *UserUpdateOne) ClearNationality() *UserUpdateOne {
	uuo.mutation.ClearNationality()
	return uuo
}

// SetLanguage sets the "language" field.
func (uuo *UserUpdateOne) SetLanguage(u user.Language) *UserUpdateOne {
	uuo.mutation.SetLanguage(u)
	return uuo
}

// SetNillableLanguage sets the "language" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableLanguage(u *user.Language) *UserUpdateOne {
	if u != nil {
		uuo.SetLanguage(*u)
	}
	return uuo
}

// SetCountryID sets the "country_id" field.
func (uuo *UserUpdateOne) SetCountryID(i int) *UserUpdateOne {
	uuo.mutation.SetCountryID(i)
	return uuo
}

// SetNillableCountryID sets the "country_id" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableCountryID(i *int) *UserUpdateOne {
	if i != nil {
		uuo.SetCountryID(*i)
	}
	return uuo
}

// SetIsWorker sets the "is_worker" field.
func (uuo *UserUpdateOne) SetIsWorker(b bool) *UserUpdateOne {
	uuo.mutation.SetIsWorker(b)
	return uuo
}

// SetNillableIsWorker sets the "is_worker" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsWorker(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsWorker(*b)
	}
	return uuo
}

// SetIsVerified sets the "is_verified" field.
func (uuo *UserUpdateOne) SetIsVerified(b bool) *UserUpdateOne {
	uuo.mutation.SetIsVerified(b)
	return uuo
}

// SetNillableIsVerified sets the "is_verified" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsVerified(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsVerified(*b)
	}
	return uuo
}

// SetIsBlocked sets the "is_blocked" field.
func (uuo *UserUpdateOne) SetIsBlocked(b bool) *UserUpdateOne {
	uuo.mutation.SetIsBlocked(b)
	return uuo
}

// SetNillableIsBlocked sets the "is_blocked" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableIsBlocked(b *bool) *UserUpdateOne {
	if b != nil {
		uuo.SetIsBlocked(*b)
	}
	return uuo
}

// SetReason sets the "reason" field.
func (uuo *UserUpdateOne) SetReason(s string) *UserUpdateOne {
	uuo.mutation.SetReason(s)
	return uuo
}

// SetNillableReason sets the "reason" field if the given value is not nil.
func (uuo *UserUpdateOne) SetNillableReason(s *string) *UserUpdateOne {
	if s != nil {
		uuo.SetReason(*s)
	}
	return uuo
}

// ClearReason clears the value of the "reason" field.
func (uuo *UserUpdateOne) ClearReason() *UserUpdateOne {
	uuo.mutation.ClearReason()
	return uuo
}

// SetCountry sets the "country" edge to the Country entity.
func (uuo *UserUpdateOne) SetCountry(c *Country) *UserUpdateOne {
	return uuo.SetCountryID(c.ID)
}

// Mutation returns the UserMutation object of the builder.
func (uuo *UserUpdateOne) Mutation() *UserMutation {
	return uuo.mutation
}

// ClearCountry clears the "country" edge to the Country entity.
func (uuo *UserUpdateOne) ClearCountry() *UserUpdateOne {
	uuo.mutation.ClearCountry()
	return uuo
}

// Where appends a list predicates to the UserUpdate builder.
func (uuo *UserUpdateOne) Where(ps ...predicate.User) *UserUpdateOne {
	uuo.mutation.Where(ps...)
	return uuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (uuo *UserUpdateOne) Select(field string, fields ...string) *UserUpdateOne {
	uuo.fields = append([]string{field}, fields...)
	return uuo
}

// Save executes the query and returns the updated User entity.
func (uuo *UserUpdateOne) Save(ctx context.Context) (*User, error) {
	uuo.defaults()
	return withHooks(ctx, uuo.sqlSave, uuo.mutation, uuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (uuo *UserUpdateOne) SaveX(ctx context.Context) *User {
	node, err := uuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (uuo *UserUpdateOne) Exec(ctx context.Context) error {
	_, err := uuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uuo *UserUpdateOne) ExecX(ctx context.Context) {
	if err := uuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (uuo *UserUpdateOne) defaults() {
	if _, ok := uuo.mutation.UpdatedAt(); !ok {
		v := user.UpdateDefaultUpdatedAt()
		uuo.mutation.SetUpdatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uuo *UserUpdateOne) check() error {
	if v, ok := uuo.mutation.FirstName(); ok {
		if err := user.FirstNameValidator(v); err != nil {
			return &ValidationError{Name: "first_name", err: fmt.Errorf(`ent: validator failed for field "User.first_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.LastName(); ok {
		if err := user.LastNameValidator(v); err != nil {
			return &ValidationError{Name: "last_name", err: fmt.Errorf(`ent: validator failed for field "User.last_name": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Phone(); ok {
		if err := user.PhoneValidator(v); err != nil {
			return &ValidationError{Name: "phone", err: fmt.Errorf(`ent: validator failed for field "User.phone": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Role(); ok {
		if err := user.RoleValidator(v); err != nil {
			return &ValidationError{Name: "role", err: fmt.Errorf(`ent: validator failed for field "User.role": %w`, err)}
		}
	}
	if v, ok := uuo.mutation.Language(); ok {
		if err := user.LanguageValidator(v); err != nil {
			return &ValidationError{Name: "language", err: fmt.Errorf(`ent: validator failed for field "User.language": %w`, err)}
		}
	}
	if uuo.mutation.CountryCleared() && len(uuo.mutation.CountryIDs()) > 0 {
		return errors.New(`ent: clearing a required unique edge "User.country"`)
	}
	return nil
}

func (uuo *UserUpdateOne) sqlSave(ctx context.Context) (_node *User, err error) {
	if err := uuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(user.Table, user.Columns, sqlgraph.NewFieldSpec(user.FieldID, field.TypeInt))
	id, ok := uuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "User.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := uuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, user.FieldID)
		for _, f := range fields {
			if !user.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != user.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := uuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uuo.mutation.UpdatedAt(); ok {
		_spec.SetField(user.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := uuo.mutation.FirstName(); ok {
		_spec.SetField(user.FieldFirstName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.LastName(); ok {
		_spec.SetField(user.FieldLastName, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Phone(); ok {
		_spec.SetField(user.FieldPhone, field.TypeString, value)
	}
	if value, ok := uuo.mutation.Role(); ok {
		_spec.SetField(user.FieldRole, field.TypeEnum, value)
	}
	if value, ok := uuo.mutation.IDType(); ok {
		_spec.SetField(user.FieldIDType, field.TypeString, value)
	}
	if uuo.mutation.IDTypeCleared() {
		_spec.ClearField(user.FieldIDType, field.TypeString)
	}
	if value, ok := uuo.mutation.IDNumber(); ok {
		_spec.SetField(user.FieldIDNumber, field.TypeString, value)
	}
	if uuo.mutation.IDNumberCleared() {
		_spec.ClearField(user.FieldIDNumber, field.TypeString)
	}
	if value, ok := uuo.mutation.IDPhoto(); ok {
		_spec.SetField(user.FieldIDPhoto, field.TypeString, value)
	}
	if uuo.mutation.IDPhotoCleared() {
		_spec.ClearField(user.FieldIDPhoto, field.TypeString)
	}
	if value, ok := uuo.mutation.ProfilePhoto(); ok {
		_spec.SetField(user.FieldProfilePhoto, field.TypeString, value)
	}
	if uuo.mutation.ProfilePhotoCleared() {
		_spec.ClearField(user.FieldProfilePhoto, field.TypeString)
	}
	if value, ok := uuo.mutation.Address(); ok {
		_spec.SetField(user.FieldAddress, field.TypeString, value)
	}
	if uuo.mutation.AddressCleared() {
		_spec.ClearField(user.FieldAddress, field.TypeString)
	}
	if value, ok := uuo.mutation.City(); ok {
		_spec.SetField(user.FieldCity, field.TypeString, value)
	}
	if uuo.mutation.CityCleared() {
		_spec.ClearField(user.FieldCity, field.TypeString)
	}
	if value, ok := uuo.mutation.Nationality(); ok {
		_spec.SetField(user.FieldNationality, field.TypeString, value)
	}
	if uuo.mutation.NationalityCleared() {
		_spec.ClearField(user.FieldNationality, field.TypeString)
	}
	if value, ok := uuo.mutation.Language(); ok {
		_spec.SetField(user.FieldLanguage, field.TypeEnum, value)
	}
	if value, ok := uuo.mutation.IsWorker(); ok {
		_spec.SetField(user.FieldIsWorker, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.IsVerified(); ok {
		_spec.SetField(user.FieldIsVerified, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.IsBlocked(); ok {
		_spec.SetField(user.FieldIsBlocked, field.TypeBool, value)
	}
	if value, ok := uuo.mutation.Reason(); ok {
		_spec.SetField(user.FieldReason, field.TypeString, value)
	}
	if uuo.mutation.ReasonCleared() {
		_spec.ClearField(user.FieldReason, field.TypeString)
	}
	if uuo.mutation.CountryCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.CountryTable,
			Columns: []string{user.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(country.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uuo.mutation.CountryIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   user.CountryTable,
			Columns: []string{user.CountryColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(country.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &User{config: uuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, uuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{user.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	uuo.mutation.done = true
	return _node, nil
}