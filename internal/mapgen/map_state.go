package mapgen

type MapState struct {
	Name      string
	Dimension Dimension
	Night     bool
	GndXRay   bool
	topY      int
}

func getAllMapStates() []MapState {
	mapStates := make([]MapState, 4)
	mapStates[0] = MapState{"day", Overworld, false, false, 100}
	mapStates[1] = MapState{"night", Overworld, true, false, 100}
	mapStates[2] = MapState{"nether", Nether, false, true, 63}
	mapStates[3] = MapState{"end", End, false, false, 100}
	return mapStates
}
