package cli_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ExampleTestSuite struct {
	suite.Suite
	fakeCliArgs []string
	oldArgs     []string
}

func (suite *ExampleTestSuite) SetupTest() {
	fmt.Println("in setup")
	fmt.Print(os.Args)
	fmt.Println("")
	suite.oldArgs = os.Args
}

func (suite *ExampleTestSuite) AfterTest(suiteName, testName string) {
	os.Args = suite.oldArgs
	fmt.Println("after test")
	fmt.Print(os.Args)
	fmt.Println("")
}

func TestExampleTestSuite(t *testing.T) {
	suite.Run(t, new(ExampleTestSuite))
}

func (suite *ExampleTestSuite) TestExample() {
	fmt.Println("in the test")
	suite.fakeCliArgs = []string{"myBin", "-n", "999"}
	os.Args = suite.fakeCliArgs
	fmt.Print(os.Args)
	fmt.Println("")
}
