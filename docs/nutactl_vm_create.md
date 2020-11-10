## nutactl vm create

Create a VM

```
nutactl vm create [FLAGS] vmname
```

### Options

```
  -c, --cluster string       Cluster Name or UUID)
  -d, --description string   Description
  -h, --help                 help for create
      --image string         Image Name or UUID
      --memoryMB int         Memory in MB
      --numCores int         Number of Cores (default 1)
      --numSockets int       Number of CPU Sockets (default 1)
      --project string       Project Name or UUID
      --root-disk-size int   Root Disk Size in MB
      --start-after-create   Start VM right after creation
      --subnet string        Subnet Name, VLAN ID or UUID
      --user-data string     Read user data from specified file
      --vm string            VM Name or UUID
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

