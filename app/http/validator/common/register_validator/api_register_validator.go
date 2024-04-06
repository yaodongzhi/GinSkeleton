package register_validator

import (
	"goskeleton/app/core/container"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/validator/api/home"
	"goskeleton/app/http/validator/api/ocean"
)

func ApiRegisterValidator() {
	//创建容器
	containers := container.CreateContainersFactory()

	var key string

	// 注册门户类表单参数验证器
	key = consts.ValidatorPrefix + "HomeNews"
	containers.Set(key, home.News{})

	key = consts.ValidatorPrefix + "Store"
	containers.Set(key, ocean.Store{})
}
