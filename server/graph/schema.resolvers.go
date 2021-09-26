package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"encoding/json"
	"fmt"
	"homework-backend/ent"
	"homework-backend/ent/question"
	"homework-backend/graph/generated"
	"homework-backend/graph/model"
	"log"
	"os"
	"reflect"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

func (r *mutationResolver) SubmitAnswer(ctx context.Context, answer *model.AnswerInput) (model.Answer, error) {
	client, err := ent.Open("sqlite3", os.Getenv("SQLITE_CONN"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	question, err := r.Query().QuestionByID(ctx, answer.QuestionID)
	if err != nil {
		return nil, err
	}
	if question == nil {
		return nil, fmt.Errorf("question %s not found", answer.QuestionID)
	}
	var answerCreate *ent.AnswerCreate
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
		// end-validation
		answerCreate = client.Answer.Create().SetQuestionID(answer.QuestionID).SetOptionID(*answer.OptionID)
	case model.TextQuestion:
		// validation
		if answer.Text == nil {
			return nil, fmt.Errorf("submitted answer is not a TextAnswer")
		}
		// end-validation
		answerCreate = client.Answer.Create().SetQuestionID(answer.QuestionID).SetBody(*answer.Text)
	default:
		return nil, fmt.Errorf("%s type answer not implemented", reflect.TypeOf(v))
	}

	a, err := answerCreate.Save(ctx)
	if err != nil {
		return nil, err
	}
	dat, err := json.Marshal(a)
	if err != nil {
		log.Println(err.Error())
	} else {
		fmt.Fprintln(os.Stdout, string(dat))
	}

	var newAnswer model.Answer
	switch question.(type) {
	case model.ChoiceQuestion:
		newAnswer = model.ChoiceAnswer{
			ID:             a.ID.String(),
			QuestionID:     a.QuestionID,
			SelectedOption: *a.OptionID,
		}
	case model.TextQuestion:
		newAnswer = model.TextAnswer{
			ID:         a.ID.String(),
			QuestionID: a.QuestionID,
			Text:       *a.Body,
		}
	}
	return newAnswer, nil
}

func (r *queryResolver) Questions(ctx context.Context) ([]model.Question, error) {
	client, err := ent.Open("sqlite3", os.Getenv("SQLITE_CONN"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	questions, err := client.Question.Query().WithOptions().All(ctx)
	if err != nil {
		return nil, err
	}
	var gqlQuestions []model.Question = make([]model.Question, 0)
	for _, question := range questions {
		switch question.Type {
		case "ChoiceQuestion":
			var gqpOpts []*model.Option = make([]*model.Option, 0)
			for _, opt := range question.Edges.Options {
				gqpOpts = append(gqpOpts, &model.Option{
					ID:     opt.ID.String(),
					Weight: opt.Weight,
					Body:   opt.Body,
				})
			}
			gqlQuestions = append(gqlQuestions, model.ChoiceQuestion{
				ID:      question.ID.String(),
				Body:    *question.Body,
				Weight:  question.Weight,
				Options: gqpOpts,
			})
		case "TextQuestion":
			gqlQuestions = append(gqlQuestions, model.TextQuestion{
				ID:     question.ID.String(),
				Weight: question.Weight,
				Body:   *question.Body,
			})
		}
	}
	return gqlQuestions, nil
}

func (r *queryResolver) QuestionByID(ctx context.Context, id string) (model.Question, error) {
	client, err := ent.Open("sqlite3", os.Getenv("SQLITE_CONN"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, err
	}
	q, err := client.Question.Query().WithOptions().Where(question.ID(uid)).Only(ctx)
	if err != nil {
		return nil, err
	}
	var question model.Question
	switch q.Type {
	case "ChoiceQuestion":
		var gqpOpts []*model.Option = make([]*model.Option, 0)
		for _, opt := range q.Edges.Options {
			gqpOpts = append(gqpOpts, &model.Option{
				ID:     opt.ID.String(),
				Weight: opt.Weight,
				Body:   opt.Body,
			})
		}
		question = model.ChoiceQuestion{
			ID:      q.ID.String(),
			Weight:  q.Weight,
			Body:    *q.Body,
			Options: gqpOpts,
		}
	case "TextQuestion":
		question = model.TextQuestion{
			ID:     q.ID.String(),
			Weight: q.Weight,
			Body:   *q.Body,
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
