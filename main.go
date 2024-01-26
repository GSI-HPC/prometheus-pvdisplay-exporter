// -*- coding: utf-8 -*-
//
// © Copyright 2024 GSI Helmholtzzentrum für Schwerionenforschung
//
// This software is distributed under
// the terms of the GNU General Public Licence version 3 (GPL Version 3),
// copied verbatim in the file "LICENCE".

package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	log "github.com/sirupsen/logrus"
)

const (
	defaultPort     = "9259"
	defaultLogLevel = "ERROR"
	version         = "0.0.1"
)

func initLogging(logLevel string) {

	if logLevel == "ERROR" {
		log.SetLevel(log.ErrorLevel)
	} else if logLevel == "WARNING" {
		log.SetLevel(log.WarnLevel)
	} else if logLevel == "INFO" {
		log.SetLevel(log.InfoLevel)
	} else if logLevel == "DEBUG" {
		log.SetLevel(log.DebugLevel)
	} else if logLevel == "TRACE" {
		log.SetLevel(log.TraceLevel)
	} else {
		log.Fatal("Not supported log level set")
	}

	log.SetOutput(os.Stdout)
}

func main() {
	port := flag.String("port", defaultPort, "The port to listen on for HTTP requests")
	logLevel := flag.String("log", defaultLogLevel, "Sets log level - ERROR, WARNING, INFO, DEBUG or TRACE")

	flag.Parse()

	initLogging(*logLevel)

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+*port, nil)
}
