package Node

import (
	"fmt"
)

func iLabelTheNode(arg1 string) error {

	if shouldFail == true {

		shouldFail = false

		if node.SetLabel(arg1) == nil {
			return fmt.Errorf("failed to throw error")
		}

		return nil

	}

	if e := node.SetLabel(arg1); e != nil {
		return e
	}
	return nil

}

func iAddTheNodeToTheTree() error {

	if e := nodeTree.AddChild(node); e != nil {
		return e
	}

	return nil

}

func dataFromNodeShouldBe(arg1, arg2 string) error {

	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}

	if c, e := nodeTree.GetChildFromLabel(arg1); e != nil {
		return e
	} else if c.data != arg2 {
		return fmt.Errorf("expected %s, got %s", arg2, c.data)
	}

	return nil

}
