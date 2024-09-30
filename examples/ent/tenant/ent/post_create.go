// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/jace996/saas/examples/ent/tenant/ent/post"
)

// PostCreate is the builder for creating a Post entity.
type PostCreate struct {
	config
	mutation *PostMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetTenantID sets the "tenant_id" field.
func (pc *PostCreate) SetTenantID(ss *sql.NullString) *PostCreate {
	pc.mutation.SetTenantID(ss)
	return pc
}

// SetTitle sets the "title" field.
func (pc *PostCreate) SetTitle(s string) *PostCreate {
	pc.mutation.SetTitle(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *PostCreate) SetDescription(s string) *PostCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *PostCreate) SetNillableDescription(s *string) *PostCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetDsn sets the "dsn" field.
func (pc *PostCreate) SetDsn(s string) *PostCreate {
	pc.mutation.SetDsn(s)
	return pc
}

// SetNillableDsn sets the "dsn" field if the given value is not nil.
func (pc *PostCreate) SetNillableDsn(s *string) *PostCreate {
	if s != nil {
		pc.SetDsn(*s)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *PostCreate) SetID(i int) *PostCreate {
	pc.mutation.SetID(i)
	return pc
}

// Mutation returns the PostMutation object of the builder.
func (pc *PostCreate) Mutation() *PostMutation {
	return pc.mutation
}

// Save creates the Post in the database.
func (pc *PostCreate) Save(ctx context.Context) (*Post, error) {
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PostCreate) SaveX(ctx context.Context) *Post {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *PostCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *PostCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pc *PostCreate) check() error {
	if _, ok := pc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Post.title"`)}
	}
	return nil
}

func (pc *PostCreate) sqlSave(ctx context.Context) (*Post, error) {
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
	if _spec.ID.Value != _node.ID {
		id := _spec.ID.Value.(int64)
		_node.ID = int(id)
	}
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *PostCreate) createSpec() (*Post, *sqlgraph.CreateSpec) {
	var (
		_node = &Post{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(post.Table, sqlgraph.NewFieldSpec(post.FieldID, field.TypeInt))
	)
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := pc.mutation.TenantID(); ok {
		_spec.SetField(post.FieldTenantID, field.TypeString, value)
		_node.TenantID = value
	}
	if value, ok := pc.mutation.Title(); ok {
		_spec.SetField(post.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(post.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.Dsn(); ok {
		_spec.SetField(post.FieldDsn, field.TypeString, value)
		_node.Dsn = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Post.Create().
//		SetTenantID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PostUpsert) {
//			SetTenantID(v+v).
//		}).
//		Exec(ctx)
func (pc *PostCreate) OnConflict(opts ...sql.ConflictOption) *PostUpsertOne {
	pc.conflict = opts
	return &PostUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *PostCreate) OnConflictColumns(columns ...string) *PostUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &PostUpsertOne{
		create: pc,
	}
}

type (
	// PostUpsertOne is the builder for "upsert"-ing
	//  one Post node.
	PostUpsertOne struct {
		create *PostCreate
	}

	// PostUpsert is the "OnConflict" setter.
	PostUpsert struct {
		*sql.UpdateSet
	}
)

// SetTenantID sets the "tenant_id" field.
func (u *PostUpsert) SetTenantID(v *sql.NullString) *PostUpsert {
	u.Set(post.FieldTenantID, v)
	return u
}

// UpdateTenantID sets the "tenant_id" field to the value that was provided on create.
func (u *PostUpsert) UpdateTenantID() *PostUpsert {
	u.SetExcluded(post.FieldTenantID)
	return u
}

// ClearTenantID clears the value of the "tenant_id" field.
func (u *PostUpsert) ClearTenantID() *PostUpsert {
	u.SetNull(post.FieldTenantID)
	return u
}

// SetTitle sets the "title" field.
func (u *PostUpsert) SetTitle(v string) *PostUpsert {
	u.Set(post.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *PostUpsert) UpdateTitle() *PostUpsert {
	u.SetExcluded(post.FieldTitle)
	return u
}

// SetDescription sets the "description" field.
func (u *PostUpsert) SetDescription(v string) *PostUpsert {
	u.Set(post.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *PostUpsert) UpdateDescription() *PostUpsert {
	u.SetExcluded(post.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *PostUpsert) ClearDescription() *PostUpsert {
	u.SetNull(post.FieldDescription)
	return u
}

// SetDsn sets the "dsn" field.
func (u *PostUpsert) SetDsn(v string) *PostUpsert {
	u.Set(post.FieldDsn, v)
	return u
}

// UpdateDsn sets the "dsn" field to the value that was provided on create.
func (u *PostUpsert) UpdateDsn() *PostUpsert {
	u.SetExcluded(post.FieldDsn)
	return u
}

// ClearDsn clears the value of the "dsn" field.
func (u *PostUpsert) ClearDsn() *PostUpsert {
	u.SetNull(post.FieldDsn)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(post.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PostUpsertOne) UpdateNewValues() *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(post.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Post.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *PostUpsertOne) Ignore() *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PostUpsertOne) DoNothing() *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PostCreate.OnConflict
// documentation for more info.
func (u *PostUpsertOne) Update(set func(*PostUpsert)) *PostUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PostUpsert{UpdateSet: update})
	}))
	return u
}

// SetTenantID sets the "tenant_id" field.
func (u *PostUpsertOne) SetTenantID(v *sql.NullString) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetTenantID(v)
	})
}

// UpdateTenantID sets the "tenant_id" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateTenantID() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateTenantID()
	})
}

