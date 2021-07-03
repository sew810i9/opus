package issue

type (
	// Store represents store data
	Store struct {
		Title       string
		Description string
	}

	// RequestCreate represents structure that will be given while creating an issue
	RequestCreate struct {
		ID          int
		Title       string
		Description string
		Reporter    int
		Assignee    int
		Priority    int
		Label       int
	}

	// RequestUpdate represents structure that will be given while creating an issue
	RequestUpdate struct {
		Title       string
		Description string
		Reporter    int
		Assignee    int
		Priority    int
		Label       int
	}
)
