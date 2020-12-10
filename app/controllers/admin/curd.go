package admin

import (
		"fmt"
		"github.com/gin-gonic/gin"
		"go-api/api/app/models/admin"
		"go-api/api/pkg/model"
		"go-api/api/pkg/utils"
		"strconv"
		"golang.org/x/crypto/bcrypt"
)
	
func Create(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)

	data["email"] = []string{"required","min:6","email"} 
	data["password"] = []string{"min:6"} 
	data["name"] = []string{"required","min:6"} 

	validate := utils.Validator(c.Request, data)
	if validate != nil{
		utils.SuccessErr(c,403,validate)
		return
	}
	var Admin admin.Admin
	pwd, _ := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), bcrypt.DefaultCost)
	Admin.Email = c.PostForm("email")
	Admin.Password = string(pwd)
	Admin.Name = c.PostForm("name")
	model.Create(&Admin)
	utils.SuccessData(c, Admin) // 返回创建成功的信息
}
func Update(c *gin.Context) {
	// 验证参数
	data := make(map[string][]string)

	data["email"] = []string{"required","min:6","email"} 
	data["password"] = []string{"min:6"} 
	data["name"] = []string{"required","min:6"} 

	validate := utils.Validator(c.Request, data)
	if validate != nil{
		utils.SuccessErr(c,403,validate)
		return
	}
	id, _ := strconv.Atoi(c.Query("id"))
	var Admin admin.Admin
	model.First(&Admin,id)
	if Admin.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}

	pwd, _ := bcrypt.GenerateFromPassword([]byte(c.PostForm("password")), bcrypt.DefaultCost)
	Admin.Email = c.PostForm("email")
	Admin.Password = string(pwd)
	Admin.Name = c.PostForm("name")
	model.UpdateOne(Admin)
	utils.SuccessData(c, Admin) // 返回创建成功的信息
}
func Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.PostForm("id"))
	var Admin admin.Admin

	model.First(&Admin,id)
	if Admin.Id == 0 {
		utils.SuccessErr(c, -1000, "数据不存在")
		return
	}
	model.DeleteOne(&Admin)
	utils.SuccessData(c, "删除成功")
}
func Info(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	var Admin admin.Admin
	model.First(&Admin,id)

	utils.SuccessData(c, Admin)
}
func Paginate(c *gin.Context) {
	var conditions []model.Condition
	Email := c.Query("email")
	if Email != "" {
		conditions = model.SetCondition(conditions,"email",fmt.Sprintf("%s%s", Email, "%"),"like")
	}
	Name := c.Query("name")
	if Name != "" {
		conditions = model.SetCondition(conditions,"name",fmt.Sprintf("%s%s", Name, "%"),"like")
	}

	var Admin []admin.Admin
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("page_size", "10"))
	lists := model.Paginate(&Admin, model.PageInfo{Page: int64(page), PageSize: int64(pageSize)}, conditions)
	utils.SuccessData(c, lists)
}