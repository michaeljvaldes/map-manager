package mapprep

import (
	"embed"
	"fmt"
	"log"

	"mapmanager/internal/io/copy"
)

//go:embed "assets/*"
var assets embed.FS

const copyPath = "assets/copy"
const templatePath = "assets/template"

func PrepareMaps(baseDirectory string) {
	log.Println("Creating maps site for deployment")
	err := copy.CopyEmbedDir(assets, copyPath, baseDirectory)
	if err != nil {
		log.Fatal(fmt.Errorf("error copying assets : %w", err))
	}
	InjectIndexData(assets, "assets", baseDirectory)
	log.Println("Finished creating maps site for deployment")
}
