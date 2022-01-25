package main

type MapConfig struct {
	Name      string
	Dimension Dimension
	Night     bool
}

func getAllMapConfigs() []MapConfig {
	mapConfigs := make([]MapConfig, 4)
	mapConfigs[0] = MapConfig{"day", Overworld, false}
	mapConfigs[1] = MapConfig{"night", Overworld, true}
	mapConfigs[2] = MapConfig{"nether", Nether, false}
	mapConfigs[3] = MapConfig{"end", End, false}
	return mapConfigs
}
