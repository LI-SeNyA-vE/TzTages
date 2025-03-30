package domain

type ImageFromBD struct {
	ID       uint64 `json:"id"`
	Uuid     string `json:"uuid"`
	FileName string `json:"file_name"`
	Data     []byte `json:"data"`
	Size     int64  `json:"size"`
}
