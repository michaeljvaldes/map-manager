package copy

import (
	"embed"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path"
)

func CopyEmbedDir(srcFS embed.FS, src string, dst string) error {
	var err error
	var fds []fs.DirEntry
	var srcfile fs.File
	var srcinfo fs.FileInfo

	srcfile, _ = srcFS.Open(src)

	if srcinfo, err = srcfile.Stat(); err != nil {
		return fmt.Errorf("error getting info for src file : %w", err)
	}

	if err = os.MkdirAll(dst, srcinfo.Mode()); err != nil {
		return fmt.Errorf("error making directory : %w", err)
	}

	if fds, err = srcFS.ReadDir(src); err != nil {
		return fmt.Errorf("error reading src directory : %w", err)
	}
	for _, fd := range fds {
		srcfp := path.Join(src, fd.Name())
		dstfp := path.Join(dst, fd.Name())

		if fd.IsDir() {
			err = CopyEmbedDir(srcFS, srcfp, dstfp)
		} else {
			err = CopyEmbedFile(srcFS, srcfp, dstfp)
		}
		if err != nil {
			return err
		}
	}
	return nil
}

func CopyEmbedFile(srcFS embed.FS, src, dst string) error {
	var err error
	var srcfd fs.File
	var dstfd *os.File
	var srcinfo os.FileInfo

	if srcfd, err = srcFS.Open(src); err != nil {
		fmt.Errorf("error opening asset "+src+" : %w", err)
	}
	defer srcfd.Close()

	if dstfd, err = os.Create(dst); err != nil {
		fmt.Errorf("error creating dst file : %w", err)
	}
	defer dstfd.Close()

	if _, err = io.Copy(dstfd, srcfd); err != nil {
		fmt.Errorf("error copying src file to dst file : %w", err)
	}
	if srcinfo, err = srcfd.Stat(); err != nil {
		fmt.Errorf("error getting info for src file : %w", err)
	}

	err = os.Chmod(dst, srcinfo.Mode())
	if err != nil {
		return fmt.Errorf("error changing mode of dst file : %w", err)
	}
	return nil
}
