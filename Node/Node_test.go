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

	flag.Parse()
	opt.Paths = flag.Args()
	result := godog.RunWithOptions("Node", func(s *godog.Suite) { FeatureContext(s) }, opt)
	if result != 0 {
		t.Errorf("godog exited with code %d", result)
	}

}

///////////////////////////////////////////////////////////////////

var node *Node = nil
var shouldFail = false

func iCreateANode() error {

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	node = &Node{}
	return nil

}

func iShouldFailTheFollowingStep() error {

	shouldFail = true
	return nil

}

func FeatureContext(s *godog.Suite) {

	s.BeforeScenario(func(i interface{}) {

		node = nil
		retrieved = ""

	})

	s.Step(`^I create a node$`, iCreateANode)
	s.Step(`^I should fail the following step$`, iShouldFailTheFollowingStep)

	//dataType
	s.Step(`^I give the node a "([^"]*)" to store$`, iGiveTheNodeAToStore)
	s.Step(`^I can retrieve the data from the node as a "([^"]*)"$`, iCanRetrieveTheDataFromTheNodeAsA)
	s.Step(`^I can\'t retrieve the data from the node as a "([^"]*)"$`, iCantRetrieveTheDataFromTheNodeAsA)

	//child
	s.Step(`^I create a child node with the label "([^"]*)"$`, iCreateAChildNodeWithTheLabel)
	s.Step(`^I give child node "([^"]*)" a value of "([^"]*)"$`, iGiveChildNodeAValueOf)
	s.Step(`^I retrieve data from child node "([^"]*)"$`, iRetrieveDataFromChildNode)
	s.Step(`^the retrieved node data should be "([^"]*)"$`, theRetrievedNodeDataShouldBe)

}
