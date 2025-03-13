package test

var (
	Videos = make([]*Creation, 0, 500)
	Users  = make([]*User, 0, 500)

	RegisterOkMap = make(map[string]*User)

	LoginOkMap        = make(map[string]*Login_OK)
	UpdateAvatarOkMap = make(map[string]*Id)
)
