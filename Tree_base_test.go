package XMLHandler

import (
	"fmt"
)

func nodeShouldHaveTheParent(arg1, arg2 string) error {

	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

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
