package uid

type Reference struct {
	Uid    ResourceUid `json:"uid"`
	Source ResourceUid `json:"source"`
	Target ResourceUid `json:"target"`
}

func NewReference(source ResourceUid, target ResourceUid, uidbuilder *UidBuilder) *Reference {
	return &Reference{
		Uid:    uidbuilder.WithResource("Res").Build(),
		Source: source,
		Target: target,
	}
}
