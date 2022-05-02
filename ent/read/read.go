// Code generated by entc, DO NOT EDIT.

package read

const (
	// Label holds the string label denoting the read type in the database.
	Label = "read"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldResult holds the string denoting the result field in the database.
	FieldResult = "result"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeSentence holds the string denoting the sentence edge name in mutations.
	EdgeSentence = "sentence"
	// Table holds the table name of the read in the database.
	Table = "reads"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "reads"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// SentenceTable is the table that holds the sentence relation/edge.
	SentenceTable = "reads"
	// SentenceInverseTable is the table name for the Sentense entity.
	// It exists in this package in order to avoid circular dependency with the "sentense" package.
	SentenceInverseTable = "sentenses"
	// SentenceColumn is the table column denoting the sentence relation/edge.
	SentenceColumn = "sentence_id"
)

// Columns holds all SQL columns for read fields.
var Columns = []string{
	FieldID,
	FieldResult,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "reads"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"sentence_id",
	"user_id",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
