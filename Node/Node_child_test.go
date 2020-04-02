package Node

import (
	"fmt"
)

func iCreateAChildNodeWithTheLabel(arg1 string) error {

	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	childNode := &Node{}
	if e := nodeTree.AddChild(childNode); e != nil {
		return e
	}
	childNode.label = arg1
	return nil

}

func iGiveChildNodeAValueOf(arg1, arg2 string) error {

	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}

	c, e := nodeTree.GetChildFromLabel(arg1)

	if shouldFail == true {

		shouldFail = false

		if e == nil {
			return fmt.Errorf("failed to throw error")
		} else {
			getError = e
		}

		return nil

	}

	if e != nil {
		return e
	}
	if e = c.SetData(arg2); e != nil {
		return e
	}

	return nil

}

func iRetrieveDataFromChildNode(arg1 string) error {

	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}

	c, e := nodeTree.GetChildFromLabel(arg1)

	if shouldFail == true {

		shouldFail = false

		if e == nil {
			return fmt.Errorf("failed to throw error")
		} else {
			getError = e
		}

		return nil

	}

	if e != nil {
		return e
	}
	retrieved = c.data

	return nil

}

func theRetrievedNodeDataShouldBe(arg1 string) error {

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	if retrieved != arg1 {
		return fmt.Errorf("expected %s, got %s", arg1, retrieved)
	}

	return nil

}

func iCreateAChildNodeInWithTheLabel(arg1, arg2 string) error {

	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}

	c := &Node{}
	n, e := nodeTree.GetChildFromLabel(arg1)
	if e != nil {
		return e
	}
	if e = n.AddChild(c); e != nil {
		return e
	}
	if e = c.SetLabel(arg2); e != nil {
		return e
	}

	return nil

}

func theChildrenArrayOfTheNodeTreeShouldHaveMembers(arg1 int) error {

	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}

	if len(nodeTree.GetChildren()) != arg1 {
		return fmt.Errorf("expected %d, got %d", arg1, len(nodeTree.GetChildren()))
	}

	return nil

}

func theChildrenArrayOfNodeShouldHaveMembers(arg1 string, arg2 int) error {

	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}

	n, e := nodeTree.GetChildFromLabel(arg1)
	if e != nil {
		return e
	}
	c, e := n.GetChildren()
	if e != nil {
		return e
	}
	if len(c) != arg2 {
		return fmt.Errorf("expected %d, got %d", arg2, len(c))
	}

	return nil

}
