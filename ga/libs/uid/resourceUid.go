package uid

import (
	"encoding/json"
	"fmt"
	"strings"
)

type ResourceUid struct {
	prefix        string
	namespace     string
	resource      string
	uid           string
	subResource   string
	subResourceId string
}

func defaultResourceUid() *ResourceUid {
	return &ResourceUid{
		prefix:        "ga",
		namespace:     "core",
		resource:      "",
		uid:           "",
		subResource:   "",
		subResourceId: "",
	}
}

func (ru *ResourceUid) getString(prefix, ns, res, uid, subRes, subId string) string {
	return strings.ToLower(fmt.Sprintf("%s:%s:%s:%s:%s:%s", prefix, ns, res, uid, subRes, subId))
}

func (ru *ResourceUid) FullUid() string {
	return ru.getString(ru.prefix, ru.namespace, ru.resource, ru.uid, ru.subResource, ru.subResourceId)
}

func (ru *ResourceUid) BaseUid() string {
	return ru.getString(ru.prefix, ru.namespace, ru.resource, ru.uid, "", "")
}

func (ru *ResourceUid) MarshalJSON() ([]byte, error) {
	return json.Marshal(ru.FullUid())
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

	ru.resource = elements[2]
	ru.uid = elements[3]
	ru.subResource = elements[4]
	ru.subResourceId = elements[5]

	return nil
}
