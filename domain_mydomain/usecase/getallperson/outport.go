package getallperson

import "your/path/project/domain_mydomain/model/repository"

// Outport of usecase
type Outport interface {
	repository.FindAllPersonRepo
}
