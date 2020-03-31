package Node

import (
	"flag"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"os"
	"testing"
)

var opt = godog.Options{Strict: true, Output: colors.Colored(os.Stdout)}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestNode(t *testing.T) {
	opt.Paths = []string{"features/dataType.feature"}
	result := godog.RunWithOptions("Node", func(suite *godog.Suite) { FeatureContext(suite) }, opt)
	if result != 0 {
		t.Errorf("godog exited with code %d", result)
	}
}

var node *Node = nil

func iCreateANode() error {

	node = &Node{}
	return nil

}

func iGiveTheNodeAToStore(arg1 string) error {

	if node == nil {
		return fmt.Errorf("no node to modify")
	}

	switch arg1 {

	case "int":
		node.SetDataFromInt(1)

	default:
		return fmt.Errorf("invalid type specified: %s", arg1)

	}

	return nil

}

func iCanRetrieveTheDataFromTheNodeAsA(arg1 string) error {

	if node == nil {
		return fmt.Errorf("no node to modify")
	}

	switch arg1 {

	case "int":
		i, e := node.GetDataAsInt()
		if e != nil {
			return e
		}
		if i != 1 {
			return fmt.Errorf("expected %d, got %d", 1, i)
		}

	}

	return nil

}

func FeatureContext(s *godog.Suite) {

	s.BeforeScenario(func(i interface{}) {
		node = nil
	})

	s.Step(`^I create a node$`, iCreateANode)
	s.Step(`^I give the node a "([^"]*)" to store$`, iGiveTheNodeAToStore)
	s.Step(`^I can retrieve the data from the node as a "([^"]*)"$`, iCanRetrieveTheDataFromTheNodeAsA)

}
