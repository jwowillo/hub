# v1.0.0 Design

# TODO: Make complete sentences.

* A `go` command with no arguments will serve a single home page with its
  template embedded in the binary.
    - `go get` will install the command. (5)
    - The command will then be able to be immediately run with no arguments. (5)
* For each request, the server opens a config file located at a fixed location,
  reads the directory from it, injects the directory into the embedded template,
  and serves the HTML page. (1, 4)
    - The config file will be a YAML file with the structure in 'Config'. (3)
    - The directory will be injected into a list which shows each websites name
      with a link to the website. (2)

## Config

```
- URL: <URL>
  name: <NAME>
- ...
```

# v1.1.0 Design

* A `LoadFavicon` function will accept a `Website`, make a request to the
  URL, and parse an image URL at a tag with `rel="icon"` if it exists into the
  `Website`. This will be called for each `Website` in `Handler`. The image will
  be injected to the right of each website link. (1)
* A shell copy config script will use `scp` to copy the config file to a passed
  remote host and directory. (2)
* A shell deploy script will use `ssh` to log into a passed remote host and
  start `hub` with `nohup` in a passed working directory. (3)
