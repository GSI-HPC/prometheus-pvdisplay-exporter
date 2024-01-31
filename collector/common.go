// -*- coding: utf-8 -*-
//
// © Copyright 2024 GSI Helmholtzzentrum für Schwerionenforschung
//
// This software is distributed under
// the terms of the GNU General Public Licence version 3 (GPL Version 3),
// copied verbatim in the file "LICENCE".

package collector

import "github.com/prometheus/client_golang/prometheus"

const (
	Namespace = "pvdisplay"
)

func createErrorMetric() prometheus.Metric {
	return prometheus.MustNewConstMetric(
		prometheus.NewDesc(
			prometheus.BuildFQName(Namespace, "", "error"),
			"Set if an error has occurred",
			nil,
			nil,
		),
		prometheus.GaugeValue,
		1,
	)
}
