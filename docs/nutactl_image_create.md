## nutactl image create

Create an image

```
nutactl image create [FLAGS] imagename
```

### Options

```
  -d, --description string   Description
  -h, --help                 help for create
  -f, --source-path string   source image path
  -s, --source-uri string    source image URI
  -t, --type string          image type (iso or disk (default "disk")
  -w, --wait                 wait to be completed (default true)
```

### Options inherited from parent commands

```
      --config string      config file to use (default $HOME/.nutactl.yaml)
      --insecure           Accept insecure TLS certificates
      --log-json           log as json
      --log-level string   log level (trace,debug,info,warn/warning,error,fatal,panic) (default "info")
```

### SEE ALSO

* [nutactl image](nutactl_image.md)	 - Manage images

