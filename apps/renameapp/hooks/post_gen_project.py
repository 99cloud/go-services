import os
import shutil
import sys


def rename_files(url, old_name, new_name):
    files = os.listdir(url)
    for f in files:
        real_path = os.path.join(url, f)
        if os.path.isfile(real_path):
            # print(os.path.abspath(real_path))
            file_path = os.path.abspath(real_path)
            temp_path = file_path + '.temp'
            with open(file_path, mode='r') as fr, open(temp_path, mode='w') as fw:
                for line in fr:
                    fw.write(line.replace(old_name, new_name))
            os.remove(file_path)
            os.rename(temp_path, file_path)

        elif os.path.isdir(real_path):
            rename_files(real_path, old_name, new_name)
        else:
            print("Another Situation!")
            pass
    return 1


root_path = os.path.abspath('..')
dir_name = "{{cookiecutter.directory_name}}"
dir_path = os.path.join(root_path, dir_name)
# print(dir_path)

rename_list = [
    "build",
    "cmd",
    "internal"
]

old_app_name = "{{ cookiecutter.old_app_slug.lower()|replace(' ', '_')|replace('-', '_')|replace('.', '_')|trim() }}"
new_app_name = "{{ cookiecutter.new_app_slug.lower()|replace(' ', '_')|replace('-', '_')|replace('.', '_')|trim() }}"

for idx in range(len(rename_list)):
    app_test_path = os.path.join(dir_path, rename_list[idx], old_app_name)
    rename_files(app_test_path, old_app_name, new_app_name)
    os.rename(app_test_path, os.path.join(dir_path, rename_list[idx], new_app_name))
