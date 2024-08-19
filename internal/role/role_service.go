package role

import (
	localization_enums "wabustock/enums/struct-enums/localization-enums"
	"wabustock/enums/struct-enums/project_module"
	"wabustock/pkg/common/localization"
)

func CreateRoleService(dto *RoleRequest) *Role {
	role, getRoleError := FindRoleByIdRepo(dto.Name)
	if getRoleError != nil {
		panic(getRoleError)
	}

	if role != nil {
		panic("Role already exists")
	}

	savedRole, saveRoleError := SaveRoleRepo(&Role{ID: dto.Name})
	if saveRoleError != nil {
		panic(saveRoleError)
	}
	return savedRole
}

func FindRoleByIdService(name *string) Role {
	role, getRoleError := FindRoleByIdRepo(name)
	if getRoleError != nil {
		panic(getRoleError)
	}

	if role == nil {
		panic(localization.GetLocalizedMessage(localization_enums.MessageCodeEnums.COLUMN_NOT_EXISTS, map[string]interface{}{
			"First":  project_module.ModuleNameEnums.ROLE,
			"Second": "Id",
		}))
	}
	return *role
}
