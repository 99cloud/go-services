import logging
import os
import shutil
import subprocess
import sys


BASE_DIR = os.path.dirname(os.path.abspath(__file__))
logging.basicConfig(level=logging.INFO)
log = logging.getLogger('manage.py')

UUID = '46ea591951824d8e9376b0f98fe4d48a'
PROJECT_NAME = 'PROJECT_NAME-' + UUID
APP_NAME = 'APP_NAME-' + UUID

def showUsage():
    print('''Usage:
    python manage.py <option>
    python manage.py startproject <project-name>
    python manage.py startapp <app-name>
    ''')
    sys.exit()


def opt_test():
    os.system(r'cp -r {}')
    for root, dirs, files in os.walk(os.path.join(BASE_DIR, 'template')):
        for name in files:
            os.system()
            print(os.path.join(root, name))
        for name in dirs:
            print(os.path.join(root, name))


def opt_startproject(projectName):
    log.info('startproject: {}'.format(projectName))


def opt_startapp(appName):
    log.info("startapp: {}".format(appName))


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
