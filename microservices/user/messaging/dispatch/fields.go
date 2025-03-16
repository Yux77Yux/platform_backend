package dispatch

const (
	LISTENER_CHANNEL_COUNT = 80
	MAX_BATCH_SIZE         = 50
	EXE_CHANNEL_COUNT      = 5

	Register      = "Register"
	RegisterCache = "RegisterCache"

	InsertUser      = "InsertUser"
	InsertUserCache = "InsertUserCache"

	UpdateUserAvatar      = "UpdateUserAvatar"
	UpdateUserAvatarCache = "UpdateUserAvatarCache"

	UpdateUserSpace      = "UpdateUserSpace"
	UpdateUserSpaceCache = "UpdateUserSpaceCache"

	UpdateUserStatus      = "UpdateUserStatus"
	UpdateUserStatusCache = "UpdateUserStatusCache"

	UpdateUserBio      = "UpdateUserBio"
	UpdateUserBioCache = "UpdateUserBioCache"

	Follow      = "Follow"
	FollowCache = "FollowCache"
)
