package Node

import (
	"fmt"
)

func iAddAChildNodeLabelledToTheNodeTree(arg1 string) error {

	newNode := Node{}
	nodeTree.AddChild(&newNode)
	newNode.label = arg1
	return nil

}

func nodeShouldHaveTheParent(arg1, arg2 string) error {

	c, e := nodeTree.GetChildFromLabel(arg1)
	if e != nil {
		return e
	}
	switch (c.parent).(type) {

	case *Node:
		if c.parent.(*Node).label != arg2 {
			return fmt.Errorf("expected %s, got %s", arg2, c.parent.(*Node).label)
		}

	case *Tree:
		if c.parent.(*Tree).Label != arg2 {
			return fmt.Errorf("expected %s, got %s", arg2, c.parent.(*Tree).Label)
		}

	}

	return nil

}
