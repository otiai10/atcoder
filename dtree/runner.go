package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Runner ...
type Runner struct {
	InputSize int
	InputV    []int
	InputU    []int
}

// Run ...
func (r *Runner) Run(stdin io.Reader, stdout io.Writer) error {
	scanner := bufio.NewScanner(stdin)
	if scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		r.InputSize = n
	}
	if scanner.Scan() {
		for _, s := range strings.Split(scanner.Text(), " ") {
			if e, err := strconv.Atoi(s); err == nil {
				r.InputV = append(r.InputV, e)
			}
		}
	}
	if scanner.Scan() {
		for _, s := range strings.Split(scanner.Text(), " ") {
			if e, err := strconv.Atoi(s); err == nil {
				r.InputU = append(r.InputU, e)
			}
		}
	}

	v := &Tree{Size: r.InputSize, Nodes: make([]Node, r.InputSize+1)}
	v.Parse(r.InputV)

	u := &Tree{Size: r.InputSize - 1, Nodes: make([]Node, r.InputSize)}
	u.Parse(r.InputU)

	leaves := v.Compare(u)
	result := fmt.Sprintf("%v", leaves)

	fmt.Fprintf(stdout, "%s", result[1:len(result)-1])

	return nil
}
