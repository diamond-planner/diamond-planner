package dp

type CommunityServiceItem struct {
	Club    Club
	Entries []ServiceEntry
}

type CommunityServiceData struct {
	User   User
	Season int
	Items  []CommunityServiceItem
}

func getCommunityServiceData(user User, season int, clubs []Club) CommunityServiceData {
	serviceData := CommunityServiceData{
		User:   user,
		Season: season,
	}

	// TODO: fill items

	return serviceData
}
