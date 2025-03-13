package messaging

const (
	ComputeSimilarCreation = "ComputeSimilarCreation"
	ComputeUser            = "ComputeUser"

	UpdateDb      = "UpdateDb"
	BatchUpdateDb = "BatchUpdateDb"
	AddCollection = "AddCollection"
	AddLike       = "AddLike"
	AddView       = "AddView"
	CancelLike    = "CancelLike"

	// Creation
	UPDATE_CREATION_ACTION_COUNT = "InteractionCount"
)

var (
	connStr         string
	ExchangesConfig = map[string]string{
		ComputeSimilarCreation: "direct",
		ComputeUser:            "direct",

		UpdateDb:      "direct",
		AddCollection: "direct",
		AddLike:       "direct",
		AddView:       "direct",
		CancelLike:    "direct",
		BatchUpdateDb: "direct",
		// Add more exchanges here
	}
)

func InitStr(_str string) {
	connStr = _str
}
