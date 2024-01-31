// -*- coding: utf-8 -*-
//
// © Copyright 2024 GSI Helmholtzzentrum für Schwerionenforschung
//
// This software is distributed under
// the terms of the GNU General Public Licence version 3 (GPL Version 3),
// copied verbatim in the file "LICENCE".

package collector

import (
	"fmt"
	"prometheus-pvdisplay-exporter/util"
	"regexp"
	"strconv"

	"github.com/prometheus/client_golang/prometheus"

	log "github.com/sirupsen/logrus"
)

const (
	CliTool = "pvdisplay"
)

var (
	pvdisplayRegex      = regexp.MustCompile(`(?m:^(?:\/dev[\/a-zA-Z0-9\-]+)\s+(?P<vg>[a-zA-Z0-9]+)\s+(?:[a-zA-Z0-9]+)\s+(?:[a-zA-Z\-]+)\s+(?P<psize>[0-9]+)B\s+(?P<pfree>[0-9]+)B$)`)
	pvdisplayVgIndex    = pvdisplayRegex.SubexpIndex("vg")
	pvdisplayPSizeIndex = pvdisplayRegex.SubexpIndex("psize")
	pvdisplayPFreeIndex = pvdisplayRegex.SubexpIndex("pfree")

	pvdisplayPSizeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(Namespace, "collector", "psize"),
		"PSize for a given VG",
		[]string{"vg"},
		nil,
	)

	pvdisplayPFreeDesc = prometheus.NewDesc(
		prometheus.BuildFQName(Namespace, "collector", "pfree"),
		"PFree for a given VG",
		[]string{"vg"},
		nil,
	)
)

type PvdisplayItem struct {
	Vg    string
	PSize float64
	PFree float64
}

type PvdisplayCollector struct{}

func NewPvdisplayCollector() prometheus.Collector {
	return &PvdisplayCollector{}
}

func (c *PvdisplayCollector) Collect(ch chan<- prometheus.Metric) {

	log.Debug("Collecting pvdisplay data")

	data, err := util.ExecuteCommandWithSudo(CliTool, "--units", "b", "-C", "--noheadings")

	if err != nil {
		log.Error(err)
		ch <- createErrorMetric()
		return
	}

	items, err := ExtractPvdisplayItems(data)

	if err != nil {
		log.Error(err)
		ch <- createErrorMetric()
		return
	}

	for _, item := range items {
		ch <- c.createMetric(pvdisplayPSizeDesc, item.Vg, item.PSize)
		ch <- c.createMetric(pvdisplayPFreeDesc, item.Vg, item.PFree)
	}
}

func (c *PvdisplayCollector) Describe(ch chan<- *prometheus.Desc) {
}

func (c *PvdisplayCollector) createMetric(desc *prometheus.Desc, vg string, val float64) prometheus.Metric {
	return prometheus.MustNewConstMetric(desc, prometheus.GaugeValue, val, vg)
}

func newPvdisplayItem(vg string, psize float64, pfree float64) (*PvdisplayItem, error) {

	if vg == "" {
		return nil, fmt.Errorf("vg argument empty")
	}

	if psize == 0 {
		return nil, fmt.Errorf("psize is 0")
	}

	return &PvdisplayItem{vg, psize, pfree}, nil
}

func ExtractPvdisplayItems(data *string) ([]PvdisplayItem, error) {

	slice := make([]PvdisplayItem, 0, 2)
	matchedItems := pvdisplayRegex.FindAllStringSubmatch(*data, -1)

	if matchedItems == nil {
		return nil, fmt.Errorf("pvdisplayRegex missmatch on data:\n%s", *data)
	}

	for _, mItem := range matchedItems {

		vg := mItem[pvdisplayVgIndex]

		psize, err := strconv.ParseFloat(mItem[pvdisplayPSizeIndex], 10)
		if err != nil {
			return nil, err
		}

		pfree, err := strconv.ParseFloat(mItem[pvdisplayPFreeIndex], 10)
		if err != nil {
			return nil, err
		}

		pvItem, err := newPvdisplayItem(vg, psize, pfree)
		if err != nil {
			return nil, err
		}

		slice = append(slice, *pvItem)
	}

	return slice, nil
}
