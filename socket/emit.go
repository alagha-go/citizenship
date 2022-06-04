package socket

import "citizenship/identity"


func EmitAddedID(ID identity.Identity) {
	data := identity.JsonMarshal(ID)
	Server.BroadcastToAll("newid", string(data))
}