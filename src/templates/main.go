package templates

import (
	"strconv"

	"github.com/fredmessias43/car-hub/src/contracts"
	"github.com/fredmessias43/car-hub/src/models"
)

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

func listToMap[V contracts.Model](modelList []V) []map[string]interface{} {
	list := models.List[V]{Value: modelList}
	return list.ToMap()
}
