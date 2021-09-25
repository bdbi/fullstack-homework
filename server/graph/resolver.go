package graph

import (
	"homework-backend/graph/model"
	"sync"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	mutex        sync.Mutex
	answers      []model.Answer
	lastAnswerID int
}
