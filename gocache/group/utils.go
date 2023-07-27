package group

// GetGroup returns the named group previously created with NewGroup,
//
//	or nil if there's no such group.
func GetGroup(name string) IGroup {
	mu.RLock()
	defer mu.RUnlock()
	return groups[name]
}
