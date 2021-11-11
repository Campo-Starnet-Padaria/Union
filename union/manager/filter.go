package manager

func (c Client) FilterByArchived(filter bool) bool {
	if filter == c.Archived {
		return true
	} else {
		return false
	}
}