package Node

import "strconv"

type Node struct {
	data string
}

func (n *Node) SetDataFromInt(i int) {

	n.data = strconv.Itoa(i)

}

func (n *Node) GetDataAsInt() (int, error) {

	if i, e := strconv.ParseInt(n.data, 0, 0); e != nil {
		return 0, e
	} else {
		return int(i), nil
	}

}
