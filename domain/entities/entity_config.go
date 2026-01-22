package entities

type Server struct {
	// Port is the server host port
	Port int `toml:"port"`

	// Host is the server host address
	Host string `toml:"host"`
}

type Database struct {
	// Host is the database host address
	Host string `toml:"host"`

	// Port is the database port
	Port string `toml:"port"`

	// Name is the database name
	Name string `toml:"name"`

	// User is the username for database connection auth
	User string `toml:"user"`

	// Password is the password for database connection auth
	Password string `toml:"password"`
}

type Paseto struct {
	// SecurityKey is the key used for token generation/validation
	SecurityKey string `toml:"paseto_security_key"`

	// UserPassSaltSecret is a key used for user password encryption
	UserPassSaltSecret string `toml:"user_pass_salt_secret"`
}

type FileStorage struct {
	// StorageFolder is the folder for storing files
	StorageFolder string `toml:"storage_folder"`
}

type SMTPConfig struct {
	Host     string `toml:"host"`
	Port     string `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
	From     string `toml:"from"`
}

func (s *SMTPConfig) Addr() string {
	return s.Host + ":" + s.Port
}

type Config struct {
	// LogDir represents the directory path of the log file
	LogDir string

	ServiceAccountKeyPath string

	// Environment is the environment where the application is running
	Environment string

	// Server is the data for server host
	Server Server

	// Database is the data for database connection
	Database Database

	// Paseto is the data for encryption
	Paseto Paseto

	// FileStorage contains data related to file storage
	FileStorage FileStorage

	SMTPConfig SMTPConfig

	IntegrationToken string `toml:"integrationToken"`
}

func (c Config) IsProduction() bool {
	return c.Environment == "production" || c.Environment == "prod"
}

func (c Config) IsHomolog() bool {
	return c.Environment == "homolog" || c.Environment == "homo"
}

func (c Config) IsLocal() bool {
	return c.Environment == "local"
}
