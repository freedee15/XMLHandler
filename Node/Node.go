package Node

import "strconv"

type Node struct {
	Label string
	data  string
}

func (n *Node) SetDataFromBool(b bool) {

	n.data = strconv.FormatBool(b)

}

func (n *Node) SetDataFromFloat(f float64) {

	n.data = strconv.FormatFloat(f, 'f', -1, 64)

}

func (n *Node) SetDataFromInt(i int) {

	n.data = strconv.Itoa(i)

}

func (n *Node) SetData(s string) {

	n.data = s

}

func (n *Node) GetDataAsBool() (bool, error) {

	if b, e := strconv.ParseBool(n.data); e != nil {
		return false, e
	} else {
		return b, nil
	}

}

func (n *Node) GetDataAsFloat() (float64, error) {

	if f, e := strconv.ParseFloat(n.data, 64); e != nil {
		return 0, e
	} else {
		return f, nil
	}

}

func (n *Node) GetDataAsInt() (int, error) {

	if i, e := strconv.Atoi(n.data); e != nil {
		return 0, e
	} else {
		return i, nil
	}

}

func (n *Node) GetData() string {

	return n.data

}
