package Converter

import (
	"fmt"
	"github.com/freedee15/XMLHandler/Node"
	"strings"
)

func convertToXML_handleChildren(n *Node.Node, indent string) (string, error) {

	if n == nil {
		return "", fmt.Errorf("nil node")
	}

	a, e := n.GetChildren()
	if e != nil {
		return "", e
	}
	nData, e := n.GetData()
	if e != nil {
		return "", e
	}
	nLabel, e := n.GetLabel()
	if e != nil {
		return "", e
	}
	var returnString string
	if len(a) == 0 {

		if strings.Replace(nData, " ", "", -1) != "" {
			returnString = fmt.Sprintf(fmt.Sprintf("%s<%s> %s </%[2]s>\n", indent, nLabel, nData))
		} else {
			returnString = fmt.Sprintf(fmt.Sprintf("%s<%s></%[2]s>\n", indent, nLabel))
		}

	} else {

		returnString = fmt.Sprintf("%s<%s>\n", indent, nLabel)
		if strings.Replace(nData, " ", "", -1) != "" {
			returnString += fmt.Sprintf("%s  %s\n", indent, nData)
		}

		for _, c := range a {

			s, e := convertToXML_handleChildren(c, fmt.Sprintf("%s  ", indent))
			if e != nil {
				return "", e
			}

			returnString += s

		}

		returnString += fmt.Sprintf("%s</%s>\n", indent, nLabel)

	}

	return returnString, nil

}

func ConvertNodeToXML(n *Node.Node) (string, error) {

	if n == nil {
		return "", fmt.Errorf("nil argument")
	}

	var i int
	var parent Node.Parent = n
	for {

		breakLoop := false
		switch parent.(type) {

		case *Node.Node:
			p, e := parent.(*Node.Node).GetParent()
			if e != nil {
				return "", e
			}
			parent = p

		case *Node.Tree:
			breakLoop = true

		default:
			return "", fmt.Errorf("this should be impossible")

		}

		if breakLoop == true {
			break
		}

		i++

	}

	var indent string
	for ; i > 0; i-- {

		indent += "  "

	}

	var returnString string
	nLabel, e := n.GetLabel()
	if e != nil {
		return "", e
	}
	nData, e := n.GetData()
	if e != nil {
		return "", e
	}
	a, e := n.GetChildren()
	if e != nil {
		return "", e
	}
	if len(a) == 0 {

		if strings.Replace(nData, " ", "", -1) != "" {
			returnString = fmt.Sprintf(fmt.Sprintf("%s<%s> %s </%[2]s>", indent, nLabel, nData))
		} else {
			returnString = fmt.Sprintf(fmt.Sprintf("%s<%s></%[2]s>", indent, nLabel))
		}

	} else {

		returnString = fmt.Sprintf("%s<%s>\n", indent, nLabel)
		if strings.Replace(nData, " ", "", -1) != "" {
			returnString += fmt.Sprintf("%s  %s\n", indent, nData)
		}
		for _, c := range a {

			s, e := convertToXML_handleChildren(c, fmt.Sprintf("%s  ", indent))
			if e != nil {
				return "", e
			}
			returnString += s

		}

		returnString += fmt.Sprintf("%s</%s>", indent, nLabel)

	}

	return returnString, nil

}

func ConvertTreeToXML(t *Node.Tree) (string, error) {

	if t == nil {
		return "", fmt.Errorf("nil tree")
	}

	returnString := fmt.Sprintf("<%s>\n", t.Label)
	for _, n := range t.GetChildren() {

		s, e := convertToXML_handleChildren(n, "  ")
		if e != nil {
			return "", e
		}
		returnString += s

	}

	returnString += fmt.Sprintf("</%s>", t.Label)

	return returnString, nil

}
