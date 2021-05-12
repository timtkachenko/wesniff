package lib

type Person struct {
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	IdNumber  string `json:"idNumber"`
}
type Document struct {
	Number  string `json:"number"`
	Type    string `json:"type"`
	Country string `json:"country"`
}
type VerifficationPost struct {
	Callback   string   `json:"callback,omitempty"`
	Person     Person   `json:"person"`
	Document   Document `json:"document,omitempty"`
	VendorData string   `json:"vendorData,omitempty"`
	Lang       string   `json:"lang,omitempty"`
	Timestamp  string   `json:"timestamp"`
}
type CreatePayload struct {
	Verification VerifficationPost `json:"verification"`
}

type SessionResponse struct {
	Status       string       `json:"status"`
	Verification Verification `json:"verification"`
}
type Verification struct {
	Id           string `json:"id"`
	Url          string `json:"url"`
	VendorData   string `json:"vendorData"`
	Host         string `json:"host"`
	Status       string `json:"status"`
	SessionToken string `json:"sessionToken"`
}
type Image struct {
	Context   string    `json:"context"`
	Content   string    `json:"content"`
	Timestamp string `json:"timestamp"`
}
type UploadPayload struct {
	Image Image `json:"image"`
}
type VerificationUpdate struct {
	Status    string    `json:"status"`
	Timestamp string `json:"timestamp"`
}
type UpdatePayload struct {
	Verification VerificationUpdate `json:"verification"`
}

type UploadResponse struct {
	Status string `json:"status"`
	Image  struct {
		Context   string      `json:"context"`
		Id        string      `json:"id"`
		Name      string      `json:"name"`
		Timestamp interface{} `json:"timestamp"`
		Size      int         `json:"size"`
		Mimetype  string      `json:"mimetype"`
		Url       string      `json:"url"`
	} `json:"image"`
}
type UpdateResponse struct {
	Status       string `json:"status"`
	Verification struct {
		Id           string `json:"id"`
		Url          string `json:"url"`
		VendorData   string `json:"vendorData"`
		Host         string `json:"host"`
		Status       string `json:"status"`
		SessionToken string `json:"sessionToken"`
	} `json:"verification"`
}
