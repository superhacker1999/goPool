package DBReader

import (
	"ex00/DBReader/jsonReader"
	"ex00/DBReader/xmlReader"
	"ex00/Entity"
	"io"
	"os"
	"strings"
)

type IDBReader interface {
	Read(stream []byte) (Entity.Recipe, error)
}

func Read(filename string) (Entity.Recipe, error) {
	var r IDBReader
	if strings.HasSuffix(filename, ".json") {
		r = JsonReader.JsonReader{}
	} else if strings.HasSuffix(filename, ".xml") {
		r = XmlReader.XmlReader{}
	}

	file, err := os.Open(filename)
	if err != nil {
		return Entity.Recipe{}, err
	}
	defer file.Close()

	data, error := io.ReadAll(file)
	if error != nil {
		return Entity.Recipe{}, err
	}
	return r.Read(data)
}
