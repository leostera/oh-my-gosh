package omg

import (
	"runtime"
	"testing"
)

var platform struct {
	cmd  string
	flag string
}

func init() {
	if runtime.GOOS == "windows" {
		platform.cmd = "cmd"
		platform.flag = "/C"
	} else {
		platform.cmd = "sh"
		platform.flag = "-c"
	}
}

/******************************************************************************/
// Fixtures
/******************************************************************************/

func fixtureCmd() string       { return platform.cmd }
func fixtureArgs() []string    { return []string{platform.flag, "exit 0"} }
func fixtureBadArgs() []string { return []string{platform.flag, "exit 2"} }

func fixtureCmdSlice(args []string) []string {
	return append([]string{fixtureCmd()}, args...)
}

// Tests

func TestRunSuccessfully(t *testing.T) {
	ok := Run(fixtureCmdSlice(fixtureArgs()))
	if ok != 0 {
		t.Fatalf("%d should be 0", ok)
	}
}

func TestRunExit(t *testing.T) {
	err := Run(fixtureCmdSlice(fixtureBadArgs()))
	if err != 2 {
		t.Fatalf("%d should not be 2", err)
	}
}

// Benchmarks

func BenchmarkRunSuccessfully(b *testing.B) {
	cmd := fixtureCmdSlice(fixtureArgs())
	for n := 0; n < b.N; n++ {
		Run(cmd)
	}
}

func BenchmarkRunExit(b *testing.B) {
	cmd := fixtureCmdSlice(fixtureBadArgs())
	for n := 0; n < b.N; n++ {
		Run(cmd)
	}
}
