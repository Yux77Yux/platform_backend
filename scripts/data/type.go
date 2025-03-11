package data

type UserProfile struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Bio    string `json:"bio"`
	Avatar string `json:"avatar"`

	Creations []*Creation `json:"creations"`
}
type User struct {
	Id     string `json:"id"`
	Name   string `json:"name"`
	Bio    string `json:"bio"`
	Avatar string `json:"avatar"`
}
type Creation struct {
	Id         string `json:"id"`
	Title      string `json:"title"`
	Bio        string `json:"bio"`
	Uid        string `json:"uid"`
	Src        string `json:"src"`
	Thumbnail  string `json:"thumbnail"`
	Duration   int32  `json:"duration"`
	CategoryId int32  `json:"categoryId"`
}
