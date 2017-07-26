/*

Copyright 2017 Travis Clarke. All rights reserved.
Use of this source code is governed by a Apache-2.0
license that can be found in the LICENSE file.

NAME:
	systemstat – display system information.

SYNOPSIS:
	systemstat [ opts... ]

OPTIONS:
	-h, --help		# print usage.
	-a, --all		# same as -c, -d, -m, -n, -p.
	-c, --cpu		# print CPU info.
	-d, --disk		# print Disk info.
	-m, --mem		# print Memory info.
	-n, --net		# print Network info.
	-p, --proc		# print Process info.
	-v, --version		# print version number.

EXAMPLES:
	systemstat -a		# list all system info.

*/

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

// VERSION - current version number
const VERSION = "v1.0.0"

// allFlag bool
type allFlag bool

func (a *allFlag) IsBoolFlag() bool {
	return true
}

func (a *allFlag) String() string {
	return "false"
}

func (a *allFlag) Set(value string) error {
	cpuStat = true
	diskStat = true
	memStat = true
	netStat = true
	procStat = true
	return nil
}

// versionFlag bool
type versionFlag bool

func (v *versionFlag) IsBoolFlag() bool {
	return true
}

func (v *versionFlag) String() string {
	return "false"
}

func (v *versionFlag) Set(value string) error {
	fmt.Printf("\n%s %v\n", bold("Version:"), VERSION)
	os.Exit(0)
	return nil
}

// Flags
var all allFlag
var version versionFlag

var cpuStat bool
var diskStat bool
var memStat bool
var netStat bool
var procStat bool

// Globals
var statusCode int
var bold = color.New(color.Bold).SprintFunc()

// init () - initialize command-line flags
func init() {
	const (
		usageAll        = "Same as --ethernet, --public."
		usageVersion    = "Print version"
		defaultEthernet = false
		usageEthernet   = "Print ethernet IP address."
		defaultPublic   = false
		usagePublic     = "Print public IP address."
	)
	// -a, --all
	flag.Var(&all, "a", "")
	flag.Var(&all, "all", usageAll)

	// -c, --cpu
	flag.BoolVar(&cpuStat, "c", false, "")
	flag.BoolVar(&cpuStat, "cpu", false, "CPU stats")

	// -d, --disk
	flag.BoolVar(&diskStat, "d", false, "")
	flag.BoolVar(&diskStat, "disk", false, "Disk stats")

	// -m, --mem
	flag.BoolVar(&memStat, "m", false, "")
	flag.BoolVar(&memStat, "mem", false, "Memory stats")

	// -n, --net
	flag.BoolVar(&netStat, "n", false, "")
	flag.BoolVar(&netStat, "net", false, "Network stats")

	// -p, --proc
	flag.BoolVar(&procStat, "p", false, "")
	flag.BoolVar(&procStat, "proc", false, "Process stats")

	// -v, --version
	flag.Var(&version, "v", "")
	flag.Var(&version, "version", usageVersion)

	// Usage
	flag.Usage = func() {
		println()
		fmt.Fprintf(os.Stdout, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
		println()
		os.Exit(statusCode)
	}
}

// main ()
func main() {
	flag.Parse()

	if !cpuStat && !diskStat && !memStat && !netStat && !procStat {
		statusCode = 0
		flag.Usage()
	} else {
		println()
		if cpuStat {
			fmt.Printf("\n%s\n%v\n", bold("CPU:"), getCPUStat())

		}
		if diskStat {
			fmt.Printf("\n%s\n%v\n", bold("Disk:"), getDiskStat())

		}
		if memStat {
			fmt.Printf("\n%s\n%v\n", bold("Mem:"), getMemStat())

		}
		if netStat {
			fmt.Printf("\n%s\n%v\n", bold("Net:"), getNetStat())

		}
		if procStat {
			fmt.Printf("\n%s\n%v\n", bold("Proc:"), getProcStat())

		}
		println()
	}
}

func checkError(err error, message string) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "There was an error retreiving", message, ":", err)
		os.Exit(1)
	}
}

// getCPUStat () string - cpu stats
func getCPUStat() string {
	// TODO: TimesStat{} // https://github.com/shirou/gopsutil/blob/master/cpu/cpu.go#L15
	cpus, err := cpu.Info() // InfoStat{} // https://github.com/shirou/gopsutil/blob/master/cpu/cpu.go#L30
	checkError(err, "cpu stat")

	if len(cpus) == 0 {
		return "no cpu was found"
	}

	cpuStats := []string{}
	for _, cpu := range cpus {
		cpuStats = append(cpuStats, fmt.Sprintf("ModelName: %v\nCores: %v\n", cpu.ModelName, cpu.Cores))
	}
	cpuStat := strings.Join(cpuStats, "\n")

	return cpuStat
}

