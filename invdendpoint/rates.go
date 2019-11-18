package invdendpoint

const RatesEndPoint = "/rates/"

type Rate struct {
	CreatedAt int64       `json:"created_at,omitempty"`
	UpdatedAt int64       `json:"updated_at,omitempty"`
	Id        string      `json:"id,omitempty"`
	IsPercent bool        `json:"is_percent,omitempty"`
	Inclusive bool        `json:"inclusive,omitempty"`
	MetaData  interface{} `json:"metadata,omitempty"`
	Name      string      `json:"name,omitempty"`
	Value     float64     `json:"value,omitempty"`
	Object    string      `json:"object,omitempty"`
}
