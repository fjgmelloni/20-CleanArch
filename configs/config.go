package configs

import "github.com/spf13/viper"

type conf struct {
    DBDriver          string `mapstructure:"DB_DRIVER"`
    DBHost            string `mapstructure:"DB_HOST"`
    DBPort            string `mapstructure:"DB_PORT"`
    DBUser            string `mapstructure:"DB_USER"`
    DBPassword        string `mapstructure:"DB_PASSWORD"`
    DBName            string `mapstructure:"DB_NAME"`
    WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
    GRPCServerPort    string `mapstructure:"GRPC_SERVER_PORT"`
    GraphQLServerPort string `mapstructure:"GRAPHQL_SERVER_PORT"`
    RabbitMQURL      string `mapstructure:"RABBITMQ_URL"`
}

func LoadConfig(path string) (*conf, error) {
    var cfg *conf
    viper.SetConfigName("app_config")
    viper.SetConfigType("env")
    viper.AddConfigPath(path)
    viper.SetConfigFile(".env")
    viper.AutomaticEnv()
    
    err := viper.ReadInConfig()
    if err != nil {
        return nil, err 
    }
    
    cfg = &conf{} 
    err = viper.Unmarshal(cfg)
    if err != nil {
        return nil, err 
    }
    
    if cfg.DBDriver == "" {
        cfg.DBDriver = "mysql"
    }
    if cfg.WebServerPort == "" {
        cfg.WebServerPort = "8080"
    }
    if cfg.GRPCServerPort == "" {
        cfg.GRPCServerPort = "50051"
    }
    if cfg.GraphQLServerPort == "" {
        cfg.GraphQLServerPort = "8081"
    }
    if cfg.RabbitMQURL == "" {
        cfg.RabbitMQURL = "amqp://guest:guest@rabbitmq:5672/"
    }
    
    return cfg, nil
}