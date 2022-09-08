package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rezaAmiri123/test-microservice/pkg/auth/tls"
	"github.com/rezaAmiri123/test-microservice/wallet_service/internal/agent"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	cli := &cli{}

	cmd := &cobra.Command{
		Use:     "wallet",
		PreRunE: cli.setupConfig,
		RunE:    cli.run,
	}
	if err := setupFlags(cmd); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Execute(); err != nil {
		log.Fatal(err)
	}
}

type cli struct {
	cfg cfg
}

type cfg struct {
	agent.Config
	GrpcServerTLSConfig tls.TLSConfig
}

func setupFlags(cmd *cobra.Command) error {
	cmd.Flags().String("config-file", "", "path to config file.")
	//cmd.Flags().String("http-server-addr", "", "http server address.")
	//cmd.Flags().Int("http-server-port", 8280, "http server port.")
	cmd.Flags().String("grpc-server-addr", "", "grpc server address.")
	cmd.Flags().Int("grpc-server-port", 8581, "grpc server port.")
	//cmd.Flags().String("auth-grpc-server-addr", "user_service", "auth grpc server address.")
	//cmd.Flags().Int("auth-grpc-server-port", 8181, "auth grpc server port.")
	cmd.Flags().String("database-type", "pgx", "database type like mysql.")
	cmd.Flags().String("database-name", "auth_db", "database name.")
	cmd.Flags().String("database-username", "postgres", "database username.")
	cmd.Flags().String("database-password", "postgres", "database password.")
	cmd.Flags().String("database-host", "api_postgesql", "database host address.")
	cmd.Flags().String("database-port", "5432", "database host address.")
	cmd.Flags().Bool("tracer-enable", true, "tracer enable mode.")
	cmd.Flags().String("tracer-service-name", "message_service", "tracer service name.")
	cmd.Flags().Bool("tracer-spans", true, "tracer spans.")
	cmd.Flags().String("tracer-host-port", "jaeger:6831", "tracer host address.")
	cmd.Flags().String("metric-service-name", "message_service", "metric service name")
	cmd.Flags().String("metric-service-host-port", ":8001", "metric service host port")

	cmd.Flags().String("grpc-server-tls-cert-file", "", "Path to server tls cert.")
	cmd.Flags().String("grpc-server-tls-key-file", "", "Path to server tls key.")
	cmd.Flags().String("grpc-server-tls-ca-file", "", "Path to server certificate authority.")

	return viper.BindPFlags(cmd.Flags())
}

func (c *cli) setupConfig(cmd *cobra.Command, args []string) error {
	var err error

	//c.cfg.HttpServerAddr = viper.GetString("http-server-addr")
	//c.cfg.HttpServerPort = viper.GetInt("http-server-port")
	//c.cfg.GRPCAuthClientAddr = viper.GetString("auth-grpc-server-addr")
	//c.cfg.GRPCAuthClientPort = viper.GetInt("auth-grpc-server-port")
	c.cfg.GRPCServerAddr = viper.GetString("grpc-server-addr")
	c.cfg.GRPCServerPort = viper.GetInt("grpc-server-port")
	c.cfg.DBConfig.Driver = viper.GetString("database-type")
	c.cfg.DBConfig.DBName = viper.GetString("database-name")
	c.cfg.DBConfig.User = viper.GetString("database-username")
	c.cfg.DBConfig.Password = viper.GetString("database-password")
	c.cfg.DBConfig.Host = viper.GetString("database-host")
	c.cfg.DBConfig.Port = viper.GetString("database-port")
	c.cfg.TracerConfig.Enable = viper.GetBool("tracer-enable")
	c.cfg.TracerConfig.ServiceName = viper.GetString("tracer-service-name")
	c.cfg.TracerConfig.LogSpans = viper.GetBool("tracer-spans")
	c.cfg.TracerConfig.HostPort = viper.GetString("tracer-host-port")
	c.cfg.MetricConfig.ServiceName = viper.GetString("metric-service-name")
	c.cfg.MetricConfig.ServiceHostPort = viper.GetString("metric-service-host-port")

	c.cfg.GrpcServerTLSConfig.CertFile = viper.GetString("grpc-server-tls-cert-file")
	c.cfg.GrpcServerTLSConfig.KeyFile = viper.GetString("grpc-server-tls-key-file")
	c.cfg.GrpcServerTLSConfig.CAFile = viper.GetString("grpc-server-tls-ca-file")

	if c.cfg.GrpcServerTLSConfig.CertFile != "" &&
		c.cfg.GrpcServerTLSConfig.KeyFile != "" {
		c.cfg.GrpcServerTLSConfig.Server = true
		c.cfg.Config.ServerTLSConfig, err = tls.SetupTLSConfig(
			c.cfg.GrpcServerTLSConfig,
		)
		if err != nil {
			return err
		}
	}

	viper.AutomaticEnv()
	configFile, err := cmd.Flags().GetString("config-file")
	if err != nil {
		return err
	}
	viper.SetConfigFile(configFile)
	if err = viper.ReadInConfig(); err != nil {
		// ti's ok if config file doesn't exist
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			return err
		}
	}
	err = viper.Unmarshal(c.cfg)
	return nil
}

func (c *cli) run(cmd *cobra.Command, args []string) error {
	var err error
	ag, err := agent.NewAgent(c.cfg.Config)
	if err != nil {
		return err
	}
	sigc := make(chan os.Signal, 1)
	signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
	<-sigc
	return ag.Shutdown()
}
