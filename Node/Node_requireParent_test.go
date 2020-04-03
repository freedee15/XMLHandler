package Node

import (
	"fmt"
)

func iCreateANode() error {

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	node = &Node{}
	return nil

}

func iLabelTheNode(arg1 string) error {

	if node == nil {
		return fmt.Errorf("no node")
	}

	if shouldFail == true {

		shouldFail = false

		if e := node.SetLabel(arg1); e == nil {
			return fmt.Errorf("failed to throw error")
		} else {
			getError = e
		}

		return nil

	}

	if e := node.SetLabel(arg1); e != nil {
		return e
	}
	return nil

}

func iAddTheNodeToTheTree() error {

	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}
	if node == nil {
		return fmt.Errorf("no node")
	}

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	if e := nodeTree.AddChild(node); e != nil {
		return e
	}

	return nil

}

func iGiveTheNodeAValueOf(arg1 string) error {

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}
	if node == nil {
		return fmt.Errorf("no node")
	}

	if e := node.SetData(arg1); e != nil {
		return e
	}
	return nil

}

func iRetrieveTheNodeDataAsA(arg1 string) error {

	if node == nil {
		return fmt.Errorf("no node")
	}

	switch arg1 {

	case "bool":
		if shouldFail == true {

			shouldFail = false

			if _, e := node.GetDataAsBool(); e == nil {
				return fmt.Errorf("failed to throw error")
			} else {
				getError = e
			}

			return nil

		}

		if _, e := node.GetDataAsBool(); e != nil {
			return e
		}

	case "float":
		if shouldFail == true {

			shouldFail = false

			if _, e := node.GetDataAsFloat(); e == nil {
				return fmt.Errorf("failed to throw error")
			} else {
				getError = e
			}

			return nil

		}

		if _, e := node.GetDataAsFloat(); e != nil {
			return e
		}

	case "int":
		if shouldFail == true {

			shouldFail = false

			if _, e := node.GetDataAsInt(); e == nil {
				return fmt.Errorf("failed to throw error")
			} else {
				getError = e
			}

			return nil

		}

		if _, e := node.GetDataAsInt(); e != nil {
			return e
		}

	case "string":
		if shouldFail == true {

			shouldFail = false

			if _, e := node.GetData(); e == nil {
				return fmt.Errorf("failed to throw error")
			} else {
				getError = e
			}

			return nil

		}

		if _, e := node.GetData(); e != nil {
			return e
		}

	}

	return nil

}

func dataFromNodeShouldBe(arg1, arg2 string) error {

	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

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

func iGetTheNodesLabel() error {

	if node == nil {
		return fmt.Errorf("no node")
	}

	if shouldFail == true {

		shouldFail = false

		if _, e := node.GetLabel(); e == nil {
			return fmt.Errorf("failed to throw error")
		} else {
			getError = e
		}

		return nil

	}

	if _, e := node.GetLabel(); e != nil {
		return e
	}

	return nil

}

func iGetTheChildrenArrayOfTheNode() error {

	if node == nil {
		return fmt.Errorf("no node")
	}

	if shouldFail == true {

		shouldFail = false

		if _, e := node.GetChildren(); e == nil {
			return fmt.Errorf("failed to throw error")
		} else {
			getError = e
		}

		return nil

	}

	if _, e := node.GetChildren(); e != nil {
		return e
	}

	return nil

}

func iGetTheParentOfTheNode() error {

	if node == nil {
		return fmt.Errorf("no node")
	}

	if shouldFail == true {

		shouldFail = false
		if _, e := node.GetParent(); e == nil {
			return fmt.Errorf("failed to throw error")
		} else {
			getError = e
		}

		return nil

	}

	if _, e := node.GetParent(); e != nil {
		return e
	}

	return nil

}
