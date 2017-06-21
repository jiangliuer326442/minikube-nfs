package minikube

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func command(command ...string) (string, error) {
	cmd := exec.Command("minikube", command...)
	if os.Geteuid() == 0 {
		cmd.SysProcAttr = &syscall.SysProcAttr{
			Credential: &syscall.Credential{},
		}

		if val, ok := os.LookupEnv("SUDO_UID"); ok {
			if uid, err := strconv.Atoi(val); err == nil {
				cmd.SysProcAttr.Credential.Uid = uint32(uid)
			}
		}

		if val, ok := os.LookupEnv("SUDO_GID"); ok {
			if gid, err := strconv.Atoi(val); err == nil {
				cmd.SysProcAttr.Credential.Gid = uint32(gid)
			}
		}
	}

	out, err := cmd.Output()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(out)), nil
}

const minikubeStatusFormat = "{{.MinikubeStatus}},{{.LocalkubeStatus}}"

// Status gets the status of a local kubernetes cluster.
func Status() (string, string, error) {
	out, err := command("status", "--format", fmt.Sprintf("%s", minikubeStatusFormat))
	if err != nil {
		return "", "", err
	}

	status := strings.Split(out[:], ",")
	switch len(status) {
	default:
		return "", "", nil
	case 1:
		return status[0], "", nil
	case 2:
		return status[0], status[1], nil
	}
}

// IP retrieves the IP address of the running cluster.
func IP() (net.IP, error) {
	out, err := command("ip")
	if err != nil {
		return nil, err
	}

	ip := net.ParseIP(out)
	return ip, nil
}
