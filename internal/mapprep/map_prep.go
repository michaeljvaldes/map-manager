package mapprep

import (
	"embed"
	"io"
	"io/fs"
	"log"
	"os"
	"path"
)

//go:embed "assets/*"
var assets embed.FS

func PrepareMaps(baseDirectory string) {
	log.Println("Creating maps site for deployment")
	copyDir("assets", baseDirectory)
	log.Println("Finished creating maps site for deployment")
}

func copyDir(src string, dst string) {
	var err error
	var fds []fs.DirEntry
	var srcfile fs.File
	var srcinfo fs.FileInfo

	srcfile, _ = assets.Open(src)

	if srcinfo, err = srcfile.Stat(); err != nil {
		handleError(err, "Error getting info for src file")
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		handleError(err, "Error making directory")
	}

	if fds, err = assets.ReadDir(src); err != nil {
		handleError(err, "Error reading src directory")
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			copyDir(srcfp, dstfp)
		} else {
			copyFile(srcfp, dstfp)
		}
	}
}

func copyFile(src, dst string) {
	var err error
	var srcfd fs.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = assets.Open(src); err != nil {
		handleError(err, "Error opening asset "+src)
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		handleError(err, "Error creating dst file")
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		handleError(err, "Error copying src file to dst file")
	}
	if srcinfo, err = srcfd.Stat(); err != nil {
		handleError(err, "Error getting info for src file ")
	}

	err = os.Chmod(dst, srcinfo.Mode())
	if err != nil {
		handleError(err, "Error changing mode of dst file")
	}
}

func handleError(err error, context string) {
	prepErr := PrepError{Err: err, Context: context}
	log.Fatal(prepErr.ErrorMessage())
}
