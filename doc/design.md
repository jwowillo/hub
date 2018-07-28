# Design

* A Go command with no arguments will serve a single home page with its template
  embedded in the binary.
    - `go get` will install this command. (5)
    - The command can then be immediately run with no arguments. (5)
* For each request, the server will open a config file located at a fixed
  location, read the directory from it, inject the directory into the embedded
  template, and serve the HTML page. (1, 4)
    - The config file will be a YAML file with the structure in 'Config'. (3)
    - The URL for each item will be used to show an iframe corresponding to the
      websites thumbnail. The name for each item will be shown under the
      thumbnail. (2)

## Config

```
directory:
  - URL: <URL>
    name: <NAME>
  - ...
```
