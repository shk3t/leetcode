package main

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}

	avgs := []float64{}
	layer := []*TreeNode{root}

	for len(layer) > 0 {
		nextLayer := []*TreeNode{}
		sum := 0
		for _, node := range layer {
			sum += node.Val
			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}
			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}
		avgs = append(avgs, float64(sum)/float64(len(layer)))
		layer = nextLayer
	}

	return avgs
}

