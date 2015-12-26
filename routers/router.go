package routers

import (
	"github.com/astaxie/beego"
	"github.com/saiyawang/hanuman/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/list", &controllers.ListController{})
	beego.Router("/new", &controllers.NewController{})
	beego.Router("/view", &controllers.ViewController{})
	beego.Router("/candidate", &controllers.CandidateController{})

	beego.Router("/getcandidatelist", &controllers.CandidateController{}, "GET:GetCandidates")
	beego.Router("/insertcandidate", &controllers.CandidateController{}, "POST:InsertOneCandidate")
	beego.Router("/getcandidatebyid", &controllers.CandidateController{}, "POST:GetCandidateByID")

	beego.Router("/insertcomment", &controllers.CommentController{}, "POST:InsertOneComment")
	beego.Router("/getcommentlist", &controllers.CommentController{}, "POST:GetComments")

	beego.Router("/insertattachment", &controllers.AttachmentController{}, "POST:InsertOneAttachment")
	beego.Router("/getattachments", &controllers.AttachmentController{}, "POST:GetAttachments")

}
