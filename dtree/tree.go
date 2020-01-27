package main

// ID ...
type ID int

// Degree ...
type Degree int

// Node ...
type Node struct {
	ID     ID
	Parent ID
	Degree Degree
}

// Tree ...
type Tree struct {
	Size  int
	Nodes []Node
	Shape map[Degree][]Node
}

// Parse ...
// TODO: Refactor
func (t *Tree) Parse(vector []int) {

	// Parse each element
	for i, p := range vector {
		// Self
		t.Nodes[i].ID = ID(i + 1)
		t.Nodes[i].Degree++
		// Parent
		t.Nodes[i].Parent = ID(p)
		t.Nodes[p-1].ID = ID(p)
		t.Nodes[p-1].Degree++

		// TODO: Fix
		if t.Nodes[p-1].Parent == 0 {
			t.Nodes[p-1].Parent = ID(i + 1)
		}
	}

	// Define shape
	t.Shape = map[Degree][]Node{}
	for _, node := range t.Nodes {
		if n, ok := t.Shape[node.Degree]; ok {
			t.Shape[node.Degree] = append(n, node)
		} else {
			t.Shape[node.Degree] = []Node{node}
		}
	}
}

// Compare ...
func (t *Tree) Compare(target *Tree) []ID {
	dict := t.findDegreeLostNodesDictionary(target)
	leaves := []ID{}
	for _, leaf := range t.Shape[Degree(1)] {
		if dict[leaf.Parent] {
			leaves = append(leaves, leaf.ID)
		}
	}
	return leaves
}

// findDegreeLostNodesDictionary ...
func (t *Tree) findDegreeLostNodesDictionary(target *Tree) map[ID]bool {
	dict := map[ID]bool{}
	for degree, nodes := range t.Shape {
		if len(nodes)-len(target.Shape[degree]) == 1 && degree != 1 {
			for _, n := range nodes {
				dict[n.ID] = true
			}
			return dict
		}
	}
	return dict
}
