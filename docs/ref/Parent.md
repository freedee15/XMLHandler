## Reference: XMLHandler.**Parent** <span style="color:#445588">interface</span>

#### Supported Types:
- [**XMLHandler.Node** <span style="color:#445588">struct</span>](/XMlHandler/docs/ref/Node)
- [**XMLHandler.Tree** <span style="color:#445588">struct</span>](/XMLHandler/docs/ref/Tree)

```go
type Parent interface {
	AddChild(child *XMLHandler.Node) error
}
```

Interface to allow both [<span style="font-weight:400">Node</span>](/XMLHandler/docs/ref/Node)s and [<span style="font-weight:400">Tree</span>](/XMLHandler/docs/ref/Tree)s to act as parents to a [<span style="font-weight:400">Node</span>](/XMLHandler/docs/ref/Node).