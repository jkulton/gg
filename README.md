# gg

Go gadget is a box of CLI utilities.

## Commands

| Command | Description |
|------|-------------|
| `downcase` | Returns a downcase representation of whatever is passed in args |
| `rot13` | Returns the ROT13 representation of whatever is passed in args |
| `sha256` | Returns hexidecimal representation of SHA256 for whatever is passed in args |
| `upcase` | Returns an upcase representation of whatever  is passed in args|
| `uuid` | Returns a random UUID v4 |
| `version` | Returns the version of gg |

## Examples

```console
$ gg uuid
f1781c30-3259-4d2f-823f-57d6487f6ca2
```

```console
$ gg downcase THIS should be downCASED
this should be downcased
```

Like most Unix-like commands the output of gg can be redirected elsewhere.

```console
$ gg uuid >> output.txt
$ gg uuid >> output.txt
$ gg uuid >> output.txt
```

```console
$ cat output.txt
2d207514-84d8-424b-8e83-1114a0a75514
b4215934-eb84-44b2-9aff-3176b98e18d1
7498efc0-61e4-4ffc-b85e-0f77d7c589fc
```

