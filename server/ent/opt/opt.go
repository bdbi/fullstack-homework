// Code generated by entc, DO NOT EDIT.

package opt

import (
	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the opt type in the database.
	Label = "opt"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBody holds the string denoting the body field in the database.
	FieldBody = "body"
	// FieldWeight holds the string denoting the weight field in the database.
	FieldWeight = "weight"
	// EdgeQuestion holds the string denoting the question edge name in mutations.
	EdgeQuestion = "question"
	// Table holds the table name of the opt in the database.
	Table = "opts"
	// QuestionTable is the table that holds the question relation/edge. The primary key declared below.
	QuestionTable = "question_options"
	// QuestionInverseTable is the table name for the Question entity.
	// It exists in this package in order to avoid circular dependency with the "question" package.
	QuestionInverseTable = "questions"
)

// Columns holds all SQL columns for opt fields.
var Columns = []string{
	FieldID,
	FieldBody,
	FieldWeight,
}

var (
	// QuestionPrimaryKey and QuestionColumn2 are the table columns denoting the
	// primary key for the question relation (M2M).
	QuestionPrimaryKey = []string{"question_id", "opt_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)