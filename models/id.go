package models

//IDListContains utilities for ID list
func IDListContains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

//PlayerIDListContains utilities for ID list
func PlayerIDListContains(slice []*PlayerData, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s.ID] = struct{}{}
	}

	_, ok := set[item]
	return ok
}
