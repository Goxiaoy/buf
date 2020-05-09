// Copyright 2020 Buf Technologies Inc.
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

package lint

import (
	"bytes"
	"context"
	"path/filepath"
	"testing"

	"github.com/bufbuild/buf/internal/pkg/app"
	"github.com/bufbuild/buf/internal/pkg/app/appproto"
	"github.com/bufbuild/buf/internal/pkg/ext/extdescriptor"
	"github.com/bufbuild/buf/internal/pkg/util/utilproto"
	"github.com/bufbuild/buf/internal/pkg/util/utilproto/utilprototesting"
	"github.com/bufbuild/buf/internal/pkg/util/utilstring"
	plugin_go "github.com/golang/protobuf/protoc-gen-go/plugin"
	"github.com/stretchr/testify/require"
)

func TestRunLint1(t *testing.T) {
	testRunLint(
		t,
		filepath.Join("testdata", "fail"),
		[]string{
			filepath.Join("testdata", "fail", "buf", "buf.proto"),
			filepath.Join("testdata", "fail", "buf", "buf_two.proto"),
		},
		"",
		[]string{
			filepath.Join("buf", "buf.proto"),
			filepath.Join("buf", "buf_two.proto"),
		},
		0,
		`
		buf/buf.proto:3:1:Files with package "other" must be within a directory "other" relative to root but were in directory "buf".
		buf/buf.proto:3:1:Package name "other" should be suffixed with a correctly formed version, such as "other.v1".
		buf/buf.proto:6:9:Field name "oneTwo" should be lower_snake_case, such as "one_two".
		buf/buf_two.proto:3:1:Files with package "other" must be within a directory "other" relative to root but were in directory "buf".
		buf/buf_two.proto:3:1:Package name "other" should be suffixed with a correctly formed version, such as "other.v1".
		buf/buf_two.proto:6:9:Field name "oneTwo" should be lower_snake_case, such as "one_two".
		`,
	)
}

func TestRunLint2(t *testing.T) {
	testRunLint(
		t,
		filepath.Join("testdata", "fail"),
		[]string{
			filepath.Join("testdata", "fail", "buf", "buf.proto"),
			filepath.Join("testdata", "fail", "buf", "buf_two.proto"),
		},
		"",
		[]string{
			filepath.Join("buf", "buf.proto"),
		},
		0,
		`
		buf/buf.proto:3:1:Files with package "other" must be within a directory "other" relative to root but were in directory "buf".
		buf/buf.proto:3:1:Package name "other" should be suffixed with a correctly formed version, such as "other.v1".
		buf/buf.proto:6:9:Field name "oneTwo" should be lower_snake_case, such as "one_two".
		`,
	)
}

func TestRunLint3(t *testing.T) {
	testRunLint(
		t,
		filepath.Join("testdata", "fail"),
		[]string{
			filepath.Join("testdata", "fail", "buf", "buf.proto"),
			filepath.Join("testdata", "fail", "buf", "buf_two.proto"),
		},
		`{"input_config":"testdata/fail/something.yaml"}`,
		[]string{
			filepath.Join("buf", "buf.proto"),
		},
		0,
		`
		buf/buf.proto:3:1:Files with package "other" must be within a directory "other" relative to root but were in directory "buf".
		`,
	)
}

func TestRunLint4(t *testing.T) {
	testRunLint(
		t,
		filepath.Join("testdata", "fail"),
		[]string{
			filepath.Join("testdata", "fail", "buf", "buf.proto"),
			filepath.Join("testdata", "fail", "buf", "buf_two.proto"),
		},
		`{"input_config":{"lint":{"use":["PACKAGE_DIRECTORY_MATCH"]}}}`,
		[]string{
			filepath.Join("buf", "buf.proto"),
		},
		0,
		`
		buf/buf.proto:3:1:Files with package "other" must be within a directory "other" relative to root but were in directory "buf".
		`,
	)
}

func TestRunLint5(t *testing.T) {
	testRunLint(
		t,
		filepath.Join("testdata", "fail"),
		[]string{
			filepath.Join("testdata", "fail", "buf", "buf.proto"),
			filepath.Join("testdata", "fail", "buf", "buf_two.proto"),
		},
		`{"input_config":{"lint":{"use":["PACKAGE_DIRECTORY_MATCH"]}},"error_format":"json"}`,
		[]string{
			filepath.Join("buf", "buf.proto"),
		},
		0,
		`
		{"path":"buf/buf.proto","start_line":3,"start_column":1,"end_line":3,"end_column":15,"type":"PACKAGE_DIRECTORY_MATCH","message":"Files with package \"other\" must be within a directory \"other\" relative to root but were in directory \"buf\"."}
		`,
	)
}

func testRunLint(
	t *testing.T,
	root string,
	realFilePaths []string,
	parameter string,
	fileToGenerate []string,
	expectedExitCode int,
	expectedErrorString string,
) {
	t.Parallel()

	testRunHandlerFunc(
		t,
		handle,
		testBuildCodeGeneratorRequest(
			t,
			root,
			realFilePaths,
			parameter,
			fileToGenerate,
		),
		expectedExitCode,
		expectedErrorString,
	)
}

func testRunHandlerFunc(
	t *testing.T,
	handlerFunc func(
		context.Context,
		app.EnvStderrContainer,
		appproto.ResponseWriter,
		*plugin_go.CodeGeneratorRequest,
	),
	request *plugin_go.CodeGeneratorRequest,
	expectedExitCode int,
	expectedErrorString string,
) {
	requestData, err := utilproto.MarshalWire(request)
	require.NoError(t, err)
	stdin := bytes.NewReader(requestData)
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)

	exitCode := app.GetExitCode(
		appproto.Run(
			context.Background(),
			app.NewContainer(
				nil,
				stdin,
				stdout,
				stderr,
			),
			handlerFunc,
		),
	)

	require.Equal(t, expectedExitCode, exitCode, utilstring.TrimLines(stderr.String()))
	if exitCode == 0 {
		response := &plugin_go.CodeGeneratorResponse{}
		require.NoError(t, utilproto.UnmarshalWire(stdout.Bytes(), response))
		require.Equal(t, utilstring.TrimLines(expectedErrorString), response.GetError(), utilstring.TrimLines(stderr.String()))
	}
}

func testBuildCodeGeneratorRequest(
	t *testing.T,
	root string,
	realFilePaths []string,
	parameter string,
	fileToGenerate []string,
) *plugin_go.CodeGeneratorRequest {
	fileDescriptorSet, err := utilprototesting.GetProtocFileDescriptorSet(
		context.Background(),
		[]string{root},
		realFilePaths,
		true,
		true,
	)
	require.NoError(t, err)
	request, err := extdescriptor.FileDescriptorSetToCodeGeneratorRequest(fileDescriptorSet, parameter, fileToGenerate...)
	require.NoError(t, err)
	return request
}