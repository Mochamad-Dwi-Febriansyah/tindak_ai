package seed

import (
	"log"
	"tindak_ai/internal/domain"

	"github.com/google/uuid"
	"gorm.io/gorm"
)


func SeedRoleAndPermission(db *gorm.DB) {
	// role 
	roles := []domain.Role{
		{ID: uuid.New(), Name: "super_admin"},
		{ID: uuid.New(), Name: "admin"},
		{ID: uuid.New(), Name: "office"},
		{ID: uuid.New(), Name: "public"},
	}

	for _, role := range roles {
		if err := db.FirstOrCreate(&role, domain.Role{Name: role.Name}).Error; err != nil {
			log.Printf("Failed to seed role %s: %v", role.Name, err)
		}
	}

	// permission
	permissions := []domain.Permission{
		{ID: uuid.New(), Action: "create", Resource: "user"},
		{ID: uuid.New(), Action: "read", Resource: "user"},
		{ID: uuid.New(), Action: "show", Resource: "user"},
		{ID: uuid.New(), Action: "update", Resource: "user"},
		{ID: uuid.New(), Action: "delete", Resource: "user"},
	}

	for _, permission := range permissions {
		if err := db.FirstOrCreate(&permission, domain.Permission{Action: permission.Action, Resource: permission.Resource}).Error; err != nil {
			log.Printf("Failed to seed permission %s: %v", permission.Action, err)
		}
	}

	// assign roles to super_admin
	var superAdminRole domain.Role
	if err := db.Where("name = ?", "super_admin").First(&superAdminRole).Error; err != nil {
		log.Printf("Failed to find super_admin role: %v", err)
		return
	}
	for _, perm := range permissions {
		var permRecord domain.Permission
		if err := db.Where("action = ? AND resource = ?", perm.Action, perm.Resource).First(&permRecord).Error; err == nil {
			rp := domain.RolePermission{
				ID:           uuid.New(),
				RoleID:       superAdminRole.ID,
				PermissionID: permRecord.ID,
			}
			db.FirstOrCreate(&rp, domain.RolePermission{
				RoleID: superAdminRole.ID,
				PermissionID: permRecord.ID,
			})
		}
	}

	// assign permission to user
	var user domain.Users
	if err := db.Where("id = ?", uuid.MustParse("8c7e7524-139a-4dc6-a89d-308619ffb6cb")).First(&user).Error; err != nil {
		log.Printf("Failed to find user permission: %v", err)
		return
	}
	for _, perm := range permissions {
		var permRecord domain.Permission
		if err := db.Where("action = ? AND resource = ?", perm.Action, perm.Resource).First(&permRecord).Error; err == nil {
			up := domain.UserPermission{
				ID:           uuid.New(),
				UserID:       user.ID,
				PermissionID: permRecord.ID,
			}
			
			db.FirstOrCreate(&up, domain.UserPermission{
				UserID: user.ID,
				PermissionID: permRecord.ID,
			})
		}
	}


}