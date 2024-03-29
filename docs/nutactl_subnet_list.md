## nutactl subnet list

List subnets

```
nutactl subnet list [FLAGS]
```

### Options

```
  -e, --external        show external subnets
  -f, --filter string   FIQL filter  (e.g. vlan_id==2711;cluster_name==mycluster, is_external==true, subnet_type==OVERLAY)
  -h, --help            help for list
  -o, --output string   json|yaml|table (default "table")
  -s, --overlay         show internal overlay subnets
```

### Options inherited from parent commands

```
      --config string      config file to use (default $HOME/.nutactl.yaml)
      --insecure           Accept insecure TLS certificates
      --log-json           log as json
      --log-level string   log level (trace,debug,info,warn/warning,error,fatal,panic) (default "info")
```

### SEE ALSO

* [nutactl subnet](nutactl_subnet.md)	 - Manage subnets

