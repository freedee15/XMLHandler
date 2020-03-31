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

func iCreateAChildNode() error {

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}
	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}

	childNode = &Node{}
	nodeTree.AddChild(childNode)

	return nil

}

func iCreateANode() error {

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	node = &Node{}
	return nil

}

func iGiveTheNodeAValueOf(arg1 string) error {

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}
	if node == nil {
		return fmt.Errorf("no node to modify")
	}

	node.SetData(arg1)
	return nil

}

func iShouldFailTheFollowingStep() error {

	shouldFail = true
	return nil

}

func FeatureContext(s *godog.Suite) {

	s.BeforeScenario(func(i interface{}) {

		nodeTree = nil
		childNode = nil
		node = nil
		retrieved = ""

	})

	s.Step(`^I create a node tree$`, iCreateANodeTree)
	s.Step(`^I create a child node$`, iCreateAChildNode)
	s.Step(`^I create a node$`, iCreateANode)
	s.Step(`^I give the node a value of "([^"]*)"$`, iGiveTheNodeAValueOf)
	s.Step(`^I should fail the following step$`, iShouldFailTheFollowingStep)

	// Node/dataType
	s.Step(`^I give the node a "([^"]*)" to store$`, iGiveTheNodeAToStore)
	s.Step(`^I can retrieve the data from the node as a "([^"]*)"$`, iCanRetrieveTheDataFromTheNodeAsA)
	s.Step(`^I can\'t retrieve the data from the node as a "([^"]*)"$`, iCantRetrieveTheDataFromTheNodeAsA)

	// Node/child
	s.Step(`^I create a child node with the label "([^"]*)"$`, iCreateAChildNodeWithTheLabel)
	s.Step(`^I give child node "([^"]*)" a value of "([^"]*)"$`, iGiveChildNodeAValueOf)
	s.Step(`^I retrieve data from child node "([^"]*)"$`, iRetrieveDataFromChildNode)
	s.Step(`^the retrieved node data should be "([^"]*)"$`, theRetrievedNodeDataShouldBe)

	// Node/requireParent
	s.Step(`^I label the node "([^"]*)"$`, iLabelTheNode)
	s.Step(`^I add the node to the tree$`, iAddTheNodeToTheTree)
	s.Step(`^data from node "([^"]*)" should be "([^"]*)"$`, dataFromNodeShouldBe)

	// Tree/noParent
	s.Step(`^I add a child node labelled "([^"]*)" to the node tree$`, iAddAChildNodeLabelledToTheNodeTree)
	s.Step(`^node "([^"]*)" should have the parent "([^"]*)"$`, nodeShouldHaveTheParent)

}
