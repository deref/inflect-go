package inflect

var DefaultInflector = NewInflector()

func init() {
	DefaultInflector.AddAcronym("id")
	DefaultInflector.AddAcronym("cpu")
}

// KebabToPublic converts a kebab-case string to a PublicGoIdentifier.
func KebabToPublic(s string) string {
	return KebabToGo(true, s)
}

// KebabToPrivate converts a kebab-case string to a privateGoIdentifier.
func KebabToPrivate(s string) string {
	return KebabToGo(false, s)
}

func KebabToGo(public bool, s string) string {
	return DefaultInflector.KebabToGo(public, s)
}

func KebabToJSVar(s string) string {
	return DefaultInflector.KebabToJS(false, s)
}
