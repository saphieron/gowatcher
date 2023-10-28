package cli_test

import (
	"os"
	"testing"

	"github.com/saphieron/gowatcher/cli"
	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
)

type ParseFlagsSuite struct {
	suite.Suite
	asserter *require.Assertions

	oldArgs []string
}

func (suite *ParseFlagsSuite) SetupSuite() {
	suite.oldArgs = os.Args
}

func (suite *ParseFlagsSuite) SetupTest() {
	suite.asserter = require.New(suite.T())

	pflag.CommandLine = pflag.NewFlagSet(os.Args[0], pflag.ExitOnError) //resets the lib
}

func (suite *ParseFlagsSuite) TearDownSuite() {
	os.Args = suite.oldArgs
}

func TestParseFlagsSuite(t *testing.T) {
	suite.Run(t, new(ParseFlagsSuite))
}

func (suite *ParseFlagsSuite) TestSettingInterval() {
	fakeCliArgs := []string{"myBin", "-n", "999", "\"ls -la\""}
	os.Args = fakeCliArgs

	options, rest := cli.ParseFlags()

	suite.asserter.Equal(float32(999.0), options.Interval)
	suite.asserter.Contains(rest, fakeCliArgs[len(fakeCliArgs)-1])
}

func (suite *ParseFlagsSuite) TestDefaultValues() {
	fakeCliArgs := []string{"myBin"}
	os.Args = fakeCliArgs

	options, rest := cli.ParseFlags()

	suite.asserter.Equal(float32(2), options.Interval)
	suite.asserter.False(options.Help)
	suite.asserter.False(options.Version)

	suite.asserter.Empty(rest)
}
