package main

import (
	"gopkg.in/macaron.v1"

	"log"
	"os"
	"strconv"

	config "bitbucket.org/novatrixbr/cryptoshortner/conf"
	conf "bitbucket.org/novatrixbr/cryptoshortner/conf/app"
	"bitbucket.org/novatrixbr/cryptoshortner/lib/auth"
)

// application entrypoint
func main() {
	app := macaron.New()
	conf.SetupMiddlewares(app)
	conf.SetupRoutes(app)
	auth.LoadAllowedApps()
	/*
		Generated using http://www.kammerl.de/ascii/AsciiSignature.php - (Font: 'starwars')
		All signatures are made with FIGlet (c) 1991, 1993, 1994 Glenn Chappell and Ian Chai
		All fonts are taken from figlet.org and jave.de.
		Please check for Font Credits the figlet font database!
		Figlet Frontend - Written by Julius Kammerl - 2005
	*/
	log.Println(".___  ___.  _______ .______        ______  __    __  .______       __   __    __       _______.     ___       ___  ")
	log.Println("|   \\/   | |   ____||   _  \\      /      ||  |  |  | |   _  \\     |  | |  |  |  |     /       |    / _ \\     / _ ")
	log.Println("|  \\  /  | |  |__   |  |_)  |    |  ,----'|  |  |  | |  |_)  |    |  | |  |  |  |    |   (----`   | | | |   | (_) |")
	log.Println("|  |\\/|  | |   __|  |      /     |  |     |  |  |  | |      /     |  | |  |  |  |     \\   \\       | | | |    > _ < ")
	log.Println("|  |  |  | |  |____ |  |\\  \\----.|  `----.|  `--'  | |  |\\  \\----.|  | |  `--'  | .----)   |      | |_| |  _| (_) |")
	log.Println("|__|  |__| |_______|| _| `._____| \\______| \\______/  | _| `._____||__|  \\______/  |_______/        \\___/  (__)___/ ")

	app.Run(port())
}

// configure http port
func port() int {
	port, err := config.Cfg.Section("").Key("http_port").Int()
	if err != nil {
		log.Fatal(err)
	}

	if forceLocal, _ := config.Cfg.Section("").Key("force_local_http_port").Bool(); forceLocal == false {
		if i, err := strconv.Atoi(os.Getenv("PORT")); err == nil {
			port = i
		}
	}

	return port
}
