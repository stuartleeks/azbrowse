from git import Repo
import shutil
import os
import re
import yaml
import json

# TODO
# - add docs to top to state goals (recreate folder structure for any files that are copied by looking up in readme.md)
# - add function to find latest version (add unit tests for this!)

class ApiVersion:
    def __init__(self, name, input_files):
        self.name = name
        self.input_files = input_files

    def get_name(self):
        return self.name

    def get_input_files(self):
        return self.input_files

    def to_json(self):
        return json.dumps(self.__dict__, ensure_ascii=False, sort_keys=True)


def get_api_version_tag(resource_provider_name, readme_contents, overrides):
    override = overrides.get(resource_provider_name)
    if override != None:
        return override

    tag_regex = re.compile("openapi-type: [a-z\\-]+\ntag: ([a-z\\-0-9]*)")
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


def get_api_version_from_readme(resource_provider_name, readme_path, version_overrides):
    if not os.path.isfile(readme_path):
        print("==> file not found: " + readme_path)
        return None
    with open(readme_path, "r", encoding="utf8") as stream:
        contents = stream.read()

    version_tag = get_api_version_tag(resource_provider_name, contents, version_overrides)
    if version_tag == None:
        print("==> no version tag found in readme: " + readme_path)
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


class ApiSet:
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
            + "/" + api_set.get_base_folder()
        )
        resource_provider_target = (
            target_folder
            + "/" + api_set.get_base_folder()
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

        # Copy the files defined in the api version
        for file in api_version.get_input_files():
            copy_file_ensure_paths(resource_provider_source, resource_provider_target, file)

        # Write api-set.json file per folder with contents to load for swagger-codegen
        api_set_filename = resource_provider_target + "/api-set.json"
        print("Writing " + api_set_filename)
        with open(api_set_filename, "w") as f:
            f.write(api_version.to_json())

def get_api_set_for_folder(spec_folder, api_folder, resource_provider_name, version_overrides):
    api_version = get_api_version_from_readme(resource_provider_name, api_folder + "/readme.md", version_overrides)
    if api_version == None:
        return None
    spec_relative_folder = api_folder[len(spec_folder) + 1 :]
    api_set = ApiSet(
        resource_provider_name,
        spec_relative_folder, 
        api_version
    )
    return api_set


def get_resource_manager_api_sets(spec_folder, version_overrides):
    rp_folders = sorted([f.path for f in os.scandir(spec_folder) if f.is_dir()])
    api_sets = []
    for folder in rp_folders:
        resource_provider_name = folder.split("/")[-1]
        got_api_set = False
        
        for api_type_folder in ["resource-manager", "data-plane"]:
            qualified_api_type_folder = folder + "/" + api_type_folder
            if not os.path.exists(qualified_api_type_folder):
                continue
            
            api_set = get_api_set_for_folder(spec_folder, qualified_api_type_folder, resource_provider_name, version_overrides)
            if api_set != None:
                api_sets.append(api_set)
                got_api_set = True
            else:
                # didn't find readme.md under (e.g.) search/data-plane/
                # look for search/data-plane/*/readme.md
                print("\n*************************************************************************************")
                print(qualified_api_type_folder)
                # sub_folders = [f.path for f in os.scandir(qualified_api_type_folder) if f.is_dir() and os.path.exists(qualified_api_type_folder + "/" + f.path + "/readme.md")]
                sub_folders = [f.path for f in os.scandir(qualified_api_type_folder) if f.is_dir()]
                for sub_folder in sub_folders:
                    print(sub_folder)
                    api_set = get_api_set_for_folder(spec_folder, sub_folder, resource_provider_name, version_overrides)
                    if api_set != None:
                        print("got api_set")
                        api_sets.append(api_set)
                        got_api_set = True


        if not got_api_set:
            print("***No api version found, ignoring: " + folder)

    return api_sets


def copy_arm_specs():
    # resource_provider_version_overrides is keyed on RP name with the value being the tag to force
    resource_provider_version_overrides = {
        "cosmos-db": "package-2019-08-preview",  # the 2019-12 version includes 2019-08-preview files that reference files not in the 2019-12 list!
        "frontdoor": "package-2019-10", # same issue as cosmos-db with spec contents referencing files not listed in the package
    }

    print(
        "\n****************************************************************************"
    )
    print("Discovering resource manager api sets:")
    api_sets = get_resource_manager_api_sets(
        "./swagger-temp/azure-rest-api-specs/specification", 
        resource_provider_version_overrides
    )

    print(
        "\n****************************************************************************"
    )
    print("Copying resource-manager files:")
    copy_api_sets_to_swagger_specs(
        api_sets,
        "./swagger-temp/azure-rest-api-specs/specification",
        "./swagger-specs",
    )
    shutil.copytree(
        "./swagger-temp/azure-rest-api-specs/specification/common-types",
        "./swagger-specs/common-types",
    )

if __name__ == "__main__":
    # print(
    #     "\n****************************************************************************"
    # )
    # print("Cloning azure-rest-api-sets repo")
    # clone_swagger_specs("swagger-temp")

    print(
        "\n****************************************************************************"
    )
    print("Deleting ")
    if os.path.exists("swagger-specs"):
        print("Deleting swagger-specs...")
        shutil.rmtree("swagger-specs")

    copy_arm_specs()


