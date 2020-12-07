package generate

import "fmt"

func getImportStr(name string) string {
	imp := `import (
		"fmt"
		"github.com/gin-gonic/gin"
		"go-api/api/app/models/%s"
		"go-api/api/pkg/model"
		"go-api/api/pkg/utils"
		"strconv"
	)
	`
	imp = fmt.Sprintf(imp, name)
	return imp
}
