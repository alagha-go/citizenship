package socket

import "citizenship/identity"


func EmitAddedID(ID identity.Identity) {
	data := identity.JsonMarshal(ID)
	Server.BroadcastToAll("newid", string(data))
}


func EmitStats(statistics Stats) {
	data := identity.JsonMarshal(statistics)
	Server.BroadcastToAll("statistics", string(data))
}