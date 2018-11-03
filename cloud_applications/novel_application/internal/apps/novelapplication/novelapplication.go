package novelapplication

import (
	dal "github.com/TheTerribleChild/cloud_appplication_portal/cloud_applications/novel_application/internal/apps/novelapplication/dal"
)

func Test() string {
	return dal.GetMapper();
}