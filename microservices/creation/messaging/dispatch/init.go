package dispatch

const (
	UpdateCount = "InteractionCount"

	LISTENER_CHANNEL_COUNT = 120
	MAX_BATCH_SIZE         = 50
	EXE_CHANNEL_COUNT      = 5
)

var (
	db SqlMethod
)

func InitDb(_db SqlMethod) {
	db = _db
}
