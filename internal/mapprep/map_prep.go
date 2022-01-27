package mapprep

import (
	"embed"
	"fmt"
	"log"

	"mapmanager/internal/io/copy"
)

//go:embed "assets/*"
var assets embed.FS

func PrepareMaps(baseDirectory string) {
	log.Println("Creating maps site for deployment")
	err := copy.CopyEmbedDir(assets, "assets", baseDirectory)
	if err != nil {
		log.Fatal(fmt.Errorf("error copying assets : %w", err))
	}
	log.Println("Finished creating maps site for deployment")
}
