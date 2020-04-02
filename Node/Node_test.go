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

//////////////////////////////////////////////////

var nodeTree *Tree = nil
var childNode *Node = nil
var getError error = nil
var node *Node = nil
var shouldFail = false
var retrieved string

func iCreateANodeTree() error {

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	nodeTree = &Tree{Label: "root"}
	return nil

}

func iShouldGetNoError() error {

	if getError != nil {
		return getError
	}

	return nil

}

func iShouldGetTheError(arg1 string) error {

	if getError == nil {
		return fmt.Errorf("got no error")
	} else if getError.Error() != arg1 {
		return fmt.Errorf("expected %s, got %s", arg1, getError.Error())
	}

	return nil

}

func theNextStepShouldFail() error {

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}
	shouldFail = true
	return nil

}

func FeatureContext(s *godog.Suite) {

	s.BeforeScenario(func(i interface{}) {

		nodeTree = nil
		childNode = nil
		getError = nil
		node = nil
		shouldFail = false
		retrieved = ""

	})

	s.Step(`^I create a node tree$`, iCreateANodeTree)
	s.Step(`^the next step should fail$`, theNextStepShouldFail)
	s.Step(`^I should get no error$`, iShouldGetNoError)
	s.Step(`^I should get the error '([^']*)'$`, iShouldGetTheError)
	s.Step(`^I should get the error "([^"]*)"$`, iShouldGetTheError)

	// Node/dataType
	s.Step(`^I create a child node$`, iCreateAChildNode)
	s.Step(`^I retrieve the child node data as a "([^"]*)"$`, iRetrieveTheChildNodeDataAsA)
	s.Step(`^I give the node a "([^"]*)" to store$`, iGiveTheNodeAToStore)
	s.Step(`^I can retrieve the data from the node as a "([^"]*)"$`, iCanRetrieveTheDataFromTheNodeAsA)
	s.Step(`^I can\'t retrieve the data from the node as a "([^"]*)"$`, iCantRetrieveTheDataFromTheNodeAsA)

	// Node/child
	s.Step(`^I create a child node with the label "([^"]*)"$`, iCreateAChildNodeWithTheLabel)
	s.Step(`^I give child node "([^"]*)" a value of "([^"]*)"$`, iGiveChildNodeAValueOf)
	s.Step(`^I retrieve data from child node "([^"]*)"$`, iRetrieveDataFromChildNode)
	s.Step(`^the retrieved node data should be "([^"]*)"$`, theRetrievedNodeDataShouldBe)
	s.Step(`^I create a child node in "([^"]*)" with the label "([^"]*)"$`, iCreateAChildNodeInWithTheLabel)
	s.Step(`^the children array of the node tree should have (\d+) members$`, theChildrenArrayOfTheNodeTreeShouldHaveMembers)
	s.Step(`^the children array of node "([^"]*)" should have (\d+) members$`, theChildrenArrayOfNodeShouldHaveMembers)

	// Node/requireParent
	s.Step(`^I create a node$`, iCreateANode)
	s.Step(`^I label the node "([^"]*)"$`, iLabelTheNode)
	s.Step(`^I add the node to the tree$`, iAddTheNodeToTheTree)
	s.Step(`^I retrieve the node data as a "([^"]*)"$`, iRetrieveTheNodeDataAsA)
	s.Step(`^I give the node a value of "([^"]*)"$`, iGiveTheNodeAValueOf)
	s.Step(`^data from node "([^"]*)" should be "([^"]*)"$`, dataFromNodeShouldBe)
	s.Step(`^I get the node\'s label$`, iGetTheNodesLabel)
	s.Step(`^I get the children array of the node$`, iGetTheChildrenArrayOfTheNode)

	// Tree/noParent
	s.Step(`^I add a child node labelled "([^"]*)" to the node tree$`, iAddAChildNodeLabelledToTheNodeTree)
	s.Step(`^node "([^"]*)" should have the parent "([^"]*)"$`, nodeShouldHaveTheParent)

}
