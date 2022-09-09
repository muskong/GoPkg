package plugin

type Node struct {
	Name    string
	Current int
	Weight  int
}

func SmoothWeight(nodes []*Node) (best *Node) {
	if len(nodes) == 0 {
		return
	}

	total := 0
	for _, node := range nodes {
		if node == nil {
			continue
		}

		total += node.Weight
		node.Current += node.Weight

		if best == nil || node.Current > best.Current {
			best = node
		}
	}

	if best == nil {
		return
	}

	best.Current -= total

	return
}
