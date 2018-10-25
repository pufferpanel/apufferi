package http

type response struct {
	Success  bool        `json:"success"`
	Message  string      `json:"msg,omitempty"`
	Data     interface{} `json:"data,omitempty"`
	Status   int         `json:"-"`
	Code     Code        `json:"code,omitempty"`
	Metadata *metadata    `json:"metadata,omitempty"`
}

type metadata struct {
	Paging *paging `json:"paging"`
}

type paging struct {
	Page    uint `json:"page,omitempty"`
	Size    uint `json:"pageSize,omitempty"`
	MaxSize uint `json:"maxSize,omitempty"`
}
