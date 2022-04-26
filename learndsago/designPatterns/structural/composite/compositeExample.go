package composite

import "fmt"

type IComposite interface {
	Perform()
}

type Leaflet struct {
	name string
}

func (leaf *Leaflet) Perform() {
	fmt.Println("Leaflet " + leaf.name)
}

type Branch struct {
	leafs    []Leaflet
	name     string
	branches []Branch
}

func (branch *Branch) Perform() {
	fmt.Println("Branch: " + branch.name)
	for _, leaf := range branch.leafs {
		leaf.Perform()
	}
	for _, branch := range branch.branches {
		branch.Perform()
	}
}

func (branch *Branch) AddLeaf(leaf Leaflet) {
	branch.leafs = append(branch.leafs, leaf)
}

func (branch *Branch) AddBranch(newBranch Branch) {
	branch.branches = append(branch.branches, newBranch)
}

func (branch *Branch) getLeaflets() []Leaflet {
	return branch.leafs
}

func Example() {
	var branch1 = &Branch{name: "b1"}
	var leaf1 = Leaflet{name: "l1"}
	var leaf2 = Leaflet{name: "l2"}
	var leaf3 = Leaflet{name: "l3"}
	var leaf4 = Leaflet{name: "l4"}
	var branch2 = Branch{name: "b2"}

	branch1.AddLeaf(leaf1)
	branch1.AddLeaf(leaf2)
	branch2.AddLeaf(leaf3)
	branch2.AddLeaf(leaf4)

	branch1.AddBranch(branch2)
	branch1.Perform()

}
