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
	"regexp"
	"strconv"
)

const (
	Namespace = "pvdisplay"
	CliTool   = "pvdisplay"
)

var (
	pvdisplayRegex      = regexp.MustCompile(`(?m:^(?:\/dev[\/a-zA-Z0-9\-]+)\s+(?P<vg>[a-zA-Z0-9]+)\s+(?:[a-zA-Z0-9]+)\s+(?:[a-zA-Z\-]+)\s+(?P<psize>[0-9]+)B\s+(?P<pfree>[0-9]+)B$)`)
	pvdisplayVgIndex    = pvdisplayRegex.SubexpIndex("vg")
	pvdisplayPSizeIndex = pvdisplayRegex.SubexpIndex("psize")
	pvdisplayPFreeIndex = pvdisplayRegex.SubexpIndex("pfree")
)

type PvdisplayItem struct {
	Vg    string
	PSize float64
	PFree float64
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

func ExtractPvdisplayItems(data string) ([]PvdisplayItem, error) {
	slice := make([]PvdisplayItem, 0, 2)
	matchedItems := pvdisplayRegex.FindAllStringSubmatch(data, -1)

	if matchedItems == nil {
		return nil, fmt.Errorf("pvdisplayRegex missmatch on data:\n%s", data)
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
