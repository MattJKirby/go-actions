package uid

type Reference struct {
	Uid    *ResourceUid
	Source *ResourceUid
	Target *ResourceUid
}

func NewReference(source *ResourceUid, target *ResourceUid, opts ...ResourceUidOption) *Reference {
	return &Reference{
		Uid:    NewResourceUid(append(opts, WithResource("Ref"))...),
		Source: source,
		Target: target,
	}
}

// func (r *Reference) SourceReference() *PartialReference {
// 	return &PartialReference{
// 		ReferenceUid: r.Uid,
// 		ResourceUid: r.source,
// 	}
// }

// func (r *Reference) TargetReference() *PartialReference {
// 	return &PartialReference{
// 		ReferenceUid: r.Uid,
// 		ResourceUid: r.target,
// 	}
// }

// type PartialReference struct {
// 	ReferenceUid *ResourceUid
// 	ResourceUid *ResourceUid
// }

// func (pr PartialReference) GetResourceId() string {
// 	return pr.ReferenceUid.GetString()
// }
