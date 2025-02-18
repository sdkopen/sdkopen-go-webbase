package sdkopen_webbase

import (
	"github.com/sdkopen/sdkopen-go-webbase/observer"
	"github.com/sdkopen/sdkopen-go-webbase/validator"
)

func Initialize() {
	validator.Initialize()
	observer.Initialize()
}
