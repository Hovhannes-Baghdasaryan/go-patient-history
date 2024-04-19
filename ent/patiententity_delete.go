// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"github.com/Hovhannes-Baghdasaryan/go-patient-history/ent/patiententity"
	"github.com/Hovhannes-Baghdasaryan/go-patient-history/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PatientEntityDelete is the builder for deleting a PatientEntity entity.
type PatientEntityDelete struct {
	config
	hooks    []Hook
	mutation *PatientEntityMutation
}

// Where appends a list predicates to the PatientEntityDelete builder.
func (ped *PatientEntityDelete) Where(ps ...predicate.PatientEntity) *PatientEntityDelete {
	ped.mutation.Where(ps...)
	return ped
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ped *PatientEntityDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, ped.sqlExec, ped.mutation, ped.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ped *PatientEntityDelete) ExecX(ctx context.Context) int {
	n, err := ped.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ped *PatientEntityDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(patiententity.Table, sqlgraph.NewFieldSpec(patiententity.FieldID, field.TypeUUID))
	if ps := ped.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ped.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ped.mutation.done = true
	return affected, err
}

// PatientEntityDeleteOne is the builder for deleting a single PatientEntity entity.
type PatientEntityDeleteOne struct {
	ped *PatientEntityDelete
}

// Where appends a list predicates to the PatientEntityDelete builder.
func (pedo *PatientEntityDeleteOne) Where(ps ...predicate.PatientEntity) *PatientEntityDeleteOne {
	pedo.ped.mutation.Where(ps...)
	return pedo
}

// Exec executes the deletion query.
func (pedo *PatientEntityDeleteOne) Exec(ctx context.Context) error {
	n, err := pedo.ped.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{patiententity.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pedo *PatientEntityDeleteOne) ExecX(ctx context.Context) {
	if err := pedo.Exec(ctx); err != nil {
		panic(err)
	}
}