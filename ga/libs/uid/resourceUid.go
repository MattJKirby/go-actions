package uid

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ResourceUid struct {
	prefix          string
	namespace       string
	resource        string
	uid             string
	subResourceType string
	subResourceId   string
}

func defaultResourceUid() *ResourceUid {
	return &ResourceUid{
		prefix:          "ga",
		namespace:       "core",
		resource:        "",
		uid:             "",
		subResourceType: "",
		subResourceId:   "",
	}
}

func NewResourceUid(opts ...ResourceUidOption) *ResourceUid {
	resourceUid := defaultResourceUid()
	for _, opt := range opts {
		opt(resourceUid)
	}
	return resourceUid
}

func (ru *ResourceUid) getString(prefix, ns, res, uid, subRes, subId string) string {
	return strings.ToLower(fmt.Sprintf("%s:%s:%s:%s:%s:%s", prefix, ns, res, uid, subRes, subId))
}

func (ru *ResourceUid) GetUid() string {
	return ru.getString(ru.prefix, ru.namespace, ru.resource, ru.uid, ru.subResourceType, ru.subResourceId)
}

func (ru *ResourceUid) GetBaseUid() string {
	return ru.getString(ru.prefix, ru.namespace, ru.resource, ru.uid, "", "")
}

func (ru *ResourceUid) MarshalJSON() ([]byte, error) {
	return json.Marshal(ru.GetUid())
}

func (ru *ResourceUid) UnmarshalJSON(data []byte) error {
	var uidValue string
	if err := json.Unmarshal(data, &uidValue); err != nil {
		return err
	}

	elements := strings.Split(uidValue, ":")
	const errorStub = "error unmarshalling uid"

	if len(elements) != 6 {
		return fmt.Errorf("%s: invalid format: '%s'", errorStub, uidValue)
	}

	if elements[0] != ru.prefix {
		return fmt.Errorf("%s: invalid prefix: got '%s', expected '%s'", errorStub, elements[0], ru.prefix)
	}

	if elements[1] != strings.ToLower(ru.namespace) {
		return fmt.Errorf("%s: unexpected namespace: got '%s', expected '%s'", errorStub, elements[1], ru.namespace)
	}

	if elements[2] != strings.ToLower(ru.resource) {
		return fmt.Errorf("%s: unexpected resource: got '%s', expected '%s'", errorStub, elements[2], ru.resource)
	}

	ru.uid = elements[3]
	ru.subResourceType = elements[4]
	ru.subResourceId = elements[5]

	return nil
}
