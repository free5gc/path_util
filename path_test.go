package path_util

import (
	"testing"

	"github.com/free5gc/path_util/logger"
)

func TestFree5gcPath(t *testing.T) {
	logger.PathLog.Infoln(Free5gcPath("free5gc/abcdef/abcdef.pem"))
}
