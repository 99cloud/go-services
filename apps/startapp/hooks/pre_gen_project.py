import os
import shutil
import sys


root_path = os.path.abspath('..')

dir_judge = "{{cookiecutter.directory_name}}"
dir_test_path = os.path.join(root_path, dir_judge, "README.md")

if not os.path.exists(dir_test_path):
    sys.exit("ERROR: Dir Name '%s' is NOT Match!" % dir_judge)

proj_judge = "{{ cookiecutter.project_slug.lower()|replace(' ', '_')|replace('-', '_')|replace('.', '_')|trim() }}"
app_judge = "{{ cookiecutter.app_slug.lower()|replace(' ', '_')|replace('-', '_')|replace('.', '_')|trim() }}"

if proj_judge != "{{ cookiecutter.project_slug }}":
    sys.exit("ERROR: The Project Name is NOT Allow!")

elif app_judge != "{{ cookiecutter.app_slug }}":
    sys.exit("ERROR: The APP Name is NOT Allow!")
