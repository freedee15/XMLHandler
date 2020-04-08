package XMLHandler

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func convertToXML_handleChildren(n *Node, indent string) (string, error) {

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

func ConvertNodeToXML(n *Node) (string, error) {

	if n == nil {
		return "", fmt.Errorf("nil argument")
	}

	var i int
	var parent Parent = n
	for {

		breakLoop := false
		switch parent.(type) {

		case *Node:
			p, e := parent.(*Node).GetParent()
			if e != nil {
				return "", e
			}
			parent = p

		case *Tree:
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

func ConvertTreeToXML(t *Tree) (string, error) {

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

func ExportTree(t *Tree, path string) error {

	if t == nil {
		return fmt.Errorf("nil tree")
	}

	o, e := ConvertTreeToXML(t)
	if e != nil {
		return e
	}
	b := []byte(fmt.Sprintf("<?xml version=\"1.0\" encoding=\"UTF-8\" ?>\n%s", o))
	if e = ioutil.WriteFile(path, b, 0644); e != nil {
		return e
	}
	return nil

}

func ImportFile(path string) (*Tree, error) {

	b, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, e
	}

	s := SortXML(string(b))

	var current Parent
	for i, c := range s {

		var bracket bool
		var end bool
		if c[0:1] == "<" && c[len(c)-1:] == ">" {
			bracket = true
			c = c[1 : len(c)-1]
			if c[0:1] == "/" {
				end = true
				c = c[1:]
			} else if c[0:1] == "?" {
				if i != 0 {
					return nil, fmt.Errorf("invalid prolog")
				}
				if strings.Contains(c, "version=") {
					if !strings.Contains(c, "version=\"1.0\"") {
						return nil, fmt.Errorf("invalid xml version")
					}
				}
				continue
			}
		}

		var breakout bool
		if bracket == true {
			if end == true {
				if current == nil {
					return nil, fmt.Errorf("invalid xml")
				}
				switch current.(type) {

				case *Node:
					l, e := current.(*Node).GetLabel()
					if e != nil {
						return nil, e
					}
					if c != l {
						return nil, fmt.Errorf("node closed without open")
					}
					a, e := current.(*Node).GetParent()
					if e != nil {
						return nil, e
					}
					current = a

				case *Tree:
					if c != current.(*Tree).Label {
						return nil, fmt.Errorf("node closed without open")
					}
					breakout = true

				}

			} else {
				if current == nil {
					current = &Tree{Label: c}
				} else {
					n := &Node{}
					if e := current.AddChild(n); e != nil {
						return nil, e
					}
					if e := n.SetLabel(c); e != nil {
						return nil, e
					}
					current = n
				}
			}
		} else {
			if current == nil {
				return nil, fmt.Errorf("invalid xml")
			}
			switch current.(type) {

			case *Tree:
				return nil, fmt.Errorf("data in tree")

			case *Node:
				if e := current.(*Node).SetData(c); e != nil {
					return nil, e
				}

			}
		}

		if breakout == true {
			break
		}

	}

	switch current.(type) {

	case *Node:
		return nil, fmt.Errorf("invalid xml")

	case *Tree:
		return current.(*Tree), nil

	default:
		return nil, fmt.Errorf("failed to return")

	}

}

func SortXML(s string) []string {

	var b = []byte(s)
	b = []byte(strings.Replace(string(b), "\n", "", -1))
	var bracket bool
	var current string
	var toReturn []string
	for _, i := range b {

		if bracket == false {

			if i != '<' {
				current += string(i)
			} else {
				if strings.Replace(current, " ", "", -1) != "" {
					for current[0:1] == " " {
						current = current[1:]
					}
					for current[len(current)-1:] == " " {
						current = current[:len(current)-1]
					}
					toReturn = append(toReturn, current)
				}
				bracket = true
				current = ""
			}

		} else {

			if i != '>' {
				current += string(i)
			} else {
				if strings.Replace(current, " ", "", -1) != "" {
					for current[0:1] == " " {
						current = current[1:]
					}
					for current[len(current)-1:] == " " {
						current = current[:len(current)-1]
					}
					toReturn = append(toReturn, fmt.Sprintf("<%s>", current))
				}
				bracket = false
				current = ""
			}

		}

	}

	return toReturn

}
