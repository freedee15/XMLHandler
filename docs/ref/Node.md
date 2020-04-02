## Reference: XMLHandler.**Node** <span style = "color:#445588">struct</span>

#### Fields:
  - [**children** <span style="color:#445588">[]*XMLHandler.Node</span>](#field_children)
  - [**data** <span style="color:#445588">string</span>](#field_data)
  - [**label** <span style="color:#445588">string</span>](#field_label)
  - [**parent** <span style="color:#445588">XMLHandler.Parent</span>](#field_parent)

#### Methods:
- [**AddChild**(child <span style="color:#445588">*XMLHandler.Node</span>) <span style="color:#445588">error</span>](#method_AddChild)
- [**GetChildFromLabel**(label <span style="color:#445588">string</span>) (<span style="color:#445588">*XMLHandler.Node</span>, <span style="color:#445588">error</span>)](#method_GetChildFromLabel)
- [**GetChildren**() (<span style="color:#445588">[]*XMLHandler.Node</span>, <span style="color:#445588">error</span>)](#method_GetChildren)
- [**GetData**() (<span style="color:#445588">string</span>, <span style="color:#445588">error</span>)](#method_GetData)
- [**GetDataAsBool**() (<span style="color:#445588">bool</span>, <span style="color:#445588">error</span>)](#method_GetDataAsBool)
- [**GetDataAsFloat**() (<span style="color:#445588">float64</span>, <span style="color:#445588">error</span>)](#method_GetDataAsFloat)
- [**GetDataAsInt**() (<span style="color:#445588">int</span>, <span style="color:#445588">error</span>)](#method_GetDataAsInt)
- [**GetLabel**() (<span style="color:#445588">string</span>, <span style="color:#445588">error</span>)](#method_GetLabel)
- [**GetParent**() (<span style="color:#445588">XMLHandler.Parent</span>, <span style="color:#445588">error</span>)](#method_GetParent)
- [**SetData**(s <span style="color:#445588">string</span>) <span style="color:#445588">error</span>](#method_SetData)
- [**SetDataFromBool**(b <span style="color:#445588">bool</span>) <span style="color:#445588">error</span>](#method_SetDataFromBool)
- [**SetDataFromFloat**(f <span style="color:#445588">float64</span>) <span style="color:#445588">error</span>](#method_SetDataFromFloat)
- [**SetDataFromInt**(i <span style="color:#445588">int</span>) <span style="color:#445588">error</span>](#method_SetDataFromInt)
- [**SetLabel**(s <span style="color:#445588">string</span>) <span style="color:#445588">error</span>](#method_SetLabel)
   
```go
type Node struct {
	children []*XMLHandler.Node
	data     string
	label    string
	parent   XMLHandler.Parent
}
```
     
A struct object to hold data like an XML node would. Access to the fields of this object will be limited until it is given a parent (and is subsequently put into a tree). To withhold access to these fields, many "getters" and "setters" are required.

### Fields

<a name="field_children"></a><span style="font-size:1.25em">**children** <span style="color:#445588">[]*[<span style='font-weight:400'>XMLHandler.Node</span>](/XMLHandler/docs/ref)</span>
><span style="font-style:normal">Stores all child nodes.</span>

---

<a name="field_data"></a><span style="font-size:1.25em">**data** <span style="color:#445588">string</span>
><span style="font-style:normal">Stores any data the node may hold. For example:</span>

```xml
<node>Some Data!</node>
```

><span style="font-style:normal">Here, `node`'s data would be `"Some Data!"`</span>

---

<a name="field_label"></a><span style="font-size:1.25em">**label** <span style="color:#445588">string</span>
><span style="font-style:normal">Holds the label of the node object it is held in. For example, an XML node formatted as `<node>` would have the label 'node'.</span>

---

<a name="field_parent"></a><span style="font-size:1.25em">**parent** [<span style='font-weight:400'>XMLHandler.Parent</span>](/XMLHandler/docs/ref/Parent)
><span style="font-style:normal">Holds the pointer to the objects parent in the hierarchy. The nodes other properties cannot be accessed until it has a parent. You can set a parent by passing the node object to a Parent object's `AddChild(child *XMLHandler.Node)` function. Both will ensure that the Node object has its `parent` set properly and that the node gets added to the Parent's `children` array.</span>

---
### Methods

<a name="method_AddChild"></a><span style="font-size:1.25em">**AddChild**(child <span style="color:#445588">*[<span style="font-weight:400">XMLHandler.Node</span>](/XMLHandler/docs/ref/Node)</span>) <span style="color:#445588">error</span>

```go
func (n *XMLHandler.Node) AddChild(child *XMLHandler.Node) error {

	if n.parent == nil {
		return fmt.Errorf("no parent")
	}

	found := false
	for _, n := range n.children {

		if n == child {
			found = true
			break
		}

	}

	if n != child.GetParent() || found == false {
		child.parent = n
		n.children = append(n.children, child)
	}
	return nil

}
```

><span style="font-style:normal">Appends parameter `child` to the node's `children` array. Returns an error if the node has no parent.</span>

---

<a name="method_GetChildFromLabel"></a><span style="font-size:1.25em">**GetChildFromLabel**(label <span style="color:#445588">string</span>) (<span style="color:#445588">*[<span style="font-weight:400">XMLHandler.Node</span>](/XMLHandler/docs/ref/Node)</span>, <span style="color:#445588">error</span>)

```go
func (n *XMLHandler.Node) GetChildFromLabel(label string) (*XMLHandler.Node, error) {

	if strings.Replace(label, " ", "", -1) == "" {
		return nil, fmt.Errorf("empty label")
	}

	for _, c := range n.children {

		if c.label == label {
			return c, nil
		}

	}

	return nil, fmt.Errorf("no child with label \"%s\"", label)

}
```

><span style="font-style:normal">Returns the first child in the node's `children` array with a label equal to `label`. Returns an error if there are no matches.</span>

---

<a name="method_GetChildren"></a><span style="font-size:1.25em">**GetChildren**() (<span style="color:#445588">[]*[<span style="font-weight:400">XMLHandler.Node</span>](/XMLHandler/docs/ref/Node)</span>, <span style="color:#445588">error</span>)</span>

```go
func (n *XMLHandler.Node) GetChildren() ([]*XMLHandler.Node, error) {

	if n.parent == nil {
		return nil, fmt.Errorf("no parent")
	}
	return n.children, nil

}
```

><span style="font-style:normal">Returns the node's `children` array. Returns an error if the node has no parent.</span>

---

<a name="method_GetData"></a><span style="font-size:1.25em">**GetData**() (<span style="color:#445588">string</span>, <span style="color:#445588">error</span>)</span>

```go
func (n *XMLHandler.Node) GetData() (string, error) {

	if n.parent == nil {
		return "", fmt.Errorf("no parent")
	}
	return n.data, nil

}
```

><span style="font-style:normal">Returns the `data` held by the node object. Returns an error if the node has no parent.</span>

---

<a name="method_GetDataAsBool"></a><span style="font-size:1.25em">**GetDataAsBool**() (<span style="color:#445588">bool</span>, <span style="color:#445588">error</span>)</span>

```go
func (n *XMLHandler.Node) GetDataAsBool() (bool, error) {

	if n.parent == nil {
		return false, fmt.Errorf("no parent")
	}
	if b, e := strconv.ParseBool(n.data); e != nil {
		return false, e
	} else {
		return b, nil
	}

}
```

><span style="font-style:normal">Returns the `data` held by the node object as a `bool` value. Return an error if the data held by the node is not `bool` compatible or if the node has no parent.</span>

---

<a name="method_GetDataAsFloat"></a><span style="font-size:1.25em">**GetDataAsFloat**() (<span style="color:#445588">float64</span>, <span style="color:#445588">error</span>)</span>

```go
func (n *Node) GetDataAsFloat() (float64, error) {

	if n.parent == nil {
		return 0, fmt.Errorf("no parent")
	}
	if f, e := strconv.ParseFloat(n.data, 64); e != nil {
		return 0, e
	} else {
		return f, nil
	}

}
```

><span style="font-style:normal">Returns the `data` held by the node object as a `float64` value. Return an error if the data held by the node is not `float64` compatible or if the node has no parent.</span>

---

<a name="method_GetDataAsInt"></a><span style="font-size:1.25em">**GetDataAsInt**() (<span style="color:#445588">int</span>, <span style="color:#445588">error</span>)</span>

```go
func (n *XMLHandler.Node) GetDataAsInt() (int, error) {

	if n.parent == nil {
		return 0, fmt.Errorf("no parent")
	}
	if i, e := strconv.Atoi(n.data); e != nil {
		return 0, e
	} else {
		return i, nil
	}

}
```

><span style="font-style:normal">Returns the `data` held by the node object as an `int` value. Return an error if the data held by the node is not `int` compatible or if the node has no parent.</span>

---

<a name="method_GetLabel"></a><span style="font-size:1.25em">**GetLabel**() (<span style="color:#445588">string</span>, <span style="color:#445588">error</span>)</span>

```go
func (n *XMLHandler.Node) GetLabel() (string, error) {

	if n.parent == nil {
		return "", fmt.Errorf("no parent")
	}
	return n.label, nil

}
```

><span style="font-style:normal">Returns the node's `label`. Returns an error if the node has no parent.</span>

---

<a name="method_GetParent"></a><span style="font-size:1.25em">**GetParent**() ([<span style="font-weight:400">XMLHandler.Parent</span>](/XMLHandler/docs/ref/Parent), <span style="color:#445588">error</span>)</span>

```go
func (n *XMLHandler.Node) GetParent() (XMLHandler.Parent, error) {

	if n.parent == nil {
		return nil, fmt.Errorf("no parent")
	}
	return n.parent, nil

}
```

><span style="font-style:normal">Returns the node's `parent`. Returns an error if the node has no parent.</span>

---

<a name="method_SetData"></a><span style="font-size:1.25em">**SetData**(s <span style="color:#445588">string</span>) <span style="color:#445588">error</span>
```go
func (n *XMLHandler.Node) SetData(s string) error {

	if n.parent == nil {
		return fmt.Errorf("no parent")
	}
	n.data = s
	return nil

}
```

><span style="font-style:normal">Sets the node's `data` to `s`. Returns an error if the node has no parent.</span>

---

<a name="method_SetDataFromBool"></a><span style="font-size:1.25em">**SetDataFromBool**(b <span style="color:#445588">bool</span>) <span style="color:#445588">error</span>
```go
func (n *XMLHandler.Node) SetDataFromBool(b bool) error {

	if n.parent == nil {
		return fmt.Errorf("no parent")
	}
	n.data = strconv.FormatBool(b)
	return nil

}
```

><span style="font-style:normal">Sets the node's `data` to a string holding the value of `b`. Returns an error if the node has no parent.</span>

---

<a name="method_SetDataFromFloat"></a><span style="font-size:1.25em">**SetDataFromFloat**(f <span style="color:#445588">float64</span>) <span style="color:#445588">error</span>
```go
func (n *XMLHandler.Node) SetDataFromFloat(f float64) error {

	if n.parent == nil {
		return fmt.Errorf("no parent")
	}
	n.data = strconv.FormatFloat(f, 'f', -1, 64)
	return nil

}
```

><span style="font-style:normal">Sets the node's `data` to a string holding the value of `f`. Returns an error if the node has no parent.</span>

---

<a name="method_SetDataFromInt"></a><span style="font-size:1.25em">**SetDataFromInt**(i <span style="color:#445588">int</span>) <span style="color:#445588">error</span>
```go
func (n *XMLHandler.Node) SetDataFromInt(i int) error {

	if n.parent == nil {
		return fmt.Errorf("no parent")
	}
	n.data = strconv.Itoa(i)
	return nil

}
```

><span style="font-style:normal">Sets the node's `data` to a string holding the value of `i`. Returns an error if the node has no parent.</span>

---

<a name="method_SetLabel"></a><span style="font-size:1.25em">**SetLabel**(s <span style="color:#445588">string</span>) <span style="color:#445588">error</span></span>
```go
func (n *XMLHandler.Node) SetLabel(s string) error {

	if n.parent == nil {
		return fmt.Errorf("no parent")
	}
	n.label = s
	return nil

}
```

><span style="font-style:normal">Sets the node's `data` to `s`. Returns an error if the node has no parent.</span>