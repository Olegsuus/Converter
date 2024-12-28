package handlers

func extensionFor(outputFormat string) string {
	switch outputFormat {
	case "pdf":
		return ".pdf"
	case "doc":
		return ".doc"
	case "docx":
		return ".docx"
	default:
		return ".unknown"
	}
}
