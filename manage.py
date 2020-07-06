import logging
import os
import re
import shutil
import subprocess
import sys


BASE_DIR = os.path.dirname(os.path.abspath(__file__))
logging.basicConfig(level=logging.INFO)
log = logging.getLogger('manage.py')

UUID = '46ea591951824d8e9376b0f98fe4d48a'
PROJECT_NAME = 'PROJECT_' + UUID
APP_NAME = 'APP_' + UUID
APP_UPPER_NAME = 'APP_UPPER_' + UUID
APP_LOWER_NAME = 'app_lower_' + UUID

def showUsage():
    print('''Usage:
    python manage.py <option>
    python manage.py startproject <project-name> <default-app-name>
    python manage.py startapp <project-name> <app-name>
    ''')
    sys.exit()


def sed(old, new, filePath):
    ignoreRegex = re.compile(r'\.((db)|(png)|(js.map))$')
    if ignoreRegex.search(filePath):
        return
    try:
        lines = [i.replace(old, new) for i in open(filePath) if not ignoreRegex.search(filePath)]
        open(filePath, 'w').writelines(lines)
    except UnicodeDecodeError as e:
        log.warning('old = {}, new = {}, filePath = {}'.format(old, new, filePath))
        log.warning(e)


def mv(old, new, filePath):
    if old in filePath:
        cmdStr = r'mv {} {}'.format(filePath, filePath.replace(old, new))
        log.debug(cmdStr)
        os.system(cmdStr)


def opt_startproject(projectName, appName):
    os.system(r'rm -rf {} && cp -r {} {} && rm -rf {} {}'.format(
        os.path.join(BASE_DIR, 'output'),
        os.path.join(BASE_DIR, 'template'),
        os.path.join(BASE_DIR, 'output'),
        os.path.join(BASE_DIR, 'output', 'dist'),
        os.path.join(BASE_DIR, 'output', 'doc')
        ))
    for root, dirs, files in os.walk(os.path.join(BASE_DIR, 'output')):
        for name in dirs:
            absPath = os.path.join(root, name)
            mv(APP_NAME, appName, absPath)
    for root, dirs, files in os.walk(os.path.join(BASE_DIR, 'output')):
        for name in files:
            absPath = os.path.join(root, name)
            sed(PROJECT_NAME, projectName, absPath)
            sed(APP_LOWER_NAME, appName.lower(), absPath)
            sed(APP_NAME, appName, absPath)
            sed(APP_UPPER_NAME, appName.upper(), absPath)
            mv(APP_NAME, appName, absPath)


def opt_startapp(projectName, appName):
    log.info("startapp: {} {}".format(projectName, appName))
    raise NotImplementedError()


def _assert_cmd_exist(cmd):
    try:
        subprocess.call(cmd)
    except Exception as e:
        log.warning('{}->{}'.format(type(e), e))
        log.error('Command "{}" not exist!'.format(cmd))
        sys.exit()


if __name__ == '__main__':
    if len(sys.argv) < 2:
        showUsage()

    selfModule = __import__(__name__)
    optFunName = 'opt_' + sys.argv[1].strip()
    if optFunName not in selfModule.__dict__:
        showUsage()

    if BASE_DIR.strip():
        os.chdir(BASE_DIR)
    try:
        selfModule.__dict__[optFunName](*sys.argv[2:])
    except TypeError as e:
        log.error('{} failed: {}'.format(optFunName, e))
        raise
