package kibana

type AbstractDocument struct {
	TypeName  string `json:"type"`
	UpdatedAt string `json:"updated_at"`
}

type IndexPatternDocument struct {
	AbstractDocument
	IndexPattern *IndexPattern `json:"index-pattern"`
}

type IndexPattern struct {
	Title         string `json:"title"`
	TimeFieldName string `json:"timeFieldName"`
	Fields        string `json:"fields"`
}

type FieldDefinition struct {
	FieldName         string `json:"name"`
	TypeName          string `json:"type"`
	Count             int32  `json:"count"`
	Scripted          bool   `json:"scripted"`
	Searchable        bool   `json:"searchable"`
	Aggregatable      bool   `json:"aggregatable"`
	ReadFromDocValues bool   `json:"readFromDocValues"`
}
