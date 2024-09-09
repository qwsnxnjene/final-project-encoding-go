package encoding

import (
	"encoding/json"
	"io"
	"os"

	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	// ниже реализуйте метод
	f, err := os.Open(j.FileInput)
	if err != nil {
		return err
	}
	defer f.Close()
	text, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	var Docker models.DockerCompose

	err = json.Unmarshal(text, &Docker)
	if err != nil {
		return err
	}

	j.DockerCompose = &Docker

	textYaml, err := yaml.Marshal(j.DockerCompose)
	if err != nil {
		return err
	}

	fileYaml, err := os.Create(j.FileOutput)
	if err != nil {
		return err
	}

	defer fileYaml.Close()
	_, err = fileYaml.Write(textYaml)
	if err != nil {
		return err
	}
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	// Ниже реализуйте метод
	f, err := os.Open(y.FileInput)
	if err != nil {
		return err
	}

	var docker models.DockerCompose = models.DockerCompose{}

	defer f.Close()
	text, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(text, &docker)
	if err != nil {
		return err
	}

	y.DockerCompose = &docker
	textJson, err := json.Marshal(y.DockerCompose)
	if err != nil {
		return err
	}

	fileJson, err := os.Create(y.FileOutput)
	if err != nil {
		return err
	}

	defer fileJson.Close()
	_, err = fileJson.Write(textJson)
	if err != nil {
		return err
	}
	return nil
}
