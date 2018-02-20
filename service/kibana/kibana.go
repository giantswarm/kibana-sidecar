// Package kibana provides a service to submit Kibana
// configuration into Elasticsearch
package kibana

import (
	"context"
	"fmt"
	"log"
	"time"

	uuid4 "github.com/nathanwinther/go-uuid4"
	"github.com/olivere/elastic"
)

const (
	// IndexName is the name of the Elasticsearch index
	// where we store our config
	IndexName = ".kibana"

	// IndexPatternName is the name pattern we assume for log data
	// indices, where * replaces the date stamp
	IndexPatternName = "filebeat-*"

	// ElasticsearchEndpointDefault is the default endpoint for elasticsearch
	ElasticsearchEndpointDefault = "http://elasticsearch:9200"

	// MaxConnectionRetries is the number of connection attempts to make
	// at most to connect to elasticsearch
	MaxConnectionRetries = 3
)

var (
	// the actual elasticsearch endpoint to use
	endpoint string
	// elasticsearch client
	client *elastic.Client
)

// Setup creates a client and checks the connection.
func Setup() error {
	endpoint = ElasticsearchEndpointDefault

	var clientErr error
	client, clientErr = elastic.NewClient(elastic.SetURL(endpoint))
	if clientErr != nil {
		return clientErr
	}

	info, code, err := client.Ping(endpoint).Do(context.Background())
	if err != nil {
		return err
	}
	log.Printf("Elasticsearch ping status: %d, version: %v", code, info.Version.Number)

	return nil
}

// CreateIndex creates the kibana index
func CreateIndex() error {
	log.Printf("Creating index %s", IndexName)

	mapping := `{
    "settings": {
      "number_of_shards": 1,
		  "number_of_replicas": 0
    },
    "mappings": {
      "doc": {
        "dynamic": "strict",
        "properties": {
          "config": {
            "dynamic": "true",
            "properties": {
              "buildNum": {
                "type": "keyword"
              },
              "defaultIndex": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 256
                  }
                }
              }
            }
          },
          "dashboard": {
            "properties": {
              "description": {
                "type": "text"
              },
              "hits": {
                "type": "integer"
              },
              "kibanaSavedObjectMeta": {
                "properties": {
                  "searchSourceJSON": {
                    "type": "text"
                  }
                }
              },
              "optionsJSON": {
                "type": "text"
              },
              "panelsJSON": {
                "type": "text"
              },
              "refreshInterval": {
                "properties": {
                  "display": {
                    "type": "keyword"
                  },
                  "pause": {
                    "type": "boolean"
                  },
                  "section": {
                    "type": "integer"
                  },
                  "value": {
                    "type": "integer"
                  }
                }
              },
              "timeFrom": {
                "type": "keyword"
              },
              "timeRestore": {
                "type": "boolean"
              },
              "timeTo": {
                "type": "keyword"
              },
              "title": {
                "type": "text"
              },
              "uiStateJSON": {
                "type": "text"
              },
              "version": {
                "type": "integer"
              }
            }
          },
          "index-pattern": {
            "properties": {
              "fieldFormatMap": {
                "type": "text"
              },
              "fields": {
                "type": "text"
              },
              "intervalName": {
                "type": "keyword"
              },
              "notExpandable": {
                "type": "boolean"
              },
              "sourceFilters": {
                "type": "text"
              },
              "timeFieldName": {
                "type": "keyword"
              },
              "title": {
                "type": "text"
              }
            }
          },
          "search": {
            "properties": {
              "columns": {
                "type": "keyword"
              },
              "description": {
                "type": "text"
              },
              "hits": {
                "type": "integer"
              },
              "kibanaSavedObjectMeta": {
                "properties": {
                  "searchSourceJSON": {
                    "type": "text"
                  }
                }
              },
              "sort": {
                "type": "keyword"
              },
              "title": {
                "type": "text"
              },
              "version": {
                "type": "integer"
              }
            }
          },
          "server": {
            "properties": {
              "uuid": {
                "type": "keyword"
              }
            }
          },
          "timelion-sheet": {
            "properties": {
              "description": {
                "type": "text"
              },
              "hits": {
                "type": "integer"
              },
              "kibanaSavedObjectMeta": {
                "properties": {
                  "searchSourceJSON": {
                    "type": "text"
                  }
                }
              },
              "timelion_chart_height": {
                "type": "integer"
              },
              "timelion_columns": {
                "type": "integer"
              },
              "timelion_interval": {
                "type": "keyword"
              },
              "timelion_other_interval": {
                "type": "keyword"
              },
              "timelion_rows": {
                "type": "integer"
              },
              "timelion_sheet": {
                "type": "text"
              },
              "title": {
                "type": "text"
              },
              "version": {
                "type": "integer"
              }
            }
          },
          "type": {
            "type": "keyword"
          },
          "updated_at": {
            "type": "date"
          },
          "url": {
            "properties": {
              "accessCount": {
                "type": "long"
              },
              "accessDate": {
                "type": "date"
              },
              "createDate": {
                "type": "date"
              },
              "url": {
                "type": "text",
                "fields": {
                  "keyword": {
                    "type": "keyword",
                    "ignore_above": 2048
                  }
                }
              }
            }
          },
          "visualization": {
            "properties": {
              "description": {
                "type": "text"
              },
              "kibanaSavedObjectMeta": {
                "properties": {
                  "searchSourceJSON": {
                    "type": "text"
                  }
                }
              },
              "savedSearchId": {
                "type": "keyword"
              },
              "title": {
                "type": "text"
              },
              "uiStateJSON": {
                "type": "text"
              },
              "version": {
                "type": "integer"
              },
              "visState": {
                "type": "text"
              }
            }
          }
        }
      }
    }
  }`

	createIndex, err := client.CreateIndex(IndexName).Body(mapping).Do(context.Background())
	if err != nil {
		return err
	}
	if !createIndex.Acknowledged {
		log.Printf("Creating index not acknowledged by Elasticsearch. Let's see...")
	}

	return nil
}

