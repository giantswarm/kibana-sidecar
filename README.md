[![CircleCI](https://circleci.com/gh/giantswarm/kibana-sidecar.svg?style=svg&circle-token=d56ef75d19e44462aa3a61cc90c6f60024756b94)](https://circleci.com/gh/giantswarm/kibana-sidecar)
[![Docker Repository on Quay](https://quay.io/repository/giantswarm/kibana-sidecar/status "Docker Repository on Quay")](https://quay.io/repository/giantswarm/kibana-sidecar)

# kibana-sidecar

A sidecar for Kibana to apply a prepared configuration

## Docs

We write documents like the following to the `.kibana` index (see the [`config`](https://github.com/giantswarm/kibana-sidecar/tree/master/config) folder for details):

### Index Pattern

ID format: `index-pattern:{uuid}`

Example:

```json
{
  "_index": ".kibana",
  "_type": "doc",
  "_id": "index-pattern:d308c4f0-157e-11e8-8862-f19ddd0b5982",
  "_source": {
    "type": "index-pattern",
    "updated_at": "2018-02-19T15:59:07.941Z",
    "index-pattern": {
      "title": "fluentd-*",
      "timeFieldName": "@timestamp",
      "fields": "[...]"
    }
  }
}
```

The `fields` attribute is a string, actually containing escaped JSON code.
It is an array of field definitions, like this:

```json
[
  {
    "name": "kubernetes.namespace_name",
    "type": "string",
    "count": 100,
    "scripted": false,
    "searchable": true,
    "aggregatable": false,
    "readFromDocValues": false
  }
]
```

### Config

ID: `config:6.1.1`

Example:

```json
{
  "_index": ".kibana",
  "_type": "doc",
  "_id": "config:6.1.1",
  "_source": {
    "type": "config",
    "updated_at": "2018-02-19T14:29:37.545Z",
    "config": {
      "buildNum": 16350,
      "defaultIndex": "d308c4f0-157e-11e8-8862-f19ddd0b5982",
      "dateFormat:tz": "UTC",
      "dateFormat": "YYYY-MM-DD HH:mm:ss.SSS",
      "dateFormat:dow": "Monday",
      "discover:sort:defaultOrder": "asc"
    }
  }
}
