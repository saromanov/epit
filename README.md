# epit

Manage of configurations

## Usage

* Create file `.epit.yml` at the root of your project

* Define stages like this

```yml
dev:
    steps:
        - name: first
          command: w
          duration: true
        - name: build
          command: date

prod:
    env:
        - TEST_APP=foobar
    command: "ls -la"
```

* Run stage which you want to execute `epit dev` 