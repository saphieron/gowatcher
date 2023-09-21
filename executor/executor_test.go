package executor_test

import (
	"context"
	"testing"

	"github.com/saphieron/gowatcher/executor"
	"github.com/stretchr/testify/suite"
)

// ################################
// #  NewExecutorSuite
// ################################

type NewExecutorSuite struct {
	suite.Suite
}

func (suite *NewExecutorSuite) SetupTest() {
}

func TestNewExecutorSuite(t *testing.T) {
	suite.Run(t, new(NewExecutorSuite))
}

const simpleCommand = "ls"

var commandArgs = []string{"-la", "./"}

func (suite *NewExecutorSuite) TestGetsExecutor() {
	cmd := executor.NewCommandExecutor(simpleCommand, commandArgs...)
	suite.Require().NotNil(cmd)
}

func (suite *NewExecutorSuite) TestGetsExecutorWithContext() {
	ctx := context.Background()
	cmd := executor.NewCommandExecutorWithContext(ctx, simpleCommand, commandArgs...)
	suite.Require().NotNil(cmd)
}
