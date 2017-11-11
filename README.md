# [systemstat](https://godoc.org/github.com/clarketm/systemstat)

Command line utility for displaying process and system information.

```shell
NAME:
    systemstat – display system information.

SYNOPSIS:
    systemstat [ opts... ]

OPTIONS:
    -h, --help          # print usage.
    -a, --all           # same as -c, -d, -m, -n, -p.
    -c, --cpu           # print cpu info.
    -d, --disk          # print disk info.
    -m, --mem           # print memory info.
    -n, --net           # print network info.
    -p, --proc          # print process info.
    -v, --version       # print version number.

EXAMPLES:
    systemstat -a       # list all system info.
```

### Credits
This tools wouldn't be possible without the exceptional process and system monitoring library [gopsutils](https://github.com/shirou/gopsutil) by [shirou](https://github.com/shirou).
