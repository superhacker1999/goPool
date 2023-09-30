package main

import (
	"ex00/Converters/JsonConverter"
	"ex00/Converters/XmlConverter"
	"ex00/DBReader"
	"fmt"
)

func main() {
	recipe, err := DBReader.Read("recipe.xml")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("base recipe - ", recipe)
	}
	jsonRecipe, err := JsonConverter.ConvertRecipeToJsonRecipe(recipe)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("json recipe - ", jsonRecipe)
	}
	xmlRecipe, err := XmlConverter.ConvertRecipeToXmlRecipe(recipe)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("xml recipe - ", xmlRecipe)
	}
}
