package Node

type Parent interface {
	AddChild(child *Node) error
}
