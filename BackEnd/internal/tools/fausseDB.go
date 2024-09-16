package tools

import (
	"time"
)

type fausseDB struct{}

var fausseDBDetails = map[string]AppDetails{
	"0": {
		Id:          0,
		Title:       "Application 1",
		Description: "Petite Description de l'application 1",
		Link:        "/application-1",
		Image:       "image/exp-appimage.jpg",
	},
	"1": {
		Id:          1,
		Title:       "Application 2",
		Description: "Petite Description de l'application 2",
		Link:        "/application-2",
		Image:       "image/exp-appimage2.jpg",
	},
	"2": {
		Id:          2,
		Title:       "Application 3",
		Description: "Petite Description de l'application 3",
		Link:        "/application-3",
		Image:       "image/exp-appimage.jpg",
	},
}

func (d *fausseDB) GetAppData(id string) *AppDetails {
	// temps de chargement de la DB
	time.Sleep(time.Second * 1)

	var appData = AppDetails{}
	appData, ok := fausseDBDetails[id]
	if !ok {
		return nil
	}
	return &appData
}

func (d *fausseDB) SetupDatabase() error {
	return nil
}
