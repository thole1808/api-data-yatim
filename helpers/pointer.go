package helpers

// Ptr mengembalikan pointer dari string (untuk keperluan assign ke *string)
func Ptr(s string) *string {
	return &s
}
