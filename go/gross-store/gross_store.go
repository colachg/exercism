package gross

// Units store the Gross Store unit measurement
func Units() map[string]int {
	return map[string]int{
		"quarter_of_a_dozen": 3,
		"half_of_a_dozen": 6,
		"dozen": 12,
		"small_gross": 120,
		"gross": 144,
		"great_gross": 1728,
	}
}

// NewBill create a new bill
func NewBill() map[string]int {
	return map[string]int{}
}

// AddItem add item to customer bill
func AddItem(bill, units map[string]int, item, unit string) bool {
	for n, _ := range units {
		if unit == n {
			bill[item] = units[n]
			return true
		}
	}
	return false
}

// RemoveItem remove item from customer bill
func RemoveItem(bill, units map[string]int, item, unit string) bool {
	_, existInBill := bill[item]
	_, existInUnits := units[unit]
	if existInBill && existInUnits {
		num := bill[item] - units[unit]
		switch {
			case num < 0: return false
			case num == 0 : delete(bill, item)
			case num > 0: bill[item] -= units[unit]
		}
		return true
	}
	return false
}

// GetItem return the quantity of item that the customer has in his/her bill
func GetItem(bill map[string]int, item string) (int, bool) {
	qty, exist := bill[item]
	return qty, exist
}
