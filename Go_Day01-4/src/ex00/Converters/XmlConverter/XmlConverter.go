package XmlConverter

import (
	"errors"
	"github.com/superhacker1999/ex00/Entity"
)

type XmlRecipe struct {
	Cakes []xmlCake `xml:"cake"`
}

type xmlCake struct {
	Name        string    `xml:"name"`
	StoveTime   string    `xml:"stovetime"`
	Ingredients []xmlItem `xml:"ingredients"`
}

type xmlItem struct {
	Items []xmlIngredient `xml:"item"`
}

type xmlIngredient struct {
	ItemName  string `xml:"itemname"`
	ItemCount string `xml:"itemcount"`
	ItemUnit  string `xml:"itemunit,omitempty"`
}

func ConvertXmlRecipeToRecipe(xmlRecipe XmlRecipe) (Entity.Recipe, error) {
	recipe := Entity.Recipe{}
	for _, xmlCake := range xmlRecipe.Cakes {
		cake := Entity.Cake{
			Name:        xmlCake.Name,
			Time:        xmlCake.StoveTime,
			Ingredients: []Entity.Ingredient{},
		}
		for _, xmlItem := range xmlCake.Ingredients {
			for _, xmlIngredient := range xmlItem.Items {
				ingredient := Entity.Ingredient{
					IngredientName:  xmlIngredient.ItemName,
					IngredientCount: xmlIngredient.ItemCount,
					IngredientUnit:  xmlIngredient.ItemUnit,
				}
				cake.Ingredients = append(cake.Ingredients, ingredient)
			}
		}
		recipe.Cakes = append(recipe.Cakes, cake)
	}
	if len(recipe.Cakes) < 1 {
		return Entity.Recipe{}, errors.New("There is no cakes in the recipe")
	}
	return recipe, nil
}

func ConvertRecipeToXmlRecipe(recipe Entity.Recipe) (XmlRecipe, error) {
	xmlRecipe := XmlRecipe{}
	for _, cake := range recipe.Cakes {
		xmlCake := xmlCake{
			Name:      cake.Name,
			StoveTime: cake.Time,
			Ingredients: []xmlItem{
				{
					Items: []xmlIngredient{},
				},
			},
		}
		for _, ingredient := range cake.Ingredients {
			xmlIngredient := xmlIngredient{
				ItemName:  ingredient.IngredientName,
				ItemCount: ingredient.IngredientCount,
				ItemUnit:  ingredient.IngredientUnit,
			}
			xmlCake.Ingredients[0].Items = append(xmlCake.Ingredients[0].Items, xmlIngredient)
		}
		xmlRecipe.Cakes = append(xmlRecipe.Cakes, xmlCake)
	}
	if len(xmlRecipe.Cakes) < 1 {
		return XmlRecipe{}, errors.New("There is no cakes in the recipe")
	}
	return xmlRecipe, nil
}
