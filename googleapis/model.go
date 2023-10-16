package googleapis

type GoogleResult struct {
	Items []struct {
		Link string `json:"link"`
	} `json:"items"`
}

type MemImage []byte
