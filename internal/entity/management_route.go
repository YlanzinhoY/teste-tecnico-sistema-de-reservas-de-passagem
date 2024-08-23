package entity

type ManagementRoute struct {
	Id          string `json:"id"`
	RouteName   string `json:"route_name"`
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
}
