package responses

type Default struct {
	ResponseCode    int    `json:"responseCode" example:"200"`
	ResponseMessage string `json:"responseMessage" example:"Berhasil mengambil data."`
}

type Data struct {
	ResponseCode    int         `json:"responseCode" example:"200"`
	ResponseMessage string      `json:"responseMessage" example:"Berhasil mengambil data."`
	Data            interface{} `json:"data"`
}
