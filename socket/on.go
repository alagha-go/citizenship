package socket

import "citizenship/identity"


func QuerryIDByName(name string) interface{} {
	var IDs []identity.Identity
	for _, ID := range identity.IDs {
		if ID.FirstName == name {
			IDs = append(IDs, ID)
			continue
		}else if ID.Name == name {
			IDs = append(IDs, ID)
			continue
		}else if ID.OtherName == name {
			IDs = append(IDs, ID)
			continue
		}
	}
	return identity.JsonMarshal(IDs)
}