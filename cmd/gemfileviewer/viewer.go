package main

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Missing fileame parameter")
		os.Exit(1)
	}
	err := outputFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func outputFile(filename string) error {
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	// Open and iterate through the files in the archive.
	tr := tar.NewReader(f)
	for {
		hdr, err := tr.Next()
		if err == io.EOF {
			break // End of archive
		}
		if err != nil {
			return err
		}
		err = nil
		var gemspecfile string
		switch hdr.Name {
		case "metadata.gz":
			fmt.Printf("====== METADATA (%s) ======\n", hdr.Name)
			err = showgz(tr)
			fmt.Println()
		case "data.tar.gz":
			fmt.Printf("====== DATA (%s) ======\n", hdr.Name)
			gemspecfile, err = showtgz(tr)
			fmt.Println()
		}
		if err != nil {
			return err
		}
		if gemspecfile != "" {
			fmt.Printf("====== GEMSPEC ======\n%s\n", gemspecfile)
		}
	}
	return nil
}

func showgz(r io.Reader) error {
	zr, err := gzip.NewReader(r)
	if err != nil {
		log.Fatal(err)
	}
	defer zr.Close()
	if _, err := io.Copy(os.Stdout, zr); err != nil {
		log.Fatal(err)
	}
	return zr.Close()
}

func showtgz(r io.Reader) (string, error) {
	var gemspecfile string

	gzf, err := gzip.NewReader(r)
	if err != nil {
		return "", err
	}
	tarReader := tar.NewReader(gzf)
	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return "", err
		}

		name := header.Name

		switch header.Typeflag {
		case tar.TypeDir:
			fmt.Println("[DIR] ", name)
		case tar.TypeReg:
			fmt.Println(name)
			if path.Ext(name) == ".gemspec" && path.Dir(name) == "." {
				gsfile, err := ioutil.ReadAll(tarReader)
				if err == nil {
					gemspecfile = string(gsfile)
				}
			}
		}
	}
	return gemspecfile, nil
}
