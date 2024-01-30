package note

func ConvertLinks(byteContent []byte) ([]byte, error) {
	var (
		content            = string(byteContent)
		updatedContent     string
		startLink, endLink string
		isFullLink         bool
	)

	for _, v := range content {
		if startLink == "" && v != '[' && v != '|' && v != ']' {
			updatedContent += string(v)
		}

		if isFullLink && v != ']' {
			continue
		}

		if startLink == "[[" {
			if v == ']' {
				switch endLink {
				case "":
					endLink += string(v)
					updatedContent += string(v)
				case "]":
					endLink += string(v)
					updatedContent += string(v)
				}

				continue
			}

			if v == '|' {
				isFullLink = true
				continue
			}

			updatedContent += string(v)
		}

		if v == '[' {
			switch endLink {
			case "":
				startLink += string(v)
				updatedContent += string(v)
			case "[":
				startLink += string(v)
				updatedContent += string(v)
			}
		}
	}

	return []byte(updatedContent), nil
}
