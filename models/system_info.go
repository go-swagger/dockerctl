// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SystemInfo system info
//
// swagger:model SystemInfo
type SystemInfo struct {

	// Hardware architecture of the host, as returned by the Go runtime
	// (`GOARCH`).
	//
	// A full list of possible values can be found in the [Go documentation](https://golang.org/doc/install/source#environment).
	//
	// Example: x86_64
	Architecture string `json:"Architecture,omitempty"`

	// Indicates if `bridge-nf-call-ip6tables` is available on the host.
	// Example: true
	BridgeNfIp6tables bool `json:"BridgeNfIp6tables,omitempty"`

	// Indicates if `bridge-nf-call-iptables` is available on the host.
	// Example: true
	BridgeNfIptables bool `json:"BridgeNfIptables,omitempty"`

	// Indicates if CPUsets (cpuset.cpus, cpuset.mems) are supported by the host.
	//
	// See [cpuset(7)](https://www.kernel.org/doc/Documentation/cgroup-v1/cpusets.txt)
	//
	// Example: true
	CPUSet bool `json:"CPUSet,omitempty"`

	// Indicates if CPU Shares limiting is supported by the host.
	// Example: true
	CPUShares bool `json:"CPUShares,omitempty"`

	// The driver to use for managing cgroups.
	//
	// Example: cgroupfs
	// Enum: [cgroupfs systemd none]
	CgroupDriver *string `json:"CgroupDriver,omitempty"`

	// The network endpoint that the Engine advertises for the purpose of
	// node discovery. ClusterAdvertise is a `host:port` combination on which
	// the daemon is reachable by other hosts.
	//
	// <p><br /></p>
	//
	// > **Note**: This field is only propagated when using standalone Swarm
	// > mode, and overlay networking using an external k/v store. Overlay
	// > networks with Swarm mode enabled use the built-in raft store, and
	// > this field will be empty.
	//
	// Example: node5.corp.example.com:8000
	ClusterAdvertise string `json:"ClusterAdvertise,omitempty"`

	// URL of the distributed storage backend.
	//
	//
	// The storage backend is used for multihost networking (to store
	// network and endpoint information) and by the node discovery mechanism.
	//
	// <p><br /></p>
	//
	// > **Note**: This field is only propagated when using standalone Swarm
	// > mode, and overlay networking using an external k/v store. Overlay
	// > networks with Swarm mode enabled use the built-in raft store, and
	// > this field will be empty.
	//
	// Example: consul://consul.corp.example.com:8600/some/path
	ClusterStore string `json:"ClusterStore,omitempty"`

	// containerd commit
	ContainerdCommit *Commit `json:"ContainerdCommit,omitempty"`

	// Total number of containers on the host.
	// Example: 14
	Containers int64 `json:"Containers,omitempty"`

	// Number of containers with status `"paused"`.
	//
	// Example: 1
	ContainersPaused int64 `json:"ContainersPaused,omitempty"`

	// Number of containers with status `"running"`.
	//
	// Example: 3
	ContainersRunning int64 `json:"ContainersRunning,omitempty"`

	// Number of containers with status `"stopped"`.
	//
	// Example: 10
	ContainersStopped int64 `json:"ContainersStopped,omitempty"`

	// Indicates if CPU CFS(Completely Fair Scheduler) period is supported by the host.
	// Example: true
	CPUCfsPeriod bool `json:"CpuCfsPeriod,omitempty"`

	// Indicates if CPU CFS(Completely Fair Scheduler) quota is supported by the host.
	// Example: true
	CPUCfsQuota bool `json:"CpuCfsQuota,omitempty"`

	// Indicates if the daemon is running in debug-mode / with debug-level logging enabled.
	// Example: true
	Debug bool `json:"Debug,omitempty"`

	// Name of the default OCI runtime that is used when starting containers.
	//
	// The default can be overridden per-container at create time.
	//
	// Example: runc
	DefaultRuntime *string `json:"DefaultRuntime,omitempty"`

	// Root directory of persistent Docker state.
	//
	// Defaults to `/var/lib/docker` on Linux, and `C:\ProgramData\docker`
	// on Windows.
	//
	// Example: /var/lib/docker
	DockerRootDir string `json:"DockerRootDir,omitempty"`

	// Name of the storage driver in use.
	// Example: overlay2
	Driver string `json:"Driver,omitempty"`

	// Information specific to the storage driver, provided as
	// "label" / "value" pairs.
	//
	// This information is provided by the storage driver, and formatted
	// in a way consistent with the output of `docker info` on the command
	// line.
	//
	// <p><br /></p>
	//
	// > **Note**: The information returned in this field, including the
	// > formatting of values and labels, should not be considered stable,
	// > and may change without notice.
	//
	// Example: [["Backing Filesystem","extfs"],["Supports d_type","true"],["Native Overlay Diff","true"]]
	DriverStatus [][]string `json:"DriverStatus"`

	// Indicates if experimental features are enabled on the daemon.
	//
	// Example: true
	ExperimentalBuild bool `json:"ExperimentalBuild,omitempty"`

	// generic resources
	GenericResources GenericResources `json:"GenericResources,omitempty"`

	// HTTP-proxy configured for the daemon. This value is obtained from the
	// [`HTTP_PROXY`](https://www.gnu.org/software/wget/manual/html_node/Proxies.html) environment variable.
	// Credentials ([user info component](https://tools.ietf.org/html/rfc3986#section-3.2.1)) in the proxy URL
	// are masked in the API response.
	//
	// Containers do not automatically inherit this configuration.
	//
	// Example: http://xxxxx:xxxxx@proxy.corp.example.com:8080
	HTTPProxy string `json:"HttpProxy,omitempty"`

	// HTTPS-proxy configured for the daemon. This value is obtained from the
	// [`HTTPS_PROXY`](https://www.gnu.org/software/wget/manual/html_node/Proxies.html) environment variable.
	// Credentials ([user info component](https://tools.ietf.org/html/rfc3986#section-3.2.1)) in the proxy URL
	// are masked in the API response.
	//
	// Containers do not automatically inherit this configuration.
	//
	// Example: https://xxxxx:xxxxx@proxy.corp.example.com:4443
	HTTPSProxy string `json:"HttpsProxy,omitempty"`

	// Unique identifier of the daemon.
	//
	// <p><br /></p>
	//
	// > **Note**: The format of the ID itself is not part of the API, and
	// > should not be considered stable.
	//
	// Example: 7TRN:IPZB:QYBB:VPBQ:UMPP:KARE:6ZNR:XE6T:7EWV:PKF4:ZOJD:TPYS
	ID string `json:"ID,omitempty"`

	// Indicates IPv4 forwarding is enabled.
	// Example: true
	IPV4Forwarding bool `json:"IPv4Forwarding,omitempty"`

	// Total number of images on the host.
	//
	// Both _tagged_ and _untagged_ (dangling) images are counted.
	//
	// Example: 508
	Images int64 `json:"Images,omitempty"`

	// Address / URL of the index server that is used for image search,
	// and as a default for user authentication for Docker Hub and Docker Cloud.
	//
	// Example: https://index.docker.io/v1/
	IndexServerAddress *string `json:"IndexServerAddress,omitempty"`

	// Name and, optional, path of the `docker-init` binary.
	//
	// If the path is omitted, the daemon searches the host's `$PATH` for the
	// binary and uses the first result.
	//
	// Example: docker-init
	InitBinary string `json:"InitBinary,omitempty"`

	// init commit
	InitCommit *Commit `json:"InitCommit,omitempty"`

	// Represents the isolation technology to use as a default for containers.
	// The supported values are platform-specific.
	//
	// If no isolation value is specified on daemon start, on Windows client,
	// the default is `hyperv`, and on Windows server, the default is `process`.
	//
	// This option is currently not used on other platforms.
	//
	// Enum: [default hyperv process]
	Isolation *string `json:"Isolation,omitempty"`

	// Indicates if the host has kernel memory limit support enabled.
	// Example: true
	KernelMemory bool `json:"KernelMemory,omitempty"`

	// Kernel version of the host.
	//
	// On Linux, this information obtained from `uname`. On Windows this
	// information is queried from the <kbd>HKEY_LOCAL_MACHINE\\SOFTWARE\\Microsoft\\Windows NT\\CurrentVersion\\</kbd>
	// registry value, for example _"10.0 14393 (14393.1198.amd64fre.rs1_release_sec.170427-1353)"_.
	//
	// Example: 4.9.38-moby
	KernelVersion string `json:"KernelVersion,omitempty"`

	// User-defined labels (key/value metadata) as set on the daemon.
	//
	// <p><br /></p>
	//
	// > **Note**: When part of a Swarm, nodes can both have _daemon_ labels,
	// > set through the daemon configuration, and _node_ labels, set from a
	// > manager node in the Swarm. Node labels are not included in this
	// > field. Node labels can be retrieved using the `/nodes/(id)` endpoint
	// > on a manager node in the Swarm.
	//
	// Example: ["storage=ssd","production"]
	Labels []string `json:"Labels"`

	// Indicates if live restore is enabled.
	//
	// If enabled, containers are kept running when the daemon is shutdown
	// or upon daemon start if running containers are detected.
	//
	// Example: false
	LiveRestoreEnabled *bool `json:"LiveRestoreEnabled,omitempty"`

	// The logging driver to use as a default for new containers.
	//
	LoggingDriver string `json:"LoggingDriver,omitempty"`

	// Total amount of physical memory available on the host, in kilobytes (kB).
	//
	// Example: 2095882240
	MemTotal int64 `json:"MemTotal,omitempty"`

	// Indicates if the host has memory limit support enabled.
	// Example: true
	MemoryLimit bool `json:"MemoryLimit,omitempty"`

	// The number of logical CPUs usable by the daemon.
	//
	// The number of available CPUs is checked by querying the operating
	// system when the daemon starts. Changes to operating system CPU
	// allocation after the daemon is started are not reflected.
	//
	// Example: 4
	NCPU int64 `json:"NCPU,omitempty"`

	// Number of event listeners subscribed.
	// Example: 30
	NEventsListener int64 `json:"NEventsListener,omitempty"`

	// The total number of file Descriptors in use by the daemon process.
	//
	// This information is only returned if debug-mode is enabled.
	//
	// Example: 64
	NFd int64 `json:"NFd,omitempty"`

	// The  number of goroutines that currently exist.
	//
	// This information is only returned if debug-mode is enabled.
	//
	// Example: 174
	NGoroutines int64 `json:"NGoroutines,omitempty"`

	// Hostname of the host.
	// Example: node5.corp.example.com
	Name string `json:"Name,omitempty"`

	// Comma-separated list of domain extensions for which no proxy should be
	// used. This value is obtained from the [`NO_PROXY`](https://www.gnu.org/software/wget/manual/html_node/Proxies.html)
	// environment variable.
	//
	// Containers do not automatically inherit this configuration.
	//
	// Example: *.local, 169.254/16
	NoProxy string `json:"NoProxy,omitempty"`

	// Generic type of the operating system of the host, as returned by the
	// Go runtime (`GOOS`).
	//
	// Currently returned values are "linux" and "windows". A full list of
	// possible values can be found in the [Go documentation](https://golang.org/doc/install/source#environment).
	//
	// Example: linux
	OSType string `json:"OSType,omitempty"`

	// Version of the host's operating system
	//
	// <p><br /></p>
	//
	// > **Note**: The information returned in this field, including its
	// > very existence, and the formatting of values, should not be considered
	// > stable, and may change without notice.
	//
	// Example: 16.04
	OSVersion string `json:"OSVersion,omitempty"`

	// Indicates if OOM killer disable is supported on the host.
	OomKillDisable bool `json:"OomKillDisable,omitempty"`

	// Name of the host's operating system, for example: "Ubuntu 16.04.2 LTS"
	// or "Windows Server 2016 Datacenter"
	//
	// Example: Alpine Linux v3.5
	OperatingSystem string `json:"OperatingSystem,omitempty"`

	// Indicates if the host kernel has PID limit support enabled.
	// Example: true
	PidsLimit bool `json:"PidsLimit,omitempty"`

	// plugins
	Plugins *PluginsInfo `json:"Plugins,omitempty"`

	// Reports a summary of the product license on the daemon.
	//
	// If a commercial license has been applied to the daemon, information
	// such as number of nodes, and expiration are included.
	//
	// Example: Community Engine
	ProductLicense string `json:"ProductLicense,omitempty"`

	// registry config
	RegistryConfig *RegistryServiceConfig `json:"RegistryConfig,omitempty"`

	// runc commit
	RuncCommit *Commit `json:"RuncCommit,omitempty"`

	// List of [OCI compliant](https://github.com/opencontainers/runtime-spec)
	// runtimes configured on the daemon. Keys hold the "name" used to
	// reference the runtime.
	//
	// The Docker daemon relies on an OCI compliant runtime (invoked via the
	// `containerd` daemon) as its interface to the Linux kernel namespaces,
	// cgroups, and SELinux.
	//
	// The default runtime is `runc`, and automatically configured. Additional
	// runtimes can be configured by the user and will be listed here.
	//
	// Example: {"custom":{"path":"/usr/local/bin/my-oci-runtime","runtimeArgs":["--debug","--systemd-cgroup=false"]},"runc":{"path":"runc"},"runc-master":{"path":"/go/bin/runc"}}
	Runtimes map[string]Runtime `json:"Runtimes,omitempty"`

	// List of security features that are enabled on the daemon, such as
	// apparmor, seccomp, SELinux, user-namespaces (userns), and rootless.
	//
	// Additional configuration options for each security feature may
	// be present, and are included as a comma-separated list of key/value
	// pairs.
	//
	// Example: ["name=apparmor","name=seccomp,profile=default","name=selinux","name=userns","name=rootless"]
	SecurityOptions []string `json:"SecurityOptions"`

	// Version string of the daemon.
	//
	// > **Note**: the [standalone Swarm API](https://docs.docker.com/swarm/swarm-api/)
	// > returns the Swarm version instead of the daemon  version, for example
	// > `swarm/1.2.8`.
	//
	// Example: 17.06.0-ce
	ServerVersion string `json:"ServerVersion,omitempty"`

	// Indicates if the host has memory swap limit support enabled.
	// Example: true
	SwapLimit bool `json:"SwapLimit,omitempty"`

	// swarm
	Swarm *SwarmInfo `json:"Swarm,omitempty"`

	// Status information about this node (standalone Swarm API).
	//
	// <p><br /></p>
	//
	// > **Note**: The information returned in this field is only propagated
	// > by the Swarm standalone API, and is empty (`null`) when using
	// > built-in swarm mode.
	//
	// Example: [["Role","primary"],["State","Healthy"],["Strategy","spread"],["Filters","health, port, containerslots, dependency, affinity, constraint, whitelist"],["Nodes","2"],[" swarm-agent-00","192.168.99.102:2376"],["  └ ID","5CT6:FBGO:RVGO:CZL4:PB2K:WCYN:2JSV:KSHH:GGFW:QOPG:6J5Q:IOZ2|192.168.99.102:2376"],["  └ Status","Healthy"],["  └ Containers","1 (1 Running, 0 Paused, 0 Stopped)"],["  └ Reserved CPUs","0 / 1"],["  └ Reserved Memory","0 B / 1.021 GiB"],["  └ Labels","kernelversion=4.4.74-boot2docker, operatingsystem=Boot2Docker 17.06.0-ce (TCL 7.2); HEAD : 0672754 - Thu Jun 29 00:06:31 UTC 2017, ostype=linux, provider=virtualbox, storagedriver=aufs"],["  └ UpdatedAt","2017-08-09T10:03:46Z"],["  └ ServerVersion","17.06.0-ce"],[" swarm-manager","192.168.99.101:2376"],["  └ ID","TAMD:7LL3:SEF7:LW2W:4Q2X:WVFH:RTXX:JSYS:XY2P:JEHL:ZMJK:JGIW|192.168.99.101:2376"],["  └ Status","Healthy"],["  └ Containers","2 (2 Running, 0 Paused, 0 Stopped)"],["  └ Reserved CPUs","0 / 1"],["  └ Reserved Memory","0 B / 1.021 GiB"],["  └ Labels","kernelversion=4.4.74-boot2docker, operatingsystem=Boot2Docker 17.06.0-ce (TCL 7.2); HEAD : 0672754 - Thu Jun 29 00:06:31 UTC 2017, ostype=linux, provider=virtualbox, storagedriver=aufs"],["  └ UpdatedAt","2017-08-09T10:04:11Z"],["  └ ServerVersion","17.06.0-ce"]]
	SystemStatus [][]string `json:"SystemStatus"`

	// Current system-time in [RFC 3339](https://www.ietf.org/rfc/rfc3339.txt)
	// format with nano-seconds.
	//
	// Example: 2017-08-08T20:28:29.06202363Z
	SystemTime string `json:"SystemTime,omitempty"`

	// List of warnings / informational messages about missing features, or
	// issues related to the daemon configuration.
	//
	// These messages can be printed by the client as information to the user.
	//
	// Example: ["WARNING: No memory limit support","WARNING: bridge-nf-call-iptables is disabled","WARNING: bridge-nf-call-ip6tables is disabled"]
	Warnings []string `json:"Warnings"`
}

