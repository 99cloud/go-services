import sys

proj_judge = "{{ cookiecutter.project_slug.lower()|replace(' ', '_')|replace('-', '_')|replace('.', '_')|trim() }}"
app_judge = "{{ cookiecutter.app_slug.lower()|replace(' ', '_')|replace('-', '_')|replace('.', '_')|trim() }}"

if proj_judge != "{{ cookiecutter.project_slug }}":
    sys.exit("ERROR: The Project Name is NOT Allow!")

elif app_judge != "{{ cookiecutter.app_slug }}":
    sys.exit("ERROR: The APP Name is NOT Allow!")
