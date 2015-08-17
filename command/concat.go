package command

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/codegangsta/cli"
	"golang.org/x/text/encoding/japanese"
	"golang.org/x/text/transform"
)

func getImportLine(fname string) ([]string, error) {
	var err error
	importLines := make([]string, 0)

	f, err := os.Open(fname)
	if err != nil {
		fmt.Fprintf(os.Stderr, "File %s could not read: %v\n", fname, err)
		os.Exit(1)
	}
	defer f.Close()

	basepath := filepath.Dir(fname)
	re := regexp.MustCompile("^// import [\"|'](.*)[\"|']")

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()

		if m := re.FindAllStringSubmatch(line, -1); m != nil {
			fpath := filepath.Join(basepath, m[0][1])
			if filepath.Ext(fpath) == "" {
				fpath = fpath + ".js"
			}
			lines, err := getImportLine(fpath)
			if err != nil {
				fmt.Errorf("%s", err)
			}

			for _, line := range lines {
				importLines = append(importLines, line)
			}
		} else {
			importLines = append(importLines, line)
		}
	}
	if err = scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "File %s scan error: %v\n", fname, err)
	}
	return importLines, err
}

func outputConcat(inpath string, outpath string, header string) error {
	var err error
	var lines []string
	lines, err = getImportLine(inpath)
	if err != nil {
		fmt.Errorf("%s", err)
	}

	fw, err := os.Create(outpath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Create out.js Error: %s", err)
	}
	defer fw.Close()
	fw.WriteString(header)
	fw.WriteString("\r\n\r\n")
	for _, line := range lines {
		sjisLine, err := utf8Toshiftjis(line)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error in encoding.")
		}
		fw.WriteString(sjisLine)
		fw.WriteString("\r\n")
	}

	return err
}

func getOutName(fname string) string {
	parent, child := filepath.Split(fname)
	ext := filepath.Ext(child)
	base := strings.TrimSuffix(child, ext)
	outDir := filepath.Join(parent, "cmd")

	d, e := os.Stat(outDir)
	if e != nil || !d.IsDir() {
		os.Mkdir(outDir, os.ModePerm)
	}

	return filepath.Join(outDir, base+".cmd")
}

func utf8Toshiftjis(utf8 string) (string, error) {
	var err error
	r := strings.NewReader(utf8)
	t := transform.NewReader(r, japanese.ShiftJIS.NewEncoder())
	ret, err := ioutil.ReadAll(t)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Encoding UTF-8 to ShiftJIS Error !")
		return "", err
	}
	return string(ret), err
}

func commandExecute(c *cli.Context, header string) error {
	var err error

	if len(c.Args()) != 1 {
		fmt.Fprintln(os.Stderr, "Input file is not found !")
		os.Exit(1)
	}

	inpath := c.Args()[0]
	if _, err = os.Stat(inpath); err != nil {
		fmt.Fprintln(os.Stderr, "Input file is not found !")
		os.Exit(1)
	}
	outpath := getOutName(inpath)

	fmt.Println("In  path:", inpath)
	fmt.Println("Out path:", outpath)

	err = outputConcat(inpath, outpath, header)
	if err != nil {
		fmt.Errorf("%s\n", err)
	}
	return err
}
