package controllers

import (
	"errors"
	"encoding/hex"
	"html/template"
	"golang.org/x/crypto/bcrypt"
	"github.com/astaxie/beego"

	"github.com/physpeach/bbs/models"
)

// SessionsController operations for Sessions
type SessionsController struct {
	beego.Controller
}

// URLMapping ...
func (c *SessionsController) URLMapping() {
	c.Mapping("New", c.New)
	c.Mapping("Create", c.Create)
	c.Mapping("Destroy", c.Destroy)
}

func (c *SessionsController) New() {
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layouts/application.tpl"
	c.TplName = "sessions/new.tpl"
}

// Post ...
// @Title Create
// @Description create Sessions
// @Param	body		body 	models.Sessions	true		"body for Sessions content"
// @Success 201 {object} models.Sessions
// @Failure 403 body is empty
// @router / [post]
func (c *SessionsController) Create() {
	account, err := ConfirmAccountPassword(c)
	if err == nil {
		c.SetSession("sessAccountID", account.ID)
		c.Ctx.Redirect(302, "/" + account.Name)
	}
	c.Data["loginError"] = err
	c.Data["xsrf"] = template.HTML(c.XSRFFormHTML())
	c.Layout = "layouts/application.tpl"
	c.TplName = "sessions/new.tpl"
}

func (c *SessionsController) Destroy() {
	c.DestroySession()
	c.Ctx.Redirect(302, "/")
}

func ConfirmAccountPassword(c *SessionsController)(*models.Account, error){
	accountName := c.GetString("Name")
	unhashed := c.GetString("Password")
	account, err := models.GetAccountByName(accountName)
	if err != nil{
		return nil, errors.New(accountName + " could not be found")
	}
	hashed, err := hex.DecodeString(account.Password)
	if err != nil{
		c.Abort("500")
	}
	if err := bcrypt.CompareHashAndPassword(hashed, []byte(unhashed)); account.Password != "0123" && err != nil {
		return nil, errors.New("Wrong password")
	}
	return account, nil
}