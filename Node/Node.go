package Node

import (
	"fmt"
	"strconv"
	"strings"
)

type Node struct {
	Label    string
	data     string
	children []*Node
}

func (n *Node) AddChild(child *Node) {

	n.children = append(n.children, child)

}

func (n *Node) GetChildFromLabel(label string) (*Node, error) {

	if strings.Replace(label, " ", "", -1) == "" {
		return nil, fmt.Errorf("empty label")
	}

	for _, c := range n.children {

		if c.Label == label {
			return c, nil
		}

	}

	return nil, fmt.Errorf("no child with label \"%s\"", label)

}

func (n *Node) GetData() string {

	return n.data

}

func (n *Node) GetDataAsBool() (bool, error) {

	if b, e := strconv.ParseBool(n.data); e != nil {
		return false, e
	} else {
		return b, nil
	}

}

func (n *Node) GetDataAsFloat() (float64, error) {

	if f, e := strconv.ParseFloat(n.data, 64); e != nil {
		return 0, e
	} else {
		return f, nil
	}

}

func (n *Node) GetDataAsInt() (int, error) {

	if i, e := strconv.Atoi(n.data); e != nil {
		return 0, e
	} else {
		return i, nil
	}

}

func (n *Node) SetData(s string) {

	n.data = s

}

func (n *Node) SetDataFromBool(b bool) {

	n.data = strconv.FormatBool(b)

}

func (n *Node) SetDataFromFloat(f float64) {

	n.data = strconv.FormatFloat(f, 'f', -1, 64)

}

func (n *Node) SetDataFromInt(i int) {

	n.data = strconv.Itoa(i)

}
