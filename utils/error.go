package utils

type HandlerError struct {
	Code    int    `json:"code" bson:"code"`
	Status  bool   `json:"status" bson:"status"`
	Message string `json:"message" bson:"message"`
	Data    string `json:"data" bson:"data"`
}

func (e HandlerError) Error() string {
	return e.Message
}
