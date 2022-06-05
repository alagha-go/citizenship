package socket

import (
	"citizenship/identity"
	"strconv"
	"strings"

	socketio "github.com/ambelovsky/gosf-socketio"
)


func QuerryIDByName(name string) interface{} {
	Names := strings.Split(name, " ")
	var IDs []identity.Identity
	if len(Names) == 1 {
		for _, ID := range identity.IDs {
			if ID.FirstName == name {
				IDs = append(IDs, ID)
				continue
			}else if ID.Surname == name {
				IDs = append(IDs, ID)
				continue
			}else if ID.OtherName == name {
				IDs = append(IDs, ID)
				continue
			}
		}
	}else if len(Names) == 2 {
		for _, ID := range identity.IDs {
			if name == ID.Name {
				IDs = append(IDs, ID)
			}else if Names[0] == ID.FirstName && Names[1] == ID.OtherName {
				IDs = append(IDs, ID)
			}
		} 
	}else if len(Names) == 3 {
		for _, ID := range identity.IDs {
			if Names[0] == ID.FirstName && Names[1] == ID.OtherName && Names[2] == ID.Surname {
				IDs = append(IDs, ID)
			}
		}
	}
	return string(identity.JsonMarshal(IDs))
}

func OnQuerry(_ *socketio.Channel, querry string) interface{} {
	_, err := strconv.Atoi(querry)
	if err != nil {
		return QuerryIDByName(querry)
	}else {
		return QuerryIDByIDNumber(querry)
	}
}


func QuerryIDByIDNumber(id string) interface{} {
	var IDs []identity.Identity
	for _, ID := range identity.IDs {
		if ID.IDNumber == id {
			IDs = append(IDs, ID)
		}
	}
	return string(identity.JsonMarshal(IDs))
}