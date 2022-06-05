package identity


func (ID *Identity) Exists() bool {
	for index := range IDs {
		if ID.IDNumber == IDs[index].IDNumber {
			return true
		}
	}
	return false
}