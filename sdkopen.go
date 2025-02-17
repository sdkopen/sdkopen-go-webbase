package sdkopen_webbase

import (
	"github.com/sdkopen/sdkopen-go-webbase/observer"
	"github.com/sdkopen/sdkopen-go-webbase/validator"
)

func InitializeSdkOpen() {
	validator.Initialize()
	observer.Initialize()
}
