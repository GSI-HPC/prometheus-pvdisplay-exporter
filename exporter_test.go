// -*- coding: utf-8 -*-
//
// © Copyright 2024 GSI Helmholtzzentrum für Schwerionenforschung
//
// This software is distributed under
// the terms of the GNU General Public Licence version 3 (GPL Version 3),
// copied verbatim in the file "LICENCE".

package main

import (
	"testing"

	"prometheus-pvdisplay-exporter/collector"
)

func TestExtractPvdisplayItems(t *testing.T) {

	pvdisplayData :=
		"/dev/disk/by-id/scsi-3600062b20854afc029a65252775956fa mds00vg lvm2 a--  19201016201216B 8205899923456B\n" +
			"/dev/sdb   mds01vg lvm2 a--  19201016201216B 8205899923456B"

	pvDisplayItems, err := collector.ExtractPvdisplayItems(&pvdisplayData)
	if err != nil {
		t.Error(err)
	}

	lenPvdisplayData := len(pvDisplayItems)
	if lenPvdisplayData != 2 {
		t.Errorf("Expected 2 items in pvDisplayItems, but got: %d", lenPvdisplayData)
	}
}

func TestParsePvdisplayItems(t *testing.T) {

	m := make(map[string]collector.PvdisplayItem) // [pvdisplayRawString]expectedPvdisplayItem
	m["/dev/disk/by-id/scsi-3600062b20854afc029a65252775956fa mds00vg lvm2 a--  19201016201216B 8205899923456B"] = collector.PvdisplayItem{Vg: "mds00vg", PSize: 19201016201216, PFree: 8205899923456}
	m["/dev/sdb   mds01vg lvm2 a--  19201016201216B 8205899923456B"] = collector.PvdisplayItem{Vg: "mds01vg", PSize: 19201016201216, PFree: 8205899923456}

	for pvdisplayData, expected := range m {

		pvDisplayItems, err := collector.ExtractPvdisplayItems(&pvdisplayData)

		if err != nil {
			t.Error(err)
		}

		lenPvdisplayData := len(pvDisplayItems)
		if len(pvDisplayItems) != 1 {
			t.Errorf("Expected 1 item in pvDisplayItems, but got: %d", lenPvdisplayData)
		}

		pvDisplayItem := pvDisplayItems[0]

		if pvDisplayItem.Vg != expected.Vg {
			t.Errorf("Expected vg %s, got %s", expected.Vg, pvDisplayItem.Vg)
		}

		if pvDisplayItem.PSize != expected.PSize {
			t.Errorf("Expected PSize %f, got %f", expected.PSize, pvDisplayItem.PSize)
		}

		if pvDisplayItem.PFree != expected.PFree {
			t.Errorf("Expected PFree %f, got %f", expected.PFree, pvDisplayItem.PFree)
		}
	}
}
