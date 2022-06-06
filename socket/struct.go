package socket


type Stats struct {
	IDNumber								int							`json:"id_number,omitempty"`
	TotalNames								int							`json:"total_names,omitempty"`
	CurrentNameIndex						int							`json:"current_name_index,omitempty"`
}