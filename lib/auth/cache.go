package auth

//LoadAllowedApps loads allowed external applications to interact with cryptoshortener
func LoadAllowedApps() {
	app := new(App)
	app.Id = "1"
	app.Name = "Admin Master Program"

	credentials := Oauth{"677574785fb264ee1c5cd87cc6593f4e5700f751bf9901096e6c18951fbd15f643712b8d1ff0", "57a34531a4241d44502dd15e"}

	DB[credentials] = app
}
