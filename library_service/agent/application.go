package agent

import (
	"github.com/rezaAmiri123/test-microservice/library_service/adapters"
	"github.com/rezaAmiri123/test-microservice/library_service/app"
	"github.com/rezaAmiri123/test-microservice/library_service/app/commands"
	"github.com/rezaAmiri123/test-microservice/library_service/app/queries"
)

func (a *Agent) setupApplication() error {
	repo, err := adapters.NewGORMArticleRepository(a.DBConfig)
	if err != nil {
		return err
	}
	application := &app.Application{
		Commands: app.Commands{
			CreateArticle: commands.NewCreateArticleHandler(repo),
		},
		Queries: app.Queries{
			GetArticleBySlug: queries.NewGetArticleHandler(repo),
		},
	}
	a.Application = application
	return nil
}
