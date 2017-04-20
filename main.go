package main

import (
	"fmt"
	"os/user"

	log "github.com/Sirupsen/logrus"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	debug        bool
	force        bool
	nfsConfig    string
	sharedFolder []string
	mountOpts    string
	useIPRange   bool
)

func main() {
	v, err := version()
	if err != nil {
		log.Error(err.Error())
	}

	kingpin.Version(v)

	kingpin.CommandLine.HelpFlag.Short('h')
	kingpin.CommandLine.VersionFlag.Short('v')

	kingpin.Flag("debug", "Enable debug mode.").Short('d').BoolVar(&debug)
	kingpin.Flag("force", "Force reconfiguration of NFS.").Short('f').BoolVar(&force)

	u, err := user.Current()
	if err != nil {
		log.Error(err.Error())
	}

	kingpin.Flag("nfs-config", "NFS configuration to use in /etc/exports.").Default(fmt.Sprintf("-alldirs -mapall=%s:%s", u.Uid, u.Gid)).Short('n').StringVar(&nfsConfig)

	kingpin.Flag("shared-folder", "Folder to share").Default("/Users").Short('s').StringsVar(&sharedFolder)
	kingpin.Flag("mount-opts", "NFS mount options").Default("noacl,async").Short('m').StringVar(&mountOpts)
	kingpin.Flag("use-ip-range", "Changes the nfs export ip to a range (e.g. -network 192.168.99.100 becomes -network 192.168.99)").Short('i').BoolVar(&useIPRange)

	kingpin.Parse()

	if debug {
		log.SetLevel(log.DebugLevel)
	}
}
