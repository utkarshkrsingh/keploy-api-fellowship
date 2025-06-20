package initializers

import "github.com/utkarshkrsingh/keploy-api-fellowship/basic-api/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.WatchList{})
}
