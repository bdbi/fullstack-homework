package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"homework-backend/graph/generated"
	"homework-backend/graph/model"
	"reflect"
	"strconv"
)

func (r *mutationResolver) SubmitAnswer(ctx context.Context, answer *model.AnswerInput) (model.Answer, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.lastAnswerID++
	questions, err := r.Query().Questions(ctx)
	if err != nil {
		return nil, err
	}
	var question model.Question
	for _, q := range questions {
		switch v := q.(type) {
		case model.ChoiceQuestion:
			if v.ID == answer.QuestionID {
				question = v
				break
			}
		case model.TextQuestion:
			if v.ID == answer.QuestionID {
				question = v
				break
			}
		}
	}
	//todo check if question == nil (question not found)
	var newAnswer model.Answer
	switch v := question.(type) {
	case model.ChoiceQuestion:
		newAnswer = &model.ChoiceAnswer{
			ID:             strconv.Itoa(r.lastAnswerID),
			QuestionID:     answer.QuestionID,
			SelectedOption: *answer.OptionID,
		}
	case model.TextQuestion:
		newAnswer = &model.TextAnswer{
			ID:         strconv.Itoa(r.lastAnswerID),
			QuestionID: answer.QuestionID,
			Text:       *answer.Text,
		}
	default:
		return nil, fmt.Errorf("%s type answer not implemented", reflect.TypeOf(v))
	}
	r.answers = append(r.answers, newAnswer)
	return newAnswer, nil
}

func (r *queryResolver) Questions(ctx context.Context) ([]model.Question, error) {
	return []model.Question{
		model.ChoiceQuestion{
			ID:     "100",
			Body:   "Where does the sun set?",
			Weight: 0.5,
			Options: []*model.Option{
				{ID: "200", Body: "East", Weight: 0},
				{ID: "201", Body: "West", Weight: 1},
			},
		},
		model.TextQuestion{
			ID:     "101",
			Body:   "What is your favourite food?",
			Weight: 1,
		},
	}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
