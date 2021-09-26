package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"homework-backend/graph/generated"
	"homework-backend/graph/model"
	"log"
	"os"
	"reflect"
	"strconv"
)

func (r *mutationResolver) SubmitAnswer(ctx context.Context, answer *model.AnswerInput) (model.Answer, error) {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	// mutex ensures integrity of ids
	answerID := r.lastAnswerID + 1
	question, err := r.Query().QuestionByID(ctx, answer.QuestionID)
	if err != nil {
		return nil, err
	}
	if question == nil {
		return nil, fmt.Errorf("question %s not found", answer.QuestionID)
	}
	var newAnswer model.Answer
	switch v := question.(type) {
	case model.ChoiceQuestion:
		// validation
		if answer.OptionID == nil {
			return nil, fmt.Errorf("submitted answer is not a ChoiceAnswer")
		}
		var validateOptions bool
		for _, opt := range v.Options {
			if opt.ID == *answer.OptionID {
				validateOptions = true
				break
			}
		}
		if !validateOptions {
			return nil, fmt.Errorf("option id %s is not a valid question option", *answer.OptionID)
		}
		// validation
		newAnswer = &model.ChoiceAnswer{
			ID:             strconv.Itoa(answerID),
			QuestionID:     answer.QuestionID,
			SelectedOption: *answer.OptionID,
		}
	case model.TextQuestion:
		// validation
		if answer.Text == nil {
			return nil, fmt.Errorf("submitted answer is not a TextAnswer")
		}
		// validation
		newAnswer = &model.TextAnswer{
			ID:         strconv.Itoa(answerID),
			QuestionID: answer.QuestionID,
			Text:       *answer.Text,
		}
	default:
		return nil, fmt.Errorf("%s type answer not implemented", reflect.TypeOf(v))
	}
	r.answers = append(r.answers, newAnswer)
	dat, err := json.Marshal(newAnswer)
	if err != nil {
		log.Println(err.Error())
	} else {
		fmt.Fprintln(os.Stdout, string(dat))
	}

	r.lastAnswerID++ // incremented at the end to avoid unnecessary increments in case of validation errors
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

func (r *queryResolver) QuestionByID(ctx context.Context, id string) (model.Question, error) {
	questions, err := r.Query().Questions(ctx)
	if err != nil {
		return nil, err
	}
	var question model.Question
search:
	for _, q := range questions {
		switch v := q.(type) {
		case model.ChoiceQuestion:
			if v.ID == id {
				question = q
				break search
			}
		case model.TextQuestion:
			if v.ID == id {
				question = q
				break search
			}
		}
	}
	return question, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
