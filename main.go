package main

import (
	"fmt"
	"os"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug         = kingpin.Flag("debug", "Enable debug mode.").Short('d').Bool()
	nfsConfig     = kingpin.Flag("nfs-config", "NFS configuration to use in /etc/exports.").Default(fmt.Sprintf("-alldirs -mapall=%d:%d", os.Getuid(), os.Getgid())).Short('n').String()
	sharedFolders = kingpin.Flag("shared-folder", "Folder to share").Default("/Users").Short('s').Strings()
	useIPRange    = kingpin.Flag("use-ip-range", "Changes the nfs export ip to a range (e.g. -network 192.168.99.100 becomes -network 192.168.99)").Short('i').Bool()
)

func parseCLI() {
	v, err := version()
	if err != nil {
		log.Error(err.Error())
	}

	kingpin.Version(v)

	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.CommandLine.VersionFlag.Short('v')

	kingpin.Parse()

	if *debug {
		log.SetLevel(log.DebugLevel)
	}
}

func main() {
	parseCLI()

	checkClusterPresence()
	checkClusterRunning()

	clusterIP := lookupMandatoryProperties()

	configureNFS(*nfsConfig, *sharedFolders, clusterIP, *useIPRange)
}
