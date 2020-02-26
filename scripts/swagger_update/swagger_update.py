from git import Repo
import shutil
import os
import re
import yaml

# last_git_message=""
# def show_git_progress(op_code, cur_count, max_count, message):
#     global last_git_message
#     if (message != "" and message != last_git_message):
#         print(message)
#         last_git_message = message

# if os.path.exists("gittemp"):
#     print("Deleting gittemp...")
#     shutil.rmtree("gittemp")

# print ("Cloning specs...")
# repo = Repo()
# repo.clone_from("git@github.com:azure/azure-rest-api-specs", "gittemp", progress=show_git_progress, multi_options=["--depth=1"])
# print ("Cloned")


# TODO
# - add docs to top to state goals (recreate folder structure for any files that are copied by looking up in readme.md)
# - add function to find latest version (add unit tests for this!)


def foo(i):
    return i + 1


class ApiVersion:
    def __init__(self, name, input_files):
        # super().__init__()
        self.name = name
        self.input_files = input_files

    def get_name(self):
        return self.name

    def get_input_files(self):
        return self.input_files


def get_api_versions_from_readme(readme_path):
    if not os.path.isfile(readme_path):
        return None
    with open(readme_path, "r") as stream:
        contents = stream.read()

    results = []

    # TODO - since this function will be invoked multiple times, move the compilation external to the function?
    yaml_start_regex = re.compile(
        "^```[\\s]*yaml \\$\\(tag\\) == '([a-z\\-0-9]*)'$", flags=re.MULTILINE
    )
    code_block_end_regex = re.compile("^[\\s]*```[\\s]*$", flags=re.MULTILINE)

    search_start = 0
    while True:
        block_start_match = yaml_start_regex.search(contents, search_start)
        if block_start_match == None:
            break
        tag = block_start_match.group(1)
        search_start = block_start_match.end()

        block_end_match = code_block_end_regex.search(contents, search_start)
        if block_end_match == None:
            break

        yaml_contents = contents[search_start : block_end_match.start()]
        search_start = block_end_match.end()

        yaml_data = yaml.load(yaml_contents, Loader=yaml.BaseLoader)
        input_files = []
        if yaml_data != None:
            input_files = yaml_data["input-file"]
        api_version = ApiVersion(tag, input_files)
        results.append(api_version)

    return results

def pick_api_version(api_versions):
    def get_name(api_version):
        return api_version.get_name()

    candidate_versions = [v for v in api_versions if v.get_name() != "all-api-versions" and not v.get_name().endswith("-preview-only")]
    if len(candidate_versions) == 0:
        return None

    sorted_versions = sorted(candidate_versions, key=get_name)
    return sorted_versions[0]

if __name__ == "__main__":
    rp_folders = sorted(
        [f.path for f in os.scandir("gittemp/specification") if f.is_dir()]
    )

    for f in rp_folders:
        readme_path = f"{f}/resource-manager/readme.md"
        api_versions = get_api_versions_from_readme(readme_path)
        if api_versions == None:
            print(f"No readme, ignoring: {f}")
            continue
        api_version = pick_api_version(api_versions)
        if api_version == None:
            print(f"No suitable api version, ignoring: {f}")
            continue

        print(f"TODO: process {f}. Version: {api_version.get_name()}")
