import pytest
from swagger_update import *


def test_get_yaml_files_from_readme_with_invalid_file_path_returns_none():
    result = get_all_api_versions_from_readme(
        "./scripts/swagger_update/test_data/does_not_exist.md"
    )
    assert (
        result == None
    ), "get_yaml_files_from_readme should return None for file that doesn't exist"


def test_get_yaml_files_from_readme_with_invalid_file_path_returns_correct_list():
    api_versions = get_all_api_versions_from_readme(
        "./scripts/swagger_update/test_data/file_with_simple_tags.md"
    )
    assert len(api_versions) == 6, "Expected 5 versions"

    api_version = api_versions[0]
    assert api_version.get_name() == "package-2019-06-preview-only"
    input_files = api_version.get_input_files()
    assert len(input_files) == 1
    assert (
        input_files[0]
        == "Microsoft.ContainerRegistry/preview/2019-06-01-preview/containerregistry_build.json"
    )

    api_version = api_versions[1]
    assert api_version.get_name() == "package-2019-05"
    input_files = api_version.get_input_files()
    assert len(input_files) == 2
    assert (
        input_files[0]
        == "Microsoft.ContainerRegistry/stable/2019-05-01/containerregistry.json"
    )
    assert (
        input_files[1]
        == "Microsoft.ContainerRegistry/stable/2019-04-01/containerregistry_build.json"
    )

    api_version = api_versions[2]
    assert api_version.get_name() == "package-2019-05-preview"
    input_files = api_version.get_input_files()
    assert len(input_files) == 2
    assert (
        input_files[0]
        == "Microsoft.ContainerRegistry/stable/2017-10-01/containerregistry.json"
    )
    assert (
        input_files[1]
        == "Microsoft.ContainerRegistry/preview/2019-05-01-preview/containerregistry_scopemap.json"
    )

    api_version = api_versions[3]
    assert api_version.get_name() == "package-2017-03"
    input_files = api_version.get_input_files()
    assert len(input_files) == 1
    assert input_files[0] == "Microsoft.ContainerRegistry/stable/2017-03-01/containerregistry.json"

    api_version = api_versions[4]
    assert api_version.get_name() == "package-2016-06-preview"
    input_files = api_version.get_input_files()
    assert len(input_files) == 1
    assert input_files[0] == "Microsoft.ContainerRegistry/preview/2016-06-27-preview/containerregistry.json"

    api_version = api_versions[5]
    assert api_version.get_name() == "all-api-versions"
    input_files = api_version.get_input_files()
    assert len(input_files) == 0
