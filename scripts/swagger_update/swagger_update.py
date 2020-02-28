from git import Repo
import shutil
import os
import re
import yaml

# TODO
# - add docs to top to state goals (recreate folder structure for any files that are copied by looking up in readme.md)
# - add function to find latest version (add unit tests for this!)

# resource_provider_version_overrides is keyed on RP name with the value being the tag to force
resource_provider_version_overrides = {
    "cosmos-db": "package-2019-08-preview",  # the 2019-12 version includes 2019-08-preview files that reference files not in the 2019-12 list!
    "frontdoor": "package-2019-10", # same issue as cosmos-db with spec contents referencing files not listed in the package
}


class ApiVersion:
    def __init__(self, name, input_files):
        # super().__init__()
        self.name = name
        self.input_files = input_files

    def get_name(self):
        return self.name

    def get_input_files(self):
        return self.input_files


def get_api_version_tag(resource_provider_name, readme_contents):
    override = resource_provider_version_overrides.get(resource_provider_name)
    if override != None:
        return override

    tag_regex = re.compile("openapi-type: arm\ntag: ([a-z\\-0-9]*)")
    match = tag_regex.search(readme_contents)
    if match == None:
        return None

    tag = match.group(1)
    return tag


def find_api_version(readme_contents, version_tag):
    code_block_end_regex = re.compile("^[\\s]*```[\\s]*$", flags=re.MULTILINE)

    # Regex to match:   ```yaml $(tag) == 'the-version-tag`
    # Also match:       ```yaml $(tag) == 'the-version-tag` || $(tag) == 'some-other-tag'
    # But don't match   ```yaml $(tag) == 'the-version-tag' && $(another-condition)
    start_match = re.search(
        "^```[\\s]*yaml [^&^\\n]*\\$\\(tag\\) == '" + version_tag + "'[^&^\\n]*$",
        readme_contents,
        flags=re.MULTILINE,
    )
    if start_match == None:
        return None

    end_match = code_block_end_regex.search(readme_contents, start_match.end())
    yaml_contents = readme_contents[start_match.end() : end_match.start()]

    yaml_data = yaml.load(yaml_contents, Loader=yaml.BaseLoader)
    input_files = []
    if yaml_data != None:
        input_files = [file.replace("\\", "/") for file in yaml_data["input-file"]]
    api_version = ApiVersion(version_tag, input_files)

    return api_version


def get_api_version_from_readme(resource_provider_name, readme_path):
    if not os.path.isfile(readme_path):
        return None
    with open(readme_path, "r", encoding="utf8") as stream:
        contents = stream.read()

    version_tag = get_api_version_tag(resource_provider_name, contents)
    if version_tag == None:
        return None

    api_version = find_api_version(contents, version_tag)
    return api_version


last_git_message = ""


def show_git_progress(op_code, cur_count, max_count, message):
    global last_git_message
    if message != "" and message != last_git_message:
        print(message)
        last_git_message = message


def clone_swagger_specs(target_folder):
    if os.path.exists(target_folder):
        print("Deleting " + target_folder + "...")
        shutil.rmtree(target_folder)

    print("Cloning specs...")
    repo = Repo()
    repo.clone_from(
        "git@github.com:azure/azure-rest-api-specs",
        target_folder,
        progress=show_git_progress,
        multi_options=["--depth=1"],
    )
    print("Cloned")
    repo.clone_from("https://github.com/azure/azure-rest-api-specs", "gittemp", progress=show_git_progress, multi_options=["--depth=1"])
    print ("Cloned")


class ResourceManagerApiSet:
    def __init__(self, resource_provider_name, base_folder, api_version):
        self.resource_provider_name = resource_provider_name
        self.base_folder = base_folder
        self.api_version = api_version
    def get_resource_provider_name(self):
        return self.resource_provider_name

    def get_base_folder(self):
        return self.base_folder

    def get_api_version(self):
        return self.api_version