// WriteIndexPattern creates the index-pattern document
func WriteIndexPattern() error {

	doc := new(IndexPatternDocument)
	doc.TypeName = "index-pattern"
	doc.UpdatedAt = time.Now().Format("2006-01-02T15:04:05.000Z")
	doc.IndexPattern = new(IndexPattern)
	doc.IndexPattern.TimeFieldName = "@timestamp"
	doc.IndexPattern.Title = IndexPatternName
	doc.IndexPattern.Fields = "[]"

	u, err := uuid4.New()
	if err != nil {
		return err
	}
	id := "index-pattern:" + u
	put1, err := client.Index().
		Index(IndexName).
		Type("doc").
		Id(id).
		BodyJson(doc).
		Do(context.Background())
	if err != nil {
		return err
	}

	fmt.Printf("Added %s to index %s", put1.Id, put1.Index)

	return nil
}

// WriteConfig writes our configuration to ElasticSearch
func WriteConfig() {

	// set up the connection (with retries)
	var err error
	for i := 0; i < MaxConnectionRetries; i++ {
		err = Setup()
		if err == nil {
			break
		} else {
			if i == (MaxConnectionRetries - 1) {
				log.Printf("Could not connect to elasticsearch after %d attempts", i)
				return
			}
		}
		log.Println("Couldn't initialize elasticsearch client yet. Wating...")
		time.Sleep(10 * time.Second)
	}

	exists, err := client.IndexExists(IndexName).Do(context.Background())
	if err != nil {
		log.Printf("Could not determine if Elasticsearch index exists at %s. Skipping.", endpoint)
		log.Println(err)
		return
	}

	if !exists {
		err = CreateIndex()
		if err != nil {
			log.Printf("Could not create index:")
			log.Println(err)
			return
		}
	}

	// now we can actually create our docs
	err = WriteIndexPattern()
	if err != nil {
		log.Println("Cloud not create index-pattern document.")
		log.Println(err)
	}

}
