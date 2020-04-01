## Reference: XMLHandler/Node.**Parent** <span style="color:#445588">interface</span>

#### Supported Types:
- [**Node** <span style="color:#445588">struct</span>](/XMlHandler/docs/ref/Node/Node)
- [**Tree** <span style="color:#445588">struct</span>](/XMLHandler/docs/ref/Node/Tree)

```go
type Parent interface {
	AddChild(child *Node) error
}
```

Interface to allow both [<span style="font-weight:400">Node</span>](/XMLHandler/docs/ref/Node/Node)s and [<span style="font-weight:400">Tree</span>](/XMLHandler/docs/ref/Node/Tree)s to act as parents to a [<span style="font-weight:400">Node</span>](/XMLHandler/docs/ref/Node/Node).