// ClearTenantID clears the value of the "tenant_id" field.
func (u *PostUpsertOne) ClearTenantID() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.ClearTenantID()
	})
}

// SetTitle sets the "title" field.
func (u *PostUpsertOne) SetTitle(v string) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateTitle() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateTitle()
	})
}

// SetDescription sets the "description" field.
func (u *PostUpsertOne) SetDescription(v string) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateDescription() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *PostUpsertOne) ClearDescription() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.ClearDescription()
	})
}

// SetDsn sets the "dsn" field.
func (u *PostUpsertOne) SetDsn(v string) *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.SetDsn(v)
	})
}

// UpdateDsn sets the "dsn" field to the value that was provided on create.
func (u *PostUpsertOne) UpdateDsn() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.UpdateDsn()
	})
}

// ClearDsn clears the value of the "dsn" field.
func (u *PostUpsertOne) ClearDsn() *PostUpsertOne {
	return u.Update(func(s *PostUpsert) {
		s.ClearDsn()
	})
}

// Exec executes the query.
func (u *PostUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PostCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PostUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *PostUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *PostUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// PostCreateBulk is the builder for creating many Post entities in bulk.
type PostCreateBulk struct {
	config
	err      error
	builders []*PostCreate
	conflict []sql.ConflictOption
}

// Save creates the Post entities in the database.
func (pcb *PostCreateBulk) Save(ctx context.Context) ([]*Post, error) {
	if pcb.err != nil {
		return nil, pcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Post, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PostMutation)
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *PostCreateBulk) SaveX(ctx context.Context) []*Post {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *PostCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *PostCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Post.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.PostUpsert) {
//			SetTenantID(v+v).
//		}).
//		Exec(ctx)
func (pcb *PostCreateBulk) OnConflict(opts ...sql.ConflictOption) *PostUpsertBulk {
	pcb.conflict = opts
	return &PostUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *PostCreateBulk) OnConflictColumns(columns ...string) *PostUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &PostUpsertBulk{
		create: pcb,
	}
}

// PostUpsertBulk is the builder for "upsert"-ing
// a bulk of Post nodes.
type PostUpsertBulk struct {
	create *PostCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(post.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *PostUpsertBulk) UpdateNewValues() *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(post.FieldID)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Post.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *PostUpsertBulk) Ignore() *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *PostUpsertBulk) DoNothing() *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the PostCreateBulk.OnConflict
// documentation for more info.
func (u *PostUpsertBulk) Update(set func(*PostUpsert)) *PostUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&PostUpsert{UpdateSet: update})
	}))
	return u
}

// SetTenantID sets the "tenant_id" field.
func (u *PostUpsertBulk) SetTenantID(v *sql.NullString) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetTenantID(v)
	})
}

// UpdateTenantID sets the "tenant_id" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateTenantID() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateTenantID()
	})
}

// ClearTenantID clears the value of the "tenant_id" field.
func (u *PostUpsertBulk) ClearTenantID() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.ClearTenantID()
	})
}

// SetTitle sets the "title" field.
func (u *PostUpsertBulk) SetTitle(v string) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateTitle() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateTitle()
	})
}

// SetDescription sets the "description" field.
func (u *PostUpsertBulk) SetDescription(v string) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateDescription() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *PostUpsertBulk) ClearDescription() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.ClearDescription()
	})
}

// SetDsn sets the "dsn" field.
func (u *PostUpsertBulk) SetDsn(v string) *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.SetDsn(v)
	})
}

// UpdateDsn sets the "dsn" field to the value that was provided on create.
func (u *PostUpsertBulk) UpdateDsn() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.UpdateDsn()
	})
}

// ClearDsn clears the value of the "dsn" field.
func (u *PostUpsertBulk) ClearDsn() *PostUpsertBulk {
	return u.Update(func(s *PostUpsert) {
		s.ClearDsn()
	})
}

// Exec executes the query.
func (u *PostUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the PostCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for PostCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *PostUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
