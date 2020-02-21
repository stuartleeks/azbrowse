package main

import (
	"testing"

	"github.com/lawrencegripper/azbrowse/pkg/swagger"
	"gotest.tools/assert"
	is "gotest.tools/assert/cmp"
)

func Test_pickVersions_ChoosesLatestNonPreview(t *testing.T) {
	versions := []swagger.APIVersion{
		{Name: "package-2019-10-preview"},
		{Name: "package-2019-08-preview"},
		{Name: "package-2019-07"},
		{Name: "package-2019-06"},
		{Name: "package-2019-05-preview"},
	}

	version := pickVersion(versions)

	assert.Assert(t, is.Equal(version.Name, "package-2019-07"))
}

func Test_pickVersions_ChoosesLatestPreviewIfNoOthers(t *testing.T) {
	versions := []swagger.APIVersion{
		{Name: "package-2019-10-preview"},
		{Name: "package-2019-08-preview"},
		{Name: "package-2019-05-preview"},
	}

	version := pickVersion(versions)

	assert.Assert(t, is.Equal(version.Name, "package-2019-10-preview"))
}
