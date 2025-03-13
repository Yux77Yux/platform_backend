package messaging

var (
	connStr string
)

func InitStr(_str string) {
	connStr = _str
}

const (
	UpdateDbCreation    = "UpdateDbCreation"
	StoreCreationInfo   = "StoreCreationInfo"
	UpdateCacheCreation = "UpdateCacheCreation"

	// Review
	PendingCreation      = "PendingCreation"      // 起点
	UpdateCreationStatus = "UpdateCreationStatus" // 终点
	DeleteCreation       = "DeleteCreation"

	// Interaction Aggregator
	UPDATE_CREATION_ACTION_COUNT = "InteractionCount" // 终点
)

var (
	ExchangesConfig = map[string]string{
		UpdateDbCreation:             "direct",
		UpdateCacheCreation:          "direct",
		StoreCreationInfo:            "direct",
		UpdateCreationStatus:         "direct",
		DeleteCreation:               "direct",
		UPDATE_CREATION_ACTION_COUNT: "direct",
		// Add more exchanges here
	}
)
