package Node

import (
	"fmt"
)

func iGiveTheNodeAToStore(arg1 string) error {

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	if childNode == nil {
		return fmt.Errorf("no node to modify")
	}

	switch arg1 {

	case "bool":
		childNode.SetDataFromBool(true)

	case "float":
		childNode.SetDataFromFloat(0.1)

	case "int":
		childNode.SetDataFromInt(1)

	case "string":
		childNode.SetData("s")

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

	if childNode == nil {
		return fmt.Errorf("no node to modify")
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
		if s := childNode.GetData(); s != "s" {
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
		return fmt.Errorf("no node to modify")
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
