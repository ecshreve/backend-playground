package entpb_test

import (
	"context"
	"testing"

	"github.com/ecshreve/backend-playground/ent/enttest"
	"github.com/ecshreve/backend-playground/ent/proto/entpb"
	"github.com/ecshreve/backend-playground/ent/todo"
	user "github.com/ecshreve/backend-playground/ent/user"
	_ "github.com/mattn/go-sqlite3"
)

func TestServiceWithEdges(t *testing.T) {
	// start by initializing an ent client connected to an in memory sqlite instance
	ctx := context.Background()
	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()

	// next, initialize the UserService. Notice we won't be opening an actual port and creating a gRPC server
	// and instead we are just calling the library code directly.
	svc := entpb.NewUserService(client)

	// next, we create a category directly using the ent client. notice we are initializing it with no relation
	// to a User.
	td := client.Todo.Create().SetText("exampletodo").SaveX(ctx)

	// next, we invoke the User service's `Create` method. Notice we are passing a list of entpb.Todo
	// instances with only the ID set.
	create, err := svc.Create(ctx, &entpb.CreateUserRequest{
		User: &entpb.User{
			Name:  "user",
			Email: "user@service.code",
			Todos: []*entpb.Todo{
				{
					Id: int64(td.ID),
				},
			},
		},
	})
	if err != nil {
		t.Fatal("failed creating user using UserService", err)
	}

	// to verify everything worked correctly, we query the todo table to check we have exactly
	// one category which is administered by the created user.
	count, err := client.Todo.
		Query().
		Where(
			todo.HasUserWith(
				user.IDEQ(int(create.Id)),
			),
		).
		Count(ctx)
	if err != nil {
		t.Fatal("failed counting categories admin by created user", err)
	}
	if count != 1 {
		t.Fatal("expected exactly one group to managed by the created user")
	}
}

// func TestGet(t *testing.T) {
// 	// start by initializing an ent client connected to an in memory sqlite instance
// 	ctx := context.Background()
// 	client := enttest.Open(t, "sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
// 	defer client.Close()

// 	// next, initialize the UserService. Notice we won't be opening an actual port and creating a gRPC server
// 	// and instead we are just calling the library code directly.
// 	svc := entpb.NewUserService(client)

// 	// next, create a user, a category and set that user to be the admin of the category
// 	user := client.User.Create().
// 		SetName("rotemtam").
// 		SetEmail("r@entgo.io").
// 		SaveX(ctx)

// 	client.Todo.Create().
// 		SetText("todo").
// 		s(user).
// 		SaveX(ctx)

// 	// next, retrieve the user without edge information
// 	get, err := svc.Get(ctx, &entpb.GetUserRequest{
// 		Id: int64(user.ID),
// 	})
// 	if err != nil {
// 		t.Fatal("failed retrieving the created user", err)
// 	}
// 	if len(get.Administered) != 0 {
// 		t.Fatal("by default edge information is not supposed to be retrieved")
// 	}

// 	// next, retrieve the user *WITH* edge information
// 	get, err = svc.Get(ctx, &entpb.GetUserRequest{
// 		Id:   int64(user.ID),
// 		View: entpb.GetUserRequest_WITH_EDGE_IDS,
// 	})
// 	if err != nil {
// 		t.Fatal("failed retrieving the created user", err)
// 	}
// 	if len(get.Administered) != 1 {
// 		t.Fatal("using WITH_EDGE_IDS edges should be returned")
// 	}
// }
