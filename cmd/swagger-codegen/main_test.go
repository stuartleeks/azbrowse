package main

import (
	"testing"

	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func Test_getVersionsForFolder_FileDoesNotExist_ReturnsError(t *testing.T) {

	_, err := getVersionsForFolder("./testdata/file_does_not_exist.md")

	assert.Assert(t, is.ErrorContains(err, "no such file or directory"))
}

func Test_getVersionsForFolder_InvalidFile_ReturnsEmptyResult(t *testing.T) {

	result, err := getVersionsForFolder("./testdata/invalid_file_format.md")

	assert.Assert(t, is.Nil(err))
	assert.Assert(t, is.Len(result, 0))
}

func Test_getVersionsForFolder_ValidFile_ReturnsCorrectVersions(t *testing.T) {

	versions, err := getVersionsForFolder("./testdata/valid_file.md")

	assert.Assert(t, is.Nil(err))
	assert.Assert(t, is.Len(versions, 4))

	version := versions[0]
	assert.Assert(t, is.Equal(version.Name, "package-2019-12-preview"))
	assert.Assert(t, is.Equal(len(version.Files), 1)) // TODO - better way to compare slices?
	assert.Assert(t, is.Equal(version.Files[0], "Microsoft.ContainerRegistry/preview/2019-12-01-preview/containerregistry.json"))

	assert.Assert(t, is.Equal(versions[1].Name, "package-2019-06-preview"))
	assert.Assert(t, is.Equal(versions[2].Name, "package-2019-06-preview-only"))
	assert.Assert(t, is.Equal(versions[3].Name, "package-2019-05"))
}
