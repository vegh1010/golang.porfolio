package documentHelper

func Bold(text string) string {
	return `<b>` + text + `</b>`
}

func Italic(text string) string {
	return `<i>` + text + `</i>`
}

func Underline(text string) string {
	return `<u>` + text + `</u>`
}

func FontStyle(text string, isBold, isUnderline, isItalic bool) string {
	if isBold {
		text = Bold(text)
	}
	if isUnderline {
		text = Underline(text)
	}
	if isItalic {
		text = Italic(text)
	}
	return text
}
