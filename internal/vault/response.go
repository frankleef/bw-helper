package vault

type Response struct {
	Success bool `json:"success"`
	Data    struct {
		Raw string `json:"raw"`
	} `json:"data"`
}
