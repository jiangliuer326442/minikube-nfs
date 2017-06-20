package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"regexp"
	"strings"

	log "github.com/Sirupsen/logrus"
	"github.com/mstrzele/minikube-nfs/nfsd"
)

func hosts(machineIP net.IP, useIPRange bool) nfsd.Hosts {
	hosts := nfsd.Hosts{}

	if useIPRange {
		ifat, err := net.InterfaceAddrs()
		if err != nil {
			return hosts
		}

		for _, ifa := range ifat {
			if _, n, err := net.ParseCIDR(ifa.String()); err == nil && n.Contains(machineIP) {
				hosts.Network = n
			}
		}
	} else {
		hosts.Names = []string{machineIP.String()}
	}

	return hosts
}

const (
	exportsBegin = "# minikube-nfs-begin #"
	exportsEnd   = "# minikube-nfs-end #"
)

func configureNFS(
	nfsConfig string,
	sharedFolders []string,
	machineIP net.IP,
	useIPRange bool,
) {
	logger := log.WithFields(log.Fields{
		"nfsConfig":     nfsConfig,
		"sharedFolders": sharedFolders,
		"machineIP":     machineIP,
		"useIPRange":    useIPRange,
	})

	export := nfsd.Export{
		Directories: sharedFolders,
		Flags:       strings.Split(nfsConfig, " "),
		Hosts:       hosts(machineIP, useIPRange),
	}

	exports, err := ioutil.ReadFile(nfsd.ExportsFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	exportsBlock := []byte(fmt.Sprintf("%s\n%s\n%s\n", exportsBegin, export, exportsEnd))

	re := regexp.MustCompile(fmt.Sprintf("%s\n.*\n%s\n", exportsBegin, exportsEnd))
	if re.Match(exports) {
		exports = re.ReplaceAll(exports, exportsBlock)
	} else {
		exports = append(exports, exportsBlock...)
	}

	if err := ioutil.WriteFile(nfsd.ExportsFile, exports, 0644); err != nil {
		log.Fatal(err.Error())
	}

	if err := nfsd.CheckExports(); err != nil {
		log.Fatal(err.Error())
	}

	if err := nfsd.Enable(); err != nil {
		log.Fatal(err.Error())
	}

	if err := nfsd.Restart(); err != nil {
		log.Fatal(err.Error())
	}

	logger.Info("Configure NFS ...")
}
