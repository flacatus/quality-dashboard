// Code generated by ent, DO NOT EDIT.

package db

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/redhat-appstudio/quality-studio/pkg/storage/ent/db/plugins"
)

// PluginsCreate is the builder for creating a Plugins entity.
type PluginsCreate struct {
	config
	mutation *PluginsMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (pc *PluginsCreate) SetName(s string) *PluginsCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetCategory sets the "category" field.
func (pc *PluginsCreate) SetCategory(s string) *PluginsCreate {
	pc.mutation.SetCategory(s)
	return pc
}

// SetLogo sets the "logo" field.
func (pc *PluginsCreate) SetLogo(s string) *PluginsCreate {
	pc.mutation.SetLogo(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PluginsCreate) SetDescription(s string) *PluginsCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetStatus sets the "status" field.
func (pc *PluginsCreate) SetStatus(s string) *PluginsCreate {
	pc.mutation.SetStatus(s)
	return pc
}

// SetID sets the "id" field.
func (pc *PluginsCreate) SetID(u uuid.UUID) *PluginsCreate {
	pc.mutation.SetID(u)
	return pc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (pc *PluginsCreate) SetNillableID(u *uuid.UUID) *PluginsCreate {
	if u != nil {
		pc.SetID(*u)
	}
	return pc
}

// Mutation returns the PluginsMutation object of the builder.
func (pc *PluginsCreate) Mutation() *PluginsMutation {
	return pc.mutation
}

// Save creates the Plugins in the database.
func (pc *PluginsCreate) Save(ctx context.Context) (*Plugins, error) {
	pc.defaults()
	return withHooks[*Plugins, PluginsMutation](ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PluginsCreate) SaveX(ctx context.Context) *Plugins {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PluginsCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PluginsCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *PluginsCreate) defaults() {
	if _, ok := pc.mutation.ID(); !ok {
		v := plugins.DefaultID()
		pc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PluginsCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`db: missing required field "Plugins.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := plugins.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`db: validator failed for field "Plugins.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Category(); !ok {
		return &ValidationError{Name: "category", err: errors.New(`db: missing required field "Plugins.category"`)}
	}
	if v, ok := pc.mutation.Category(); ok {
		if err := plugins.CategoryValidator(v); err != nil {
			return &ValidationError{Name: "category", err: fmt.Errorf(`db: validator failed for field "Plugins.category": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Logo(); !ok {
		return &ValidationError{Name: "logo", err: errors.New(`db: missing required field "Plugins.logo"`)}
	}
	if v, ok := pc.mutation.Logo(); ok {
		if err := plugins.LogoValidator(v); err != nil {
			return &ValidationError{Name: "logo", err: fmt.Errorf(`db: validator failed for field "Plugins.logo": %w`, err)}
		}
	}
	if _, ok := pc.mutation.Description(); !ok {
		return &ValidationError{Name: "description", err: errors.New(`db: missing required field "Plugins.description"`)}
	}
	if _, ok := pc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`db: missing required field "Plugins.status"`)}
	}
	if v, ok := pc.mutation.Status(); ok {
		if err := plugins.StatusValidator(v); err != nil {
			return &ValidationError{Name: "status", err: fmt.Errorf(`db: validator failed for field "Plugins.status": %w`, err)}
		}
	}
	return nil
}

func (pc *PluginsCreate) sqlSave(ctx context.Context) (*Plugins, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PluginsCreate) createSpec() (*Plugins, *sqlgraph.CreateSpec) {
	var (
		_node = &Plugins{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: plugins.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: plugins.FieldID,
			},
		}
	)
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(plugins.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Category(); ok {
		_spec.SetField(plugins.FieldCategory, field.TypeString, value)
		_node.Category = value
	}
	if value, ok := pc.mutation.Logo(); ok {
		_spec.SetField(plugins.FieldLogo, field.TypeString, value)
		_node.Logo = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(plugins.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.Status(); ok {
		_spec.SetField(plugins.FieldStatus, field.TypeString, value)
		_node.Status = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Plugins.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PluginsUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (pc *PluginsCreate) OnConflict(opts ...sql.ConflictOption) *PluginsUpsertOne {
	pc.conflict = opts
	return &PluginsUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Plugins.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PluginsCreate) OnConflictColumns(columns ...string) *PluginsUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PluginsUpsertOne{
		create: pc,
	}
}

type (
	// PluginsUpsertOne is the builder for "upsert"-ing
	//  one Plugins node.
	PluginsUpsertOne struct {
		create *PluginsCreate
	}

	// PluginsUpsert is the "OnConflict" setter.
	PluginsUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *PluginsUpsert) SetName(v string) *PluginsUpsert {
	u.Set(plugins.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PluginsUpsert) UpdateName() *PluginsUpsert {
	u.SetExcluded(plugins.FieldName)
	return u
}

// SetCategory sets the "category" field.
func (u *PluginsUpsert) SetCategory(v string) *PluginsUpsert {
	u.Set(plugins.FieldCategory, v)
	return u
}

// UpdateCategory sets the "category" field to the value that was provided on create.
func (u *PluginsUpsert) UpdateCategory() *PluginsUpsert {
	u.SetExcluded(plugins.FieldCategory)
	return u
}

// SetLogo sets the "logo" field.
func (u *PluginsUpsert) SetLogo(v string) *PluginsUpsert {
	u.Set(plugins.FieldLogo, v)
	return u
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *PluginsUpsert) UpdateLogo() *PluginsUpsert {
	u.SetExcluded(plugins.FieldLogo)
	return u
}

// SetDescription sets the "description" field.
func (u *PluginsUpsert) SetDescription(v string) *PluginsUpsert {
	u.Set(plugins.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *PluginsUpsert) UpdateDescription() *PluginsUpsert {
	u.SetExcluded(plugins.FieldDescription)
	return u
}

// SetStatus sets the "status" field.
func (u *PluginsUpsert) SetStatus(v string) *PluginsUpsert {
	u.Set(plugins.FieldStatus, v)
	return u
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *PluginsUpsert) UpdateStatus() *PluginsUpsert {
	u.SetExcluded(plugins.FieldStatus)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Plugins.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(plugins.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PluginsUpsertOne) UpdateNewValues() *PluginsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(plugins.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Plugins.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PluginsUpsertOne) Ignore() *PluginsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PluginsUpsertOne) DoNothing() *PluginsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PluginsCreate.OnConflict
// documentation for more info.
func (u *PluginsUpsertOne) Update(set func(*PluginsUpsert)) *PluginsUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PluginsUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *PluginsUpsertOne) SetName(v string) *PluginsUpsertOne {
	return u.Update(func(s *PluginsUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PluginsUpsertOne) UpdateName() *PluginsUpsertOne {
	return u.Update(func(s *PluginsUpsert) {
		s.UpdateName()
	})
}

// SetCategory sets the "category" field.
func (u *PluginsUpsertOne) SetCategory(v string) *PluginsUpsertOne {
	return u.Update(func(s *PluginsUpsert) {
		s.SetCategory(v)
	})
}

// UpdateCategory sets the "category" field to the value that was provided on create.
func (u *PluginsUpsertOne) UpdateCategory() *PluginsUpsertOne {
	return u.Update(func(s *PluginsUpsert) {
		s.UpdateCategory()
	})
}

// SetLogo sets the "logo" field.
func (u *PluginsUpsertOne) SetLogo(v string) *PluginsUpsertOne {
	return u.Update(func(s *PluginsUpsert) {
		s.SetLogo(v)
	})
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *PluginsUpsertOne) UpdateLogo() *PluginsUpsertOne {
	return u.Update(func(s *PluginsUpsert) {
		s.UpdateLogo()
	})
}

// SetDescription sets the "description" field.
func (u *PluginsUpsertOne) SetDescription(v string) *PluginsUpsertOne {
	return u.Update(func(s *PluginsUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *PluginsUpsertOne) UpdateDescription() *PluginsUpsertOne {
	return u.Update(func(s *PluginsUpsert) {
		s.UpdateDescription()
	})
}

// SetStatus sets the "status" field.
func (u *PluginsUpsertOne) SetStatus(v string) *PluginsUpsertOne {
	return u.Update(func(s *PluginsUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *PluginsUpsertOne) UpdateStatus() *PluginsUpsertOne {
	return u.Update(func(s *PluginsUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *PluginsUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for PluginsCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PluginsUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PluginsUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("db: PluginsUpsertOne.ID is not supported by MySQL driver. Use PluginsUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PluginsUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PluginsCreateBulk is the builder for creating many Plugins entities in bulk.
type PluginsCreateBulk struct {
	config
	builders []*PluginsCreate
	conflict []sql.ConflictOption
}

// Save creates the Plugins entities in the database.
func (pcb *PluginsCreateBulk) Save(ctx context.Context) ([]*Plugins, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Plugins, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PluginsMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PluginsCreateBulk) SaveX(ctx context.Context) []*Plugins {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PluginsCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PluginsCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Plugins.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PluginsUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (pcb *PluginsCreateBulk) OnConflict(opts ...sql.ConflictOption) *PluginsUpsertBulk {
	pcb.conflict = opts
	return &PluginsUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Plugins.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PluginsCreateBulk) OnConflictColumns(columns ...string) *PluginsUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PluginsUpsertBulk{
		create: pcb,
	}
}

// PluginsUpsertBulk is the builder for "upsert"-ing
// a bulk of Plugins nodes.
type PluginsUpsertBulk struct {
	create *PluginsCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Plugins.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(plugins.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PluginsUpsertBulk) UpdateNewValues() *PluginsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(plugins.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Plugins.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PluginsUpsertBulk) Ignore() *PluginsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PluginsUpsertBulk) DoNothing() *PluginsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PluginsCreateBulk.OnConflict
// documentation for more info.
func (u *PluginsUpsertBulk) Update(set func(*PluginsUpsert)) *PluginsUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PluginsUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *PluginsUpsertBulk) SetName(v string) *PluginsUpsertBulk {
	return u.Update(func(s *PluginsUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *PluginsUpsertBulk) UpdateName() *PluginsUpsertBulk {
	return u.Update(func(s *PluginsUpsert) {
		s.UpdateName()
	})
}

// SetCategory sets the "category" field.
func (u *PluginsUpsertBulk) SetCategory(v string) *PluginsUpsertBulk {
	return u.Update(func(s *PluginsUpsert) {
		s.SetCategory(v)
	})
}

// UpdateCategory sets the "category" field to the value that was provided on create.
func (u *PluginsUpsertBulk) UpdateCategory() *PluginsUpsertBulk {
	return u.Update(func(s *PluginsUpsert) {
		s.UpdateCategory()
	})
}

// SetLogo sets the "logo" field.
func (u *PluginsUpsertBulk) SetLogo(v string) *PluginsUpsertBulk {
	return u.Update(func(s *PluginsUpsert) {
		s.SetLogo(v)
	})
}

// UpdateLogo sets the "logo" field to the value that was provided on create.
func (u *PluginsUpsertBulk) UpdateLogo() *PluginsUpsertBulk {
	return u.Update(func(s *PluginsUpsert) {
		s.UpdateLogo()
	})
}

// SetDescription sets the "description" field.
func (u *PluginsUpsertBulk) SetDescription(v string) *PluginsUpsertBulk {
	return u.Update(func(s *PluginsUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *PluginsUpsertBulk) UpdateDescription() *PluginsUpsertBulk {
	return u.Update(func(s *PluginsUpsert) {
		s.UpdateDescription()
	})
}

// SetStatus sets the "status" field.
func (u *PluginsUpsertBulk) SetStatus(v string) *PluginsUpsertBulk {
	return u.Update(func(s *PluginsUpsert) {
		s.SetStatus(v)
	})
}

// UpdateStatus sets the "status" field to the value that was provided on create.
func (u *PluginsUpsertBulk) UpdateStatus() *PluginsUpsertBulk {
	return u.Update(func(s *PluginsUpsert) {
		s.UpdateStatus()
	})
}

// Exec executes the query.
func (u *PluginsUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("db: OnConflict was set for builder %d. Set it on the PluginsCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("db: missing options for PluginsCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PluginsUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
