package Node

import (
	"fmt"
)

func iGiveTheNodeAToStore(arg1 string) error {

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	if node == nil {
		return fmt.Errorf("no node to modify")
	}

	switch arg1 {

	case "bool":
		node.SetDataFromBool(true)

	case "float":
		node.SetDataFromFloat(0.1)

	case "int":
		node.SetDataFromInt(1)

	case "string":
		node.SetData("s")

	default:
		return fmt.Errorf("invalid type specified: %s", arg1)

	}

	return nil

}

func iCanRetrieveTheDataFromTheNodeAsA(arg1 string) error {

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	if node == nil {
		return fmt.Errorf("no node to modify")
	}

	switch arg1 {

	case "bool":
		b, e := node.GetDataAsBool()
		if e != nil {
			return e
		}
		if b != true {
			return fmt.Errorf("expected %t, got %t", true, b)
		}

	case "float":
		f, e := node.GetDataAsFloat()
		if e != nil {
			return e
		}
		if f != 0.1 {
			return fmt.Errorf("expected %f, got %f", 0.1, f)
		}

	case "int":
		i, e := node.GetDataAsInt()
		if e != nil {
			return e
		}
		if i != 1 {
			return fmt.Errorf("expected %d, got %d", 1, i)
		}

	case "string":
		if s := node.GetData(); s != "s" {
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

	if node == nil {
		return fmt.Errorf("no node to modify")
	}

	switch arg1 {

	case "bool":
		if _, e := node.GetDataAsBool(); e == nil {
			return fmt.Errorf("failed to throw error")
		}

	case "float":
		if _, e := node.GetDataAsFloat(); e == nil {
			return fmt.Errorf("failed to throw error")
		}

	case "int":
		if _, e := node.GetDataAsInt(); e == nil {
			return fmt.Errorf("failed to throw error")
		}

	default:
		return fmt.Errorf("invalid type specified: %s", arg1)

	}

	return nil

}
