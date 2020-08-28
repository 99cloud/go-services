import os
import shutil
import sys


root_path = os.path.abspath('.')

dir_name = input("Please Enter the Dir Name [Output]: ") or "Output"
dir_path = os.path.join(root_path, dir_name)

if not os.path.exists(dir_path):
    sys.exit("The Dir '%s' is NOT Under the Path: %s" % (dir_name, root_path))

app_name = input("Please Enter the App Name: ")
app_judge = app_name.lower().strip().replace(' ', '_').replace('-', '_').replace('.', '_')

app_list = [
    "build",
    "cmd",
    "internal"
]

for idx in range(len(app_list)):
    app_test_path = os.path.join(dir_path, app_list[idx], app_judge)
    if os.path.exists(app_test_path):
        shutil.rmtree(app_test_path)
        print("The App '%s' Under '%s' is DELETED!" % (
            app_test_path.rsplit('/', 2)[-1], app_test_path.rsplit('/', 2)[-2]
        ))
    else:
        sys.exit("ERROR: App Name '%s' is NOT Match!" % app_judge)
