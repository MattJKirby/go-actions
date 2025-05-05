package uid

type Reference struct {
	Uid    *ResourceUid `json:"uid"`
	Source *ResourceUid `json:"source"`
	Target *ResourceUid `json:"target"`
}

func NewReference(source *ResourceUid, target *ResourceUid, opts ...ResourceUidOption) *Reference {
	return &Reference{
		Uid:    NewResourceUid(append(opts, WithResource("Ref"))...),
		Source: source,
		Target: target,
	}
}
