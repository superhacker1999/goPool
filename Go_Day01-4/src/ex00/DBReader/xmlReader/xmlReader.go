package XmlReader

import (
	"encoding/xml"
	"ex00/Converters/XmlConverter"
	"ex00/Entity"
)

type XmlReader struct {
}

func (reader XmlReader) Read(stream []byte) (Entity.Recipe, error) {
	var xmlRec XmlConverter.XmlRecipe

	err := xml.Unmarshal(stream, &xmlRec)
	if err != nil {
		return Entity.Recipe{}, err
	}
	return XmlConverter.ConvertXmlRecipeToRecipe(xmlRec)
}
