# prometheus-pvdisplay-exporter

## Building

### Binary

`go build`

### RPM Package

`rpm/build.sh`

> Must be executed from the project directory

## Requirements

CLI tool `pvdisplay` installed on the target host.

## Metrics

All metrics are prefixed with "pvdisplay_".

| Metric | Labels | Description                  |
| ------ | ------ | ---------------------------- |
| error  | -      | Set if an error has occurred |
| psize  | vg     | PSize for given VG           |
| pfree  | vg     | PFree for given VG           |