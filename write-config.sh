#!/bin/bash

# expects these environment variables:
#
# - `ELASTICSEARCH_ENDPOINT`: URL to elasticsearch

set -u

# Our index name
INDEX=.kibana

# Check env variable
if [ "${ELASTICSEARCH_ENDPOINT}" == "" ]; then
  echo "ERROR: Environment variable ELASTICSEARCH_ENDPOINT is not set. Exiting."
  exit 1
fi

while true
do

  # Wait for Elasticsearch to become available
  OKAY="unset"
  curl -s --head "${ELASTICSEARCH_ENDPOINT}"|grep -q "HTTP/1.1 200"
  OKAY=$?
  while [ "$OKAY" != "0" ]; do
    echo "Waiting for Elasticsearch to become available"
    sleep 10
    curl -s --head "${ELASTICSEARCH_ENDPOINT}"|grep -q "HTTP/1.1 200"
    OKAY=$?
  done


  # Check if index exists
  curl -s --head "${ELASTICSEARCH_ENDPOINT}/${INDEX}"|grep -q "HTTP/1.1 200"
  OKAY=$?

  if [ "$OKAY" != "0" ]; then
    # Create Index
    echo "Creating index .kibana"
    curl -s -XPUT "${ELASTICSEARCH_ENDPOINT}/${INDEX}" \
      -H 'Content-Type: application/json' \
      -d @config/index-mapping.json
    echo ""
  fi

  # Write index pattern document
  echo "Writing index-pattern document"
  curl -s -XPUT \
    -H 'Content-Type: application/json' \
    -d @config/indexpattern.json \
    "${ELASTICSEARCH_ENDPOINT}/${INDEX}/doc/index-pattern:giantswarm"
  echo ""

  # Write config document
  echo "Writing config document"
  curl -s -XPUT \
    -H 'Content-Type: application/json' \
    -d @config/config.json \
    "${ELASTICSEARCH_ENDPOINT}/${INDEX}/doc/config:${KIBANA_VERSION}"
  echo ""

  sleep 3600
done
