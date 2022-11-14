// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// InteractionsColumns holds the columns for the "interactions" table.
	InteractionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"DISLIKE", "LIKE", "RESQUEAK", "UNDO_DISLIKE", "UNDO_LIKE", "UNDO_RESQUEAK"}},
		{Name: "squeak_interactions", Type: field.TypeInt, Nullable: true},
		{Name: "user_interactions", Type: field.TypeInt, Nullable: true},
	}
	// InteractionsTable holds the schema information for the "interactions" table.
	InteractionsTable = &schema.Table{
		Name:       "interactions",
		Columns:    InteractionsColumns,
		PrimaryKey: []*schema.Column{InteractionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "interactions_squeaks_interactions",
				Columns:    []*schema.Column{InteractionsColumns[2]},
				RefColumns: []*schema.Column{SqueaksColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "interactions_users_interactions",
				Columns:    []*schema.Column{InteractionsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PoolsColumns holds the columns for the "pools" table.
	PoolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "amount", Type: field.TypeInt, SchemaType: map[string]string{"postgres": "numeric(78, 0)"}},
		{Name: "shares", Type: field.TypeInt, SchemaType: map[string]string{"postgres": "numeric(78, 0)"}},
		{Name: "block_number", Type: field.TypeInt, SchemaType: map[string]string{"postgres": "numeric(78, 0)"}},
		{Name: "score", Type: field.TypeInt},
		{Name: "squeak_pool", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// PoolsTable holds the schema information for the "pools" table.
	PoolsTable = &schema.Table{
		Name:       "pools",
		Columns:    PoolsColumns,
		PrimaryKey: []*schema.Column{PoolsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pools_squeaks_pool",
				Columns:    []*schema.Column{PoolsColumns[5]},
				RefColumns: []*schema.Column{SqueaksColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PoolPassesColumns holds the columns for the "pool_passes" table.
	PoolPassesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "shares", Type: field.TypeInt, SchemaType: map[string]string{"postgres": "numeric(78, 0)"}},
		{Name: "pool_pool_passes", Type: field.TypeInt, Nullable: true},
		{Name: "user_pool_passes", Type: field.TypeInt, Nullable: true},
	}
	// PoolPassesTable holds the schema information for the "pool_passes" table.
	PoolPassesTable = &schema.Table{
		Name:       "pool_passes",
		Columns:    PoolPassesColumns,
		PrimaryKey: []*schema.Column{PoolPassesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "pool_passes_pools_pool_passes",
				Columns:    []*schema.Column{PoolPassesColumns[2]},
				RefColumns: []*schema.Column{PoolsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "pool_passes_users_pool_passes",
				Columns:    []*schema.Column{PoolPassesColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RolesColumns holds the columns for the "roles" table.
	RolesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "title", Type: field.TypeString, Size: 32},
		{Name: "hash", Type: field.TypeString},
	}
	// RolesTable holds the schema information for the "roles" table.
	RolesTable = &schema.Table{
		Name:       "roles",
		Columns:    RolesColumns,
		PrimaryKey: []*schema.Column{RolesColumns[0]},
	}
	// SqueaksColumns holds the columns for the "squeaks" table.
	SqueaksColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "block_number", Type: field.TypeInt, SchemaType: map[string]string{"postgres": "numeric(78, 0)"}},
		{Name: "content", Type: field.TypeString, Size: 256},
		{Name: "user_created", Type: field.TypeInt, Nullable: true},
		{Name: "user_owned", Type: field.TypeInt, Nullable: true},
	}
	// SqueaksTable holds the schema information for the "squeaks" table.
	SqueaksTable = &schema.Table{
		Name:       "squeaks",
		Columns:    SqueaksColumns,
		PrimaryKey: []*schema.Column{SqueaksColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "squeaks_users_created",
				Columns:    []*schema.Column{SqueaksColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "squeaks_users_owned",
				Columns:    []*schema.Column{SqueaksColumns[4]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "address", Type: field.TypeString},
		{Name: "username", Type: field.TypeString, Size: 32},
		{Name: "status", Type: field.TypeEnum, Enums: []string{"UNKNOWN", "ACTIVE", "SUSPENDED", "BANNED"}},
		{Name: "level", Type: field.TypeInt8, Default: 1},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// UserRolesColumns holds the columns for the "user_roles" table.
	UserRolesColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "role_id", Type: field.TypeInt},
	}
	// UserRolesTable holds the schema information for the "user_roles" table.
	UserRolesTable = &schema.Table{
		Name:       "user_roles",
		Columns:    UserRolesColumns,
		PrimaryKey: []*schema.Column{UserRolesColumns[0], UserRolesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_roles_user_id",
				Columns:    []*schema.Column{UserRolesColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_roles_role_id",
				Columns:    []*schema.Column{UserRolesColumns[1]},
				RefColumns: []*schema.Column{RolesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// UserFollowingColumns holds the columns for the "user_following" table.
	UserFollowingColumns = []*schema.Column{
		{Name: "user_id", Type: field.TypeInt},
		{Name: "follower_id", Type: field.TypeInt},
	}
	// UserFollowingTable holds the schema information for the "user_following" table.
	UserFollowingTable = &schema.Table{
		Name:       "user_following",
		Columns:    UserFollowingColumns,
		PrimaryKey: []*schema.Column{UserFollowingColumns[0], UserFollowingColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "user_following_user_id",
				Columns:    []*schema.Column{UserFollowingColumns[0]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "user_following_follower_id",
				Columns:    []*schema.Column{UserFollowingColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		InteractionsTable,
		PoolsTable,
		PoolPassesTable,
		RolesTable,
		SqueaksTable,
		UsersTable,
		UserRolesTable,
		UserFollowingTable,
	}
)

func init() {
	InteractionsTable.ForeignKeys[0].RefTable = SqueaksTable
	InteractionsTable.ForeignKeys[1].RefTable = UsersTable
	PoolsTable.ForeignKeys[0].RefTable = SqueaksTable
	PoolPassesTable.ForeignKeys[0].RefTable = PoolsTable
	PoolPassesTable.ForeignKeys[1].RefTable = UsersTable
	SqueaksTable.ForeignKeys[0].RefTable = UsersTable
	SqueaksTable.ForeignKeys[1].RefTable = UsersTable
	UserRolesTable.ForeignKeys[0].RefTable = UsersTable
	UserRolesTable.ForeignKeys[1].RefTable = RolesTable
	UserFollowingTable.ForeignKeys[0].RefTable = UsersTable
	UserFollowingTable.ForeignKeys[1].RefTable = UsersTable
}
