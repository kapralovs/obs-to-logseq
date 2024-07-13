package note

func ConvertLinks(byteContent []byte) ([]byte, error) {
	var (
		content                            = string(byteContent)
		updatedContent, startLink, endLink string
		isFullLink                         bool
	)

	for _, v := range content {
		switch v {
		case '[':
			if startLink == "" || startLink == "[" {
				startLink += string(v)
				updatedContent += string(v)
			}
		case '|':
			if startLink == "[[" {
				isFullLink = true
			}
		case ']':
			if startLink == "[[" {
				if endLink == "" || endLink == "]" {
					endLink += string(v)
				}

				updatedContent += string(v)
			}
		default:
			if isFullLink {
				if endLink == "]]" {
					updatedContent += string(v)
					startLink = ""
					endLink = ""
					isFullLink = false
				}

				continue
			}

			updatedContent += string(v)
		}
	}

	return []byte(updatedContent), nil
}
