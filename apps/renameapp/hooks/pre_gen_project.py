import os
import sys


root_path = os.path.abspath('..')
# root_list = [x.lower() for x in os.listdir(root_path)]

dir_judge = "{{cookiecutter.directory_name}}"
dir_test_path = os.path.join(root_path, dir_judge, "README.md")

if not os.path.exists(dir_test_path):
    sys.exit("ERROR: Dir Name '%s' is NOT Match!" % dir_judge)

old_app_judge = "{{ cookiecutter.old_app_slug.lower()|replace(' ', '_')|replace('-', '_')|replace('.', '_')|trim() }}"

if old_app_judge != "{{ cookiecutter.old_app_slug }}":
    sys.exit("ERROR: The APP Name is NOT Allow!")

new_app_judge = "{{ cookiecutter.new_app_slug.lower()|replace(' ', '_')|replace('-', '_')|replace('.', '_')|trim() }}"

if new_app_judge != "{{ cookiecutter.new_app_slug }}":
    sys.exit("ERROR: The APP Name is NOT Allow!")

app_list = [
    "build",
    "cmd",
    "internal"
]

for idx in range(len(app_list)):

    old_app_path = os.path.join(root_path, dir_judge, app_list[idx], old_app_judge)
    if not os.path.exists(old_app_path):
        sys.exit("ERROR: App Name '%s' is NOT Match!" % old_app_judge)

    new_app_path = os.path.join(root_path, dir_judge, app_list[idx], new_app_judge)
    if os.path.exists(new_app_path):
        sys.exit("ERROR: App Name '%s' is Already Have!" % new_app_judge)
