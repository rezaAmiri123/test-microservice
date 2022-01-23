package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rezaAmiri123/test-microservice/library_service/agent"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	cli := &cli{}

	cmd := &cobra.Command{
		Use:     "library",
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
}

func setupFlags(cmd *cobra.Command) error {
	cmd.Flags().String("config-file", "", "path to config file.")
	cmd.Flags().String("http-server-addr", "", "http server address.")
	cmd.Flags().Int("http-server-port", 8280, "http server port.")
	cmd.Flags().String("grpc-server-addr", "", "grpc server address.")
	cmd.Flags().Int("grpc-server-port", 8281, "grpc server port.")
	cmd.Flags().String("auth-grpc-server-addr", "user_service", "auth grpc server address.")
	cmd.Flags().Int("auth-grpc-server-port", 8181, "auth grpc server port.")
	cmd.Flags().String("database-type", "pgx", "database type like mysql.")
	cmd.Flags().String("database-name", "auth_db", "database name.")
	cmd.Flags().String("database-username", "postgres", "database username.")
	cmd.Flags().String("database-password", "postgres", "database password.")
	cmd.Flags().String("database-host", "postgesql", "database host address.")
	cmd.Flags().String("database-port", "5432", "database host address.")
	cmd.Flags().Bool("tracer-enable", true, "tracer enable mode.")
	cmd.Flags().String("tracer-service-name", "library_service", "tracer service name.")
	cmd.Flags().Bool("tracer-spans", true, "tracer spans.")
	cmd.Flags().String("tracer-host-port", "jaeger:6831", "tracer host address.")
	cmd.Flags().String("metric-service-name", "library_service", "metric service name")
	cmd.Flags().String("metric-service-host-port", ":8001", "metric service host port")
	cmd.Flags().StringSlice("kafka-service-brokers", []string{"kafka1:9092"}, "kafka service brokers")
	cmd.Flags().String("kafka-service-group-id", "library_microservice_consumer", "metric service host port")
	cmd.Flags().Bool("kafka-service-init-topics", true, "metric service host port")

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

	c.cfg.HttpServerAddr = viper.GetString("http-server-addr")
	c.cfg.HttpServerPort = viper.GetInt("http-server-port")
	c.cfg.GRPCAuthClientAddr = viper.GetString("auth-grpc-server-addr")
	c.cfg.GRPCAuthClientPort = viper.GetInt("auth-grpc-server-port")
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
	// c.cfg.KafkaConfig.Kafka.Brokers = viper.GetStringSlice("kafka-service-brokers")
	// c.cfg.KafkaConfig.Kafka.GroupID = viper.GetString("kafka-service-group-id")
	// c.cfg.KafkaConfig.Kafka.InitTopics = viper.GetBool("kafka-service-init-topics")

	// c.cfg.KafkaConfig.KafkaTopics.UserCreate.TopicName = kafka.CreateUserTopic
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
