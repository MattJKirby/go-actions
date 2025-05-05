package uid

type UidBuilder struct {
	*ResourceUid
}

func NewUidBuilder() *UidBuilder {
	return &UidBuilder{
		ResourceUid: defaultResourceUid(),
	}
}

func (ub *UidBuilder) FromParent(parent ResourceUid) *UidBuilder {
	ub.ResourceUid = &parent
	return ub
}

func (ub *UidBuilder) WithNamespace(ns string) *UidBuilder {
	ub.namespace = ns
	return ub
}

func (ub *UidBuilder) WithResource(res string) *UidBuilder {
	ub.resource = res
	return ub
}

func (ub *UidBuilder) WithUid(uid string) *UidBuilder {
	ub.uid = uid
	return ub
}

func (ub *UidBuilder) WithSubResource(res string) *UidBuilder {
	ub.subResource = res
	return ub
}

func (ub *UidBuilder) WithSubResourceId(id string) *UidBuilder {
	ub.subResourceId = id
	return ub
}

func (ub *UidBuilder) Build() ResourceUid {
	uid := defaultResourceUid()
	uid.namespace = ub.namespace
	uid.resource =       ub.resource
	uid.uid = ub.uid   

	uid.subResource =   ub.subResource
	uid.subResourceId =  ub.subResourceId

	return *uid
}