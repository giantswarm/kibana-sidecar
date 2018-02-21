package kibana

// AbstractDocument defines fields used in everal document types
type AbstractDocument struct {
	TypeName  string `json:"type"`
	UpdatedAt string `json:"updated_at"`
}

// IndexPatternDocument allows to create a document of type "index-pattern"
//
// {
// 	"type": "index-pattern",
// 	"updated_at": "2018-02-19T15:59:07.941Z",
// 	"index-pattern": {
// 		"title": "fluentd-*",
// 		"timeFieldName": "@timestamp",
// 		"fields": "[...]"
// 	}
// }
//
type IndexPatternDocument struct {
	AbstractDocument
	IndexPattern *IndexPattern `json:"index-pattern"`
}

// IndexPattern is a sub-structure used in IndexPatternDocument
type IndexPattern struct {
	Title         string `json:"title"`
	TimeFieldName string `json:"timeFieldName"`
	Fields        string `json:"fields"`
}

// FieldDefinition is a sub-structure used in IndexPatternDocument
type FieldDefinition struct {
	FieldName         string `json:"name"`
	TypeName          string `json:"type"`
	Count             int32  `json:"count"`
	Scripted          bool   `json:"scripted"`
	Searchable        bool   `json:"searchable"`
	Aggregatable      bool   `json:"aggregatable"`
	ReadFromDocValues bool   `json:"readFromDocValues"`
}

// ConfigDocument allows to create a document of type "config"
//
// {
// 	"type": "config",
// 	"updated_at": "2018-02-19T14:29:37.545Z",
// 	"config": {
// 		"buildNum": 16350,
// 		"defaultIndex": "d308c4f0-157e-11e8-8862-f19ddd0b5982",
// 		"dateFormat:tz": "UTC",
// 		"dateFormat": "YYYY-MM-DD HH:mm:ss.SSS",
// 		"dateFormat:dow": "Monday",
// 		"discover:sort:defaultOrder": "asc"
// 	}
// }
//
type ConfigDocument struct {
	AbstractDocument
	Config *ConfigDefinition `json:"config"`
}

type ConfigDefinition struct {
	BuildNum     int32  `json:"buildNum"`
	DefaultIndex string `json:"defaultIndex"`
}
