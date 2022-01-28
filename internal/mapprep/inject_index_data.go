package mapprep

import (
	"embed"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"time"
)

const indexTemplatePath = "assets/index.html"

func InjectIndexData(srcFS embed.FS, srcPath string, dstDir string) {
	tmpl := template.Must(template.ParseFS(srcFS, indexTemplatePath))
	output, err := os.Create(filepath.Join(dstDir, filepath.Base(indexTemplatePath)))
	if err != nil {
		log.Fatal(fmt.Errorf("error creating empty index.html : %w", err))
	}
	data := ConstructIndexData(time.Now())
	err = tmpl.Execute(output, data)
	if err != nil {
		log.Fatal(fmt.Errorf("error injecting data into index.html template : %w", err))
	}
}
