package tinkoff


func serializeBool(b bool) string {
	if b {
		return "true"
	}
	return "false"
}
