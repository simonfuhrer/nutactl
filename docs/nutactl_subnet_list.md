## nutactl subnet list

List subnets

### Synopsis

List subnets

```
nutactl subnet list [FLAGS]
```

### Options

```
  -f, --filter string   FIQL filter  (e.g. vlan_id==2711;cluster_name==mycluster)
  -h, --help            help for list
  -o, --output string   json|yaml|table (default "table")
```

### Options inherited from parent commands

```
      --config string      config file (default is $HOME/.nutactl.yaml)
      --insecure           Accept insecure TLS certificates
      --log-json           log as json
      --log-level string   log level (trace,debug,info,warn/warning,error,fatal,panic) (default "info")
```

### SEE ALSO

* [nutactl subnet](nutactl_subnet.md)	 - Manage subnets

###### Auto generated by spf13/cobra on 15-Oct-2020