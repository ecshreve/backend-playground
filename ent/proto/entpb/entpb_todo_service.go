// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	base64 "encoding/base64"
	entproto "entgo.io/contrib/entproto"
	runtime "entgo.io/contrib/entproto/runtime"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	fmt "fmt"
	ent "github.com/ecshreve/backend-playground/ent"
	todo "github.com/ecshreve/backend-playground/ent/todo"
	user "github.com/ecshreve/backend-playground/ent/user"
	empty "github.com/golang/protobuf/ptypes/empty"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	regexp "regexp"
	strconv "strconv"
	strings "strings"
)

// TodoService implements TodoServiceServer
type TodoService struct {
	client *ent.Client
	UnimplementedTodoServiceServer
}

// NewTodoService returns a new TodoService
func NewTodoService(client *ent.Client) *TodoService {
	return &TodoService{
		client: client,
	}
}

var protoIdentNormalizeRegexpTodo_Status = regexp.MustCompile(`[^a-zA-Z0-9_]+`)

func protoIdentNormalizeTodo_Status(e string) string {
	return protoIdentNormalizeRegexpTodo_Status.ReplaceAllString(e, "_")
}

func toProtoTodo_Status(e todo.Status) Todo_Status {
	if v, ok := Todo_Status_value[strings.ToUpper("STATUS_"+protoIdentNormalizeTodo_Status(string(e)))]; ok {
		return Todo_Status(v)
	}
	return Todo_Status(0)
}

func toEntTodo_Status(e Todo_Status) todo.Status {
	if v, ok := Todo_Status_name[int32(e)]; ok {
		entVal := map[string]string{
			"STATUS_IN_PROGRESS": "IN_PROGRESS",
			"STATUS_COMPLETED":   "COMPLETED",
		}[v]
		return todo.Status(entVal)
	}
	return ""
}

// toProtoTodo transforms the ent type to the pb type
func toProtoTodo(e *ent.Todo) (*Todo, error) {
	v := &Todo{}
	created_at := timestamppb.New(e.CreatedAt)
	v.CreatedAt = created_at
	id := int64(e.ID)
	v.Id = id
	priority := int64(e.Priority)
	v.Priority = priority
	status := toProtoTodo_Status(e.Status)
	v.Status = status
	text := e.Text
	v.Text = text
	updated_at := timestamppb.New(e.UpdatedAt)
	v.UpdatedAt = updated_at
	for _, edg := range e.Edges.User {
		id := int64(edg.ID)
		v.User = append(v.User, &User{
			Id: id,
		})
	}
	return v, nil
}

// toProtoTodoList transforms a list of ent type to a list of pb type
func toProtoTodoList(e []*ent.Todo) ([]*Todo, error) {
	var pbList []*Todo
	for _, entEntity := range e {
		pbEntity, err := toProtoTodo(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements TodoServiceServer.Create
func (svc *TodoService) Create(ctx context.Context, req *CreateTodoRequest) (*Todo, error) {
	todo := req.GetTodo()
	m, err := svc.createBuilder(todo)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoTodo(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Get implements TodoServiceServer.Get
func (svc *TodoService) Get(ctx context.Context, req *GetTodoRequest) (*Todo, error) {
	var (
		err error
		get *ent.Todo
	)
	id := int(req.GetId())
	switch req.GetView() {
	case GetTodoRequest_VIEW_UNSPECIFIED, GetTodoRequest_BASIC:
		get, err = svc.client.Todo.Get(ctx, id)
	case GetTodoRequest_WITH_EDGE_IDS:
		get, err = svc.client.Todo.Query().
			Where(todo.ID(id)).
			WithUser(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoTodo(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements TodoServiceServer.Update
func (svc *TodoService) Update(ctx context.Context, req *UpdateTodoRequest) (*Todo, error) {
	todo := req.GetTodo()
	todoID := int(todo.GetId())
	m := svc.client.Todo.UpdateOneID(todoID)
	todoPriority := int(todo.GetPriority())
	m.SetPriority(todoPriority)
	todoStatus := toEntTodo_Status(todo.GetStatus())
	m.SetStatus(todoStatus)
	todoText := todo.GetText()
	m.SetText(todoText)
	todoUpdatedAt := runtime.ExtractTime(todo.GetUpdatedAt())
	m.SetUpdatedAt(todoUpdatedAt)
	for _, item := range todo.GetUser() {
		user := int(item.GetId())
		m.AddUserIDs(user)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoTodo(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Delete implements TodoServiceServer.Delete
func (svc *TodoService) Delete(ctx context.Context, req *DeleteTodoRequest) (*empty.Empty, error) {
	var err error
	id := int(req.GetId())
	err = svc.client.Todo.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements TodoServiceServer.List
func (svc *TodoService) List(ctx context.Context, req *ListTodoRequest) (*ListTodoResponse, error) {
	var (
		err      error
		entList  []*ent.Todo
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Todo.Query().
		Order(ent.Desc(todo.FieldID)).
		Limit(pageSize + 1)
	if req.GetPageToken() != "" {
		bytes, err := base64.StdEncoding.DecodeString(req.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		token, err := strconv.ParseInt(string(bytes), 10, 32)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		pageToken := int(token)
		listQuery = listQuery.
			Where(todo.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListTodoRequest_VIEW_UNSPECIFIED, ListTodoRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListTodoRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithUser(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
			}).
			All(ctx)
	}
	switch {
	case err == nil:
		var nextPageToken string
		if len(entList) == pageSize+1 {
			nextPageToken = base64.StdEncoding.EncodeToString(
				[]byte(fmt.Sprintf("%v", entList[len(entList)-1].ID)))
			entList = entList[:len(entList)-1]
		}
		protoList, err := toProtoTodoList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &ListTodoResponse{
			TodoList:      protoList,
			NextPageToken: nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements TodoServiceServer.BatchCreate
func (svc *TodoService) BatchCreate(ctx context.Context, req *BatchCreateTodosRequest) (*BatchCreateTodosResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.TodoCreate, len(requests))
	for i, req := range requests {
		todo := req.GetTodo()
		var err error
		bulk[i], err = svc.createBuilder(todo)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.Todo.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoTodoList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreateTodosResponse{
			Todos: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *TodoService) createBuilder(todo *Todo) (*ent.TodoCreate, error) {
	m := svc.client.Todo.Create()
	todoCreatedAt := runtime.ExtractTime(todo.GetCreatedAt())
	m.SetCreatedAt(todoCreatedAt)
	todoPriority := int(todo.GetPriority())
	m.SetPriority(todoPriority)
	todoStatus := toEntTodo_Status(todo.GetStatus())
	m.SetStatus(todoStatus)
	todoText := todo.GetText()
	m.SetText(todoText)
	todoUpdatedAt := runtime.ExtractTime(todo.GetUpdatedAt())
	m.SetUpdatedAt(todoUpdatedAt)
	for _, item := range todo.GetUser() {
		user := int(item.GetId())
		m.AddUserIDs(user)
	}
	return m, nil
}
