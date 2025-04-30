package uid

import (
	"fmt"
	"go-actions/ga/app/config"
	"strings"
)

type ResourceUid struct {
	namespace string
	resource string
	uid string
}

func defaultResourceUid(config *config.GlobalConfig) *ResourceUid {
	return &ResourceUid{
		namespace: "core",
		resource: "",
		uid: config.UidGenerator.GenerateUid(),
	}
}

func NewResourceUid(config *config.GlobalConfig, opts ...ResourceUidOption) *ResourceUid {
	resourceUid := defaultResourceUid(config)
	for _,opt := range opts {
		opt(resourceUid)
	}
	return resourceUid
}

func (ru *ResourceUid) getUidValue(resourceType, resourceName string) string {
	return strings.ToLower(fmt.Sprintf("ga:%s:%s:%s:%s:%s", ru.namespace, ru.resource, ru.uid, resourceType, resourceName))
}

func (ru *ResourceUid) GetString() string {
  return ru.getUidValue("", "")
}

func (ru *ResourceUid) GetSecondaryUid(resourceType, resourceName string) string {
	return ru.getUidValue(resourceType, resourceName)
}