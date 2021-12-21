package dto

type Client struct {
	BaseURL string
}

type Forum struct {
	ForumId      int
	ForumName    string
	TopicKeyword string
	Users        []string
}

type ForumsResponse struct {
	ForumName    string   `json:"ForumName"`
	TopicKeyword string   `json:"InterestName"`
	Users        []string `json:"UserName"`
}

type User struct {
	UserName  string
	UserType  int
	UserMail  string
	Password  string
	Interests []string
}
