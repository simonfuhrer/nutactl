## nutactl vm list

List all VM

```
nutactl vm list [FLAGS]
```

### Options

```
  -c, --cluster string   filter vms by cluster
  -f, --filter string    FIQL filter (e.g. vm_name==srv.*, ip_addresses==192.168.10.59, power_state==off)
  -h, --help             help for list
  -o, --output string    json|yaml|table (default "table")
```

### Options inherited from parent commands

```
      --config string      config file to use (default $HOME/.nutactl.yaml)
      --insecure           Accept insecure TLS certificates
      --log-json           log as json
      --log-level string   log level (trace,debug,info,warn/warning,error,fatal,panic) (default "info")
```

### SEE ALSO

* [nutactl vm](nutactl_vm.md)	 - Manage vms

