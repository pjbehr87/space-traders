package service

import stapi "github.com/pjbehr87/space-traders/st-api"

func WpHasMarketplace(wp stapi.Waypoint) bool {
	for _, trait := range wp.Traits {
		if trait.Symbol == "MARKETPLACE" {
			return true
		}
	}
	return false
}

func MarketHasFuel(market stapi.Market) (bool, *int) {
	for _, good := range market.TradeGoods {
		if good.Symbol == "FUEL" {
			price := int(good.PurchasePrice)
			return true, &price
		}
	}
	return false, nil
}
