package Converter

import (
	"flag"
	"fmt"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
	"github.com/freedee15/XMLHandler/Node"
	"os"
	"testing"
)

var opt = godog.Options{Strict: true, Output: colors.Colored(os.Stdout)}

func init() {

	godog.BindFlags("godog.", flag.CommandLine, &opt)

}

func TestConverter(t *testing.T) {

	flag.Parse()
	opt.Paths = flag.Args()
	result := godog.RunWithOptions("Converter", func(s *godog.Suite) { FeatureContext(s) }, opt)
	if result != 0 {
		t.Errorf("godog exited with code %d", result)
	}

}

//////////////////////////////////////////////////

var nodeTree *Node.Tree = nil
var selected *Node.Node = nil
var output string = ""

func iCreateANodeTreeLabelled(arg1 string) error {

	nodeTree = &Node.Tree{Label: arg1}
	return nil

}

func iCreateAChildNodeOfNodeTreeLabelled(arg1 string) error {

	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}

	n := &Node.Node{}

	if e := nodeTree.AddChild(n); e != nil {
		return e
	}

	if e := n.SetLabel(arg1); e != nil {
		return e
	}

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

func FeatureContext(s *godog.Suite) {

	s.BeforeScenario(func(i interface{}) {

		nodeTree = nil
		selected = nil
		output = ""

	})

	s.Step(`^I create a node tree labelled "([^"]*)"$`, iCreateANodeTreeLabelled)
	s.Step(`^I create a child node of node tree labelled "([^"]*)"$`, iCreateAChildNodeOfNodeTreeLabelled)
	s.Step(`^I select child node "([^"]*)" of node tree$`, iSelectChildNodeOfNodeTree)
	s.Step(`^I select child node "([^"]*)" of selected node$`, iSelectChildNodeOfSelectedNode)
	s.Step(`^I set the selected node\'s data to "([^"]*)"$`, iSetTheSelectedNodesDataTo)

	//convertToXML
	s.Step(`^I create a node tree$`, iCreateANodeTree)
	s.Step(`^I create a child node of node "([^"]*)" labelled "([^"]*)"$`, iCreateAChildNodeOfNodeLabelled)
	s.Step(`^I convert the selected node to XML$`, iConvertTheSelectedNodeToXML)
	s.Step(`^I convert the node tree to XML$`, iConvertTheNodeTreeToXML)
	s.Step(`^the output should be:$`, theOutputShouldBe)

	//export
	s.Step(`^I create a child node of selected node labelled "([^"]*)"$`, iCreateAChildNodeOfSelectedNodeLabelled)
	s.Step(`^I export the node tree to "([^"]*)"$`, iExportTheNodeTreeTo)
	s.Step(`^the file "([^"]*)" should read:$`, theFileShouldRead)

}
