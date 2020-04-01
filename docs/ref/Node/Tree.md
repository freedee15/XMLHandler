## Reference: XMLHandler/Node.**Tree** <span style="color:#445588">struct</span>

#### Fields
- [**children** <span style="color:#445588">[]*Node</span>](#field_children)
- [**Label** <span style="color:#445588">string</span>](#field_Label)

#### Methods
- [**AddChild**(child <span style="color:#445588">*Node</span>) <span style="color:#445588">error</span>](#method_AddChild)
- [**GetChildFromLabel**(label <span style="color:#445588">string</span>) (<span style="color:#445588">*Node</span>, <span style="color:#445588">error</span>)](#method_GetChildFromLabel)
- [**GetChildren**() (<span style="color:#445588">[]*Node</span>, <span style="color:#445588">error</span>)](#method_GetChildren)

```go
type Tree struct {
	children []*Node
	Label    string
}
```
A structure that acts as a base for all other nodes. This object needs no parent.

### Fields

<a name="field_children"></a><span style="font-size:1.25em">**children** <span style="color:#445588">[]*[<span style='font-weight:400'>Node</span>](/XMLHandler/docs/ref/Node/Node)</span>
><span style="font-style:normal">Stores all child nodes.</span>

---

<a name="field_Label"></a><span style="font-size:1.25em">**label** <span style="color:#445588">string</span>
><span style="font-style:normal">Holds the label of the tree object it is held in. For example, an XML node formatted as `<tree>` would have the label 'tree'.</span>

---
### Methods

<a name="method_AddChild"></a><span style="font-size:1.25em">**AddChild**(child <span style="color:#445588">*[<span style="font-weight:400">Node</span>](/XMLHandler/docs/ref/Node/Node)</span>) <span style="color:#445588">error</span>

```go
func (t *Tree) AddChild(child *Node) error {

	if t != child.parent {
		child.parent = t
		t.children = append(t.children, child)
	} else {
		return fmt.Errorf("\"%s\" is already parent of node \"%s\"", t.Label, child.label)
	}

	return nil

}
```

><span style="font-style:normal">Appends parameter `child` to the tree's `children` array.</span>

---

<a name="method_GetChildFromLabel"></a><span style="font-size:1.25em">**GetChildFromLabel**(label <span style="color:#445588">string</span>) (<span style="color:#445588">*[<span style="font-weight:400">Node</span>](/XMLHandler/docs/ref/Node/Node)</span>, <span style="color:#445588">error</span>)

```go
func (t *Tree) GetChildFromLabel(label string) (*Node, error) {

	if strings.Replace(label, " ", "", -1) == "" {
		return nil, fmt.Errorf("empty label")
	}

	for _, c := range t.children {

		if c.label == label {
			return c, nil
		}

	}

	return nil, fmt.Errorf("no child with label \"%s\"", label)

}
```

><span style="font-style:normal">Returns the first child in the tree's `children` array with a label equal to `label`. Returns an error if there are no matches.</span>

---

<a name="method_GetChildren"></a><span style="font-size:1.25em">**GetChildren**() (<span style="color:#445588">[]*[<span style="font-weight:400">Node</span>](/XMLHandler/docs/ref/Node/Node)</span>, <span style="color:#445588">error</span>)</span>

```go
func (t *Tree) GetChildren() []*Node {

	return t.children

}
```

><span style="font-style:normal">Returns the tree's `children` array.</span>
