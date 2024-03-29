package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rezaAmiri123/test-microservice/api_service/agent"
	"github.com/rezaAmiri123/test-microservice/pkg/auth/tls"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	cli := &cli{}

	cmd := &cobra.Command{
		Use:     "api",
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
	AuthGrpcClientTLSConfig    tls.TLSConfig
	LibraryGrpcClientTLSConfig tls.TLSConfig
	MessageGrpcClientTLSConfig tls.TLSConfig
}

func setupFlags(cmd *cobra.Command) error {
	cmd.Flags().String("config-file", "", "path to config file.")
	cmd.Flags().Bool("debug", false, "debug service")
	cmd.Flags().String("http-server-addr", "", "http server address.")
	cmd.Flags().Int("http-server-port", 8380, "http server port.")
	cmd.Flags().String("grpc-server-addr", "", "grpc server address.")
	//cmd.Flags().Int("grpc-server-port", 8281, "grpc server port.")
	cmd.Flags().String("auth-grpc-server-addr", "user_service", "auth grpc server address.")
	cmd.Flags().Int("auth-grpc-server-port", 8181, "auth grpc server port.")
	cmd.Flags().String("library-grpc-server-addr", "library_service", "library grpc server address.")
	cmd.Flags().Int("library-grpc-server-port", 8281, "library grpc server port.")
	cmd.Flags().String("message-grpc-server-addr", "message_service", "message grpc server address.")
	cmd.Flags().Int("message-grpc-server-port", 8481, "message grpc server port.")
	//cmd.Flags().String("database-type", "pgx", "database type like mysql.")
	//cmd.Flags().String("database-name", "auth_db", "database name.")
	//cmd.Flags().String("database-username", "postgres", "database username.")
	//cmd.Flags().String("database-password", "postgres", "database password.")
	//cmd.Flags().String("database-host", "postgesql", "database host address.")
	//cmd.Flags().String("database-port", "5432", "database host address.")
	cmd.Flags().Bool("tracer-enable", true, "tracer enable mode.")
	cmd.Flags().String("tracer-service-name", "api_service", "tracer service name.")
	cmd.Flags().Bool("tracer-spans", true, "tracer spans.")
	cmd.Flags().String("tracer-host-port", "jaeger:6831", "tracer host address.")
	cmd.Flags().String("metric-service-name", "api_service", "metric service name")
	cmd.Flags().String("metric-service-host-port", ":8001", "metric service host port")
	cmd.Flags().StringSlice("kafka-service-brokers", []string{"kafka:9092"}, "kafka service brokers")
	cmd.Flags().String("kafka-service-group-id", "api_microservice_consumer", "metric service host port")
	cmd.Flags().Bool("kafka-service-init-topics", true, "metric service host port")

	cmd.Flags().String("auth-grpc-client-tls-cert-file", "", "Path to server tls cert.")
	cmd.Flags().String("auth-grpc-client-tls-key-file", "", "Path to server tls key.")
	cmd.Flags().String("auth-grpc-client-tls-ca-file", "", "Path to server certificate authority.")

	cmd.Flags().String("library-grpc-client-tls-cert-file", "", "Path to server tls cert.")
	cmd.Flags().String("library-grpc-client-tls-key-file", "", "Path to server tls key.")
	cmd.Flags().String("library-grpc-client-tls-ca-file", "", "Path to server certificate authority.")

	cmd.Flags().String("message-grpc-client-tls-cert-file", "", "Path to server tls cert.")
	cmd.Flags().String("message-grpc-client-tls-key-file", "", "Path to server tls key.")
	cmd.Flags().String("message-grpc-client-tls-ca-file", "", "Path to server certificate authority.")

	return viper.BindPFlags(cmd.Flags())
}

