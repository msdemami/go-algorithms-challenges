package clock
import  "fmt"

type Clock struct {
    minute int 
}

func New(h, m int) Clock {
	total_minute := h * 60 + m 
    final_time := ((total_minute % 1440) + 1440) % 1440
    return Clock {minute : final_time}
}

func (c Clock) Add(m int) Clock {
	return New(0,c.minute + m) 
}

func (c Clock) Subtract(m int) Clock {
	return New(0,c.minute - m)
}

func (c Clock) String() string {
	h := c.minute / 60
    m := c.minute % 60 
    return fmt.Sprintf("%02d:%02d",h,m)
}