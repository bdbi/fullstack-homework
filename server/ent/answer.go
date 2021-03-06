// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"homework-backend/ent/answer"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
)

// Answer is the model entity for the Answer schema.
type Answer struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// QuestionID holds the value of the "question_id" field.
	QuestionID string `json:"question_id,omitempty"`
	// Body holds the value of the "body" field.
	Body *string `json:"body,omitempty"`
	// OptionID holds the value of the "option_id" field.
	OptionID *string `json:"option_id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Answer) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case answer.FieldQuestionID, answer.FieldBody, answer.FieldOptionID:
			values[i] = new(sql.NullString)
		case answer.FieldID:
			values[i] = new(uuid.UUID)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Answer", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Answer fields.
func (a *Answer) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case answer.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				a.ID = *value
			}
		case answer.FieldQuestionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field question_id", values[i])
			} else if value.Valid {
				a.QuestionID = value.String
			}
		case answer.FieldBody:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field body", values[i])
			} else if value.Valid {
				a.Body = new(string)
				*a.Body = value.String
			}
		case answer.FieldOptionID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field option_id", values[i])
			} else if value.Valid {
				a.OptionID = new(string)
				*a.OptionID = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Answer.
// Note that you need to call Answer.Unwrap() before calling this method if this Answer
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Answer) Update() *AnswerUpdateOne {
	return (&AnswerClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Answer entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Answer) Unwrap() *Answer {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Answer is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Answer) String() string {
	var builder strings.Builder
	builder.WriteString("Answer(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", question_id=")
	builder.WriteString(a.QuestionID)
	if v := a.Body; v != nil {
		builder.WriteString(", body=")
		builder.WriteString(*v)
	}
	if v := a.OptionID; v != nil {
		builder.WriteString(", option_id=")
		builder.WriteString(*v)
	}
	builder.WriteByte(')')
	return builder.String()
}

// Answers is a parsable slice of Answer.
type Answers []*Answer

func (a Answers) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
