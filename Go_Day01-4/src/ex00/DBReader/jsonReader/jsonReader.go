package JsonReader

import (
	"encoding/json"
	"ex00/Converters/JsonConverter"
	"ex00/Entity"
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
