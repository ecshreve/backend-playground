// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/backend-playground/ent/predicate"
	"github.com/ecshreve/backend-playground/ent/todo"
	"github.com/ecshreve/backend-playground/ent/user"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeTodo = "Todo"
	TypeUser = "User"
)

// TodoMutation represents an operation that mutates the Todo nodes in the graph.
type TodoMutation struct {
	config
	op            Op
	typ           string
	id            *int
	created_at    *time.Time
	updated_at    *time.Time
	text          *string
	status        *todo.Status
	priority      *int
	addpriority   *int
	clearedFields map[string]struct{}
	user          map[int]struct{}
	removeduser   map[int]struct{}
	cleareduser   bool
	done          bool
	oldValue      func(context.Context) (*Todo, error)
	predicates    []predicate.Todo
}

var _ ent.Mutation = (*TodoMutation)(nil)

// todoOption allows management of the mutation configuration using functional options.
type todoOption func(*TodoMutation)

// newTodoMutation creates new mutation for the Todo entity.
func newTodoMutation(c config, op Op, opts ...todoOption) *TodoMutation {
	m := &TodoMutation{
		config:        c,
		op:            op,
		typ:           TypeTodo,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withTodoID sets the ID field of the mutation.
func withTodoID(id int) todoOption {
	return func(m *TodoMutation) {
		var (
			err   error
			once  sync.Once
			value *Todo
		)
		m.oldValue = func(ctx context.Context) (*Todo, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Todo.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withTodo sets the old Todo of the mutation.
func withTodo(node *Todo) todoOption {
	return func(m *TodoMutation) {
		m.oldValue = func(context.Context) (*Todo, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m TodoMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m TodoMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *TodoMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *TodoMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Todo.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *TodoMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *TodoMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *TodoMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *TodoMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *TodoMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *TodoMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetText sets the "text" field.
func (m *TodoMutation) SetText(s string) {
	m.text = &s
}

// Text returns the value of the "text" field in the mutation.
func (m *TodoMutation) Text() (r string, exists bool) {
	v := m.text
	if v == nil {
		return
	}
	return *v, true
}

// OldText returns the old "text" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldText(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldText is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldText requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldText: %w", err)
	}
	return oldValue.Text, nil
}

// ResetText resets all changes to the "text" field.
func (m *TodoMutation) ResetText() {
	m.text = nil
}

// SetStatus sets the "status" field.
func (m *TodoMutation) SetStatus(t todo.Status) {
	m.status = &t
}

// Status returns the value of the "status" field in the mutation.
func (m *TodoMutation) Status() (r todo.Status, exists bool) {
	v := m.status
	if v == nil {
		return
	}
	return *v, true
}

// OldStatus returns the old "status" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldStatus(ctx context.Context) (v todo.Status, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldStatus is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldStatus requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldStatus: %w", err)
	}
	return oldValue.Status, nil
}

// ResetStatus resets all changes to the "status" field.
func (m *TodoMutation) ResetStatus() {
	m.status = nil
}

// SetPriority sets the "priority" field.
func (m *TodoMutation) SetPriority(i int) {
	m.priority = &i
	m.addpriority = nil
}

// Priority returns the value of the "priority" field in the mutation.
func (m *TodoMutation) Priority() (r int, exists bool) {
	v := m.priority
	if v == nil {
		return
	}
	return *v, true
}

// OldPriority returns the old "priority" field's value of the Todo entity.
// If the Todo object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *TodoMutation) OldPriority(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldPriority is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldPriority requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldPriority: %w", err)
	}
	return oldValue.Priority, nil
}

// AddPriority adds i to the "priority" field.
func (m *TodoMutation) AddPriority(i int) {
	if m.addpriority != nil {
		*m.addpriority += i
	} else {
		m.addpriority = &i
	}
}

// AddedPriority returns the value that was added to the "priority" field in this mutation.
func (m *TodoMutation) AddedPriority() (r int, exists bool) {
	v := m.addpriority
	if v == nil {
		return
	}
	return *v, true
}

// ResetPriority resets all changes to the "priority" field.
func (m *TodoMutation) ResetPriority() {
	m.priority = nil
	m.addpriority = nil
}

// AddUserIDs adds the "user" edge to the User entity by ids.
func (m *TodoMutation) AddUserIDs(ids ...int) {
	if m.user == nil {
		m.user = make(map[int]struct{})
	}
	for i := range ids {
		m.user[ids[i]] = struct{}{}
	}
}

// ClearUser clears the "user" edge to the User entity.
func (m *TodoMutation) ClearUser() {
	m.cleareduser = true
}

// UserCleared reports if the "user" edge to the User entity was cleared.
func (m *TodoMutation) UserCleared() bool {
	return m.cleareduser
}

// RemoveUserIDs removes the "user" edge to the User entity by IDs.
func (m *TodoMutation) RemoveUserIDs(ids ...int) {
	if m.removeduser == nil {
		m.removeduser = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.user, ids[i])
		m.removeduser[ids[i]] = struct{}{}
	}
}

// RemovedUser returns the removed IDs of the "user" edge to the User entity.
func (m *TodoMutation) RemovedUserIDs() (ids []int) {
	for id := range m.removeduser {
		ids = append(ids, id)
	}
	return
}

// UserIDs returns the "user" edge IDs in the mutation.
func (m *TodoMutation) UserIDs() (ids []int) {
	for id := range m.user {
		ids = append(ids, id)
	}
	return
}

// ResetUser resets all changes to the "user" edge.
func (m *TodoMutation) ResetUser() {
	m.user = nil
	m.cleareduser = false
	m.removeduser = nil
}

// Where appends a list predicates to the TodoMutation builder.
func (m *TodoMutation) Where(ps ...predicate.Todo) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the TodoMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *TodoMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Todo, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *TodoMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *TodoMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Todo).
func (m *TodoMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *TodoMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.created_at != nil {
		fields = append(fields, todo.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, todo.FieldUpdatedAt)
	}
	if m.text != nil {
		fields = append(fields, todo.FieldText)
	}
	if m.status != nil {
		fields = append(fields, todo.FieldStatus)
	}
	if m.priority != nil {
		fields = append(fields, todo.FieldPriority)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *TodoMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case todo.FieldCreatedAt:
		return m.CreatedAt()
	case todo.FieldUpdatedAt:
		return m.UpdatedAt()
	case todo.FieldText:
		return m.Text()
	case todo.FieldStatus:
		return m.Status()
	case todo.FieldPriority:
		return m.Priority()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *TodoMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case todo.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case todo.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case todo.FieldText:
		return m.OldText(ctx)
	case todo.FieldStatus:
		return m.OldStatus(ctx)
	case todo.FieldPriority:
		return m.OldPriority(ctx)
	}
	return nil, fmt.Errorf("unknown Todo field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TodoMutation) SetField(name string, value ent.Value) error {
	switch name {
	case todo.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case todo.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case todo.FieldText:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetText(v)
		return nil
	case todo.FieldStatus:
		v, ok := value.(todo.Status)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetStatus(v)
		return nil
	case todo.FieldPriority:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetPriority(v)
		return nil
	}
	return fmt.Errorf("unknown Todo field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *TodoMutation) AddedFields() []string {
	var fields []string
	if m.addpriority != nil {
		fields = append(fields, todo.FieldPriority)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *TodoMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case todo.FieldPriority:
		return m.AddedPriority()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *TodoMutation) AddField(name string, value ent.Value) error {
	switch name {
	case todo.FieldPriority:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddPriority(v)
		return nil
	}
	return fmt.Errorf("unknown Todo numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *TodoMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *TodoMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *TodoMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Todo nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *TodoMutation) ResetField(name string) error {
	switch name {
	case todo.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case todo.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case todo.FieldText:
		m.ResetText()
		return nil
	case todo.FieldStatus:
		m.ResetStatus()
		return nil
	case todo.FieldPriority:
		m.ResetPriority()
		return nil
	}
	return fmt.Errorf("unknown Todo field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *TodoMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.user != nil {
		edges = append(edges, todo.EdgeUser)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *TodoMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case todo.EdgeUser:
		ids := make([]ent.Value, 0, len(m.user))
		for id := range m.user {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *TodoMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removeduser != nil {
		edges = append(edges, todo.EdgeUser)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *TodoMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case todo.EdgeUser:
		ids := make([]ent.Value, 0, len(m.removeduser))
		for id := range m.removeduser {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *TodoMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareduser {
		edges = append(edges, todo.EdgeUser)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *TodoMutation) EdgeCleared(name string) bool {
	switch name {
	case todo.EdgeUser:
		return m.cleareduser
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *TodoMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown Todo unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *TodoMutation) ResetEdge(name string) error {
	switch name {
	case todo.EdgeUser:
		m.ResetUser()
		return nil
	}
	return fmt.Errorf("unknown Todo edge %s", name)
}

// UserMutation represents an operation that mutates the User nodes in the graph.
type UserMutation struct {
	config
	op               Op
	typ              string
	id               *int
	created_at       *time.Time
	updated_at       *time.Time
	name             *string
	email            *string
	avatar_image_url *string
	clearedFields    map[string]struct{}
	todos            map[int]struct{}
	removedtodos     map[int]struct{}
	clearedtodos     bool
	done             bool
	oldValue         func(context.Context) (*User, error)
	predicates       []predicate.User
}

var _ ent.Mutation = (*UserMutation)(nil)

// userOption allows management of the mutation configuration using functional options.
type userOption func(*UserMutation)

// newUserMutation creates new mutation for the User entity.
func newUserMutation(c config, op Op, opts ...userOption) *UserMutation {
	m := &UserMutation{
		config:        c,
		op:            op,
		typ:           TypeUser,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUserID sets the ID field of the mutation.
func withUserID(id int) userOption {
	return func(m *UserMutation) {
		var (
			err   error
			once  sync.Once
			value *User
		)
		m.oldValue = func(ctx context.Context) (*User, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().User.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUser sets the old User of the mutation.
func withUser(node *User) userOption {
	return func(m *UserMutation) {
		m.oldValue = func(context.Context) (*User, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UserMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UserMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UserMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UserMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().User.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetCreatedAt sets the "created_at" field.
func (m *UserMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *UserMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *UserMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUpdatedAt sets the "updated_at" field.
func (m *UserMutation) SetUpdatedAt(t time.Time) {
	m.updated_at = &t
}

// UpdatedAt returns the value of the "updated_at" field in the mutation.
func (m *UserMutation) UpdatedAt() (r time.Time, exists bool) {
	v := m.updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldUpdatedAt returns the old "updated_at" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUpdatedAt: %w", err)
	}
	return oldValue.UpdatedAt, nil
}

// ResetUpdatedAt resets all changes to the "updated_at" field.
func (m *UserMutation) ResetUpdatedAt() {
	m.updated_at = nil
}

// SetName sets the "name" field.
func (m *UserMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UserMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UserMutation) ResetName() {
	m.name = nil
}

// SetEmail sets the "email" field.
func (m *UserMutation) SetEmail(s string) {
	m.email = &s
}

// Email returns the value of the "email" field in the mutation.
func (m *UserMutation) Email() (r string, exists bool) {
	v := m.email
	if v == nil {
		return
	}
	return *v, true
}

// OldEmail returns the old "email" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldEmail(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldEmail is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldEmail requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldEmail: %w", err)
	}
	return oldValue.Email, nil
}

// ResetEmail resets all changes to the "email" field.
func (m *UserMutation) ResetEmail() {
	m.email = nil
}

// SetAvatarImageURL sets the "avatar_image_url" field.
func (m *UserMutation) SetAvatarImageURL(s string) {
	m.avatar_image_url = &s
}

// AvatarImageURL returns the value of the "avatar_image_url" field in the mutation.
func (m *UserMutation) AvatarImageURL() (r string, exists bool) {
	v := m.avatar_image_url
	if v == nil {
		return
	}
	return *v, true
}

// OldAvatarImageURL returns the old "avatar_image_url" field's value of the User entity.
// If the User object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UserMutation) OldAvatarImageURL(ctx context.Context) (v *string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldAvatarImageURL is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldAvatarImageURL requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAvatarImageURL: %w", err)
	}
	return oldValue.AvatarImageURL, nil
}

// ClearAvatarImageURL clears the value of the "avatar_image_url" field.
func (m *UserMutation) ClearAvatarImageURL() {
	m.avatar_image_url = nil
	m.clearedFields[user.FieldAvatarImageURL] = struct{}{}
}

// AvatarImageURLCleared returns if the "avatar_image_url" field was cleared in this mutation.
func (m *UserMutation) AvatarImageURLCleared() bool {
	_, ok := m.clearedFields[user.FieldAvatarImageURL]
	return ok
}

// ResetAvatarImageURL resets all changes to the "avatar_image_url" field.
func (m *UserMutation) ResetAvatarImageURL() {
	m.avatar_image_url = nil
	delete(m.clearedFields, user.FieldAvatarImageURL)
}

// AddTodoIDs adds the "todos" edge to the Todo entity by ids.
func (m *UserMutation) AddTodoIDs(ids ...int) {
	if m.todos == nil {
		m.todos = make(map[int]struct{})
	}
	for i := range ids {
		m.todos[ids[i]] = struct{}{}
	}
}

// ClearTodos clears the "todos" edge to the Todo entity.
func (m *UserMutation) ClearTodos() {
	m.clearedtodos = true
}

// TodosCleared reports if the "todos" edge to the Todo entity was cleared.
func (m *UserMutation) TodosCleared() bool {
	return m.clearedtodos
}

// RemoveTodoIDs removes the "todos" edge to the Todo entity by IDs.
func (m *UserMutation) RemoveTodoIDs(ids ...int) {
	if m.removedtodos == nil {
		m.removedtodos = make(map[int]struct{})
	}
	for i := range ids {
		delete(m.todos, ids[i])
		m.removedtodos[ids[i]] = struct{}{}
	}
}

// RemovedTodos returns the removed IDs of the "todos" edge to the Todo entity.
func (m *UserMutation) RemovedTodosIDs() (ids []int) {
	for id := range m.removedtodos {
		ids = append(ids, id)
	}
	return
}

// TodosIDs returns the "todos" edge IDs in the mutation.
func (m *UserMutation) TodosIDs() (ids []int) {
	for id := range m.todos {
		ids = append(ids, id)
	}
	return
}

// ResetTodos resets all changes to the "todos" edge.
func (m *UserMutation) ResetTodos() {
	m.todos = nil
	m.clearedtodos = false
	m.removedtodos = nil
}

// Where appends a list predicates to the UserMutation builder.
func (m *UserMutation) Where(ps ...predicate.User) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UserMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UserMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.User, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UserMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UserMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (User).
func (m *UserMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UserMutation) Fields() []string {
	fields := make([]string, 0, 5)
	if m.created_at != nil {
		fields = append(fields, user.FieldCreatedAt)
	}
	if m.updated_at != nil {
		fields = append(fields, user.FieldUpdatedAt)
	}
	if m.name != nil {
		fields = append(fields, user.FieldName)
	}
	if m.email != nil {
		fields = append(fields, user.FieldEmail)
	}
	if m.avatar_image_url != nil {
		fields = append(fields, user.FieldAvatarImageURL)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UserMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case user.FieldCreatedAt:
		return m.CreatedAt()
	case user.FieldUpdatedAt:
		return m.UpdatedAt()
	case user.FieldName:
		return m.Name()
	case user.FieldEmail:
		return m.Email()
	case user.FieldAvatarImageURL:
		return m.AvatarImageURL()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UserMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case user.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case user.FieldUpdatedAt:
		return m.OldUpdatedAt(ctx)
	case user.FieldName:
		return m.OldName(ctx)
	case user.FieldEmail:
		return m.OldEmail(ctx)
	case user.FieldAvatarImageURL:
		return m.OldAvatarImageURL(ctx)
	}
	return nil, fmt.Errorf("unknown User field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) SetField(name string, value ent.Value) error {
	switch name {
	case user.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case user.FieldUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUpdatedAt(v)
		return nil
	case user.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case user.FieldEmail:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetEmail(v)
		return nil
	case user.FieldAvatarImageURL:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAvatarImageURL(v)
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UserMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UserMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UserMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown User numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UserMutation) ClearedFields() []string {
	var fields []string
	if m.FieldCleared(user.FieldAvatarImageURL) {
		fields = append(fields, user.FieldAvatarImageURL)
	}
	return fields
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UserMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UserMutation) ClearField(name string) error {
	switch name {
	case user.FieldAvatarImageURL:
		m.ClearAvatarImageURL()
		return nil
	}
	return fmt.Errorf("unknown User nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UserMutation) ResetField(name string) error {
	switch name {
	case user.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case user.FieldUpdatedAt:
		m.ResetUpdatedAt()
		return nil
	case user.FieldName:
		m.ResetName()
		return nil
	case user.FieldEmail:
		m.ResetEmail()
		return nil
	case user.FieldAvatarImageURL:
		m.ResetAvatarImageURL()
		return nil
	}
	return fmt.Errorf("unknown User field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UserMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.todos != nil {
		edges = append(edges, user.EdgeTodos)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UserMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeTodos:
		ids := make([]ent.Value, 0, len(m.todos))
		for id := range m.todos {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UserMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	if m.removedtodos != nil {
		edges = append(edges, user.EdgeTodos)
	}
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UserMutation) RemovedIDs(name string) []ent.Value {
	switch name {
	case user.EdgeTodos:
		ids := make([]ent.Value, 0, len(m.removedtodos))
		for id := range m.removedtodos {
			ids = append(ids, id)
		}
		return ids
	}
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UserMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedtodos {
		edges = append(edges, user.EdgeTodos)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UserMutation) EdgeCleared(name string) bool {
	switch name {
	case user.EdgeTodos:
		return m.clearedtodos
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UserMutation) ClearEdge(name string) error {
	switch name {
	}
	return fmt.Errorf("unknown User unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UserMutation) ResetEdge(name string) error {
	switch name {
	case user.EdgeTodos:
		m.ResetTodos()
		return nil
	}
	return fmt.Errorf("unknown User edge %s", name)
}
