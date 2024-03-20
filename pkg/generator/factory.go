package generator

import (
	"github.com/white43/SEP401-pdf-generator/pkg/jobs"
	"log"
)

type FactoryInterface interface {
	FactoryMethod(job jobs.Job) ProcessorInterface
}

type Factory struct {
	chromiumURL string
}

func NewFactory(url string) FactoryInterface {
	return Factory{url}
}

func (f Factory) FactoryMethod(job jobs.Job) ProcessorInterface {
	switch job.Type {
	case "html":
		return NewHTMLProcessor(f.chromiumURL)
	case "url":
		return NewURLProcessor(f.chromiumURL)
	}

	log.Fatalln("Unknown job type " + job.Type)

	return nil
}
