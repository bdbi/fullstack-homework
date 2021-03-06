// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AnswersColumns holds the columns for the "answers" table.
	AnswersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "question_id", Type: field.TypeString},
		{Name: "body", Type: field.TypeString, Nullable: true},
		{Name: "option_id", Type: field.TypeString, Nullable: true},
	}
	// AnswersTable holds the schema information for the "answers" table.
	AnswersTable = &schema.Table{
		Name:       "answers",
		Columns:    AnswersColumns,
		PrimaryKey: []*schema.Column{AnswersColumns[0]},
	}
	// OptsColumns holds the columns for the "opts" table.
	OptsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "body", Type: field.TypeString},
		{Name: "weight", Type: field.TypeFloat64},
	}
	// OptsTable holds the schema information for the "opts" table.
	OptsTable = &schema.Table{
		Name:       "opts",
		Columns:    OptsColumns,
		PrimaryKey: []*schema.Column{OptsColumns[0]},
	}
	// QuestionsColumns holds the columns for the "questions" table.
	QuestionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID},
		{Name: "weight", Type: field.TypeFloat64},
		{Name: "type", Type: field.TypeString},
		{Name: "body", Type: field.TypeString, Nullable: true},
	}
	// QuestionsTable holds the schema information for the "questions" table.
	QuestionsTable = &schema.Table{
		Name:       "questions",
		Columns:    QuestionsColumns,
		PrimaryKey: []*schema.Column{QuestionsColumns[0]},
	}
	// QuestionOptionsColumns holds the columns for the "question_options" table.
	QuestionOptionsColumns = []*schema.Column{
		{Name: "question_id", Type: field.TypeUUID},
		{Name: "opt_id", Type: field.TypeUUID},
	}
	// QuestionOptionsTable holds the schema information for the "question_options" table.
	QuestionOptionsTable = &schema.Table{
		Name:       "question_options",
		Columns:    QuestionOptionsColumns,
		PrimaryKey: []*schema.Column{QuestionOptionsColumns[0], QuestionOptionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "question_options_question_id",
				Columns:    []*schema.Column{QuestionOptionsColumns[0]},
				RefColumns: []*schema.Column{QuestionsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "question_options_opt_id",
				Columns:    []*schema.Column{QuestionOptionsColumns[1]},
				RefColumns: []*schema.Column{OptsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AnswersTable,
		OptsTable,
		QuestionsTable,
		QuestionOptionsTable,
	}
)

func init() {
	QuestionOptionsTable.ForeignKeys[0].RefTable = QuestionsTable
	QuestionOptionsTable.ForeignKeys[1].RefTable = OptsTable
}
