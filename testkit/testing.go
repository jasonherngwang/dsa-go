package testkit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Include desired assertions in exports
var Equal = assert.Equal

type TestCase[Input any] struct {
	Name     string
	Args     Input
	Expected any
}

func Run_Tests[Input any](t *testing.T, tests []TestCase[Input], testFunc func(TestCase[Input])) {
	for _, test := range tests {
		t.Run(test.Name, func(*testing.T) { testFunc(test) })
	}
}
