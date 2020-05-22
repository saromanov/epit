package epit

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExecStage(t *testing.T) {
	st := Stage{}
	assert.Error(t, execStage(st))
	assert.NoError(t, execStage(Stage{
		Command: "w",
	}))
}
