package swagger

import (
	"net/http"
	_ "mail-phone-auth/docs"
	httpSwagger "github.com/swaggo/http-swagger"
)

type Swagger struct {
	router *http.ServeMux
}

func New(router *http.ServeMux) *Swagger {
	swagger := Swagger{
		router: router,
	}
	router.Handle("GET /swagger/", swagger.Handler())

	return &swagger
}

func (swagger *Swagger) Handler() http.HandlerFunc {	
	return httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
		swagger.AfterScript(),
	)
}

func (swagger *Swagger) AfterScript() func(*httpSwagger.Config) {
	return httpSwagger.AfterScript(`
		
		document.querySelector('.download-url-wrapper').remove()
		document.querySelector('.topbar').style.backgroundColor = '#62a03f'

		deleteLink()
		function deleteLink() {
			console.log('tick')
			if (document.querySelector('span.url')) {
				document.querySelector('span.url').remove()
			} else {
				setTimeout(deleteLink, 100)
			}
		}
	`)		
}

