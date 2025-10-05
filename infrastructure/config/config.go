package config

type Config struct {
	Database Database
	Server   Server
	Crypt    struct {
		SymmetricKey string `toml:"symmetric_key"`
	}
}
