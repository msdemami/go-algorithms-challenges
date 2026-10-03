package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	// تبدیل متن به زمان با الگوی ۱/۲/۲۰۰۶
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	// تبدیل متن با الگوی متنی ماه (January)
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	// مقایسه با زمان حال
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	// تبدیل متن با الگوی دارای نام روز هفته (Monday)
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	// استخراج ساعت و بررسی شرط بعدازظهر (بین ۱۲ تا ۱۸)
	return t.Hour() >= 12 && t.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	// استفاده مجدد از تابعی که در وظیفه اول نوشتیم
	t := Schedule(date)
	// تبدیل مجدد زمان به متن با الگوی دلخواه (t.Format)
	formattedTime := t.Format("Monday, January 2, 2006, at 15:04")
	return fmt.Sprintf("You have an appointment on %s.", formattedTime)
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	// گرفتن سال فعلی
	currentYear := time.Now().Year()
	// ساخت یک تاریخ جدید: ۱۵ سپتامبرِ سالِ فعلی، ساعت ۰۰:۰۰:۰۰
	return time.Date(currentYear, time.September, 15, 0, 0, 0, 0, time.UTC)
}