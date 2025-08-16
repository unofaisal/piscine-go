package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	menu := map[string]food{
		"burger":  {preptime: 15},
		"chips":   {preptime: 10},
		"nuggets": {preptime: 12},
	}
	if v, ok := menu[order]; ok {
		return v.preptime
	}
	return 404
}
