package main

import (
	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
	_ "wabustock/docs"
	global_gin_context "wabustock/global/global-gin-context"
	"wabustock/initializers"
	"wabustock/internal/auth"
	"wabustock/internal/role"
	temporary_attachments "wabustock/internal/temporary-attachments"
	"wabustock/internal/user"
	"wabustock/pkg/common/database"
	"wabustock/pkg/common/localization"
	"wabustock/pkg/common/tenant"
	"wabustock/pkg/middleware"
	audit_middleware "wabustock/pkg/middleware/audit-middleware"
	cors_middleware "wabustock/pkg/middleware/cors-middleware"
	lang_middleware "wabustock/pkg/middleware/lang-middleware"
	"wabustock/pkg/utils"
	"wabustock/pkg/utils/paseto-token"
)

const (
	projectID  = "your-project-id"  // FILL IN WITH YOURS
	bucketName = "blackpearlbucket" // FILL IN WITH YOURS
)

func init() {
	print("Here in init")
	initializers.LoadEnvironments()
	localization.InitLocalizationManager()
	global_gin_context.NewGlobalGinContext()
	database.ConnectToDB()
	tenant.MigrateAll()

	os.Setenv("GOOGLE_APPLICATION_CREDENTIALS", "/home/lazybot/Desktop/continual-mind-432410-g5-df1fc7f32718.json") // FILL IN WITH YOUR FILE PATH

	//migrateToPublicTenantError := tenant.MigrateTenantPublicTable(database.DB)
	//if migrateToPublicTenantError != nil {
	//	panic(migrateToPublicTenantError)
	//}
	//
	//schemaList, getSchemaError := database.GetAllSchemasRepo(database.DB)
	//if getSchemaError != nil {
	//	panic(getSchemaError)
	//}
	//
	//for _, schema := range schemaList {
	//	if schema == tenant_constants.PublicTenant {
	//		continue
	//	}
	//	migrateError := tenant.MigrateTenantTables(database.DB, tenant.Tenant{Name: schema})
	//	if migrateError != nil {
	//		return
	//	}
	//}

	//tenantPublic := generic_models.Tenant{
	//	ID:       "Public",
	//	SchemaName: "public",
	//}
	//
	//if err := database.CreateTenantSchema(database.DB, tenantPublic); err != nil {
	//	log.Fatal("Failed to create tenant schema:", err)
	//}
	//
	//if err := MigrateTenantPublicTable(database.DB, tenantPublic); err != nil {
	//	log.Fatal("Failed to migrate tenant tables:", err)
	//}
	//tenant := generic_models.Tenant{
	//	ID:       "TenantA",
	//	SchemaName: "tenant_a_schema",
	//}
	//
	//if err := database.CreateTenantSchema(database.DB, tenant); err != nil {
	//	log.Fatal("Failed to create tenant schema:", err)
	//}
	//
	//if err := MigrateTenantTables(database.DB, tenant); err != nil {
	//	log.Fatal("Failed to migrate tenant tables:", err)
	//}
	//
	//err := SaveTenantDetails(database.DB, tenant)
	//
	//if err != nil {
	//	panic(err)
	//}

	//db.AutoMigrate(&temporary_attachments.TemporaryAttachments{})
	//
	//db.AutoMigrate(&user.BaseUser{}, &generic_models.AuditModel{})

}

func main() {
	r := gin.Default()
	validate := validator.New()
	tokenMaker, err := paseto_token.NewPaseto("abcdefghijkl12345678901234567890")
	if err != nil {
		panic("Couldnt open tokenmaker " + err.Error())
	}

	// Initialize the Google Cloud Storage client using Gin's context.
	r.Use(func(c *gin.Context) {
		client, err := storage.NewClient(c.Request.Context())
		if err != nil {
			log.Fatalf("Failed to create storage client: %v", err)
		}

		// Store the client in Gin's context for use in handlers.
		c.Set("storageClient", client)

		utils.Uploader = &utils.ClientUploader{
			Cl:         client,
			BucketName: bucketName,
			//ProjectID:  projectID,
			UploadPath: "test-files/",
		}
		// Make sure to close the client after the request is processed.
		defer client.Close()

		c.Next()
	})

	paseto_token.TokenMaker = tokenMaker

	// middlewares
	r.Use(cors_middleware.CorsMiddleware())
	r.Use(middleware.RecoveryMiddleware())
	r.Use(lang_middleware.LocalizationMiddleware(localization.InitBundle()))
	//r.Use(tenant_middleware.TenantMiddleware(database.DB))

	// payload validations
	payloadValidations()

	//routes
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) // for swagger

	//r.GET("/ws", func(c *gin.Context) {
	//	socket_config.ServeWs(hub, c.Writer, c.Request)
	//})

	// Register the audit log callbacks and perform migrations
	errVal := audit_middleware.RegisterCallbacks(database.DB)
	if errVal != nil {
		panic("failed to register audit log callbacks")
	}

	// Registering routes
	user.UserRoutes(r, validate)
	auth.AuthRoutes(r, validate)
	temporary_attachments.TempAttachmentsRoutes(r, validate)
	tenant.TenantRoutes(r, validate)
	role.RoleRoutes(r, validate)

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
