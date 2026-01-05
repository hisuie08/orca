package compose

import "maps"

func (n NetworkSpec) Equal(u NetworkSpec) bool {
	l := n.Name == u.Name && n.Driver == u.Driver && n.External == u.External
	m := maps.Equal(n.Labels, u.Labels)
	return l && m
}