func (c *cli) setupConfig(cmd *cobra.Command, args []string) error {
	var err error

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
	c.cfg.Debug = viper.GetBool("debug")
	c.cfg.HttpServerAddr = viper.GetString("http-server-addr")
	c.cfg.HttpServerPort = viper.GetInt("http-server-port")
	c.cfg.GRPCAuthClientAddr = viper.GetString("auth-grpc-server-addr")
	c.cfg.GRPCAuthClientPort = viper.GetInt("auth-grpc-server-port")
	c.cfg.GRPCLibraryClientAddr = viper.GetString("library-grpc-server-addr")
	c.cfg.GRPCLibraryClientPort = viper.GetInt("library-grpc-server-port")
	c.cfg.GRPCMessageClientAddr = viper.GetString("message-grpc-server-addr")
	c.cfg.GRPCMessageClientPort = viper.GetInt("message-grpc-server-port")

	//c.cfg.GRPCServerAddr = viper.GetString("grpc-server-addr")
	//c.cfg.GRPCServerPort = viper.GetInt("grpc-server-port")
	//c.cfg.DBConfig.Driver = viper.GetString("database-type")
	//c.cfg.DBConfig.DBName = viper.GetString("database-name")
	//c.cfg.DBConfig.User = viper.GetString("database-username")
	//c.cfg.DBConfig.Password = viper.GetString("database-password")
	//c.cfg.DBConfig.Host = viper.GetString("database-host")
	//c.cfg.DBConfig.Port = viper.GetString("database-port")
	c.cfg.TracerConfig.Enable = viper.GetBool("tracer-enable")
	c.cfg.TracerConfig.ServiceName = viper.GetString("tracer-service-name")
	c.cfg.TracerConfig.LogSpans = viper.GetBool("tracer-spans")
	c.cfg.TracerConfig.HostPort = viper.GetString("tracer-host-port")
	c.cfg.MetricConfig.ServiceName = viper.GetString("metric-service-name")
	c.cfg.MetricConfig.ServiceHostPort = viper.GetString("metric-service-host-port")
	c.cfg.KafkaConfig.Brokers = viper.GetStringSlice("kafka-service-brokers")
	c.cfg.KafkaConfig.GroupID = viper.GetString("kafka-service-group-id")
	c.cfg.KafkaConfig.InitTopics = viper.GetBool("kafka-service-init-topics")

	// c.cfg.KafkaConfig.KafkaTopics.UserCreate.TopicName = kafka.CreateUserTopic

	c.cfg.AuthGrpcClientTLSConfig.CertFile = viper.GetString("auth-grpc-client-tls-cert-file")
	c.cfg.AuthGrpcClientTLSConfig.KeyFile = viper.GetString("auth-grpc-client-tls-key-file")
	c.cfg.AuthGrpcClientTLSConfig.CAFile = viper.GetString("auth-grpc-client-tls-ca-file")

	if c.cfg.AuthGrpcClientTLSConfig.CertFile != "" &&
		c.cfg.AuthGrpcClientTLSConfig.KeyFile != "" {
		c.cfg.AuthGrpcClientTLSConfig.Server = true
		c.cfg.Config.GRPCAuthClientTLSConfig, err = tls.SetupTLSConfig(
			c.cfg.AuthGrpcClientTLSConfig,
		)
		if err != nil {
			return err
		}
	}

	c.cfg.LibraryGrpcClientTLSConfig.CertFile = viper.GetString("library-grpc-client-tls-cert-file")
	c.cfg.LibraryGrpcClientTLSConfig.KeyFile = viper.GetString("library-grpc-client-tls-key-file")
	c.cfg.LibraryGrpcClientTLSConfig.CAFile = viper.GetString("library-grpc-client-tls-ca-file")

	if c.cfg.LibraryGrpcClientTLSConfig.CertFile != "" &&
		c.cfg.LibraryGrpcClientTLSConfig.KeyFile != "" {
		c.cfg.LibraryGrpcClientTLSConfig.Server = true
		c.cfg.Config.GRPCLibraryClientTLSConfig, err = tls.SetupTLSConfig(
			c.cfg.LibraryGrpcClientTLSConfig,
		)
		if err != nil {
			return err
		}
	}

	c.cfg.MessageGrpcClientTLSConfig.CertFile = viper.GetString("message-grpc-client-tls-cert-file")
	c.cfg.MessageGrpcClientTLSConfig.KeyFile = viper.GetString("message-grpc-client-tls-key-file")
	c.cfg.MessageGrpcClientTLSConfig.CAFile = viper.GetString("message-grpc-client-tls-ca-file")

	if c.cfg.MessageGrpcClientTLSConfig.CertFile != "" &&
		c.cfg.MessageGrpcClientTLSConfig.KeyFile != "" {
		c.cfg.MessageGrpcClientTLSConfig.Server = true
		c.cfg.Config.GRPCMessageClientTLSConfig, err = tls.SetupTLSConfig(
			c.cfg.MessageGrpcClientTLSConfig,
		)
		if err != nil {
			return err
		}
	}

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
