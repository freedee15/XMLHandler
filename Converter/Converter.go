package Converter

import (
	"XMLHandler/Node"
	"fmt"
)

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

	var returnString string

	for ; i > 0; i-- {

		returnString += "  "

	}

	nLabel, e := n.GetLabel()
	if e != nil {
		return "", e
	}
	nData, e := n.GetData()
	if e != nil {
		return "", e
	}
	returnString += "<" + nLabel + ">" + nData + "</" + nLabel + ">"
	return returnString, nil

}
