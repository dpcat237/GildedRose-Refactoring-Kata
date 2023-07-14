package gildedrose_test

import (
	"testing"

	"github.com/emilybache/gildedrose-refactoring-kata/gildedrose"
)

func Test_UpdateQuality(t *testing.T) {
	type checkFn func(*testing.T, gildedrose.Item)
	type requestData struct {
		days uint32
		item gildedrose.Item
	}

	isCorrectResponse := func(want gildedrose.Item) checkFn {
		return func(t *testing.T, got gildedrose.Item) {
			if want.SellIn != got.SellIn {
				t.Fatalf("want SellIn: %v, got: %v", want.SellIn, got.SellIn)
			}
			if want.Quality != got.Quality {
				t.Fatalf("want Quality: %v, got: %v", want.Quality, got.Quality)
			}
		}
	}

	tests := []struct {
		name   string
		req    requestData
		checks []checkFn
	}{
		{
			name: "Decrease both - 1 day",
			req: requestData{
				days: 1,
				item: gildedrose.Item{
					Name:    "+5 Dexterity Vest",
					SellIn:  10,
					Quality: 20,
				},
			},
			checks: []checkFn{
				isCorrectResponse(gildedrose.Item{
					SellIn:  9,
					Quality: 19,
				}),
			},
		},
		{
			name: "Decrease both - 10 days",
			req: requestData{
				days: 10,
				item: gildedrose.Item{
					Name:    "+5 Dexterity Vest",
					SellIn:  10,
					Quality: 20,
				},
			},
			checks: []checkFn{
				isCorrectResponse(gildedrose.Item{
					SellIn:  0,
					Quality: 10,
				}),
			},
		},
		{
			name: "Decrease both - 22 days",
			req: requestData{
				days: 22,
				item: gildedrose.Item{
					Name:    "+5 Dexterity Vest",
					SellIn:  10,
					Quality: 20,
				},
			},
			checks: []checkFn{
				isCorrectResponse(gildedrose.Item{
					SellIn:  -12,
					Quality: 0,
				}),
			},
		},
		{
			name: "Decrease SellIn and  increase Quality - 4 days",
			req: requestData{
				days: 4,
				item: gildedrose.Item{
					Name:    "Aged Brie",
					SellIn:  2,
					Quality: 0,
				},
			},
			checks: []checkFn{
				isCorrectResponse(gildedrose.Item{
					SellIn:  -2,
					Quality: 6,
				}),
			},
		},
		{
			name: "Decrease SellIn and  increase Quality - 30 days",
			req: requestData{
				days: 30,
				item: gildedrose.Item{
					Name:    "Aged Brie",
					SellIn:  2,
					Quality: 0,
				},
			},
			checks: []checkFn{
				isCorrectResponse(gildedrose.Item{
					SellIn:  -28,
					Quality: 50,
				}),
			},
		},
		{
			name: "Decrease both - 2 days",
			req: requestData{
				days: 2,
				item: gildedrose.Item{
					Name:    "Elixir of the Mongoose",
					SellIn:  5,
					Quality: 7,
				},
			},
			checks: []checkFn{
				isCorrectResponse(gildedrose.Item{
					SellIn:  3,
					Quality: 5,
				}),
			},
		},
		{
			name: "Decrease both - 10 days",
			req: requestData{
				days: 10,
				item: gildedrose.Item{
					Name:    "Elixir of the Mongoose",
					SellIn:  5,
					Quality: 7,
				},
			},
			checks: []checkFn{
				isCorrectResponse(gildedrose.Item{
					SellIn:  -5,
					Quality: 0,
				}),
			},
		},
		{
			name: "No changes - 1 example",
			req: requestData{
				days: 2,
				item: gildedrose.Item{
					Name:    "Sulfuras, Hand of Ragnaros",
					SellIn:  0,
					Quality: 80,
				},
			},
			checks: []checkFn{
				isCorrectResponse(gildedrose.Item{
					SellIn:  0,
					Quality: 80,
				}),
			},
		},
		{
			name: "No changes - 2 example",
			req: requestData{
				days: 2,
				item: gildedrose.Item{
					Name:    "Sulfuras, Hand of Ragnaros",
					SellIn:  -1,
					Quality: 80,
				},
			},
			checks: []checkFn{
				isCorrectResponse(gildedrose.Item{
					SellIn:  -1,
					Quality: 80,
				}),
			},
		},
		{
			name: "Decrease SellIn and temporary increase Quality - 2 days",
			req: requestData{
				days: 2,
				item: gildedrose.Item{
					Name:    "Backstage passes to a TAFKAL80ETC concert",
					SellIn:  15,
					Quality: 20,
				},
			},
			checks: []checkFn{
				isCorrectResponse(gildedrose.Item{
					SellIn:  13,
					Quality: 22,
				}),
			},
		},
		{
			name: "Decrease SellIn and Quality in consequence - 16 days",
			req: requestData{
				days: 16,
				item: gildedrose.Item{
					Name:    "Backstage passes to a TAFKAL80ETC concert",
					SellIn:  15,
					Quality: 20,
				},
			},
			checks: []checkFn{
				isCorrectResponse(gildedrose.Item{
					SellIn:  -1,
					Quality: 0,
				}),
			},
		},
		{
			name: "Decrease SellIn and increase Quality - 3 days",
			req: requestData{
				days: 3,
				item: gildedrose.Item{
					Name:    "Backstage passes to a TAFKAL80ETC concert",
					SellIn:  10,
					Quality: 49,
				},
			},
			checks: []checkFn{
				isCorrectResponse(gildedrose.Item{
					SellIn:  7,
					Quality: 50,
				}),
			},
		},
		{
			name: "Decrease SellIn and Quality in consequence - 11 days",
			req: requestData{
				days: 11,
				item: gildedrose.Item{
					Name:    "Backstage passes to a TAFKAL80ETC concert",
					SellIn:  10,
					Quality: 49,
				},
			},
			checks: []checkFn{
				isCorrectResponse(gildedrose.Item{
					SellIn:  -1,
					Quality: 0,
				}),
			},
		},
		{
			name: "Decrease both - 3 day",
			req: requestData{
				days: 3,
				item: gildedrose.Item{
					Name:    "Conjured Mana Cake",
					SellIn:  3,
					Quality: 6,
				},
			},
			checks: []checkFn{
				isCorrectResponse(gildedrose.Item{
					SellIn:  0,
					Quality: 3,
				}),
			},
		},
		{
			name: "Decrease both - 8 day",
			req: requestData{
				days: 8,
				item: gildedrose.Item{
					Name:    "Conjured Mana Cake",
					SellIn:  3,
					Quality: 6,
				},
			},
			checks: []checkFn{
				isCorrectResponse(gildedrose.Item{
					SellIn:  -5,
					Quality: 0,
				}),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := tt.req.item
			for day := uint32(0); day < tt.req.days; day++ {
				gildedrose.UpdateQuality([]*gildedrose.Item{&i})
			}
			for _, ch := range tt.checks {
				ch(t, i)
			}
		})
	}
}
