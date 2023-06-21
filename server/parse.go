package server

import (
	"github.com/evepraisal/go-evepraisal/parsers"
)

func Parse(text string) []*ParsedItem {
	result, _ := parsers.AllParser(parsers.StringToInput(text))
	parser_results := result.(*parsers.MultiParserResult).Results
	items_map := make(map[string]int64, len(parser_results))

	for _, sub_result := range parser_results {
		switch r := sub_result.(type) {
		default:
			panic("Unreachable")
		case *parsers.AssetList:
			for _, item := range r.Items {
				items_map[item.Name] += item.Quantity
			}
		case *parsers.CargoScan:
			for _, item := range r.Items {
				items_map[item.Name] += item.Quantity
			}
		case *parsers.Contract:
			for _, item := range r.Items {
				items_map[item.Name] += item.Quantity
			}
		case *parsers.EFT:
			items_map[r.Ship] += 1
			for _, item := range r.Items {
				items_map[item.Name] += item.Quantity
			}
		case *parsers.Fitting:
			for _, item := range r.Items {
				items_map[item.Name] += item.Quantity
			}
		case *parsers.Industry:
			for _, item := range r.Items {
				items_map[item.Name] += item.Quantity
			}
		case *parsers.Listing:
			for _, item := range r.Items {
				items_map[item.Name] += item.Quantity
			}
		case *parsers.LootHistory:
			for _, item := range r.Items {
				items_map[item.Name] += item.Quantity
			}
		case *parsers.PI:
			for _, item := range r.Items {
				items_map[item.Name] += item.Quantity
			}
		case *parsers.SurveyScan:
			for _, item := range r.Items {
				items_map[item.Name] += item.Quantity
			}
		case *parsers.ViewContents:
			for _, item := range r.Items {
				items_map[item.Name] += item.Quantity
			}
		case *parsers.MiningLedger:
			for _, item := range r.Items {
				items_map[item.Name] += item.Quantity
			}
		case *parsers.MoonLedger:
			for _, item := range r.Items {
				items_map[item.Name] += item.Quantity
			}
		case *parsers.HeuristicResult:
			for _, item := range r.Items {
				items_map[item.Name] += item.Quantity
			}
		case *parsers.DScan:
			for _, item := range r.Items {
				items_map[item.Name] += 1
			}
		case *parsers.Compare:
			for _, item := range r.Items {
				items_map[item.Name] += 1
			}
		case *parsers.Wallet:
			for _, item := range r.ItemizedTransactions {
				items_map[item.Name] += item.Quantity
			}
		case *parsers.Killmail:
			for _, item := range r.Dropped {
				items_map[item.Name] += item.Quantity
			}
			for _, item := range r.Destroyed {
				items_map[item.Name] += item.Quantity
			}
		}
	}

	items := make([]*ParsedItem, 0, len(items_map))

	for name, quantity := range items_map {
		if quantity > 0 {
			items = append(items, &ParsedItem{
				Name:     name,
				Quantity: quantity,
			})
		}
	}

	return items
}
