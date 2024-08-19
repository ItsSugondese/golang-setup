package tenant

type TenantRequest struct {
	Name string `json:"name" binding:"required"`
}
