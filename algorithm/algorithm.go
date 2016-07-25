package algorithm

// Algorithm that'll make it possible to choose the data
type Algorithm interface {
	Choose(key string, nodes int) []int
}
