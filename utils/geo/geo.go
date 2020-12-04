package geo

func Format(country, region, city, space string) string {
	var geo string
	if country != "" {
		geo = country
	}
	if region != "" && region != country {
		geo += space + region
	}
	if city != "" && city != country && city != region {
		geo += space + city
	}
	return geo
}
