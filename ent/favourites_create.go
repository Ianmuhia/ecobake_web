// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"ecobake/ent/favourites"
	"ecobake/ent/product"
	"ecobake/ent/user"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// FavouritesCreate is the builder for creating a Favourites entity.
type FavouritesCreate struct {
	config
	mutation *FavouritesMutation
	hooks    []Hook
}

// SetCreatedAt sets the "created_at" field.
func (fc *FavouritesCreate) SetCreatedAt(t time.Time) *FavouritesCreate {
	fc.mutation.SetCreatedAt(t)
	return fc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (fc *FavouritesCreate) SetNillableCreatedAt(t *time.Time) *FavouritesCreate {
	if t != nil {
		fc.SetCreatedAt(*t)
	}
	return fc
}

// SetUpdatedAt sets the "updated_at" field.
func (fc *FavouritesCreate) SetUpdatedAt(t time.Time) *FavouritesCreate {
	fc.mutation.SetUpdatedAt(t)
	return fc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (fc *FavouritesCreate) SetNillableUpdatedAt(t *time.Time) *FavouritesCreate {
	if t != nil {
		fc.SetUpdatedAt(*t)
	}
	return fc
}

// SetDeletedAt sets the "deleted_at" field.
func (fc *FavouritesCreate) SetDeletedAt(t time.Time) *FavouritesCreate {
	fc.mutation.SetDeletedAt(t)
	return fc
}

// SetNillableDeletedAt sets the "deleted_at" field if the given value is not nil.
func (fc *FavouritesCreate) SetNillableDeletedAt(t *time.Time) *FavouritesCreate {
	if t != nil {
		fc.SetDeletedAt(*t)
	}
	return fc
}

// SetProductID sets the "product" edge to the Product entity by ID.
func (fc *FavouritesCreate) SetProductID(id int) *FavouritesCreate {
	fc.mutation.SetProductID(id)
	return fc
}

// SetNillableProductID sets the "product" edge to the Product entity by ID if the given value is not nil.
func (fc *FavouritesCreate) SetNillableProductID(id *int) *FavouritesCreate {
	if id != nil {
		fc = fc.SetProductID(*id)
	}
	return fc
}

// SetProduct sets the "product" edge to the Product entity.
func (fc *FavouritesCreate) SetProduct(p *Product) *FavouritesCreate {
	return fc.SetProductID(p.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (fc *FavouritesCreate) SetUserID(id int) *FavouritesCreate {
	fc.mutation.SetUserID(id)
	return fc
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (fc *FavouritesCreate) SetNillableUserID(id *int) *FavouritesCreate {
	if id != nil {
		fc = fc.SetUserID(*id)
	}
	return fc
}

// SetUser sets the "user" edge to the User entity.
func (fc *FavouritesCreate) SetUser(u *User) *FavouritesCreate {
	return fc.SetUserID(u.ID)
}

// Mutation returns the FavouritesMutation object of the builder.
func (fc *FavouritesCreate) Mutation() *FavouritesMutation {
	return fc.mutation
}

// Save creates the Favourites in the database.
func (fc *FavouritesCreate) Save(ctx context.Context) (*Favourites, error) {
	var (
		err  error
		node *Favourites
	)
	fc.defaults()
	if len(fc.hooks) == 0 {
		if err = fc.check(); err != nil {
			return nil, err
		}
		node, err = fc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*FavouritesMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = fc.check(); err != nil {
				return nil, err
			}
			fc.mutation = mutation
			if node, err = fc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(fc.hooks) - 1; i >= 0; i-- {
			if fc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = fc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, fc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Favourites)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from FavouritesMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FavouritesCreate) SaveX(ctx context.Context) *Favourites {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FavouritesCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FavouritesCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (fc *FavouritesCreate) defaults() {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		v := favourites.DefaultCreatedAt
		fc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FavouritesCreate) check() error {
	if _, ok := fc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Favourites.created_at"`)}
	}
	return nil
}

func (fc *FavouritesCreate) sqlSave(ctx context.Context) (*Favourites, error) {
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (fc *FavouritesCreate) createSpec() (*Favourites, *sqlgraph.CreateSpec) {
	var (
		_node = &Favourites{config: fc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: favourites.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: favourites.FieldID,
			},
		}
	)
	if value, ok := fc.mutation.CreatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: favourites.FieldCreatedAt,
		})
		_node.CreatedAt = value
	}
	if value, ok := fc.mutation.UpdatedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: favourites.FieldUpdatedAt,
		})
		_node.UpdatedAt = value
	}
	if value, ok := fc.mutation.DeletedAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: favourites.FieldDeletedAt,
		})
		_node.DeletedAt = value
	}
	if nodes := fc.mutation.ProductIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   favourites.ProductTable,
			Columns: []string{favourites.ProductColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: product.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.product_favourites = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   favourites.UserTable,
			Columns: []string{favourites.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_favourites = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FavouritesCreateBulk is the builder for creating many Favourites entities in bulk.
type FavouritesCreateBulk struct {
	config
	builders []*FavouritesCreate
}

// Save creates the Favourites entities in the database.
func (fcb *FavouritesCreateBulk) Save(ctx context.Context) ([]*Favourites, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Favourites, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FavouritesMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FavouritesCreateBulk) SaveX(ctx context.Context) []*Favourites {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FavouritesCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FavouritesCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
