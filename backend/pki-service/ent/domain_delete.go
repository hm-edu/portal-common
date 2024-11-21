// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/hm-edu/pki-service/ent/domain"
	"github.com/hm-edu/pki-service/ent/predicate"
)

// DomainDelete is the builder for deleting a Domain entity.
type DomainDelete struct {
	config
	hooks    []Hook
	mutation *DomainMutation
}

// Where appends a list predicates to the DomainDelete builder.
func (dd *DomainDelete) Where(ps ...predicate.Domain) *DomainDelete {
	dd.mutation.Where(ps...)
	return dd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dd *DomainDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, dd.sqlExec, dd.mutation, dd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (dd *DomainDelete) ExecX(ctx context.Context) int {
	n, err := dd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dd *DomainDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(domain.Table, sqlgraph.NewFieldSpec(domain.FieldID, field.TypeInt))
	if ps := dd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, dd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	dd.mutation.done = true
	return affected, err
}

// DomainDeleteOne is the builder for deleting a single Domain entity.
type DomainDeleteOne struct {
	dd *DomainDelete
}

// Where appends a list predicates to the DomainDelete builder.
func (ddo *DomainDeleteOne) Where(ps ...predicate.Domain) *DomainDeleteOne {
	ddo.dd.mutation.Where(ps...)
	return ddo
}

// Exec executes the deletion query.
func (ddo *DomainDeleteOne) Exec(ctx context.Context) error {
	n, err := ddo.dd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{domain.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ddo *DomainDeleteOne) ExecX(ctx context.Context) {
	if err := ddo.Exec(ctx); err != nil {
		panic(err)
	}
}
