package config

var (
	// ElasticsearchEndpointDefault is the default endpoint for elasticsearch
	ElasticsearchEndpointDefault = "http://elasticsearch:9200"

	// ElasticsearchEndpoint is the actual elasticsearch endpoint to be used
	// as it will be configured via command line flag
	ElasticsearchEndpoint = ""

	// IndexName is the name of the Elasticsearch index
	// where we store the kibana config
	IndexName = ".kibana"

	// IndexPatternName is the name pattern we assume for log data
	// indices, where * replaces the date stamp
	IndexPatternName = "filebeat-*"

	// TimeFieldName is the field Kibana assumes in index documents carrying
	// the date/time information for a log entry
	TimeFieldName = "@timestamp"

	// MaxConnectionRetries is the number of connection attempts to make
	// at most to connect to elasticsearch (per cycle)
	MaxConnectionRetries = 5
)
