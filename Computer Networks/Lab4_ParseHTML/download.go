package main

import (
	"github.com/antchfx/htmlquery"
	"github.com/mgutz/logxi/v1"
)

type DataItem struct {
	Name, Code, Type, TimeUpdated, Capitalization, Price string
}

func search() []DataItem {
	doc, err := htmlquery.LoadURL("https://investfunds.ru/crypto/")
	if err != nil{
		log.Error("failed to load page investfunds.ru", "error", err)
		return nil
	}

	namesEls := htmlquery.Find(doc, "//div[contains(@class, 'fixed_table')]//table//tbody//tr//td[position()=2]//a/text()")
	dataEls := htmlquery.Find(doc, "//div[contains(@class, 'roll_table')]//table//tbody//tr//td[position()<=5]//div/text()")
	if len(namesEls) * 5 != len(dataEls){
		log.Error("page structure invalid", "namesEls length", len(namesEls), "dataEls length", len(dataEls))
		return nil
	}

	var result []DataItem
	for i := 0; i< len(namesEls); i++{
		result = append(result, DataItem{
			Name:           namesEls[i].Data,
			Code:           dataEls[i*5].Data,
			Type:           dataEls[i*5+1].Data,
			TimeUpdated:    dataEls[i*5+2].Data,
			Capitalization: dataEls[i*5+3].Data,
			Price:          dataEls[i*5+4].Data,
		})
	}
	return result
}
