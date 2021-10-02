package utils

func ConvertTo254(phone string) string {
	if len(phone) > 0 {
		if string(phone[0]) == "0" {
			return "254" + phone[1:]
		} else if string(phone[0]) == "+" {
			return phone[1:]
		} else {
			return phone
		}
	} else {
		return "0000000000"
	}
}
