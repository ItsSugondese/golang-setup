package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	_ "wabustock/docs"
	generic_models "wabustock/generics/generic-models"
	"wabustock/initializers"
	"wabustock/internal/auth"
	temporary_attachments "wabustock/internal/temporary-attachments"
	"wabustock/internal/user"
	"wabustock/pkg/common/database"
	"wabustock/pkg/common/localization"
	"wabustock/pkg/middleware"
	audit_middleware "wabustock/pkg/middleware/audit-middleware"
	cors_middleware "wabustock/pkg/middleware/cors-middleware"
	lang_middleware "wabustock/pkg/middleware/lang-middleware"
)

func init() {
	print("Here in init")
	initializers.LoadEnvironments()
	db := database.ConnectToDB()
	db.AutoMigrate(&temporary_attachments.TemporaryAttachments{})

	db.AutoMigrate(&user.BaseUser{}, &generic_models.AuditModel{})

	// Register the audit log callbacks and perform migrations
	errVal := audit_middleware.RegisterCallbacks(db)
	if errVal != nil {
		panic("failed to register audit log callbacks")
	}

}

func main() {
	r := gin.Default()
	validate := validator.New()

	// middlewares
	r.Use(cors_middleware.CorsMiddleware())
	r.Use(middleware.RecoveryMiddleware())
	r.Use(lang_middleware.LocalizationMiddleware(localization.InitBundle()))

	// payload validations
	payloadValidations()

	//routes
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // for swagger

	//r.GET("/ws", func(c *gin.Context) {
	//	socket_config.ServeWs(hub, c.Writer, c.Request)
	//})
	// Registering routes

	user.UserRoutes(r, validate)
	auth.AuthRoutes(r, validate)
	temporary_attachments.TempAttachmentsRoutes(r, validate)

	log.Println("_____________")
	// Serve static files from the images directory
	r.Static("/images", "./images")

	r.Run()
}

func payloadValidations() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("validUserType", user.ValidUserType)
		v.RegisterValidation("validGenderType", user.ValidGenderType)
	}
}

func getLocalizedMessage(langTag string, bundle *i18n.Bundle) string {
	// Create a Localizer for the specified language
	localizer := i18n.NewLocalizer(bundle, langTag)

	// Localize the "hello_world" message
	localizedMessage, _ := localizer.Localize(&i18n.LocalizeConfig{
		MessageID: "welcome",
		DefaultMessage: &i18n.Message{
			ID:    "welcome",
			Other: "Hello there",
		},
	})

	return localizedMessage
}
