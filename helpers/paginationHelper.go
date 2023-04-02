package helpers

func CalculateOffset(page int, perPage int) int {
	offset := (page - 1) * perPage
	return offset
}
