package luhn
import   "strings"

func Valid(id string) bool {
	cleanid := strings.ReplaceAll(id, " ", "")
	if len(cleanid) <= 1 {
		return false
	}

	sum := 0
	
	// می‌پرسیم آیا طول کل شماره زوج است؟ (باقیمانده‌‌اش بر 2 صفر می‌شود؟)
	isEvenLength := len(cleanid)%2 == 0

	// حلقه‌ی کلاسیک رو به جلو (دقیقا مثل کدی که خودت نوشته بودی)
	for i := 0; i < len(cleanid); i++ {
		
		// فیلتر امنیتی: اگر حرف یا علامتی داخلش بود نامعتبر است
		if cleanid[i] < '0' || cleanid[i] > '9' {
			return false
		}

		// تبدیل کاراکتر به عدد ریاضی
		num := int(cleanid[i] - '0')

		// تشخیص اینکه آیا الان نوبت دو برابر کردن این عدد هست یا نه
		// اگر طول شماره زوج است و ایندکسِ ما هم زوج است -> دو برابر کن
		// یا اگر طول شماره فرد است و ایندکسِ ما هم فرد است -> دو برابر کن
		if (isEvenLength && i%2 == 0) || (!isEvenLength && i%2 != 0) {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}

		// انداختن عدد نهایی در قلک
		sum += num
	}

	// بررسی بخش‌پذیری بر 10
	return sum%10 == 0
}