package main

import (
	"bytes"
	"os"
	"testing"

	. "github.com/otiai10/mint"
)

func TestMain(m *testing.M) {
	code := m.Run()
	os.Exit(code)
}

type testcase struct {
	Label    string
	Input    string
	Expected string
}

var cases = []testcase{
	testcase{
		Input:    "10\n2 9 2 2 2 2 6 7 10 11\n2 3 4 10 4 5 4 4 4",
		Expected: "8 11",
	},
	testcase{
		Input:    "5\n6 1 1 1 1\n5 5 5 5",
		Expected: "2 3 4 5 6",
	},
	testcase{
		Input:    "5\n4 3 4 5 6\n2 3 4 5",
		Expected: "1",
	},
}

func TestRunner_Run(t *testing.T) {

	for _, tc := range cases {
		stdin := bytes.NewBuffer([]byte(tc.Input))
		stdout := bytes.NewBuffer(nil)
		runner := new(Runner)
		err := runner.Run(stdin, stdout)
		Expect(t, err).ToBe(nil)
		Expect(t, stdout.String()).ToBe(tc.Expected)
	}

}
