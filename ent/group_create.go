// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/tiagoposse/connect/ent/group"
	"github.com/tiagoposse/connect/ent/user"
	"github.com/tiagoposse/connect/internal/types"
	"github.com/tiagoposse/go-auth/authorization"
)

// GroupCreate is the builder for creating a Group entity.
type GroupCreate struct {
	config
	mutation *GroupMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetName sets the "name" field.
func (gc *GroupCreate) SetName(s string) *GroupCreate {
	gc.mutation.SetName(s)
	return gc
}

// SetScopes sets the "scopes" field.
func (gc *GroupCreate) SetScopes(a authorization.Scopes) *GroupCreate {
	gc.mutation.SetScopes(a)
	return gc
}

// SetCidr sets the "cidr" field.
func (gc *GroupCreate) SetCidr(t types.Cidr) *GroupCreate {
	gc.mutation.SetCidr(t)
	return gc
}

// SetRules sets the "rules" field.
func (gc *GroupCreate) SetRules(t []types.Rule) *GroupCreate {
	gc.mutation.SetRules(t)
	return gc
}

// SetID sets the "id" field.
func (gc *GroupCreate) SetID(s string) *GroupCreate {
	gc.mutation.SetID(s)
	return gc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (gc *GroupCreate) SetNillableID(s *string) *GroupCreate {
	if s != nil {
		gc.SetID(*s)
	}
	return gc
}

// AddUserIDs adds the "users" edge to the User entity by IDs.
func (gc *GroupCreate) AddUserIDs(ids ...string) *GroupCreate {
	gc.mutation.AddUserIDs(ids...)
	return gc
}

// AddUsers adds the "users" edges to the User entity.
func (gc *GroupCreate) AddUsers(u ...*User) *GroupCreate {
	ids := make([]string, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return gc.AddUserIDs(ids...)
}

// Mutation returns the GroupMutation object of the builder.
func (gc *GroupCreate) Mutation() *GroupMutation {
	return gc.mutation
}

// Save creates the Group in the database.
func (gc *GroupCreate) Save(ctx context.Context) (*Group, error) {
	gc.defaults()
	return withHooks(ctx, gc.sqlSave, gc.mutation, gc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (gc *GroupCreate) SaveX(ctx context.Context) *Group {
	v, err := gc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gc *GroupCreate) Exec(ctx context.Context) error {
	_, err := gc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gc *GroupCreate) ExecX(ctx context.Context) {
	if err := gc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (gc *GroupCreate) defaults() {
	if _, ok := gc.mutation.Scopes(); !ok {
		v := group.DefaultScopes
		gc.mutation.SetScopes(v)
	}
	if _, ok := gc.mutation.ID(); !ok {
		v := group.DefaultID()
		gc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (gc *GroupCreate) check() error {
	if _, ok := gc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Group.name"`)}
	}
	if _, ok := gc.mutation.Scopes(); !ok {
		return &ValidationError{Name: "scopes", err: errors.New(`ent: missing required field "Group.scopes"`)}
	}
	if _, ok := gc.mutation.Cidr(); !ok {
		return &ValidationError{Name: "cidr", err: errors.New(`ent: missing required field "Group.cidr"`)}
	}
	if _, ok := gc.mutation.Rules(); !ok {
		return &ValidationError{Name: "rules", err: errors.New(`ent: missing required field "Group.rules"`)}
	}
	return nil
}

func (gc *GroupCreate) sqlSave(ctx context.Context) (*Group, error) {
	if err := gc.check(); err != nil {
		return nil, err
	}
	_node, _spec := gc.createSpec()
	if err := sqlgraph.CreateNode(ctx, gc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Group.ID type: %T", _spec.ID.Value)
		}
	}
	gc.mutation.id = &_node.ID
	gc.mutation.done = true
	return _node, nil
}

func (gc *GroupCreate) createSpec() (*Group, *sqlgraph.CreateSpec) {
	var (
		_node = &Group{config: gc.config}
		_spec = sqlgraph.NewCreateSpec(group.Table, sqlgraph.NewFieldSpec(group.FieldID, field.TypeString))
	)
	_spec.OnConflict = gc.conflict
	if id, ok := gc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := gc.mutation.Name(); ok {
		_spec.SetField(group.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := gc.mutation.Scopes(); ok {
		_spec.SetField(group.FieldScopes, field.TypeOther, value)
		_node.Scopes = value
	}
	if value, ok := gc.mutation.Cidr(); ok {
		_spec.SetField(group.FieldCidr, field.TypeString, value)
		_node.Cidr = value
	}
	if value, ok := gc.mutation.Rules(); ok {
		_spec.SetField(group.FieldRules, field.TypeJSON, value)
		_node.Rules = value
	}
	if nodes := gc.mutation.UsersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   group.UsersTable,
			Columns: []string{group.UsersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Group.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (gc *GroupCreate) OnConflict(opts ...sql.ConflictOption) *GroupUpsertOne {
	gc.conflict = opts
	return &GroupUpsertOne{
		create: gc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gc *GroupCreate) OnConflictColumns(columns ...string) *GroupUpsertOne {
	gc.conflict = append(gc.conflict, sql.ConflictColumns(columns...))
	return &GroupUpsertOne{
		create: gc,
	}
}

type (
	// GroupUpsertOne is the builder for "upsert"-ing
	//  one Group node.
	GroupUpsertOne struct {
		create *GroupCreate
	}

	// GroupUpsert is the "OnConflict" setter.
	GroupUpsert struct {
		*sql.UpdateSet
	}
)

// SetName sets the "name" field.
func (u *GroupUpsert) SetName(v string) *GroupUpsert {
	u.Set(group.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GroupUpsert) UpdateName() *GroupUpsert {
	u.SetExcluded(group.FieldName)
	return u
}

// SetScopes sets the "scopes" field.
func (u *GroupUpsert) SetScopes(v authorization.Scopes) *GroupUpsert {
	u.Set(group.FieldScopes, v)
	return u
}

// UpdateScopes sets the "scopes" field to the value that was provided on create.
func (u *GroupUpsert) UpdateScopes() *GroupUpsert {
	u.SetExcluded(group.FieldScopes)
	return u
}

// SetCidr sets the "cidr" field.
func (u *GroupUpsert) SetCidr(v types.Cidr) *GroupUpsert {
	u.Set(group.FieldCidr, v)
	return u
}

// UpdateCidr sets the "cidr" field to the value that was provided on create.
func (u *GroupUpsert) UpdateCidr() *GroupUpsert {
	u.SetExcluded(group.FieldCidr)
	return u
}

// SetRules sets the "rules" field.
func (u *GroupUpsert) SetRules(v []types.Rule) *GroupUpsert {
	u.Set(group.FieldRules, v)
	return u
}

// UpdateRules sets the "rules" field to the value that was provided on create.
func (u *GroupUpsert) UpdateRules() *GroupUpsert {
	u.SetExcluded(group.FieldRules)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(group.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GroupUpsertOne) UpdateNewValues() *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(group.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Group.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *GroupUpsertOne) Ignore() *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupUpsertOne) DoNothing() *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupCreate.OnConflict
// documentation for more info.
func (u *GroupUpsertOne) Update(set func(*GroupUpsert)) *GroupUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *GroupUpsertOne) SetName(v string) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateName() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateName()
	})
}

// SetScopes sets the "scopes" field.
func (u *GroupUpsertOne) SetScopes(v authorization.Scopes) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetScopes(v)
	})
}

// UpdateScopes sets the "scopes" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateScopes() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateScopes()
	})
}

// SetCidr sets the "cidr" field.
func (u *GroupUpsertOne) SetCidr(v types.Cidr) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetCidr(v)
	})
}

