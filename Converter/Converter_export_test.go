package Converter

import (
	"fmt"
	"github.com/cucumber/godog/gherkin"
	"github.com/freedee15/XMLHandler/Node"
	"io/ioutil"
	"os"
)

func iCreateAChildNodeOfSelectedNodeLabelled(arg1 string) error {

	if selected == nil {
		return fmt.Errorf("no selected node")
	}

	c := &Node.Node{}
	if e := selected.AddChild(c); e != nil {
		return e
	}
	if e := c.SetLabel(arg1); e != nil {
		return e
	}

	return nil

}

func iExportTheNodeTreeTo(arg1 string) error {

	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}

	if e := ExportTree(nodeTree, arg1); e != nil {
		return e
	}

	return nil

}

func theFileShouldRead(arg1 string, arg2 *gherkin.DocString) error {

	f, e := ioutil.ReadFile(arg1)
	if e != nil {
		return e
	}
	if string(f) != arg2.Content {
		return fmt.Errorf("expected:\n %s\n got:\n %s", arg2.Content, string(f))
	}
	if e = os.Remove(arg1); e != nil {
		return e
	}

	return nil

}
