package tools

import (
	generated "github.com/Yux77Yux/platform_backend/generated/creation"
)

func GetSpaceCreationCountType(byWhat generated.ByCount) string {
	typeStr := ""
	switch byWhat {
	case generated.ByCount_VIEWS:
		typeStr = "ByViews"
	case generated.ByCount_LIKES:
		typeStr = "ByLikes"
	case generated.ByCount_COLLECTIONS:
		typeStr = "ByCollections"
	default:
		typeStr = "ByPublished_Time"
	}
	return typeStr
}

func GetUserCreationsCountType(byWhat generated.ByCount) string {
	typeStr := ""
	switch byWhat {
	case generated.ByCount_VIEWS:
		typeStr = "views"
	case generated.ByCount_LIKES:
		typeStr = "likes"
	case generated.ByCount_COLLECTIONS:
		typeStr = "saves"
	default:
		typeStr = "publish_time"
	}
	return typeStr
}