// UpdateCidr sets the "cidr" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateCidr() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateCidr()
	})
}

// SetRules sets the "rules" field.
func (u *GroupUpsertOne) SetRules(v []types.Rule) *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.SetRules(v)
	})
}

// UpdateRules sets the "rules" field to the value that was provided on create.
func (u *GroupUpsertOne) UpdateRules() *GroupUpsertOne {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateRules()
	})
}

// Exec executes the query.
func (u *GroupUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *GroupUpsertOne) ID(ctx context.Context) (id string, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: GroupUpsertOne.ID is not supported by MySQL driver. Use GroupUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *GroupUpsertOne) IDX(ctx context.Context) string {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// GroupCreateBulk is the builder for creating many Group entities in bulk.
type GroupCreateBulk struct {
	config
	err      error
	builders []*GroupCreate
	conflict []sql.ConflictOption
}

// Save creates the Group entities in the database.
func (gcb *GroupCreateBulk) Save(ctx context.Context) ([]*Group, error) {
	if gcb.err != nil {
		return nil, gcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(gcb.builders))
	nodes := make([]*Group, len(gcb.builders))
	mutators := make([]Mutator, len(gcb.builders))
	for i := range gcb.builders {
		func(i int, root context.Context) {
			builder := gcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*GroupMutation)
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
					_, err = mutators[i+1].Mutate(root, gcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = gcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, gcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, gcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (gcb *GroupCreateBulk) SaveX(ctx context.Context) []*Group {
	v, err := gcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (gcb *GroupCreateBulk) Exec(ctx context.Context) error {
	_, err := gcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (gcb *GroupCreateBulk) ExecX(ctx context.Context) {
	if err := gcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Group.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.GroupUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (gcb *GroupCreateBulk) OnConflict(opts ...sql.ConflictOption) *GroupUpsertBulk {
	gcb.conflict = opts
	return &GroupUpsertBulk{
		create: gcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (gcb *GroupCreateBulk) OnConflictColumns(columns ...string) *GroupUpsertBulk {
	gcb.conflict = append(gcb.conflict, sql.ConflictColumns(columns...))
	return &GroupUpsertBulk{
		create: gcb,
	}
}

// GroupUpsertBulk is the builder for "upsert"-ing
// a bulk of Group nodes.
type GroupUpsertBulk struct {
	create *GroupCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(group.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *GroupUpsertBulk) UpdateNewValues() *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(group.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Group.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *GroupUpsertBulk) Ignore() *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *GroupUpsertBulk) DoNothing() *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the GroupCreateBulk.OnConflict
// documentation for more info.
func (u *GroupUpsertBulk) Update(set func(*GroupUpsert)) *GroupUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&GroupUpsert{UpdateSet: update})
	}))
	return u
}

// SetName sets the "name" field.
func (u *GroupUpsertBulk) SetName(v string) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateName() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateName()
	})
}

// SetScopes sets the "scopes" field.
func (u *GroupUpsertBulk) SetScopes(v authorization.Scopes) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetScopes(v)
	})
}

// UpdateScopes sets the "scopes" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateScopes() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateScopes()
	})
}

// SetCidr sets the "cidr" field.
func (u *GroupUpsertBulk) SetCidr(v types.Cidr) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetCidr(v)
	})
}

// UpdateCidr sets the "cidr" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateCidr() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateCidr()
	})
}

// SetRules sets the "rules" field.
func (u *GroupUpsertBulk) SetRules(v []types.Rule) *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.SetRules(v)
	})
}

// UpdateRules sets the "rules" field to the value that was provided on create.
func (u *GroupUpsertBulk) UpdateRules() *GroupUpsertBulk {
	return u.Update(func(s *GroupUpsert) {
		s.UpdateRules()
	})
}

// Exec executes the query.
func (u *GroupUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the GroupCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for GroupCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *GroupUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
