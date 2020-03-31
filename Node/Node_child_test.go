package Node

import (
	"fmt"
)

var retrieved string

func iCreateAChildNodeWithTheLabel(arg1 string) error {

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	childNode := &Node{}
	nodeTree.AddChild(childNode)
	childNode.label = arg1
	return nil

}

func iGiveChildNodeAValueOf(arg1, arg2 string) error {

	c, e := nodeTree.GetChildFromLabel(arg1)

	if shouldFail == true {

		shouldFail = false

		if e == nil {
			return fmt.Errorf("failed to throw error")
		}

		return nil

	}

	if e != nil {
		return e
	}
	c.SetData(arg2)

	return nil

}

func iRetrieveDataFromChildNode(arg1 string) error {

	c, e := nodeTree.GetChildFromLabel(arg1)

	if shouldFail == true {

		shouldFail = false

		if e == nil {
			return fmt.Errorf("failed to throw error")
		}

		return nil

	}

	if e != nil {
		return e
	}
	retrieved = c.GetData()

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
