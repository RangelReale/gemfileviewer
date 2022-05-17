package main

import (
	"archive/tar"
	"compress/gzip"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path"
)

var (
	outputFileNo   = flag.Int("n", -1, "output numbered file N")
	gemspecFile    = ""
	fileOutputName = ""
	fileOutput     = ""
)

func main() {
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Missing fileame parameter")
		os.Exit(1)
	}
	err := outputFile(flag.Arg(0))
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
		switch hdr.Name {
		case "metadata.gz":
			fmt.Printf("====== METADATA (%s) ======\n", hdr.Name)
			err = showgz(tr)
			fmt.Println()
		case "data.tar.gz":
			fmt.Printf("====== DATA (%s) ======\n", hdr.Name)
			err = showtgz(tr)
			fmt.Println()
		}
		if err != nil {
			return err
		}
	}

	if gemspecFile != "" {
		fmt.Printf("====== GEMSPEC ======\n%s\n", gemspecFile)
	}

	if fileOutput != "" {
		fmt.Printf("====== FILE OUTPUT (%s) ======\n%s\n", fileOutputName, fileOutput)
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

func showtgz(r io.Reader) error {
	gzf, err := gzip.NewReader(r)
	if err != nil {
		return err
	}
	tarReader := tar.NewReader(gzf)
	ct := 0
	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return err
		}

		name := header.Name

		ct++
		fmt.Printf("[%d] ", ct)

		switch header.Typeflag {
		case tar.TypeDir:
			fmt.Println("[DIR] ", name)
		case tar.TypeReg:
			fmt.Println(name)
			if path.Ext(name) == ".gemspec" && path.Dir(name) == "." {
				gsfile, err := ioutil.ReadAll(tarReader)
				if err == nil {
					gemspecFile = string(gsfile)
				}
			}
			if ct == *outputFileNo {
				ofile, err := ioutil.ReadAll(tarReader)
				if err == nil {
					fileOutputName = name
					fileOutput = string(ofile)
				}
			}
		}
	}
	return nil
}
