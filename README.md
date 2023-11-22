## ARCHIVED

This repo was an attempt at using the swagger spec prior to an official client has been released.  Please use the official client that can be found [here](https://github.com/grafana/grafana-openapi-client-go).



---

![Go workflow](https://github.com/esnet/grafana-swagger-api-golang/actions/workflows/go.yml/badge.svg)


# Grafana HTTP API Client for Go

This library provides a low-level client to access Grafana [HTTP API](https://grafana.com/docs/grafana/latest/http_api/).

## Generating New GoModels

`GRAFANA_TARGET_VERSION=v9.3.6 make generate`


This project is heavily inspired by the work from @papagian which can be found on [here](https://github.com/esnet/grafana-swagger-api-golang/tree/papagian/generate-client-from-swagger)

Initially, API version will match the grafana version.
