package src

import (
	"context"
	"io"
)

// Application ...
type Application struct {
	CodeRunner
}

// NewApplication creates a new Application
func NewApplication(codeRunner CodeRunner) Application {
	return Application{
		CodeRunner: codeRunner,
	}
}

// RunGo ...
func (a *Application) RunGo(ctx context.Context, in Input) (Output, error) {
	body, err := a.CodeRunner.RunGo(ctx, in.Src)
	if err != nil {
		return Output{}, err
	}
	return Output{
		Body:  body,
		Error: "",
	}, nil
}

// RunRuby ...
func (a *Application) RunRuby(ctx context.Context, in Input) (Output, error) {
	body, err := a.CodeRunner.RunRuby(ctx, in.Src)
	if err != nil {
		return Output{}, err
	}
	return Output{
		Body:  body,
		Error: "",
	}, nil
}

// CodeRunner ...
type CodeRunner interface {
	RunGo(context.Context, io.Reader) (string, error)
	RunRuby(context.Context, io.Reader) (string, error)
}

// Input ...
type Input struct {
	Src io.Reader
}

// Output ...
type Output struct {
	Error string `json:"error"`
	Body  string `json:"body"`
}
