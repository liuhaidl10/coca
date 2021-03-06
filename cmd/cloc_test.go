package cmd

import (
	"github.com/phodal/coca/cocatest/testcase"
	"testing"
)

func TestCloc(t *testing.T) {
	analysis := []testcase.CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p .",
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []testcase.CmdTestCase{{
		Name:   "cloc",
		Cmd:    "cloc",
		Golden: "",
	}}
	RunTestCmd(t, tests)
}

func TestClocByDirectory(t *testing.T) {
	abs := "../_fixtures/cloc"

	analysis := []testcase.CmdTestCase{{
		Name:   "analysis",
		Cmd:    "analysis -p " + abs,
		Golden: "",
	}}
	RunTestCmd(t, analysis)

	tests := []testcase.CmdTestCase{{
		Name:   "cloc",
		Cmd:    "cloc " + abs + " --by-directory --include-ext=java",
		Golden: "testdata/cloc_directory.txt",
	}}
	RunTestCmd(t, tests)
}