package socket

import (
	"citizenship/identity"
	"strconv"

	socketio "github.com/ambelovsky/gosf-socketio"
)


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

func OnQuerry(_ *socketio.Channel, querry string) interface{} {
	_, err := strconv.Atoi(querry)
	if err != nil {
		return QuerryIDByName(querry)
	}else if len(querry) < 8 {
		return nil
	}else {
		return QuerryIDByIDName(querry)
	}
}


func QuerryIDByIDName(id string) interface{} {
	var IDs []identity.Identity
	for _, ID := range identity.IDs {
		if ID.IDNumber == id {
			IDs = append(IDs, ID)
			return identity.JsonMarshal(IDs)
		}
	}
	return IDs
}