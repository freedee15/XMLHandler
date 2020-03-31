package Node

import (
	"fmt"
)

func iLabelTheNode(arg1 string) error {

	if shouldFail == true {

		shouldFail = false

		if node.SetLabel("s") == nil {
			return fmt.Errorf("failed to throw error")
		}

		return nil

	}

	if e := node.SetLabel("s"); e != nil {
		return e
	}
	return nil

}
