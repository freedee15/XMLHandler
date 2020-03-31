package Node

import (
	"flag"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"os"
	"testing"
)

var node *Node = nil

var opt = godog.Options{Strict: true, Output: colors.Colored(os.Stdout)}

func init() {

	godog.BindFlags("godog.", flag.CommandLine, &opt)

}

func TestNode(t *testing.T) {

	flag.Parse()
	opt.Paths = flag.Args()
	result := godog.RunWithOptions("Node", func(s *godog.Suite) { FeatureContext(s) }, opt)
	if result != 0 {
		t.Errorf("godog exited with code %d", result)
	}

}

func FeatureContext(s *godog.Suite) {

	s.Step(`^I create a node$`, iCreateANode)
	s.Step(`^I give the node a "([^"]*)" to store$`, iGiveTheNodeAToStore)
	s.Step(`^I can retrieve the data from the node as a "([^"]*)"$`, iCanRetrieveTheDataFromTheNodeAsA)
	s.Step(`^I can\'t retrieve the data from the node as a "([^"]*)"$`, iCantRetrieveTheDataFromTheNodeAsA)

}
