package Entity

type Recipe struct {
	Cakes []Cake ``
}

type Cake struct {
	Name        string       ``
	Time        string       ``
	Ingredients []Ingredient ``
}

type Ingredient struct {
	IngredientName  string ``
	IngredientCount string ``
	IngredientUnit  string ``
}
