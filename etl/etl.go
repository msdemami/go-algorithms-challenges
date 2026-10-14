package etl


import "strings"

func Transform(in map[int][]string) map[string]int {
	// ۱. ساختن مپ جدید برای خروجی
	out := make(map[string]int)

	// ۲. حلقه زدن روی مپ ورودی (گرفتن امتیاز و لیست حروف)
	for score, letters := range in {
		
		// ۳. حلقه زدن روی تک‌تک حروف اون لیست
		for _, letter := range letters {
			
			// ۴. تبدیل حرف به حروف کوچک و ذخیره امتیاز برای اون حرف
			lowercaseLetter := strings.ToLower(letter)
			out[lowercaseLetter] = score
		}
	}

	// برگرداندن مپ جدید
	return out
}