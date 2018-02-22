#!/bin/bash

# expects these environment variables:
#
# - `ELASTICSEARCH_ENDPOINT`: URL to elasticsearch

while true
do

  echo "Waiting 60 sec for Elasticsearch to come up."
  sleep 60

  # Delete Index
  echo "Deleting index .kibana"
  curl -s -XDELETE "$ELASTICSEARCH_ENDPOINT/.kibana" || echo "Index didn't exist, not deleted."
  echo ""

  # Create Index
  echo "Creating index .kibana"
  curl -s -XPUT "$ELASTICSEARCH_ENDPOINT/.kibana" \
    -H 'Content-Type: application/json' \
    -d @/config/index-mapping.json
  echo ""

  # Write index pattern document
  echo "Writing index-pattern document"
  curl -s -XPUT \
    -H 'Content-Type: application/json' \
    -d @/config/indexpattern.json \
    "$ELASTICSEARCH_ENDPOINT/.kibana/doc/index-pattern:giantswarm"
  echo ""

  # Write config document
  echo "Writing config document"
  curl -s -XPUT \
    -H 'Content-Type: application/json' \
    -d @/config/config.json \
    "$ELASTICSEARCH_ENDPOINT/.kibana/doc/config:6.1.1"
  echo ""

	echo "Waiting for an hour"
	sleep 3600
done
