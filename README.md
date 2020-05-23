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


## Running on Docker

```sh
docker run -e STAGE=dev -v $(pwd):/app -w /app motorcode/epit
```

where variable STAGE shoild contains stage from your `.epit.yml` file

You can create own Dockerfile with base `motorcode/epit` image