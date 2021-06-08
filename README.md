# je

[![CircleCI](https://circleci.com/gh/wooos/je.svg?style=shield)](https://circleci.com/gh/woos/je)

Command line editor for json.

## Usage

```
Command line editor for json.

Usage
  je [command]

Available Commands:
  completion  Generate auto completions script for goe for the specified shell (bash)
  delete      Delete the keys of json
  help        Help about any command
  update      update a json file
  version     print current version information

Flags:
  -h, --help   help for je

  Use "je [command] --help" for more information about a command.
```

## Example

The below is content of `source.json`.

```json
{
  "name": "je",
  "version": "v1.0",
  "detail": {
    "n": "je",
    "v": "v1.0"
  },
  "users": [
    "a",
    "b"
  ]
}

```

1. modify `name` value `je` to `jee`

```
$ je update --set name=jee source.json
```

2. modify `detail.n` value `je` to `jee`

```
$ je update --set detail.n=jee source.json
```

3. modify `users` values `a` to `aa`

```
$ je update --set users[0]=aa source.json
```

4. modify `name` value `je` to `jee` and modify `n` value `je` to `jee`

```
$ je update --set name=jee,detail.n=jee source.json
```

5. delete key `name`

```
$ je delete --keys name source.json
```

6. delete key `detail.n`

```
$ je delete --keys detail.n source.json
```

