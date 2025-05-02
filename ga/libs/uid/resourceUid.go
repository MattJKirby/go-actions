package uid

import (
	"encoding/json"
	"fmt"
	"go-actions/ga/app/config"
	"strings"
)

type ResourceUid struct {
	prefix    string
	namespace string
	resource  string
	uid       string
}

func defaultResourceUid(config *config.GlobalConfig) *ResourceUid {
	return &ResourceUid{
		prefix:    "ga",
		namespace: "core",
		resource:  "",
		uid:       config.UidGenerator.GenerateUid(),
	}
}

func NewResourceUid(config *config.GlobalConfig, opts ...ResourceUidOption) *ResourceUid {
	resourceUid := defaultResourceUid(config)
	for _, opt := range opts {
		opt(resourceUid)
	}
	return resourceUid
}

func (ru *ResourceUid) getUidValue(resourceType, resourceName string) string {
	return strings.ToLower(fmt.Sprintf("%s:%s:%s:%s:%s:%s", ru.prefix, ru.namespace, ru.resource, ru.uid, resourceType, resourceName))
}

func (ru *ResourceUid) GetString() string {
	return ru.getUidValue("", "")
}

func (ru *ResourceUid) GetSecondaryId(resourceType string) string {
	return ru.getUidValue(resourceType, "")
}

func (ru *ResourceUid) GetSecondaryUid(resourceType, resourceName string) string {
	return ru.getUidValue(resourceType, resourceName)
}

func (ru *ResourceUid) MarshalJSON() ([]byte, error) {
	return json.Marshal(ru.GetString())
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

	return nil
}
