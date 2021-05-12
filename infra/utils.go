package infra

import (
	"github.com/gobuffalo/packr/v2"
)

func getPackrFile(name string) ([]byte, error) {
	box := packr.New("views", "../views")
	file, err := box.Find(name + ".html")
	if err != nil {
		return nil, err
	}

	return file, nil
}
