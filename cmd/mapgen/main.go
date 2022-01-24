package main

func main() {

	worldPath := "C://dev//go//minecraft-mapper//World_of_Duane//"
	outputPath := "C://dev//go//minecraft-mapper//assets//test5//"
	dimension := Overworld
	night := false
	generateMap(worldPath, outputPath, dimension, night)

}

/*
command

./unmined-cli.exe image render --trim --world="C:/Users/micha/AppData/Roaming/.minecraft/saves/World of Duane" --output="C:/dev/go/minecraft-mapper/assets/test3/World_of_Duane.png"

./unmined-cli.exe web render --world="C:/Users/micha/AppData/Roaming/.minecraft/saves/World of Duane" --output="C:/dev/go/minecraft-mapper/assets/test5/day/"
./unmined-cli.exe web render --world="C:/Users/micha/AppData/Roaming/.minecraft/saves/World of Duane" --output="C:/dev/go/minecraft-mapper/assets/test4/night/" --night=true
./unmined-cli.exe web render --world="C:/Users/micha/AppData/Roaming/.minecraft/saves/World of Duane" --output="C:/dev/go/minecraft-mapper/assets/test4/nether/" --dimension=nether
./unmined-cli.exe web render --world="C:/Users/micha/AppData/Roaming/.minecraft/saves/World of Duane" --output="C:/dev/go/minecraft-mapper/assets/test4/end/" --dimension=end



"/C:/Users/micha/AppData/Roaming/.minecraft/saves/World of Duane"





C:/dev/go/minecraft-mapper/third_party/unmined/unmined-cli.exe web render --world="C://dev//go//minecraft-mapper//World_of_Duane//" --output="C://dev/go//minecraft-mapper//assets//test5//" --dimension=overworld --night=false --log-level=error

*/
