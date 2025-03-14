package messaging

const (
	New_review      = "New_review"
	Comment_review  = "Comment_review"
	User_review     = "User_review"
	Creation_review = "Creation_review"
	PendingCreation = "PendingCreation"

	Update      = "Update"
	BatchUpdate = "BatchUpdate"

	// USER
	USER_APPROVE  = "UpdateUserStatus"
	USER_REJECTED = "UpdateUserStatus"

	// CREATION
	CREATION_APPROVE  = "UpdateCreationStatus"
	CREATION_REJECTED = "UpdateCreationStatus"
	CREATION_DELETED  = "DeleteCreation"

	// COMMENT
	COMMENT_REJECTED = "DeleteComment"
	COMMENT_DELETED  = "DeleteComment"
)

var (
	connStr         string
	ExchangesConfig = map[string]string{
		New_review:      "direct",
		PendingCreation: "direct",
		Update:          "direct",
		BatchUpdate:     "direct",
		// Add more exchanges here
	}
)

func InitStr(_str string) {
	connStr = _str
}
