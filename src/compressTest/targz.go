package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	fmt.Println("heloo")
	pattern := `D:\leetcode\*.java`
	matches, _ := filepath.Glob(pattern)
	for _, v := range matches {

		cmd := fmt.Sprintf("tar -rvf %s", v)
		fmt.Println(cmd)
		//command := exec.Command("/bin/bash", "-c", cmd)
		//bytes, e := command.Output()
		//if e!=nil{
		//	fmt.Println(e)
		//}
		//fmt.Println(string(bytes))

	}
	err := tartar("test.tar.gz", matches)
	if err != nil {
		fmt.Println(err)
	}

}
func tartar(tarName string, paths []string) (err error) {
	tarFile, err := os.Create(tarName)

	// enable compression if file ends in .gz
	tw := tar.NewWriter(tarFile)
	if strings.HasSuffix(tarName, ".gz") {
		gz := gzip.NewWriter(tarFile)
		defer gz.Close()
		tw = tar.NewWriter(gz)
	}
	defer tw.Close()

	for _, path := range paths {
		f, e := os.Stat(path)
		if e != nil {
			fmt.Print(e)
			return e
		}
		hdr := new(tar.Header)
		hdr.Name = path
		hdr.Size = f.Size()
		hdr.Mode = int64(f.Mode())
		hdr.ModTime = f.ModTime()

		if err != nil {
			return err
		}
		if err := tw.WriteHeader(hdr); err != nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		ff, err := os.Open(path)
		if err != nil {
			return err
		}

		_, e = io.Copy(tw, ff)
		if e != nil {
			return e
		}
	}
	return nil

}
