package JsonConverter

import (
	"errors"
	"github.com/superhacker1999/ex00/Entity"
)

type JsonRecipe struct {
	Cakes []jsonCake `json:"cake"`
}

type jsonCake struct {
	Name        string           `json:"name"`
	Time        string           `json:"time"`
	Ingredients []jsonIngredient `json:"ingredients"`
}

type jsonIngredient struct {
	IngredientName  string `json:"ingredient_name"`
	IngredientCount string `json:"ingredient_count"`
	IngredientUnit  string `json:"ingredient_unit,omitempty"`
}

func ConvertJsonRecipeToRecipe(jsonRecipe JsonRecipe) (Entity.Recipe, error) {
	recipe := Entity.Recipe{}
	for _, jsonCake := range jsonRecipe.Cakes {
		cake := Entity.Cake{
			Name:        jsonCake.Name,
			Time:        jsonCake.Time,
			Ingredients: []Entity.Ingredient{},
		}
		for _, jsonIngredient := range jsonCake.Ingredients {
			ingredient := Entity.Ingredient{
				IngredientName:  jsonIngredient.IngredientName,
				IngredientCount: jsonIngredient.IngredientCount,
				IngredientUnit:  jsonIngredient.IngredientUnit,
			}
			cake.Ingredients = append(cake.Ingredients, ingredient)
		}
		recipe.Cakes = append(recipe.Cakes, cake)
	}
	if len(recipe.Cakes) < 1 {
		return Entity.Recipe{}, errors.New("There is no cakes in the recipe")
	}
	return recipe, nil
}

func ConvertRecipeToJsonRecipe(recipe Entity.Recipe) (JsonRecipe, error) {
	jsonRecipe := JsonRecipe{}
	for _, cake := range recipe.Cakes {
		jsonCake := jsonCake{
			Name:        cake.Name,
			Time:        cake.Time,
			Ingredients: []jsonIngredient{},
		}
		for _, ingredient := range cake.Ingredients {
			jsonIngredient := jsonIngredient{
				IngredientName:  ingredient.IngredientName,
				IngredientCount: ingredient.IngredientCount,
				IngredientUnit:  ingredient.IngredientUnit,
			}
			jsonCake.Ingredients = append(jsonCake.Ingredients, jsonIngredient)
		}
		jsonRecipe.Cakes = append(jsonRecipe.Cakes, jsonCake)
	}
	if len(jsonRecipe.Cakes) < 1 {
		return JsonRecipe{}, errors.New("There is no cakes in the recipe")
	}
	return jsonRecipe, nil
}
