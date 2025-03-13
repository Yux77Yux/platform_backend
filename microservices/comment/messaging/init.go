package messaging

const (
	PublishComment = "PublishComment"
	DeleteComment  = "DeleteComment"
)

var (
	connStr         string
	ExchangesConfig = map[string]string{
		"PublishComment": "direct",
		"DeleteComment":  "direct",
	}
)

func InitStr(_str string) {
	connStr = _str
}
