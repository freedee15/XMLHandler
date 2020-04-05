package XMLHandler

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

func TestXMLHandler(t *testing.T) {

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
var selected *Node = nil
var output string = ""

func iConvertTheNodeTreeToXML() error {

	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}

	o, e := ConvertTreeToXML(nodeTree)
	if e != nil {
		return e
	}
	output = o

	return nil

}

func iCreateAChildNodeOfNodeTreeLabelled(arg1 string) error {

	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}

	n := &Node{}

	if e := nodeTree.AddChild(n); e != nil {
		return e
	}

	if e := n.SetLabel(arg1); e != nil {
		return e
	}

	return nil

}

func iCreateANodeTree() error {

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	nodeTree = &Tree{Label: "root"}
	return nil

}

func iCreateANodeTreeLabelled(arg1 string) error {

	nodeTree = &Tree{Label: arg1}
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

func iSelectChildNodeOfNodeTree(arg1 string) error {

	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}

	s, e := nodeTree.GetChildFromLabel(arg1)
	if e != nil {
		return e
	}
	selected = s

	return nil

}

func iSelectChildNodeOfSelectedNode(arg1 string) error {

	if selected == nil {
		return fmt.Errorf("no selected node")
	}

	s, e := selected.GetChildFromLabel(arg1)
	if e != nil {
		return e
	}
	selected = s

	return nil

}

func iSetTheSelectedNodesDataTo(arg1 string) error {

	if selected == nil {
		return fmt.Errorf("no selected node")
	}

	if e := selected.SetData(arg1); e != nil {
		return e
	}

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

func FeatureContext(s *godog.Suite) {

	s.BeforeScenario(func(i interface{}) {

		nodeTree = nil
		childNode = nil
		getError = nil
		node = nil
		shouldFail = false
		retrieved = ""
		selected = nil
		output = ""

	})

	s.Step(`^I convert the node tree to XML$`, iConvertTheNodeTreeToXML)
	s.Step(`^I create a child node of node tree labelled "([^"]*)"$`, iCreateAChildNodeOfNodeTreeLabelled)
	s.Step(`^I create a node tree$`, iCreateANodeTree)
	s.Step(`^I create a node tree labelled "([^"]*)"$`, iCreateANodeTreeLabelled)
	s.Step(`^the next step should fail$`, theNextStepShouldFail)
	s.Step(`^I select child node "([^"]*)" of node tree$`, iSelectChildNodeOfNodeTree)
	s.Step(`^I select child node "([^"]*)" of selected node$`, iSelectChildNodeOfSelectedNode)
	s.Step(`^I set the selected node\'s data to "([^"]*)"$`, iSetTheSelectedNodesDataTo)
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
	s.Step(`^I get the parent of the node$`, iGetTheParentOfTheNode)

	// Tree/noParent
	s.Step(`^node "([^"]*)" should have the parent "([^"]*)"$`, nodeShouldHaveTheParent)

	//convert
	s.Step(`^I create a child node of node "([^"]*)" labelled "([^"]*)"$`, iCreateAChildNodeOfNodeLabelled)
	s.Step(`^I convert the selected node to XML$`, iConvertTheSelectedNodeToXML)
	s.Step(`^the output should be:$`, theOutputShouldBe)

	//io
	s.Step(`^I create a child node of selected node labelled "([^"]*)"$`, iCreateAChildNodeOfSelectedNodeLabelled)
	s.Step(`^I export the node tree to "([^"]*)"$`, iExportTheNodeTreeTo)
	s.Step(`^the file "([^"]*)" should read:$`, theFileShouldRead)

}
