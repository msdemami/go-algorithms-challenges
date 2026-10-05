package collatzconjecture


import "errors"

func CollatzConjecture(n int) (int, error) {
	// هندل کردن اعداد غیرمجاز (صفر و منفی)
	if n <= 0 {
		return 0, errors.New("n must be positive")
	}

	var i int
	
	// بخش اول حلقه خالی است چون i بالا تعریف شده
	for ; n != 1; i++ {
		if n%2 == 0 {
			n /= 2
		} else {
			n = (n * 3) + 1
		}
	}
	
	// برگرداندن هر دو خروجی مورد نیاز
	return i, nil
}