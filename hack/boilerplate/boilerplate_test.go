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
	"bytes"
	"strings"
	"testing"
)

func Test_fileExtension(t *testing.T) {
	tests := []struct {
		name, file, extension string
	}{
		{
			name:      "Simple extension",
			file:      "test.go",
			extension: "go",
		},
		{
			name:      "dot in between",
			file:      "script.test.sh",
			extension: "sh",
		},
		{
			name:      "no extension",
			file:      "script",
			extension: "",
		},
		{
			name:      "dot in folder structure",
			file:      "k8s.io/script",
			extension: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ext := fileExtension(test.file)
			if ext != test.extension {
				t.Errorf("file extension should be %q, is: %q", test.extension, ext)
			}
		})
	}
}

var expected = `Boilerplate header is wrong for:
test/fail.go
test/fail.py`

func Test_run(t *testing.T) {
	var buf bytes.Buffer

	failed, err := run(&buf, "./test")
	if err != nil {
		t.Errorf("unexpected error: %v", err)
		return
	}
	if !failed {
		t.Errorf("should have failed")
	}
	output := strings.TrimSpace(buf.String())
	if output != expected {
		t.Errorf("unexpected messages printed:\n%s\nshould be:\n%s", output, expected)
	}
}
