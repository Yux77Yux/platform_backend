package test

var (
	Videos = make([]*Creation, 0, 500)
	Users  = make([]*User, 0, 500)

	RegisterOkMap = make(map[string]*Register_OK)

	LoginOkMap        = make(map[string]*Login_OK)
	UpdateAvatarOkMap = make(map[string]*Id)
	UpdateSpaceOkMap  = make(map[string]*Id)

	UploadOkMap = make(map[string]*Creation_OK)

	LoginOKMapIdInDb        = make(map[int64]*Login_OK)
	GetVideosOkMapIdInDb    = make(map[int64]*CreationInfo_OK)
	PendingVideoStatusOkMap = make(map[int64]*CreationInfo_OK)
)
