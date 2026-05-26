package campaign

type Repository interface {
	Save(campaign *Campaign) error
}

func (r Repository) On(s string, matcher mock.argumentMatcher) {
	panic("unimplemented")
}
