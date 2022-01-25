package mapdeploy

import (
	"archive/zip"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

func ZipDirectory(srcDir string, zipFile string) {

	srcDir += string(filepath.Separator)
	// Get a Buffer to Write To
	outFile, err := os.Create(zipFile)
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)

	// Add some files to the archive.
	addFiles(w, srcDir, "")

	if err != nil {
		fmt.Println(err)
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}

}

func addFiles(w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := ioutil.ReadDir(filepath.Clean(filepath.FromSlash(basePath)))
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		fmt.Println(basePath + file.Name())
		if !file.IsDir() {
			dat, err := ioutil.ReadFile(basePath + file.Name())
			if err != nil {
				fmt.Println(err)
			}

			// Add some files to the archive.
			f, err := w.Create(baseInZip + file.Name())
			if err != nil {
				fmt.Println(err)
			}
			_, err = f.Write(dat)
			if err != nil {
				fmt.Println(err)
			}

		} else if file.IsDir() {

			// Recurse
			newBase := basePath + file.Name() + "/"
			fmt.Println("Recursing and Adding SubDir: " + file.Name())
			fmt.Println("Recursing and Adding SubDir: " + newBase)

			addFiles(w, newBase, baseInZip+file.Name()+"/")
		}
	}
}

// func ZipDirectory(srcDir string, zipFile string) {
// 	file, err := os.Create(zipFile)
// 	if err != nil {
// 		panic(err)
// 	}
// 	defer file.Close()

// 	w := zip.NewWriter(file)
// 	defer w.Close()

// 	walker := func(path string, info os.FileInfo, err error) error {
// 		fmt.Printf("Crawling: %#v\n", path)
// 		if err != nil {
// 			return err
// 		}
// 		if info.IsDir() {
// 			return nil
// 		}
// 		file, err := os.Open(path)
// 		if err != nil {
// 			return err
// 		}
// 		defer file.Close()

// 		// Ensure that `path` is not absolute; it should not start with "/".
// 		// This snippet happens to work because I don't use
// 		// absolute paths, but ensure your real-world code
// 		// transforms path into a zip-root relative path.
// 		relPath, err := filepath.Rel(srcDir, path)
// 		if err != nil {
// 			return err
// 		}

// 		f, err := w.Create(relPath)
// 		if err != nil {
// 			return err
// 		}

// 		_, err = io.Copy(f, file)
// 		if err != nil {
// 			return err
// 		}

// 		return nil
// 	}
// 	err = filepath.Walk(srcDir, walker)
// 	if err != nil {
// 		panic(err)
// 	}
// }