def get_resource_manager_api_sets(base_folder):
    rp_folders = sorted([f.path for f in os.scandir(base_folder) if f.is_dir()])
    api_sets = []
    for folder in rp_folders:
        resource_provider_name = folder.split("/")[-1]
        readme_path = folder + "/resource-manager/readme.md"
        api_version = get_api_version_from_readme(resource_provider_name, readme_path)

        if api_version == None:
            print("***No api version found, ignoring: " + folder)
            continue

        spec_relative_folder = folder[len(base_folder) + 1 :]
        print(spec_relative_folder + ", using api-version " + api_version.get_name())

        api_set = ResourceManagerApiSet(
            resource_provider_name,
            spec_relative_folder + "/resource-manager",
            api_version,
        )
        api_sets.append(api_set)

    return api_sets


def get_folder_for_file(file):
    return file[0 : file.rfind("/")]


def copy_file_ensure_paths(source_base, target_base, file):
    source_file = source_base + "/" + file
    target_file = target_base + "/" + file
    print("--> " + target_file)
    target_folder = get_folder_for_file(target_file)
    os.makedirs(target_folder, exist_ok=True)
    shutil.copy(source_file, target_file)


def copy_child_folder_if_exists(source_base, target_base, relative_path):
    source_path = source_base + "/" + relative_path
    if os.path.exists(source_path):
        target_path = target_base + "/" + relative_path
        print("---> " + relative_path)
        shutil.copytree(source_path, target_path)


def copy_api_sets_to_swagger_specs(api_sets, source_folder, target_folder):
    for api_set in api_sets:
        print("\nCopying " + api_set.get_resource_provider_name())
        api_version = api_set.get_api_version()

        resource_provider_source = (
            source_folder
            + "/"
            + api_set.get_resource_provider_name()
            + "/resource-manager"
        )
        resource_provider_target = (
            target_folder
            + "/"
            + api_set.get_resource_provider_name()
            + "/resource-manager"
        )

        # The core of this method is to copy the files defined by the api version
        # There are some places where additional files that aren't listed in the definition tend to live
        # For now we're handling the separate cases (e.g. `common`)
        # We _could_ load the specs and scan for linked files and build out the list that way
        # Doing that would remove the need for these additional checks
        # as well as fixing the problem with definitions referenced back in other folders as with comsmos-db etc


        # Look for `common` folder under the `resource-manager` folder
        copy_child_folder_if_exists(
            resource_provider_source,
            resource_provider_target,
            "/common",
        )

        # Look for `common` folders under the resource type folder
        resource_type_folders = set(
            [x[0 : x.index("/")] for x in api_version.get_input_files()]
        )
        for resource_type_folder in resource_type_folders:
            copy_child_folder_if_exists(
                resource_provider_source,
                resource_provider_target,
                resource_type_folder + "/common",
            )

        # Look for `entityTypes` or `definitions` folders under api versions
        api_version_folders = set(
            [x[0 : x.rfind("/")] for x in api_version.get_input_files()]
        )
        for api_version_folder in api_version_folders:
            copy_child_folder_if_exists(
                resource_provider_source,
                resource_provider_target,
                api_version_folder + "/entityTypes",
            )
            copy_child_folder_if_exists(
                resource_provider_source,
                resource_provider_target,
                api_version_folder + "/definitions",
            )
            if os.path.exists(resource_provider_source + "/" + api_version_folder + "/common.json"):
                copy_file_ensure_paths(resource_provider_source, resource_provider_target, api_version_folder + "/common.json")


        # Copy the files definte in the api version
        for file in api_version.get_input_files():
            copy_file_ensure_paths(resource_provider_source, resource_provider_target, file)

    # TODO write json file per folder with contents to load?


if __name__ == "__main__":
    # clone_swagger_specs("swagger-temp")

    if os.path.exists("swagger-specs"):
        print("Deleting swagger-specs...")
        shutil.rmtree("swagger-specs")

    print(
        "\n****************************************************************************"
    )
    print("Discovering api sets:")
    api_sets = get_resource_manager_api_sets(
        "./swagger-temp/azure-rest-api-specs/specification"
    )

    print(
        "\n****************************************************************************"
    )
    print("Copying files:")
    copy_api_sets_to_swagger_specs(
        api_sets,
        "./swagger-temp/azure-rest-api-specs/specification",
        "./swagger-specs",
    )

    shutil.copytree(
        "./swagger-temp/azure-rest-api-specs/specification/common-types",
        "./swagger-specs/common-types",
    )
