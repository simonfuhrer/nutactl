## nutactl vm update

Update a VM

```
nutactl vm update [FLAGS] VM
```

### Options

```
  -d, --description string   Description
  -h, --help                 help for update
      --memoryMB int         Memory in MB
  -n, --name string          New VM name
      --numCores int         Number of Cores (default 1)
      --numSockets int       Number of CPU Sockets (default 1)
      --project string       Project Name or UUID
      --root-disk-size int   Root Disk Size in MB
      --subnet string        Subnet Name, VLAN ID or UUID
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

