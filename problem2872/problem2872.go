package problem2872

/* There is an undirected tree with n nodes labeled from 0 to n - 1. 
You are given the integer n and a 2D integer array edges of length n - 1, where edges[i] = [ai, bi] indicates 
that there is an edge between nodes ai and bi in the tree.
You are also given a 0-indexed integer array values of length n, where values[i] is the value associated with 
the ith node, and an integer k.
A valid split of the tree is obtained by removing any set of edges, possibly empty, from the tree such that
the resulting components all have values that are divisible by k, where the value of a connected component 
is the sum of the values of its nodes.
Return the maximum number of components in any valid split. */

func MaxKDivisibleComponents(n int, edges [][]int, values []int, k int) int {
	adj := make([][]int, n)
	for _, e := range edges {
		u, v := e[0], e[1]
		adj[u] = append(adj[u], v)
		adj[v] = append(adj[v], u)
	}

	components := 0

	var dfs func(u, parent int) int64
	dfs = func(u, parent int) int64 {
		subtreeRem := int64(values[u] % k)

		for _, v := range adj[u] {
			if v == parent {
				continue
			}
			childRem := dfs(v, u)
			subtreeRem = (subtreeRem + childRem) % int64(k)
		}

		if subtreeRem == 0 {
			components++
			return 0
		}
		return subtreeRem
	}

	dfs(0, -1)
	return components
}
