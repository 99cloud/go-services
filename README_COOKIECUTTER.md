# Using Cookiecutter to Create the Go-Services

## Install the Cookiecutter

```bash
pip install -U pip
pip install cookiecutter
```

## About Go-Services

Firstly, we need to switch to the repository's root directory.

```console
cd [root_path]/go-services
```

### Setup a Project

```bash
cookiecutter ./projects
```

### Add an App

```bash
cookiecutter ./apps/startapp -f
```

### Rename an App

```bash
cookiecutter ./apps/renameapp -f
```

## Test

- We need to switch to the project's directory.

```console
# For Example
cd Output
```

- Run the Go-Services

```console
# For Example
APP_NAME="app_lisa"
go run cmd/$APP_NAME/main.go
```

- Curl Testcase

```bash
curl --location --request POST 'http://127.0.0.1:8080/liveness/v1/echo' \
--header 'Content-Type: application/json' \
--data-raw '
{
    "Message": "OK!"
}
'
```

- Curl Result

```console
{
 "Message": "OK!"
}
```

## Other Information

### About Conflict Code

Since the Cookiecutter uses `{{ cookiecutter.[variable_name] }}` for the variable, this may conflict with some coding styles of the original file (eg. `.go` and `.map` ) and can be resolved using the following approach.

```console
{% raw %}
... ambiguous code segments ...
{% endraw %}
```

### About Rename an App

When we use the following statement to determine the files that need to be modified when renaming the app, we find that the vast majority of the files are overwritten, so we use the method of traversing through all the files and modifying them to complete the task of renaming the app

```bash
grep "{{cookiecutter.app_slug}}" -r apps/startapp
```

### Delete an App

When you need to delete an app during testing

```console
cd [root_path]/go-services
python apps/delete_app.py
```
