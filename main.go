package main

import (
	"github.com/MasDev-12/mechta.testapi/config"
	"github.com/MasDev-12/mechta.testapi/servers"
	"log"
)

//TIP To run your code, right-click the code and select <b>Run</b>. Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.

func main() {
	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Recovered from panic: %v", r)
		}
	}()
	dbSetting, err := config.LoadSettingsDb("tsconfig.json")
	if err != nil {
		panic(err)
	}
	serverSetting, err := config.LoadSettingServer("tsconfig.json")
	if err != nil {
		panic(err)
	}
	argon2Setting, err := config.LoadSettingArgon2("tsconfig.json")
	if err != nil {
		panic(err)
	}
	swaggerSetting, err := config.LoadSettingsSwagger("tsconfig.json")
	if err != nil {
		panic(err)
	}
	restServer := servers.NewRestServer(serverSetting, dbSetting, argon2Setting, swaggerSetting)
	if err := restServer.Start(); err != nil {
		log.Fatalf("Failed to start REST server: %v", err)
	}
}

//TIP See GoLand help at <a href="https://www.jetbrains.com/help/go/">jetbrains.com/help/go/</a>.
// Also, you can try interactive lessons for GoLand by selecting 'Help | Learn IDE Features' from the main menu.
