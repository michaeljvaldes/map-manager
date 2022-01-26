package mapgen

type MapState struct {
	Name      string
	Dimension Dimension
	Night     bool
}

func getAllMapStates() []MapState {
	mapStates := make([]MapState, 4)
	mapStates[0] = MapState{"day", Overworld, false}
	mapStates[1] = MapState{"night", Overworld, true}
	mapStates[2] = MapState{"nether", Nether, false}
	mapStates[3] = MapState{"end", End, false}
	return mapStates
}
