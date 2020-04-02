package XMLHandler

import "fmt"

//child

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

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
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

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
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

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
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

//dataType

func iCreateAChildNode() error {

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

	childNode = &Node{}
	if e := nodeTree.AddChild(childNode); e != nil {
		return e
	}

	return nil

}

func iRetrieveTheChildNodeDataAsA(arg1 string) error {

	if childNode == nil {
		return fmt.Errorf("no child node")
	}

	switch arg1 {

	case "bool":
		if shouldFail == true {

			shouldFail = false

			if _, e := childNode.GetDataAsBool(); e == nil {
				return fmt.Errorf("failed to throw error")
			} else {
				getError = e
			}

			return nil

		}

		if _, e := childNode.GetDataAsBool(); e != nil {
			return e
		}

	case "float":
		if shouldFail == true {

			shouldFail = false

			if _, e := childNode.GetDataAsFloat(); e == nil {
				return fmt.Errorf("failed to throw error")
			} else {
				getError = e
			}

			return nil

		}

		if _, e := childNode.GetDataAsFloat(); e != nil {
			return e
		}

	case "int":
		if shouldFail == true {

			shouldFail = false

			if _, e := childNode.GetDataAsInt(); e == nil {
				return fmt.Errorf("failed to throw error")
			} else {
				getError = e
			}

			return nil

		}

		if _, e := childNode.GetDataAsInt(); e != nil {
			return e
		}

	case "string":
		if shouldFail == true {

			shouldFail = false

			if _, e := childNode.GetData(); e == nil {
				return fmt.Errorf("failed to throw error")
			} else {
				getError = e
			}

			return nil

		}

		if _, e := childNode.GetData(); e != nil {
			return e
		}

	}

	return nil

}

func iGiveTheNodeAToStore(arg1 string) error {

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	if childNode == nil {
		return fmt.Errorf("no child node")
	}

	switch arg1 {

	case "bool":
		if e := childNode.SetDataFromBool(true); e != nil {
			return e
		}

	case "float":
		if e := childNode.SetDataFromFloat(0.1); e != nil {
			return e
		}

	case "int":
		if e := childNode.SetDataFromInt(1); e != nil {
			return e
		}

	case "string":
		if e := childNode.SetData("s"); e != nil {
			return e
		}

	default:
		return fmt.Errorf("invalid type specified: %s", arg1)

	}

	return nil

}

func iCanRetrieveTheDataFromTheNodeAsA(arg1 string) error {

	if childNode == nil {
		return fmt.Errorf("no child node")
	}

	if shouldFail == true {

		shouldFail = false

		var e error

		switch arg1 {

		case "bool":
			_, e = childNode.GetDataAsBool()

		case "float":
			_, e = childNode.GetDataAsFloat()

		case "int":
			_, e = childNode.GetDataAsInt()

		default:
			return fmt.Errorf("invalid type specified: %s", arg1)
		}

		if e == nil {
			return fmt.Errorf("failed to throw error")
		} else {
			getError = e
		}

	}

	switch arg1 {

	case "bool":
		b, e := childNode.GetDataAsBool()
		if e != nil {
			return e
		}
		if b != true {
			return fmt.Errorf("expected %t, got %t", true, b)
		}

	case "float":
		f, e := childNode.GetDataAsFloat()
		if e != nil {
			return e
		}
		if f != 0.1 {
			return fmt.Errorf("expected %f, got %f", 0.1, f)
		}

	case "int":
		i, e := childNode.GetDataAsInt()
		if e != nil {
			return e
		}
		if i != 1 {
			return fmt.Errorf("expected %d, got %d", 1, i)
		}

	case "string":
		if s, _ := childNode.GetData(); s != "s" {
			return fmt.Errorf("expected %s, got %s", "s", s)
		}

	default:
		return fmt.Errorf("invalid type specified: %s", arg1)

	}

	return nil

}

func iCantRetrieveTheDataFromTheNodeAsA(arg1 string) error {

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	if childNode == nil {
		return fmt.Errorf("no child node")
	}

	switch arg1 {

	case "bool":
		if _, e := childNode.GetDataAsBool(); e == nil {
			return fmt.Errorf("failed to throw error")
		}

	case "float":
		if _, e := childNode.GetDataAsFloat(); e == nil {
			return fmt.Errorf("failed to throw error")
		}

	case "int":
		if _, e := childNode.GetDataAsInt(); e == nil {
			return fmt.Errorf("failed to throw error")
		}

	default:
		return fmt.Errorf("invalid type specified: %s", arg1)

	}

	return nil

}

//requireParent

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
