package dota

type Message struct {
	HTML string `json:"html"`
	Text string `json:"text"`
	Subject string `json:"subject"`
	FromEmail string `json:"from_email"`
	FromName string `json:"from_name"`
	To []To `json:"to"`
}

type MandrillPost struct {
	Key string `json:"key"`
	Message Message `json:"message"`

}

type To struct {
	Email string `json:"email"`
	Name string `json:"name"`
}