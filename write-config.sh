#!/bin/bash

# expects these environment variables:
#
# - `ELASTICSEARCH_ENDPOINT`: URL to elasticsearch

# Our index name
INDEX=.kibana

while true
do

  echo "Waiting 3 sec for Elasticsearch to come up."
  sleep 3

  # Check if index exists
  curl -s --head "$ELASTICSEARCH_ENDPOINT/$INDEX"|grep -q "HTTP/1.1 200"
  OKAY=$?

  if [ "$OKAY" != "0" ]; then
    # Create Index
    echo "Creating index .kibana"
    curl -s -XPUT "$ELASTICSEARCH_ENDPOINT/$INDEX" \
      -H 'Content-Type: application/json' \
      -d @config/index-mapping.json
    echo ""
  fi

  # Write index pattern document
  echo "Writing index-pattern document"
  curl -s -XPUT \
    -H 'Content-Type: application/json' \
    -d @config/indexpattern.json \
    "$ELASTICSEARCH_ENDPOINT/$INDEX/doc/index-pattern:giantswarm"
  echo ""

  # Write config document
  echo "Writing config document"
  curl -s -XPUT \
    -H 'Content-Type: application/json' \
    -d @config/config.json \
    "$ELASTICSEARCH_ENDPOINT/$INDEX/doc/config:6.1.1"
  echo ""

	sleep 3600
done
