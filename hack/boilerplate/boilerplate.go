/*
Copyright YEAR The XXX Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// regular expressions
var (
	now       = time.Now()
	thisYear  = now.Year()
	startYear = 2014

	// Search for "YEAR" which exists in the boilerplate,
	// but shouldn't in the real thing.
	yearRe = regexp.MustCompile("YEAR")

	// finds all dates from startYear to the current Year.
	dateRe = regexp.MustCompile(getDateRegex())

	// strip // +build \n\n build constraints
	goBuildConstraintsRe = regexp.MustCompile(`^(// \+build.*\n)+\n`)

	// strip #!.* from shell scripts
	shebangRe = regexp.MustCompile(`^(#!.*\n)\n*`)

	// Search for generated files
	generatedRe = regexp.MustCompile(`DO NOT EDIT`)
)

// flags
var (
	boilerplateDir string
	verbose        bool
)

const (
	rootDir = "."
	// this file extension is used for Re files
	generatedGoFileExtension  = "generatego"
	generatedBzlFileExtension = "generatebzl"
)

// hardcoded settings
var (
	// skipped files and directories
	skipped = map[string]struct{}{
		"bin":                           struct{}{},
		"Godeps":                        struct{}{},
		".git":                          struct{}{},
		"vendor":                        struct{}{},
		"_gopath":                       struct{}{},
		"_output":                       struct{}{},
		"cluster/env.sh":                struct{}{},
		"test/e2e/generated/bindata.go": struct{}{},
		"staging/src/k8s.io/kubectl/pkg/generated/bindata.go": struct{}{},
		"hack/boilerplate/test":                               struct{}{},
		"pkg/apis/kubeadm/v1beta1/bootstraptokenstring.go":    struct{}{},
		"pkg/apis/kubeadm/v1beta1/types.go":                   struct{}{},
		"pkg/apis/kubeadm/v1beta1/zz_generated.deepcopy.go":   struct{}{},
		// third_party folders
		"third_party": struct{}{},
		"staging/src/k8s.io/apimachinery/third_party":   struct{}{},
		"staging/src/k8s.io/client-go/third_party":      struct{}{},
		"staging/src/k8s.io/code-generator/third_party": struct{}{},
	}

	// list all the files contain 'DO NOT EDIT', but are not generated
	skippedUngeneratedFiles = map[string]struct{}{
		"hack/boilerplate/boilerplate.py": struct{}{},
		"hack/lib/swagger.sh":             struct{}{},
	}
)

func main() {
	flag.StringVar(&boilerplateDir, "boilerplate-dir", "./hack/boilerplate", "Directory containing the boilerplate files for file extensions.")
	flag.BoolVar(&verbose, "verbose", false, "give verbose output regarding why a file does not pass.")
	flag.Parse()

	failed, err := run(os.Stdout, rootDir)
	if err != nil {
		fmt.Fprintf(os.Stderr, err.Error()+"\n")
		os.Exit(1)
	}
	if failed {
		// checks print the files that have not passed directly
		// so we can exit here without printing another notice to the user
		os.Exit(1)
	}
}

func run(out io.Writer, rootDir string) (failed bool, err error) {
	boilerplateMap, err := getBoilerplateForExtensions()
	if err != nil {
		return false, fmt.Errorf("getting boilerplate files: %v", err)
	}

	files := flag.Args()
	if len(files) == 0 {
		files, err = getFiles(rootDir, boilerplateMap)
		if err != nil {
			return false, fmt.Errorf("getting files to check: %v", err)
		}
	}

	var failedFiles []string
	for _, file := range files {
		ok, err := filePasses(file, boilerplateMap, out)
		if err != nil {
			return false, fmt.Errorf("checking file: %v", err)
		}
		if !ok {
			failedFiles = append(failedFiles, file)
		}
	}

	failed = len(failedFiles) > 0
	if failed {
		fmt.Fprintln(out, "Boilerplate header is wrong for:")
	}
	for _, failedFile := range failedFiles {
		fmt.Fprintln(out, failedFile)
	}
	return
}

func getFiles(rootDir string, extensions map[string]string) (files []string, err error) {
	err = filepath.Walk(rootDir,
		func(currentPath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			checkPath := strings.TrimPrefix(currentPath, "./")
			for skip := range skipped {
				if strings.HasPrefix(checkPath, skip) {
					return nil
				}
			}

			// only record files
			if info.IsDir() {
				return nil
			}

			// Skip files that we don't have a boilerplate for
			ext := fileExtension(currentPath)
			base := path.Base(currentPath)
			_, extFound := extensions[ext]
			_, baseFound := extensions[base]
			if !extFound && !baseFound {
				return nil
			}

			files = append(files, currentPath)
			return nil
		})

	return
}

func isGenerated(filename string, content []byte) bool {
	if _, ok := skippedUngeneratedFiles[filename]; ok {
		return false
	}
	return generatedRe.Match(content)
}

func filePasses(filename string, boilerplateMap map[string]string, out io.Writer) (bool, error) {
	fileContent, err := ioutil.ReadFile(filename)
	if err != nil {
		return false, fmt.Errorf("opening file: %v", err)
	}

	// determine if the file is automatically generated
	generated := isGenerated(filename, fileContent)

	// determine extension to use to lookup the boilerplate header
	extension := fileExtension(filename)
	if generated {
		switch extension {
		case "go":
			extension = generatedGoFileExtension
		case "bzl":
			extension = generatedBzlFileExtension
		}
	}
	if extension == "" {
		extension = path.Base(filename)
	}
	boilerplate, ok := boilerplateMap[extension]
	if !ok {
		return false, fmt.Errorf("no boilerplate registered for extension: %q", extension)
	}

	// remove extra content from the top of files
	// fileContent := string(fileContent)
	switch extension {
	case "go", "generatego":
		fileContent = goBuildConstraintsRe.ReplaceAll(fileContent, nil)
	case "sh", "py":
		fileContent = shebangRe.ReplaceAll(fileContent, nil)
	}

	// if our test file is smaller than the reference it surely fails!
	if len(fileContent) < len(boilerplate) {
		if verbose {
			fmt.Fprintf(out, "%s: is smaller than reference (%d < %d)\n",
				filename, len(fileContent), len(boilerplate))
		}
		return false, nil
	}

	// trim the file to the same length as our boilerplate header
	fileContent = fileContent[0:len(boilerplate)]

	// is YEAR in the file
	if yearRe.Match(fileContent) {
		if verbose {
			if generated {
				fmt.Fprintf(out, "%s: has the YEAR field, but it should not be in generated file\n", filename)
			} else {
				fmt.Fprintf(out, "%s: has the YEAR field, but missing the year of date\n", filename)
			}
		}
		return false, nil
	}

	if !generated {
		// replace the actual year eg. "2014" with "YEAR" to compare with the boilerplate file
		fileContent = dateRe.ReplaceAll(fileContent, []byte("YEAR"))
	}

	// check if the file header matches the boilerplate
	actual := strings.TrimSpace(string(fileContent))
	expected := strings.TrimSpace(boilerplate)
	if actual != expected {
		if verbose {
			fmt.Printf("tested %q with %q extension\n", filename, extension)
			fmt.Fprintf(out, "%s: does not match reference\nis:     %q\nshould: %q\n\n", filename, actual, expected)
		}
		return false, nil
	}

	return true, nil
}

func fileExtension(filename string) string {
	base := path.Base(filename)
	i := strings.LastIndex(base, ".")
	if i < 0 {
		// no dot found in the filename
		return ""
	}

	return base[i+1 : len(base)]
}

// getBoilerplateForExtensions reads the boilerplate.*.txt files in the directory
// and returns a map of file extension to the files content
func getBoilerplateForExtensions() (map[string]string, error) {
	boilerplate := map[string]string{}

	matches, err := filepath.Glob(path.Join(boilerplateDir, "boilerplate.*.txt"))
	if err != nil {
		return nil, fmt.Errorf("finding files via glob: %v", err)
	}
	for _, match := range matches {
		parts := strings.Split(path.Base(match), ".")
		if len(parts) != 3 {
			return nil, fmt.Errorf("wrong filename for boilerplate file: %q should be \"boilerplate.EXTENSION.txt\"", match)
		}

		content, err := ioutil.ReadFile(match)
		if err != nil {
			return nil, fmt.Errorf("reading file: %v", err)
		}

		// map file extension to the boilerplate for the file
		boilerplate[parts[1]] = string(content)
	}

	return boilerplate, nil
}

// getDateRegex returns a regex like "(2014|2015|2016|2017|2018)"
// containing all years from 2014 until the current year.
func getDateRegex() string {
	var dates []string
	for i := startYear; i <= thisYear; i++ {
		dates = append(dates, strconv.Itoa(i))
	}
	return "(" + strings.Join(dates, "|") + ")"
}
