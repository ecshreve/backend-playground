// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// TodosColumns holds the columns for the "todos" table.
	TodosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "text", Type: field.TypeString, Size: 2147483647},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"IN_PROGRESS", "COMPLETED"}, Default: "IN_PROGRESS"},
		{Name: "priority", Type: field.TypeInt, Default: 0},
	}
	// TodosTable holds the schema information for the "todos" table.
	TodosTable = &schema.Table{
		Name:       "todos",
		Columns:    TodosColumns,
		PrimaryKey: []*schema.Column{TodosColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
		{Name: "name", Type: field.TypeString},
		{Name: "email", Type: field.TypeString},
		{Name: "avatar_image_url", Type: field.TypeString, Nullable: true},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// TodoUserColumns holds the columns for the "todo_user" table.
	TodoUserColumns = []*schema.Column{
		{Name: "todo_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// TodoUserTable holds the schema information for the "todo_user" table.
	TodoUserTable = &schema.Table{
		Name:       "todo_user",
		Columns:    TodoUserColumns,
		PrimaryKey: []*schema.Column{TodoUserColumns[0], TodoUserColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "todo_user_todo_id",
				Columns:    []*schema.Column{TodoUserColumns[0]},
				RefColumns: []*schema.Column{TodosColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "todo_user_user_id",
				Columns:    []*schema.Column{TodoUserColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		TodosTable,
		UsersTable,
		TodoUserTable,
	}
)

func init() {
	TodoUserTable.ForeignKeys[0].RefTable = TodosTable
	TodoUserTable.ForeignKeys[1].RefTable = UsersTable
}
