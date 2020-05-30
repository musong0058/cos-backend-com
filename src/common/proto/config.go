package proto

type CommonEnvConfig struct {
	ChainAddr string       `conf:"chain_addr"`
	TmpDir    string       `conf:"tmp_dir"`
	PG        PGConfig     `conf:"pg"`
	Redis     RedisConfig  `conf:"redis"`
	Jaeger    JaegerConfig `conf:"jaeger"`
}

type ServiceEndpoint struct {
	Account      string `conf:"account"`
	Cores        string `conf:"cores"`
	Notification string `conf:"notification"`
	Eth          string `conf:"eth"`
	Web3         string `conf:"web3"`
}

type PGConfig struct {
	Master  string `conf:"master"`
	MaxOpen int    `conf:"max_open"`
	MaxIdle int    `conf:"max_idle"`
}

type RedisConfig struct {
	Addr     string `conf:"addr"`
	Secret   string `conf:"secret"`
	PoolSize int    `conf:"pool_size"`
}

type JaegerConfig struct {
	ServiceName string `conf:"service_name"`
	AgentAddr   string `conf:"agent_addr"`
	User        string `conf:"user"`
	Password    string `conf:"password"`
	Tags        string `conf:"tags"`
}
