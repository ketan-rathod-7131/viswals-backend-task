package postgres

// DbConfig required for connection to the postgresql database
type DbConfig struct {
	ConnectionString string `json:"connection_string" yaml:"connection_string"`
	MaxConnRetries   int    `json:"max_conn_retries" yaml:"max_conn_retries"`
	MaxOpenConns     int    `json:"max_open_conns" yaml:"max_open_conns"`
	MigrationsPath   string `json:"migrations_path" yaml:"migrations_path"`
}
