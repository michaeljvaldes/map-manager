package zip

import (
	"archive/zip"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

func ZipDirectory(srcDir string, zipFile string) {

	srcDir += string(filepath.Separator)
	// Get a Buffer to Write To
	outFile, err := os.Create(zipFile)
	if err != nil {
		handleError(err, "Error creating zip file")
	}
	defer outFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)

	// Add some files to the archive.
	addFiles(w, srcDir, "")

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		handleError(err, "Error closing zip file")
	}

}

func addFiles(w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := ioutil.ReadDir(filepath.Clean(filepath.FromSlash(basePath)))
	if err != nil {
		handleError(err, "Error reading source dir for zip")
	}

	for _, file := range files {
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				handleError(err, "Error reading source file for zip")
			}

			// Add some files to the archive.
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				handleError(err, "Error creating file in zip")
			}
			_, err = f.Write(dat)
			if err != nil {
				handleError(err, "Error writing to file in zip")
			}

		} else if file.IsDir() {
			// Recurse
			newBase := basePath + file.Name() + "/"
			addFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}
}

func handleError(err error, context string) {
	prepErr := ZipError{Err: err, Context: context}
	log.Fatal(prepErr.ErrorMessage())
}
