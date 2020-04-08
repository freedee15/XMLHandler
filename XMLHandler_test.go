package XMLHandler

import (
	"fmt"
	"github.com/cucumber/godog/gherkin"
	"io/ioutil"
	"os"
)

//convert

func iCreateAChildNodeOfNodeLabelled(arg1, arg2 string) error {

	if nodeTree == nil {
		return fmt.Errorf("no node tree")
	}

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	c := &Node{}

	n, e := nodeTree.GetChildFromLabel(arg1)
	if e != nil {
		return e
	}
	if e = n.AddChild(c); e != nil {
		return e
	}
	if e = c.SetLabel(arg2); e != nil {
		return e
	}

	return nil

}

func iConvertTheSelectedNodeToXML() error {

	if selected == nil {
		return fmt.Errorf("no selected node")
	}

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	o, e := ConvertNodeToXML(selected)
	if e != nil {
		return e
	}
	output = o

	return nil

}

func theOutputShouldBe(arg1 *gherkin.DocString) error {

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	if output != arg1.Content {
		return fmt.Errorf("expected:\n%s\ngot:\n%s", arg1.Content, output)
	}

	return nil

}

//io

func iCreateAChildNodeOfSelectedNodeLabelled(arg1 string) error {

	if selected == nil {
		return fmt.Errorf("no selected node")
	}

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	c := &Node{}
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

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	if e := ExportTree(nodeTree, arg1); e != nil {
		return e
	}

	return nil

}

func theFileShouldRead(arg1 string, arg2 *gherkin.DocString) error {

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

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

func iImportTheFile(arg1 string) error {

	if shouldFail == true {

		shouldFail = false

		_, e := ImportFile(arg1)
		if e == nil {
			return fmt.Errorf("failed to throw error")
		} else {
			getError = e
		}

		return nil

	}

	t, e := ImportFile(arg1)
	if e != nil {
		return nil
	}
	if t == nil {
		return fmt.Errorf("nil tree")
	}
	nodeTree = t

	return nil

}

func iKeepTheSelectedNodesData() error {

	if selected == nil {
		return fmt.Errorf("no node selected")
	}

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	s, e := selected.GetData()
	if e != nil {
		return e
	}
	keep += s

	return nil

}

func theKeptDataShouldBe(arg1 string) error {

	if shouldFail == true {
		shouldFail = false
		return fmt.Errorf("this step does not know how to fail")
	}

	if keep != arg1 {
		return fmt.Errorf("expected %s, got %s", arg1, keep)
	}

	return nil

}
