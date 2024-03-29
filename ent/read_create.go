// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"pingpong2/ent/read"
	"pingpong2/ent/sentense"
	"pingpong2/ent/user"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ReadCreate is the builder for creating a Read entity.
type ReadCreate struct {
	config
	mutation *ReadMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (rc *ReadCreate) SetUserID(i int) *ReadCreate {
	rc.mutation.SetUserID(i)
	return rc
}

// SetSentenceID sets the "sentence_id" field.
func (rc *ReadCreate) SetSentenceID(i int) *ReadCreate {
	rc.mutation.SetSentenceID(i)
	return rc
}

// SetResult sets the "result" field.
func (rc *ReadCreate) SetResult(i int) *ReadCreate {
	rc.mutation.SetResult(i)
	return rc
}

// SetUser sets the "user" edge to the User entity.
func (rc *ReadCreate) SetUser(u *User) *ReadCreate {
	return rc.SetUserID(u.ID)
}

// SetSentence sets the "sentence" edge to the Sentense entity.
func (rc *ReadCreate) SetSentence(s *Sentense) *ReadCreate {
	return rc.SetSentenceID(s.ID)
}

// Mutation returns the ReadMutation object of the builder.
func (rc *ReadCreate) Mutation() *ReadMutation {
	return rc.mutation
}

// Save creates the Read in the database.
func (rc *ReadCreate) Save(ctx context.Context) (*Read, error) {
	var (
		err  error
		node *Read
	)
	if len(rc.hooks) == 0 {
		if err = rc.check(); err != nil {
			return nil, err
		}
		node, err = rc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ReadMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rc.check(); err != nil {
				return nil, err
			}
			rc.mutation = mutation
			if node, err = rc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rc.hooks) - 1; i >= 0; i-- {
			if rc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rc *ReadCreate) SaveX(ctx context.Context) *Read {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *ReadCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *ReadCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *ReadCreate) check() error {
	if _, ok := rc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Read.user_id"`)}
	}
	if _, ok := rc.mutation.SentenceID(); !ok {
		return &ValidationError{Name: "sentence_id", err: errors.New(`ent: missing required field "Read.sentence_id"`)}
	}
	if _, ok := rc.mutation.Result(); !ok {
		return &ValidationError{Name: "result", err: errors.New(`ent: missing required field "Read.result"`)}
	}
	if _, ok := rc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user", err: errors.New(`ent: missing required edge "Read.user"`)}
	}
	if _, ok := rc.mutation.SentenceID(); !ok {
		return &ValidationError{Name: "sentence", err: errors.New(`ent: missing required edge "Read.sentence"`)}
	}
	return nil
}

func (rc *ReadCreate) sqlSave(ctx context.Context) (*Read, error) {
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (rc *ReadCreate) createSpec() (*Read, *sqlgraph.CreateSpec) {
	var (
		_node = &Read{config: rc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: read.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: read.FieldID,
			},
		}
	)
	if value, ok := rc.mutation.Result(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: read.FieldResult,
		})
		_node.Result = value
	}
	if nodes := rc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   read.UserTable,
			Columns: []string{read.UserColumn},
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
		_node.UserID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rc.mutation.SentenceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   read.SentenceTable,
			Columns: []string{read.SentenceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: sentense.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SentenceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// ReadCreateBulk is the builder for creating many Read entities in bulk.
type ReadCreateBulk struct {
	config
	builders []*ReadCreate
}

// Save creates the Read entities in the database.
func (rcb *ReadCreateBulk) Save(ctx context.Context) ([]*Read, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Read, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ReadMutation)
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
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *ReadCreateBulk) SaveX(ctx context.Context) []*Read {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *ReadCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *ReadCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