// Validate validates this system info
func (m *SystemInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCgroupDriver(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContainerdCommit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGenericResources(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitCommit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIsolation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlugins(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegistryConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuncCommit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRuntimes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSwarm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var systemInfoTypeCgroupDriverPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cgroupfs","systemd","none"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		systemInfoTypeCgroupDriverPropEnum = append(systemInfoTypeCgroupDriverPropEnum, v)
	}
}

const (

	// SystemInfoCgroupDriverCgroupfs captures enum value "cgroupfs"
	SystemInfoCgroupDriverCgroupfs string = "cgroupfs"

	// SystemInfoCgroupDriverSystemd captures enum value "systemd"
	SystemInfoCgroupDriverSystemd string = "systemd"

	// SystemInfoCgroupDriverNone captures enum value "none"
	SystemInfoCgroupDriverNone string = "none"
)

// prop value enum
func (m *SystemInfo) validateCgroupDriverEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, systemInfoTypeCgroupDriverPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SystemInfo) validateCgroupDriver(formats strfmt.Registry) error {
	if swag.IsZero(m.CgroupDriver) { // not required
		return nil
	}

	// value enum
	if err := m.validateCgroupDriverEnum("CgroupDriver", "body", *m.CgroupDriver); err != nil {
		return err
	}

	return nil
}

func (m *SystemInfo) validateContainerdCommit(formats strfmt.Registry) error {
	if swag.IsZero(m.ContainerdCommit) { // not required
		return nil
	}

	if m.ContainerdCommit != nil {
		if err := m.ContainerdCommit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ContainerdCommit")
			}
			return err
		}
	}

	return nil
}

