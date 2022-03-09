## nutactl routing-policy create

Create a routing policy

```
nutactl routing-policy create [FLAGS] routing policy
```

### Options

```
      --action string          PERMIT or DENY (default "PERMIT")
      --destination string     ALL, INTERNET or CIDR
  -h, --help                   help for create
      --isbidirectional        Additionally Create Policy in reverse direction
      --priority int32         priority of rule (between 10-1000
      --protocol-type string   any of 'ALL', 'TCP', 'UDP', 'ICMP', 'PROTOCOL_NUMBER'  (default "ALL")
      --source string          ALL, INTERNET or CIDR
      --vpc string             vpc uuid or name
```

### Options inherited from parent commands

```
      --config string      config file to use (default $HOME/.nutactl.yaml)
      --insecure           Accept insecure TLS certificates
      --log-json           log as json
      --log-level string   log level (trace,debug,info,warn/warning,error,fatal,panic) (default "info")
```

### SEE ALSO

* [nutactl routing-policy](nutactl_routing-policy.md)	 - Manage routing policies

