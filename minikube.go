package main

import (
	"net"

	log "github.com/Sirupsen/logrus"
	"github.com/mstrzele/minikube-nfs/minikube"
)

func checkClusterPresence() {
	clusterState, _, err := minikube.Status()
	if err != nil {
		log.Fatal(err.Error())
	}

	logger := log.WithFields(log.Fields{"clusterState": clusterState})

	if clusterState != "" {
		logger.Info("machine presence ...")
	} else {
		logger.Fatal("Could not find the cluster.")
	}
}

func checkClusterRunning() {
	clusterState, _, err := minikube.Status()
	if err != nil {
		log.Fatal(err.Error())
	}

	logger := log.WithFields(log.Fields{"clusterState": clusterState})

	if clusterState == "Running" {
		logger.Info("machine running ...")
	} else {
		logger.Fatal("The cluster is not running!")
	}
}

func lookupMandatoryProperties() net.IP {
	clusterIP, err := minikube.IP()
	if err != nil {
		log.Fatal(err.Error())
	}

	logger := log.WithFields(log.Fields{"clusterIP": clusterIP})

	if clusterIP != nil {
		logger.Info("Lookup mandatory properties ...")
		return clusterIP
	}

	return nil
}
