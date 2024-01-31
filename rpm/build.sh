#!/bin/bash
#
# -*- coding: utf-8 -*-
#
# © Copyright 2024 GSI Helmholtzzentrum für Schwerionenforschung
#
# This software is distributed under
# the terms of the GNU General Public Licence version 3 (GPL Version 3),
# copied verbatim in the file "LICENCE".

set -e

function do_checks {

  info_msg="The build script must be executed from the projects base directory!"

  if [ -z "$VERSION" ]; then
    echo "ERROR: Build failed! VERSION file not found" >&2
    echo "INFO: $info_msg"
    exit 1
  fi

}

export VERSION=$(cat VERSION)
export BUILD_DIR=$HOME/rpmbuild
export PKG_DIR=prometheus-pvdisplay-exporter-$VERSION

go build

sed -i "s/VERSION/$(cat VERSION)/" rpm/prometheus-pvdisplay-exporter.spec
mkdir -p $BUILD_DIR/{BUILD,RPMS,SOURCES,SPECS,SRPMS}
mkdir -p $BUILD_DIR/SOURCES/$PKG_DIR/usr/sbin
mkdir -p $BUILD_DIR/SOURCES/$PKG_DIR/usr/lib/systemd/system
cp rpm/prometheus-pvdisplay-exporter.spec $BUILD_DIR/SPECS/
cp systemd/prometheus-pvdisplay-exporter.service $BUILD_DIR/SOURCES/$PKG_DIR/usr/lib/systemd/system/
cp prometheus-pvdisplay-exporter $BUILD_DIR/SOURCES/$PKG_DIR/usr/sbin/
cd $BUILD_DIR/SOURCES
tar -czvf $PKG_DIR.tar.gz $PKG_DIR
cd $BUILD_DIR
echo build dir is $BUILD_DIR
ls -la $BUILD_DIR/SOURCES
rpmbuild -ba $BUILD_DIR/SPECS/prometheus-pvdisplay-exporter.spec

