package pricing

func Price(stock string) float32 {

	priceMap := map[string]float32{
		"facebook": 238.23,
		"moderna":  136.20,
		"walmart":  151.25,
	}
	return priceMap[stock]
}