func (m *SystemInfo) validateGenericResources(formats strfmt.Registry) error {
	if swag.IsZero(m.GenericResources) { // not required
		return nil
	}

	if err := m.GenericResources.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("GenericResources")
		}
		return err
	}

	return nil
}

func (m *SystemInfo) validateInitCommit(formats strfmt.Registry) error {
	if swag.IsZero(m.InitCommit) { // not required
		return nil
	}

	if m.InitCommit != nil {
		if err := m.InitCommit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InitCommit")
			}
			return err
		}
	}

	return nil
}

var systemInfoTypeIsolationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["default","hyperv","process"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		systemInfoTypeIsolationPropEnum = append(systemInfoTypeIsolationPropEnum, v)
	}
}

const (

	// SystemInfoIsolationDefault captures enum value "default"
	SystemInfoIsolationDefault string = "default"

	// SystemInfoIsolationHyperv captures enum value "hyperv"
	SystemInfoIsolationHyperv string = "hyperv"

	// SystemInfoIsolationProcess captures enum value "process"
	SystemInfoIsolationProcess string = "process"
)

// prop value enum
func (m *SystemInfo) validateIsolationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, systemInfoTypeIsolationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SystemInfo) validateIsolation(formats strfmt.Registry) error {
	if swag.IsZero(m.Isolation) { // not required
		return nil
	}

	// value enum
	if err := m.validateIsolationEnum("Isolation", "body", *m.Isolation); err != nil {
		return err
	}

	return nil
}

