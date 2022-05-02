// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ReadsColumns holds the columns for the "reads" table.
	ReadsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "result", Type: field.TypeInt},
		{Name: "sentence_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// ReadsTable holds the schema information for the "reads" table.
	ReadsTable = &schema.Table{
		Name:       "reads",
		Columns:    ReadsColumns,
		PrimaryKey: []*schema.Column{ReadsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "reads_sentenses_reads",
				Columns:    []*schema.Column{ReadsColumns[2]},
				RefColumns: []*schema.Column{SentensesColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "reads_users_reads",
				Columns:    []*schema.Column{ReadsColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// SentensesColumns holds the columns for the "sentenses" table.
	SentensesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "chinese", Type: field.TypeString},
		{Name: "pinyin", Type: field.TypeString},
		{Name: "english", Type: field.TypeString},
	}
	// SentensesTable holds the schema information for the "sentenses" table.
	SentensesTable = &schema.Table{
		Name:       "sentenses",
		Columns:    SentensesColumns,
		PrimaryKey: []*schema.Column{SentensesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ReadsTable,
		SentensesTable,
		UsersTable,
	}
)

func init() {
	ReadsTable.ForeignKeys[0].RefTable = SentensesTable
	ReadsTable.ForeignKeys[1].RefTable = UsersTable
}
