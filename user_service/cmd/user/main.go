package main

import (
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
