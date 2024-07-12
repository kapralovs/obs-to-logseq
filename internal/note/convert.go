package note

func ConvertLinks(byteContent []byte) ([]byte, error) {
	var (
		content                            = string(byteContent)
		updatedContent, startLink, endLink string
		isFullLink                         bool
	)

	for _, v := range content {
		if startLink == "" &&
			v != '[' && v != '|' && v != ']' {
			updatedContent += string(v)
		}

		if isFullLink && v != ']' {
			continue
		}

		if startLink == "[[" {
			if v == ']' {
				if endLink == "" || endLink == "]" {
					endLink += string(v)
				}
				updatedContent += string(v)

				continue
			}

			if v == '|' {
				isFullLink = true
				continue
			}

			updatedContent += string(v)
		}

		if v == '[' {
			switch startLink {
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
