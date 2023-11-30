package templates

import "strconv"

func getEndpoint(entityPlural string, verb string, id int) string {
	endpoints := map[string]string{
		"index":  "/" + entityPlural + "",
		"show":   "/" + entityPlural + "/" + strconv.Itoa(id) + "",
		"edit":   "/" + entityPlural + "/" + strconv.Itoa(id) + "/edit",
		"create": "/" + entityPlural + "/create",
		"store":  "/" + entityPlural + "",
		"update": "/" + entityPlural + "/" + strconv.Itoa(id) + "",
		"delete": "/" + entityPlural + "/" + strconv.Itoa(id) + "",
	}

	return endpoints[verb]
}
