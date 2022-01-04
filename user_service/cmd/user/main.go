package main

import (
	"github.com/rezaAmiri123/test-microservice/pkg/kafka"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rezaAmiri123/test-microservice/user_service/agent"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func main() {
	cli := &cli{}

	cmd := &cobra.Command{
		Use:     "users",
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
	cmd.Flags().Int("http-server-port", 8080, "http server port.")
	cmd.Flags().String("auth-server-addr", "localhost", "auth server address.")
	cmd.Flags().Int("auth-server-port", 8081, "auth server port.")
	//cmd.Flags().String("grpc-server-addr", "localhost", "grpc server address.")
	//cmd.Flags().Int("grpc-server-port", 8081, "grpc server port.")
	cmd.Flags().String("database-type", "mysql", "database type like mysql.")
	cmd.Flags().String("database-name", "go", "database name.")
	cmd.Flags().String("database-username", "go", "database username.")
	cmd.Flags().String("database-password", "go", "database password.")
	cmd.Flags().String("database-host", "db.local", "database host address.")
	cmd.Flags().String("database-port", "3306", "database host address.")
	cmd.Flags().Bool("tracer-enable", true, "tracer enable mode.")
	cmd.Flags().String("tracer-service-name", "users_service", "tracer service name.")
	cmd.Flags().Bool("tracer-spans", true, "tracer spans.")
	cmd.Flags().String("tracer-host-port", "jaeger:6831", "tracer host address.")
	cmd.Flags().String("metric-service-name", "users_service", "metric service name")
	cmd.Flags().String("metric-service-host-port", ":8001", "metric service host port")
	cmd.Flags().StringSlice("kafka-service-brokers", []string{"kafka1:9092"}, "kafka service brokers")
	cmd.Flags().String("kafka-service-group-id", "user_microservice_consumer", "metric service host port")
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
	//c.cfg.GRPCServerAddr = viper.GetString("grpc-server-addr")
	//c.cfg.GRPCServerPort = viper.GetInt("grpc-server-port")
	c.cfg.DBConfig.Type = viper.GetString("database-type")
	c.cfg.DBConfig.Name = viper.GetString("database-name")
	c.cfg.DBConfig.User = viper.GetString("database-username")
	c.cfg.DBConfig.Pass = viper.GetString("database-password")
	c.cfg.DBConfig.Host = viper.GetString("database-host")
	c.cfg.DBConfig.Port = viper.GetString("database-port")
	c.cfg.TracerConfig.Enable = viper.GetBool("tracer-enable")
	c.cfg.TracerConfig.ServiceName = viper.GetString("tracer-service-name")
	c.cfg.TracerConfig.LogSpans = viper.GetBool("tracer-spans")
	c.cfg.TracerConfig.HostPort = viper.GetString("tracer-host-port")
	c.cfg.MetricConfig.ServiceName = viper.GetString("metric-service-name")
	c.cfg.MetricConfig.ServiceHostPort = viper.GetString("metric-service-host-port")
	c.cfg.KafkaConfig.Kafka.Brokers = viper.GetStringSlice("kafka-service-brokers")
	c.cfg.KafkaConfig.Kafka.GroupID = viper.GetString("kafka-service-group-id")
	c.cfg.KafkaConfig.Kafka.InitTopics = viper.GetBool("kafka-service-init-topics")

	c.cfg.KafkaConfig.KafkaTopics.UserCreate.TopicName = kafka.CreateUserTopic
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
