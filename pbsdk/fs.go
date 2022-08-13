package pbsdk

/*
#include "inkview.h"
#cgo CFLAGS: -pthread
#cgo LDFLAGS: -pthread -lpthread -linkview
*/
import "C"

const (
	FlashDir     = C.FLASHDIR
	SDCardDir    = C.SDCARDDIR
	GlobalConfig = C.GLOBALCONFIGFILE
	ConfigDir    = C.CONFIGPATH
	LangDir      = C.USERLANGPATH
)