func (m *SystemInfo) validatePlugins(formats strfmt.Registry) error {
	if swag.IsZero(m.Plugins) { // not required
		return nil
	}

	if m.Plugins != nil {
		if err := m.Plugins.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Plugins")
			}
			return err
		}
	}

	return nil
}

func (m *SystemInfo) validateRegistryConfig(formats strfmt.Registry) error {
	if swag.IsZero(m.RegistryConfig) { // not required
		return nil
	}

	if m.RegistryConfig != nil {
		if err := m.RegistryConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RegistryConfig")
			}
			return err
		}
	}

	return nil
}

func (m *SystemInfo) validateRuncCommit(formats strfmt.Registry) error {
	if swag.IsZero(m.RuncCommit) { // not required
		return nil
	}

	if m.RuncCommit != nil {
		if err := m.RuncCommit.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RuncCommit")
			}
			return err
		}
	}

	return nil
}

func (m *SystemInfo) validateRuntimes(formats strfmt.Registry) error {
	if swag.IsZero(m.Runtimes) { // not required
		return nil
	}

	for k := range m.Runtimes {

		if err := validate.Required("Runtimes"+"."+k, "body", m.Runtimes[k]); err != nil {
			return err
		}
		if val, ok := m.Runtimes[k]; ok {
			if err := val.Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *SystemInfo) validateSwarm(formats strfmt.Registry) error {
	if swag.IsZero(m.Swarm) { // not required
		return nil
	}

	if m.Swarm != nil {
		if err := m.Swarm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Swarm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this system info based on the context it is used
func (m *SystemInfo) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateContainerdCommit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGenericResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInitCommit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePlugins(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRegistryConfig(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRuncCommit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateRuntimes(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSwarm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SystemInfo) contextValidateContainerdCommit(ctx context.Context, formats strfmt.Registry) error {

	if m.ContainerdCommit != nil {
		if err := m.ContainerdCommit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ContainerdCommit")
			}
			return err
		}
	}

	return nil
}

func (m *SystemInfo) contextValidateGenericResources(ctx context.Context, formats strfmt.Registry) error {

	if err := m.GenericResources.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("GenericResources")
		}
		return err
	}

	return nil
}

func (m *SystemInfo) contextValidateInitCommit(ctx context.Context, formats strfmt.Registry) error {

	if m.InitCommit != nil {
		if err := m.InitCommit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("InitCommit")
			}
			return err
		}
	}

	return nil
}

