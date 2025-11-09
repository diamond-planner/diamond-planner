package bsm

type Person struct {
	Type        string `json:"type"`
	ID          int    `json:"id"`
	NamePrefix  string `json:"name_prefix"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	NameSuffix  string `json:"name_suffix"`
	Gender      string `json:"gender"`
	HumanGender string `json:"human_gender"`
}
