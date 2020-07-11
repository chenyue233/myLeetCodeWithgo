package other

func pigai(money int) int {
	// 钱可购买数量
	beer := money / 2
	// 瓶盖数量
	cover := beer
	// 酒瓶数量
	bott := cover
	// 当前已购买瓶酒数量
	tmpBeer := beer
	for {
		if (cover) >= 4 {
			cover = cover - 4
			tmpBeer = tmpBeer + 1
			cover += 1
			bott += 1
		}
		if bott >= 2 {
			bott -= 2
			tmpBeer += 1
			cover += 1
			bott += 1
		}
		if cover < 4 && bott < 2 {
			break
		}
	}
	return tmpBeer
}