func (m *SystemInfo) contextValidatePlugins(ctx context.Context, formats strfmt.Registry) error {

	if m.Plugins != nil {
		if err := m.Plugins.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Plugins")
			}
			return err
		}
	}

	return nil
}

func (m *SystemInfo) contextValidateRegistryConfig(ctx context.Context, formats strfmt.Registry) error {

	if m.RegistryConfig != nil {
		if err := m.RegistryConfig.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RegistryConfig")
			}
			return err
		}
	}

	return nil
}

func (m *SystemInfo) contextValidateRuncCommit(ctx context.Context, formats strfmt.Registry) error {

	if m.RuncCommit != nil {
		if err := m.RuncCommit.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("RuncCommit")
			}
			return err
		}
	}

	return nil
}

func (m *SystemInfo) contextValidateRuntimes(ctx context.Context, formats strfmt.Registry) error {

	for k := range m.Runtimes {

		if val, ok := m.Runtimes[k]; ok {
			if err := val.ContextValidate(ctx, formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *SystemInfo) contextValidateSwarm(ctx context.Context, formats strfmt.Registry) error {

	if m.Swarm != nil {
		if err := m.Swarm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("Swarm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SystemInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SystemInfo) UnmarshalBinary(b []byte) error {
	var res SystemInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
