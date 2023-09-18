package domain

import (
	"archive/tar"
	"compress/gzip"
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

func ArchiveFiles(files []string, path string) error {
	if err := checkDirectory(path); err != nil {
		return err
	}

	errs := make(chan error, 1)
	wg.Add(len(files))
	for _, file := range files {
		go func(file string) {
			errs <- archive(file, path)
		}(file)
	}
	wg.Wait()

	if err := <-errs; err != nil {
		return err
	}

	return nil
}

func checkDirectory(path string) error {
	if path != "" {
		fileInfo, err := os.Stat(path)
		if err != nil {
			return err
		}
		if is := fileInfo.IsDir(); !is {
			return errors.New("the specified path is not a directory")
		}
	}
	return nil
}

func archive(filename string, path string) error {
	defer wg.Done()

	if _, err := os.Stat(filename); err != nil {
		return errors.Wrap(err, "error writing archive")
	}

	out, err := os.Create(setName(filename, path) + ".tar.gz")
	if err != nil {
		return errors.Wrap(err, "error creating archive")
	}
	defer out.Close()

	err = createArchive(filename, out)
	if err != nil {
		return errors.Wrap(err, "error writing archive")
	}

	return nil
}

func setName(filename string, path string) string {
	if !strings.HasSuffix(path, "/") {
		path += "/"
	}
	filename = filepath.Base(filename)
	filename = strings.TrimSuffix(filename, filepath.Ext(filename))
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	return path + filename + "_" + timestamp
}

func createArchive(file string, buf io.Writer) error {
	gw := gzip.NewWriter(buf)
	defer gw.Close()
	tw := tar.NewWriter(gw)
	defer tw.Close()

	err := addToArchive(tw, file)
	if err != nil {
		return err
	}
	return nil
}

func addToArchive(tw *tar.Writer, filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	info, err := file.Stat()
	if err != nil {
		return err
	}

	header, err := tar.FileInfoHeader(info, info.Name())
	if err != nil {
		return err
	}

	header.Name = filename
	err = tw.WriteHeader(header)
	if err != nil {
		return err
	}

	_, err = io.Copy(tw, file)
	if err != nil {
		return err
	}
	return nil
}
