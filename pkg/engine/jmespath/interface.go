package jmespath

import (
	"github.com/kyverno/go-community-jmespath/pkg/interpreter"
	"github.com/kyverno/kyverno/pkg/config"
)

type Query interface {
	Search(interface{}) (interface{}, error)
}

type Interface interface {
	Query(string) (Query, error)
	Search(string, interface{}) (interface{}, error)
}

type implementation struct {
	functionCaller interpreter.FunctionCaller
}

func New(configuration config.Configuration) Interface {
	return newImplementation(configuration)
}

func (i implementation) Query(query string) (Query, error) {
	return newJMESPath(query, i.functionCaller)
}

func (i implementation) Search(query string, data interface{}) (interface{}, error) {
	return newExecution(i.functionCaller, query, data)
}
