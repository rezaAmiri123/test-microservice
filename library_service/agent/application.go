package agent

import (
	"github.com/rezaAmiri123/test-microservice/library_service/adapters/pg"
	"github.com/rezaAmiri123/test-microservice/library_service/app"
	"github.com/rezaAmiri123/test-microservice/library_service/app/commands"
	"github.com/rezaAmiri123/test-microservice/library_service/app/queries"
	"github.com/rezaAmiri123/test-microservice/pkg/db/postgres"
)

func (a *Agent) setupApplication() error {
	dbConn, err := postgres.NewPsqlDB(a.DBConfig)
	if err != nil {
		return err
	}

	//repo, err := adapters.NewGORMArticleRepository(a.DBConfig)
	repo := pg.NewPGArticleRepository(dbConn)

	application := &app.Application{
		Commands: app.Commands{
			CreateArticle: commands.NewCreateArticleHandler(repo),
		},
		Queries: app.Queries{
			GetArticleBySlug: queries.NewGetArticleBySlugHandler(repo),
			GetArticles:      queries.NewGetArticlesHandler(repo),
		},
	}
	a.Application = application
	return nil
}
