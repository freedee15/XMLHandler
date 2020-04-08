package XMLHandler

type Parent interface {
	AddChild(child *Node) error
}
