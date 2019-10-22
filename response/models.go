package response

import "github.com/pufferpanel/apufferi/v4"

type Error struct {
	Error *apufferi.Error `json:"error"`
}

type Metadata struct {
	Paging *Paging `json:"paging"`
}

type Paging struct {
	Page    uint `json:"page,omitempty"`
	Size    uint `json:"pageSize,omitempty"`
	MaxSize uint `json:"maxSize,omitempty"`
	Total   uint `json:"total,omitempty"`
}
