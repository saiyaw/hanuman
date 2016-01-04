package routers

import (
	"github.com/astaxie/beego"
	"github.com/saiyawang/hanuman/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/list", &controllers.ListController{})
	beego.Router("/new", &controllers.NewController{})
	beego.Router("/admin", &controllers.AdminController{})
	beego.Router("/candidate", &controllers.CandidateController{})

	beego.Router("/getcandidatelist", &controllers.CandidateController{}, "GET:GetCandidates")
	beego.Router("/insertcandidate", &controllers.CandidateController{}, "POST:InsertOneCandidate")
	beego.Router("/updatecandidate", &controllers.CandidateController{}, "POST:UpdateOneCandidate")
	beego.Router("/getcandidatebyid", &controllers.CandidateController{}, "POST:GetCandidateByID")

	beego.Router("/insertcomment", &controllers.CommentController{}, "POST:InsertOneComment")
	beego.Router("/getcommentlist", &controllers.CommentController{}, "POST:GetComments")

	beego.Router("/insertattachment", &controllers.AttachmentController{}, "POST:InsertOneAttachment")
	beego.Router("/getattachments", &controllers.AttachmentController{}, "POST:GetAttachments")

	beego.Router("/insertpost", &controllers.PostController{}, "POST:InsertOnePost")
	beego.Router("/getpostlist", &controllers.PostController{}, "GET:GetPostList")

	beego.Router("/insertlabel", &controllers.LabelController{}, "POST:InsertOneLabel")
	beego.Router("/getlabellist", &controllers.LabelController{}, "GET:GetLabelList")
	beego.Router("/getlabel", &controllers.LabelController{}, "GET:GetLabel")
	beego.Router("/deletelabel", &controllers.LabelController{}, "POST:DeleteOneLabel")

	beego.Router("/insertcandidatelabel", &controllers.CandidateLabelController{}, "POST:InsertCandidateLabel")
	beego.Router("/getcandidatelabellist", &controllers.CandidateLabelController{}, "POST:GetCandidateLabelList")
	beego.Router("/deletecandidatelabel", &controllers.CandidateLabelController{}, "POST:DeleteCandidateLabel")
	beego.Router("/deletecandidatelabelbylabel", &controllers.CandidateLabelController{}, "POST:DeleteCandidateLabelByLabel")

}
