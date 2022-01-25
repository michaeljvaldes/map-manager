package mapprep

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
)

//go:embed "assets/*"
var assets embed.FS

func PrepareMaps(baseDirectory string) {
	copyAssets(baseDirectory)
}

func copyAssets(baseDirectory string) error {
	return copyDir("assets", baseDirectory)
}

func copyDir(src string, dst string) error {
	var err error
	var fds []fs.DirEntry
	var srcfile fs.File
	var srcinfo fs.FileInfo

	srcfile, _ = assets.Open(src)

	if srcinfo, err = srcfile.Stat(); err != nil {
		return err
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return err
	}

	if fds, err = assets.ReadDir(src); err != nil {
		return err
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			if err = copyDir(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		} else {
			if err = copyFile(srcfp, dstfp); err != nil {
				fmt.Println(err)
			}
		}
	}
	return nil
}

func copyFile(src, dst string) error {
	var err error
	var srcfd fs.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = assets.Open(src); err != nil {
		return err
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		return err
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		return err
	}
	if srcinfo, err = srcfd.Stat(); err != nil {
		return err
	}
	return os.Chmod(dst, srcinfo.Mode())
}
