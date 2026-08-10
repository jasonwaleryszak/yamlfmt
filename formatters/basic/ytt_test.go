// Copyright 2026 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package basic_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestYttExamples exercises ytt syntax using reduced cases taken from ytt's
// first-party examples. Keeping the cases as formatter goldens makes it easy to
// add syntax without coupling the test suite to a ytt checkout or executable.
func TestYttExamples(t *testing.T) {
	testCases := []string{
		"block_scalar",
		"control_flow",
		"inline_expression",
		"overlay_annotations",
		"schema_annotation",
		"sequence_item_comments",
		"text_templated_strings",
	}

	formatter, err := factory.NewFormatter(nil)
	require.NoError(t, err)

	for _, name := range testCases {
		t.Run(name, func(t *testing.T) {
			folder := filepath.Join("testdata", "ytt", name)
			input, err := os.ReadFile(filepath.Join(folder, "input.yaml"))
			require.NoError(t, err)
			expected, err := os.ReadFile(filepath.Join(folder, "expected.yaml"))
			require.NoError(t, err)

			actual, err := formatter.Format(input)
			require.NoError(t, err)
			require.Equal(t, string(expected), string(actual))
		})
	}
}
