package controllers

import (
	"blog/models"
	"blog/syserror"
	"errors"
	"github.com/jinzhu/gorm"
)

type NoteController struct {
	BaseController
}

///note
// @router /new [get]
func (this *NoteController) Index()  {
	this.Data["key"] = this.UUID()
 	this.TplName = "note_new.html"
}

func (this *NoteController) NextPrepare() {
	this.MustLogin()
	if this.User.Role != 0 {
		this.Abort500(errors.New("权限不足"))
	}
}
// @router /save/:key [post]
func (receiver *NoteController) Save()  {
	key := receiver.Ctx.Input.Param("key")
	note, err := models.QueryNoteByKey(key)
	title := receiver.GetMustString("title","请输入标题")
	content := receiver.GetMustString("content","请输入文章内容")

	var n models.Note
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			n = models.Note{
				Key: key,
				Title:title,
				Content: content,
				UserID: int(receiver.User.ID),
				User: receiver.User,
			}
			if err := models.SaveNote(&n); err != nil{
				receiver.Abort500(syserror.NewError("保存失败", err))
			}
		} else  {
			receiver.Abort500(syserror.NewError("保存失败", err))
		}
	} else {

		if  err := models.SaveNote(&note); err != nil{
			receiver.Abort500(syserror.NewError("保存失败", err))
		}
	}
	receiver.JsonOK("保存成僧","/")

 }