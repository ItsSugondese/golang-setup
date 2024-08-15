package main

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
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
	tenant_middleware "wabustock/pkg/middleware/tenant-middleware"
)

func init() {
	print("Here in init")
	initializers.LoadEnvironments()
	database.ConnectToDB()
	//tenantPublic := generic_models.Tenant{
	//	Name:       "Public",
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
	//	Name:       "TenantA",
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

	// middlewares
	r.Use(cors_middleware.CorsMiddleware())
	r.Use(middleware.RecoveryMiddleware())
	r.Use(lang_middleware.LocalizationMiddleware(localization.InitBundle()))
	r.Use(tenant_middleware.TenantMiddleware(database.DB))

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

// MigrateTenantTables migrates the necessary tables within the tenant schema
func MigrateTenantTables(db *gorm.DB, tenant generic_models.Tenant) error {
	if err := tenant_middleware.SetTenantSchema(db, tenant.SchemaName); err != nil {
		return err
	}

	//Automigrate tenant-specific tables
	if err := db.AutoMigrate(&temporary_attachments.TemporaryAttachments{}); err != nil { // Add more models as needed
		return err
	}

	if err := db.AutoMigrate(&user.BaseUser{}, &generic_models.AuditModel{}); err != nil { // Add more models as needed
		return err
	}

	return nil
}

func MigrateTenantPublicTable(db *gorm.DB, tenant generic_models.Tenant) error {
	if err := tenant_middleware.SetTenantSchema(db, tenant.SchemaName); err != nil {
		return err
	}

	//Automigrate tenant-specific tables
	if err := db.AutoMigrate(&generic_models.Tenant{}); err != nil { // Add more models as needed
		return err
	}
	return nil
}

func SaveTenantDetails(db *gorm.DB, tenant generic_models.Tenant) error {
	if err := tenant_middleware.SetTenantSchema(db, "public"); err != nil {
		return err
	}
	result := database.DB.Create(&tenant)
	return result.Error
}
