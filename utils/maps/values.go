package maps

import "github.com/platx/go-nova-poshta/utils/slices"

func Values[M ~map[K]V, K comparable, V any](ms ...M) []V {
	total := slices.Reduce(ms, func(acc int, m M) int { return acc + len(m) }, 0)
	r := make([]V, 0, total)
	for _, m := range ms {
		for _, v := range m {
			r = append(r, v)
		}
	}
	return r
}
