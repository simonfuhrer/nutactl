## nutactl subnet update

Update an subnet

```
nutactl subnet update [FLAGS] SUBNET
```

### Options

```
  -d, --description string    Description
      --dns-servers strings   Default DNS Servers seperated with a comma
      --domain string         Default Domainname
      --gateway string        Default Gateway IP
  -h, --help                  help for update
      --ip-pool strings       Start address to end address seperated with a comma
      --ip-range string       Network CIDR
  -n, --name string           New project name
      --type string           VLAN or OVERLAY (default "VLAN")
      --vlan-id int           VlanID
      --vpc string            VPC Name or UUID
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

