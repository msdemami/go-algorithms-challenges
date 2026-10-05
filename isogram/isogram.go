package isogram

import "strings"

func IsIsogram(word string) bool {
	// ۱. تبدیل کل کلمه به حروف کوچک تا حروف بزرگ و کوچک یکی حساب شوند
	myword := strings.ToLower(word)

	// ۲. ساختن دفترچه حضور و غیاب (نقشه)
	seen := make(map[rune]bool)

	// ۳. حرکت روی حروف کلمه با استفاده از range
	for _, char := range myword {
		
		// اگر کاراکتر ما فاصله یا خط تیره بود، از آن رد می‌شویم
		if char == ' ' || char == '-' {
			continue
		}

		// ۴. چک می‌کنیم آیا این حرف قبلاً در دفترچه تیک خورده است؟
		// نوشتن seen[char] دقیقاً معادل seen[char] == true است
		if seen[char]==true {
			return false // حرف تکراری است! پس کلمه ایزوگرام نیست
		}

		// ۵. اگر حرف جدید بود، جلوی اسمش در دفترچه تیک می‌‌زنیم
		seen[char] = true
	}

	// ۶. اگر حلقه تا آخر رفت و هیچ حرف تکراری‌ای پیدا نشد
	return true
}