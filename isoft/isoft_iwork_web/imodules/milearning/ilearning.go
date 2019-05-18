package milearning

import (
	"github.com/astaxie/beego/orm"
	"isoft/isoft_iwork_web/models/cms"
	"isoft/isoft_iwork_web/models/common"
	"isoft/isoft_iwork_web/models/iblog"
	"isoft/isoft_iwork_web/models/ifile"
	"isoft/isoft_iwork_web/models/ilearning"
	"isoft/isoft_iwork_web/models/monitor"
	"isoft/isoft_iwork_web/models/share"
)

func RegisterModel() {
	orm.RegisterModel(new(iblog.Catalog))
	orm.RegisterModel(new(iblog.Blog))

	orm.RegisterModel(new(ilearning.Course))
	orm.RegisterModel(new(ilearning.CourseVideo))
	orm.RegisterModel(new(ilearning.Favorite))
	orm.RegisterModel(new(ilearning.CommentTheme))
	orm.RegisterModel(new(ilearning.CommentReply))
	orm.RegisterModel(new(ilearning.Note))

	orm.RegisterModel(new(ifile.IFile))

	orm.RegisterModel(new(cms.Configuration))
	orm.RegisterModel(new(cms.CommonLink))

	orm.RegisterModel(new(share.Share))

	orm.RegisterModel(new(common.History))

	orm.RegisterModel(new(monitor.HeartBeat2))
	orm.RegisterModel(new(monitor.HeartBeatDetail))
}
