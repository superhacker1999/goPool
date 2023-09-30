package JsonReader

import (
	"encoding/json"
	"github.com/superhacker1999/ex00/Converters/JsonConverter"
	"github.com/superhacker1999/ex00/Entity"
)

type JsonReader struct {
}

func (reader JsonReader) Read(stream []byte) (Entity.Recipe, error) {
	var jsonRec JsonConverter.JsonRecipe

	err := json.Unmarshal(stream, &jsonRec)
	if err != nil {
		return Entity.Recipe{}, err
	}
	return JsonConverter.ConvertJsonRecipeToRecipe(jsonRec)
}
