## Reference: <span style="color:#445588">package</span> **XMLHandler**

#### Functions:
- [**ConvertNodeToXML**(n <span style="color:#445588">*XMLHandler.Node</span>) (<span style="color:#445588">string</span>, <span style="color:#445588">error</span>)](#func_ConvertNodeToXML)
- [**ConvertTreeToXML**(t <span style="color:#445588">*XMLHandler.Tree</span>) (<span style="color:#445588">string</span>, <span style="color:#445588">error</span>)](#func_ConvertTreeToXML)
- [**ExportTree**(t <span style="color:#445588">*XMLHandler.Tree</span>, path <span style="color:#445588">string</span>) <span style="color:#445588">error</span>](#func_ExportTree)
- [**ImportFile**(path <span style="color:#445588">string</span>) (<span style="color:#445588">*XMLHandler.Tree</span>, <span style="color:#445588">error</span>)](#func_ImportFile)
- [**SortXML**(s <span style="color:#445588">string</span>) <span style="color:#445588">[]string</span>](#func_SortXML)

#### Objects:
- [**Node** <span style="color:#445588">struct</span>](Node)
- [**Parent** <span style="color:#445588">struct</span>](Parent)
- [**Tree** <span style="color:#445588">struct</span>](Tree)

### Functions:

<a name="func_ConvertNodeToXML"></a><span style="font-size:1.25em">**ConvertNodeToXML**(n <span style="color:#445588">*</span>[XMLHandler.<span style="font-weight:400">Node</span>](/XMLHandler/docs/ref/Node)) (<span style="color:#445588">string</span>, <span style="color:#445588">error</span>)</span>

```go
func ConvertNodeToXML(n *XMLHandler.Node) (string, error) {

	if n == nil {
		return "", fmt.Errorf("nil argument")
	}

	var i int
	var parent XMLHandler.Parent = n
	for {

		breakLoop := false
		switch parent.(type) {

		case *XMLHandler.Node:
			p, e := parent.(*Node).GetParent()
			if e != nil {
				return "", e
			}
			parent = p

		case *XMLHandler.Tree:
			breakLoop = true

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
```

><span style="font-style:normal">Converts a [<span style="font-weight:400">Node</span>](/XMLHandler/docs/ref/Node) and its children to XML-compatible output. This can then be manipulated by the user in any way a string can be manipulated.</span>

---

<a name="func_ConvertTreeToXML"></a><span style="font-size:1.25em">**ConvertTreeToXML**(t <span style="color:#445588">*</span>[XMLHandler.<span style="font-weight:400">Tree</span>](/XMLHandler/docs/ref/Tree)) (<span style="color:#445588">string</span>, <span style="color:#445588">error</span>)</span>

```go
func ConvertTreeToXML(t *XMLHandler.Tree) (string, error) {

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
```

><span style="font-style:normal">Converts a [<span style="font-weight:400">Tree</span>](/XMLHandler/docs/ref/Tree) and its children to XML-compatible output. This can then be manipulated by the user in any way a string can be manipulated.</span>

---

<a name="func_ExportTree"></a><span style="font-size:1.25em">**ExportTree**(t <span style="color:#445588">*</span>[XMLHandler.<span style="font-weight:400">Tree</span>](/XMLHandler/docs/ref/Tree), path <span style="color:#445588">string</span>) <span style="color:#445588">error</span></span>

```go
func ExportTree(t *XMLHandler.Tree, path string) error {

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
```

><span style="font-style:normal">Outputs a [<span style="font-weight:400">Tree</span>](/XMLHandler/docs/ref/Tree) to an XML file located at `path`.

---

<a name="func_ImportFile"></a><span style="font-size:1.25em">**ImportFile**(path <span style="color:#445588">string</span>) (<span style="color:#445588">*</span>[XMLHandler.<span style="font-weight:400">Tree</span>](/XMLHandler/docs/ref/Tree), <span style="color:#445588">error</span>)</span>

```go
func ImportFile(path string) (*XMLHandler.Tree, error) {

	b, e := ioutil.ReadFile(path)
	if e != nil {
		return nil, e
	}

	s := SortXML(string(b))

	var current XMLHandler.Parent
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

				case *XMLHandler.Node:
					l, e := current.(*XMLHandler.Node).GetLabel()
					if e != nil {
						return nil, e
					}
					if c != l {
						return nil, fmt.Errorf("node closed without open")
					}
					a, e := current.(*XMLHandler.Node).GetParent()
					if e != nil {
						return nil, e
					}
					current = a

				case *XMLHandler.Tree:
					if c != current.(*XMLHandler.Tree).Label {
						return nil, fmt.Errorf("node closed without open")
					}
					breakout = true

				}

			} else {
				if current == nil {
					current = &XMLHandler.Tree{Label: c}
				} else {
					n := &XMLHandler.Node{}
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

			case *XMLHandler.Tree:
				return nil, fmt.Errorf("data in tree")

			case *XMLHandler.Node:
				if e := current.(*XMLHandler.Node).SetData(c); e != nil {
					return nil, e
				}

			}
		}

		if breakout == true {
			break
		}

	}

	switch current.(type) {

	case *XMLHandler.Node:
		return nil, fmt.Errorf("invalid xml")

	case *XMLHandler.Tree:
		return current.(*Tree), nil

	default:
		return nil, fmt.Errorf("failed to return")

	}

}
```

><span style="font-style:normal">Imports an external XML file from `path`. Returns a [<span style="font-weight:400">Tree</span>](/XMLHandler/docs/ref/Tree) containing the data of the XML file if successful.</span>

---

<a name="func_SortXML"></a><span style="font-size:1.25em">**SortXML**(s <span style="color:#445588">string</span>) <span style="color:#445588">[]string</span></span>

```go
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
```

><span style="font-style:normal">Sorts an XML file (passed as a string) and separates each field. For example:</span>

```xml
<?xml version="1.0" encoding="UTF-8"?>
<XMLTree>
    <NewNode> foo </NewNode>
    <OtherNode></OtherNode>
</XMLTree>
```

><span style="font-style:normal">When this function was called on this XML tree, the output would contain:</span>

```go
["<?xml version= \"1.0\" encoding=\"UTF-8\"?>", "<XMLTree>", "<NewNode>", "foo", "</NewNode>", "<OtherNode>", "</OtherNode>", "</XMLTree>"]
```