package Converter

import (
	"XMLHandler/Node"
	"flag"
	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
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

func FeatureContext(s *godog.Suite) {

	s.BeforeScenario(func(i interface{}) {

		nodeTree = nil
		selected = nil
		output = ""

	})

	// individual
	s.Step(`^I create a node tree$`, iCreateANodeTree)
	s.Step(`^I create a child node of node tree labelled "([^"]*)"$`, iCreateAChildNodeOfNodeTreeLabelled)
	s.Step(`^I create a child node of node "([^"]*)" labelled "([^"]*)"$`, iCreateAChildNodeOfNodeLabelled)
	s.Step(`^I select child node "([^"]*)" of node tree$`, iSelectChildNodeOfNodeTree)
	s.Step(`^I select child node "([^"]*)" of selected node$`, iSelectChildNodeOfSelectedNode)
	s.Step(`^I convert the selected node to XML$`, iConvertTheSelectedNodeToXML)
	s.Step(`^The output should be "([^"]*)"$`, theOutputShouldBe)
	s.Step(`^I set the selected node\'s data to "([^"]*)"$`, iSetTheSelectedNodesDataTo)

}