// getDiskStat () string - disk stats
func getDiskStat() string {
	// TODO: UsageStat{} // https://github.com/shirou/gopsutil/blob/master/disk/disk.go#L15
	// TODO: IOCountersStat{} // https://github.com/shirou/gopsutil/blob/master/disk/disk.go#L35
	disks, err := disk.Partitions(true) // PartitionStat{} // https://github.com/shirou/gopsutil/blob/master/disk/disk.go#L28
	checkError(err, "disk stat")

	if len(disks) == 0 {
		return "no disk was found"
	}

	diskStats := []string{}
	for _, disk := range disks {
		if disk.Fstype != "autofs" {
			diskStats = append(diskStats, fmt.Sprintf("Device: %v\nMountpoint: %v\n", disk.Device, disk.Mountpoint))
		}
	}
	diskStat := strings.Join(diskStats, "\n")

	return diskStat
}

// TODO: getDockerStat () string - docker stats
// func getDockerStat() string { }

// getHostStat () string - host stats
func getHostStat() string {
	// TODO: UserStat{} // https://github.com/shirou/gopsutil/blob/master/host/host.go#L32
	// TODO: TemperatureStat{} // https://github.com/shirou/gopsutil/blob/master/host/host.go#L39
	host, err := host.Info() // InfoStat{} // https://github.com/shirou/gopsutil/blob/master/host/host.go#L17
	checkError(err, "host stat")

	hostStat := fmt.Sprintf("OS: %v\nHostname: %v\nUptime: %v\nBootTime: %v", host.OS, host.Hostname, host.Uptime, host.BootTime)

	return hostStat
}

// TODO: getLoadStat () string - load stats
// func getDockerStat() string {
//// TODO: AvgStat{} // https://github.com/shirou/gopsutil/blob/master/load/load.go#L15
//// TODO: MiscStat{} // https://github.com/shirou/gopsutil/blob/master/load/load.go#L26
// }

// getMemStat () string - mem stats
func getMemStat() string {
	// TODO: SwapMemoryStat{} // https://github.com/shirou/gopsutil/blob/master/mem/mem.go#L63
	mem, err := mem.VirtualMemory() // VirtualMemoryStat{}
	checkError(err, "mem stat")

	memStat := fmt.Sprintf("Total: %v\nFree: %v\nUsedPercent: %f%%\n", mem.Total, mem.Free, mem.UsedPercent)

	return memStat
}

// getNetStat () string - ne stats
func getNetStat() string {
	nets, err := net.IOCounters(true) // IOCountersStat{}
	checkError(err, "net stat")

	if len(nets) == 0 {
		return "no net was found"
	}

	netStats := []string{}
	for _, net := range nets {
		if net.PacketsSent > 0 && net.PacketsRecv > 0 {
			netStats = append(netStats, fmt.Sprintf("Interface: %v\nPacketsSent: %v\nPacketsRecv: %v\n", net.Name, net.PacketsSent, net.PacketsRecv))
		}
	}
	netStat := strings.Join(netStats, "\n")

	return netStat
}

func getProcStat() string {
	pids, err := process.Pids() // Connections{}
	checkError(err, "pid stat")

	if len(pids) == 0 {
		return "no pid was found"
	}

	var procStat string
	for _, pid := range pids {
		if pid == 1 {
			proc, err := process.NewProcess(pid)
			checkError(err, "proc stat pid:1")
			name, err := proc.Name()
			checkError(err, "proc stat name")
			mem, err := proc.MemoryInfo()
			checkError(err, "proc stat mem info")
			times, err := proc.Times()
			checkError(err, "proc stat times")
			children, err := proc.Children()
			checkError(err, "proc stat children")
			procStat = fmt.Sprintf("Init(pid=1):\n"+
				" Name: %v\n"+
				" ResidentSetSize: %v\n"+ // how much RAM the process is using.
				" VirtualMemorySize: %v\n"+ // how much memory a process has available for its execution.
				" SwapSize: %v\n"+ // how much Swap the process is using.
				" UserTime: %v\n"+
				" SystemTime: %v\n"+
				" Children: %v\n", name, mem.RSS, mem.VMS, mem.Swap, times.User, times.System, len(children))
			break
		}
	}

	return procStat
}
