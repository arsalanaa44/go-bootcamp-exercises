package intj

type Inta []int

func (in Inta) Len() int {
	return len(in)
}
func (in Inta) Less(i, j int) bool {
	return in[i] < in[j]
}
func (in Inta) Swap(i, j int) {
	in[i], in[j] = in[j], in[i]
}
