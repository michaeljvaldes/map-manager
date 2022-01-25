package mapgen

type Dimension int

const (
	Overworld Dimension = iota
	Nether
	End
)

func (d Dimension) toString() string {
	return []string{"overworld", "nether", "end"}[d]
}
