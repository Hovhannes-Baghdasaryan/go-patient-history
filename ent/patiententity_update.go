// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"github.com/Hovhannes-Baghdasaryan/go-patient-history/ent/patiententity"
	"github.com/Hovhannes-Baghdasaryan/go-patient-history/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PatientEntityUpdate is the builder for updating PatientEntity entities.
type PatientEntityUpdate struct {
	config
	hooks    []Hook
	mutation *PatientEntityMutation
}

// Where appends a list predicates to the PatientEntityUpdate builder.
func (peu *PatientEntityUpdate) Where(ps ...predicate.PatientEntity) *PatientEntityUpdate {
	peu.mutation.Where(ps...)
	return peu
}

// SetName sets the "name" field.
func (peu *PatientEntityUpdate) SetName(s string) *PatientEntityUpdate {
	peu.mutation.SetName(s)
	return peu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (peu *PatientEntityUpdate) SetNillableName(s *string) *PatientEntityUpdate {
	if s != nil {
		peu.SetName(*s)
	}
	return peu
}

// SetSurname sets the "surname" field.
func (peu *PatientEntityUpdate) SetSurname(s string) *PatientEntityUpdate {
	peu.mutation.SetSurname(s)
	return peu
}

// SetNillableSurname sets the "surname" field if the given value is not nil.
func (peu *PatientEntityUpdate) SetNillableSurname(s *string) *PatientEntityUpdate {
	if s != nil {
		peu.SetSurname(*s)
	}
	return peu
}

// SetPatronymic sets the "patronymic" field.
func (peu *PatientEntityUpdate) SetPatronymic(s string) *PatientEntityUpdate {
	peu.mutation.SetPatronymic(s)
	return peu
}

// SetNillablePatronymic sets the "patronymic" field if the given value is not nil.
func (peu *PatientEntityUpdate) SetNillablePatronymic(s *string) *PatientEntityUpdate {
	if s != nil {
		peu.SetPatronymic(*s)
	}
	return peu
}

// SetAge sets the "age" field.
func (peu *PatientEntityUpdate) SetAge(i int) *PatientEntityUpdate {
	peu.mutation.ResetAge()
	peu.mutation.SetAge(i)
	return peu
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (peu *PatientEntityUpdate) SetNillableAge(i *int) *PatientEntityUpdate {
	if i != nil {
		peu.SetAge(*i)
	}
	return peu
}

// AddAge adds i to the "age" field.
func (peu *PatientEntityUpdate) AddAge(i int) *PatientEntityUpdate {
	peu.mutation.AddAge(i)
	return peu
}

// SetGender sets the "gender" field.
func (peu *PatientEntityUpdate) SetGender(pa patiententity.Gender) *PatientEntityUpdate {
	peu.mutation.SetGender(pa)
	return peu
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (peu *PatientEntityUpdate) SetNillableGender(pa *patiententity.Gender) *PatientEntityUpdate {
	if pa != nil {
		peu.SetGender(*pa)
	}
	return peu
}

// SetCountry sets the "country" field.
func (peu *PatientEntityUpdate) SetCountry(s string) *PatientEntityUpdate {
	peu.mutation.SetCountry(s)
	return peu
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (peu *PatientEntityUpdate) SetNillableCountry(s *string) *PatientEntityUpdate {
	if s != nil {
		peu.SetCountry(*s)
	}
	return peu
}

// Mutation returns the PatientEntityMutation object of the builder.
func (peu *PatientEntityUpdate) Mutation() *PatientEntityMutation {
	return peu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (peu *PatientEntityUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, peu.sqlSave, peu.mutation, peu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (peu *PatientEntityUpdate) SaveX(ctx context.Context) int {
	affected, err := peu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (peu *PatientEntityUpdate) Exec(ctx context.Context) error {
	_, err := peu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (peu *PatientEntityUpdate) ExecX(ctx context.Context) {
	if err := peu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (peu *PatientEntityUpdate) check() error {
	if v, ok := peu.mutation.Age(); ok {
		if err := patiententity.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "PatientEntity.age": %w`, err)}
		}
	}
	if v, ok := peu.mutation.Gender(); ok {
		if err := patiententity.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "PatientEntity.gender": %w`, err)}
		}
	}
	return nil
}

func (peu *PatientEntityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := peu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(patiententity.Table, patiententity.Columns, sqlgraph.NewFieldSpec(patiententity.FieldID, field.TypeUUID))
	if ps := peu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := peu.mutation.Name(); ok {
		_spec.SetField(patiententity.FieldName, field.TypeString, value)
	}
	if value, ok := peu.mutation.Surname(); ok {
		_spec.SetField(patiententity.FieldSurname, field.TypeString, value)
	}
	if value, ok := peu.mutation.Patronymic(); ok {
		_spec.SetField(patiententity.FieldPatronymic, field.TypeString, value)
	}
	if value, ok := peu.mutation.Age(); ok {
		_spec.SetField(patiententity.FieldAge, field.TypeInt, value)
	}
	if value, ok := peu.mutation.AddedAge(); ok {
		_spec.AddField(patiententity.FieldAge, field.TypeInt, value)
	}
	if value, ok := peu.mutation.Gender(); ok {
		_spec.SetField(patiententity.FieldGender, field.TypeEnum, value)
	}
	if value, ok := peu.mutation.Country(); ok {
		_spec.SetField(patiententity.FieldCountry, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, peu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patiententity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	peu.mutation.done = true
	return n, nil
}

// PatientEntityUpdateOne is the builder for updating a single PatientEntity entity.
type PatientEntityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PatientEntityMutation
}

// SetName sets the "name" field.
func (peuo *PatientEntityUpdateOne) SetName(s string) *PatientEntityUpdateOne {
	peuo.mutation.SetName(s)
	return peuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (peuo *PatientEntityUpdateOne) SetNillableName(s *string) *PatientEntityUpdateOne {
	if s != nil {
		peuo.SetName(*s)
	}
	return peuo
}

// SetSurname sets the "surname" field.
func (peuo *PatientEntityUpdateOne) SetSurname(s string) *PatientEntityUpdateOne {
	peuo.mutation.SetSurname(s)
	return peuo
}

// SetNillableSurname sets the "surname" field if the given value is not nil.
func (peuo *PatientEntityUpdateOne) SetNillableSurname(s *string) *PatientEntityUpdateOne {
	if s != nil {
		peuo.SetSurname(*s)
	}
	return peuo
}

// SetPatronymic sets the "patronymic" field.
func (peuo *PatientEntityUpdateOne) SetPatronymic(s string) *PatientEntityUpdateOne {
	peuo.mutation.SetPatronymic(s)
	return peuo
}

// SetNillablePatronymic sets the "patronymic" field if the given value is not nil.
func (peuo *PatientEntityUpdateOne) SetNillablePatronymic(s *string) *PatientEntityUpdateOne {
	if s != nil {
		peuo.SetPatronymic(*s)
	}
	return peuo
}

// SetAge sets the "age" field.
func (peuo *PatientEntityUpdateOne) SetAge(i int) *PatientEntityUpdateOne {
	peuo.mutation.ResetAge()
	peuo.mutation.SetAge(i)
	return peuo
}

// SetNillableAge sets the "age" field if the given value is not nil.
func (peuo *PatientEntityUpdateOne) SetNillableAge(i *int) *PatientEntityUpdateOne {
	if i != nil {
		peuo.SetAge(*i)
	}
	return peuo
}

// AddAge adds i to the "age" field.
func (peuo *PatientEntityUpdateOne) AddAge(i int) *PatientEntityUpdateOne {
	peuo.mutation.AddAge(i)
	return peuo
}

// SetGender sets the "gender" field.
func (peuo *PatientEntityUpdateOne) SetGender(pa patiententity.Gender) *PatientEntityUpdateOne {
	peuo.mutation.SetGender(pa)
	return peuo
}

// SetNillableGender sets the "gender" field if the given value is not nil.
func (peuo *PatientEntityUpdateOne) SetNillableGender(pa *patiententity.Gender) *PatientEntityUpdateOne {
	if pa != nil {
		peuo.SetGender(*pa)
	}
	return peuo
}

// SetCountry sets the "country" field.
func (peuo *PatientEntityUpdateOne) SetCountry(s string) *PatientEntityUpdateOne {
	peuo.mutation.SetCountry(s)
	return peuo
}

// SetNillableCountry sets the "country" field if the given value is not nil.
func (peuo *PatientEntityUpdateOne) SetNillableCountry(s *string) *PatientEntityUpdateOne {
	if s != nil {
		peuo.SetCountry(*s)
	}
	return peuo
}

// Mutation returns the PatientEntityMutation object of the builder.
func (peuo *PatientEntityUpdateOne) Mutation() *PatientEntityMutation {
	return peuo.mutation
}

// Where appends a list predicates to the PatientEntityUpdate builder.
func (peuo *PatientEntityUpdateOne) Where(ps ...predicate.PatientEntity) *PatientEntityUpdateOne {
	peuo.mutation.Where(ps...)
	return peuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (peuo *PatientEntityUpdateOne) Select(field string, fields ...string) *PatientEntityUpdateOne {
	peuo.fields = append([]string{field}, fields...)
	return peuo
}

// Save executes the query and returns the updated PatientEntity entity.
func (peuo *PatientEntityUpdateOne) Save(ctx context.Context) (*PatientEntity, error) {
	return withHooks(ctx, peuo.sqlSave, peuo.mutation, peuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (peuo *PatientEntityUpdateOne) SaveX(ctx context.Context) *PatientEntity {
	node, err := peuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (peuo *PatientEntityUpdateOne) Exec(ctx context.Context) error {
	_, err := peuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (peuo *PatientEntityUpdateOne) ExecX(ctx context.Context) {
	if err := peuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (peuo *PatientEntityUpdateOne) check() error {
	if v, ok := peuo.mutation.Age(); ok {
		if err := patiententity.AgeValidator(v); err != nil {
			return &ValidationError{Name: "age", err: fmt.Errorf(`ent: validator failed for field "PatientEntity.age": %w`, err)}
		}
	}
	if v, ok := peuo.mutation.Gender(); ok {
		if err := patiententity.GenderValidator(v); err != nil {
			return &ValidationError{Name: "gender", err: fmt.Errorf(`ent: validator failed for field "PatientEntity.gender": %w`, err)}
		}
	}
	return nil
}

func (peuo *PatientEntityUpdateOne) sqlSave(ctx context.Context) (_node *PatientEntity, err error) {
	if err := peuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(patiententity.Table, patiententity.Columns, sqlgraph.NewFieldSpec(patiententity.FieldID, field.TypeUUID))
	id, ok := peuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PatientEntity.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := peuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, patiententity.FieldID)
		for _, f := range fields {
			if !patiententity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != patiententity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := peuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := peuo.mutation.Name(); ok {
		_spec.SetField(patiententity.FieldName, field.TypeString, value)
	}
	if value, ok := peuo.mutation.Surname(); ok {
		_spec.SetField(patiententity.FieldSurname, field.TypeString, value)
	}
	if value, ok := peuo.mutation.Patronymic(); ok {
		_spec.SetField(patiententity.FieldPatronymic, field.TypeString, value)
	}
	if value, ok := peuo.mutation.Age(); ok {
		_spec.SetField(patiententity.FieldAge, field.TypeInt, value)
	}
	if value, ok := peuo.mutation.AddedAge(); ok {
		_spec.AddField(patiententity.FieldAge, field.TypeInt, value)
	}
	if value, ok := peuo.mutation.Gender(); ok {
		_spec.SetField(patiententity.FieldGender, field.TypeEnum, value)
	}
	if value, ok := peuo.mutation.Country(); ok {
		_spec.SetField(patiententity.FieldCountry, field.TypeString, value)
	}
	_node = &PatientEntity{config: peuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, peuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{patiententity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	peuo.mutation.done = true
	return _node, nil
}