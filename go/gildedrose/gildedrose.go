package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		updateItemQuality(items[i])
	}
}

func (i *Item) decreaseQuality(q int) {
	if i.Quality > 0 {
		i.Quality = i.Quality - q
	}
}

func (i *Item) decreaseSellIn(s int) {
	i.SellIn = i.SellIn - s
}

func (i *Item) increaseQuality(q int) {
	if i.Quality < 50 {
		i.Quality = i.Quality + q
	}
}

func updateItemQuality(item *Item) {
	if item.Name != "Aged Brie" && item.Name != "Backstage passes to a TAFKAL80ETC concert" &&
		item.Name != "Sulfuras, Hand of Ragnaros" {
		item.decreaseQuality(1)
	}
	if item.Name == "Aged Brie" || item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		item.increaseQuality(1)
	}
	if item.SellIn < 11 && item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		item.increaseQuality(1)
	}
	if item.SellIn < 6 && item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		item.increaseQuality(1)
	}
	if item.Name != "Sulfuras, Hand of Ragnaros" {
		item.decreaseSellIn(1)
	}

	if item.SellIn >= 0 {
		return
	}
	if item.Name == "Aged Brie" {
		item.increaseQuality(1)
		return
	}
	if item.Name != "Aged Brie" && item.Name != "Backstage passes to a TAFKAL80ETC concert" &&
		item.Name != "Sulfuras, Hand of Ragnaros" {
		item.decreaseQuality(1)
		return
	}
	if item.Name == "Backstage passes to a TAFKAL80ETC concert" {
		item.decreaseQuality(item.Quality)
	}
}
