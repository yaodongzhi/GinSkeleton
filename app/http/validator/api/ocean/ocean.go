package ocean

import (
	"github.com/gin-gonic/gin"
	"goskeleton/app/global/consts"
	"goskeleton/app/http/controller/api"
	"goskeleton/app/http/validator/core/data_transfer"
	"goskeleton/app/utils/response"
)

type Store struct {
	SpreadId      string `json:"spread_id" form:"spread_id"`         //  外推编号
	Acid          string `json:"acid" form:"acid"`                   //  广告主id
	Gid           string `json:"gid" form:"gid"`                     //  广告组id
	Aid           string `json:"aid" form:"aid"`                     //  广告计划id
	AidName       string `json:"aid_name" form:"aid_name"`           //  广告计划名称
	Cid           string `json:"cid" form:"cid"`                     //  创意id
	CidName       string `json:"cid_name" form:"cid_name"`           //  创意名称
	CampaignId    string `json:"campaign_id" form:"campaign_id"`     //  广告组id
	CampaignName  string `json:"campaign_name" form:"campaign_name"` //  广告组名称
	Ctype         string `json:"ctype" form:"ctype"`                 //  创意样式
	AdvertiserId  string `json:"advertiser_id" form:"advertiser_id"` //  广告主id
	Csite         string `json:"csite" form:"csite"`                 //  广告投放位置
	ConvertId     string `json:"convert_id" form:"convert_id"`       //  转化id
	RequestId     string `json:"request_id" form:"request_id"`       //  请求下发id
	Sl            string `json:"sl" form:"sl"`                       //  请求的语言
	Imei          string `json:"imei" form:"imei"`                   //  imei md5加密
	Idfa          string `json:"idfa" form:"idfa"`                   //  ios的设备id
	IdfaMd5       string `json:"idfa_md5" form:"idfa_md5"`           //  ios的设备id md5
	Androidid     string `json:"androidid" form:"androidid"`         //  安卓设备的原值
	Oaid          string `json:"oaid" form:"oaid"`                   //  oaid 原值
	OaidMd5       string `json:"oaid_md5" form:"oaid_md5"`           //  oaid md5
	Os            string `json:"os" form:"os"`                       //  操作系统，0：安卓，1：ios，3：其他
	Mac           string `json:"mac" form:"mac"`                     //  移动设备mac地址,转换成大写字母,去掉“:”，并且取md5摘要后的结果
	Mac1          string `json:"mac1" form:"mac1"`                   //  移动设备 mac 地址,转换成大写字母,并且取md5摘要后的结果，32位
	Ip            string `json:"ip" form:"ip"`                       //  正常情况下，全量下发ipv4地址。极少数情况如ipv4无法取数，则下发 ipv6
	Ipv4          string `json:"ipv4" form:"ipv4"`                   //  ipv4
	Ipv6          string `json:"ipv6" form:"ipv6"`                   //  ipv6
	Ua            string `json:"ua" form:"ua"`
	Geo           string `json:"geo" form:"geo"`                       //  位置信息，包含三部分:latitude（纬度），longitude（经度）以及precise（确切信息,精度）
	Ts            string `json:"ts" form:"ts"`                         //  客户端发生广告点击事件的时间，以毫秒为单位时间戳
	CallbackParam string `json:"callback_param" form:"callback_param"` //  一些跟广告信息相关的回调参数，内容是一个加密字符串，在调用事件回传接口的时候会用到
	CallbackUrl   string `json:"callback_url" form:"callback_url"`     //  直接把调用事件回传接口的url生成出来，广告主可以直接使用
	Model         string `json:"model" form:"model"`                   //  手机型号
	UnionSite     string `json:"union_site" form:"union_site"`         //  对外广告位编码
	Caid          string `json:"caid" form:"caid"`                     //  中国广告协会互联网广告标识，包含最新两个版本的caid和版本号，url encode之后的json字符串
	CaidMd5       string `json:"caid_md5" form:"caid_md5"`             //  caid md5
	CreateTime    int64  `json:"create_time" form:"create_time"`       //  创建时间
	Status        int64  `json:"status" form:"status"`                 //  上报状态 1已上报 0未上报
	UploadTime    int64  `json:"upload_time" form:"upload_time"`       //  上报时间
	Uid           string `json:"uid" form:"uid"`                       //  uid
	IsRuning      int64  `json:"is_runing" form:"is_runing"`           //  是否已经跑过匹配，1为是
	Type          int64  `json:"type" form:"type"`                     //  类型 1普通回传 2落地页
	TrackId       string `json:"track_id" form:"track_id"`
	Clickid       string `json:"clickid" form:"clickid"` //  落地页id
	Adid          string `json:"adid" form:"adid"`
	Creativeid    string `json:"creativeid" form:"creativeid"`
	Creativetype  string `json:"creativetype" form:"creativetype"`
	Productid     string `json:"productid" form:"productid"`
	Outerid       string `json:"outerid" form:"outerid"`
	PromotionId   string `json:"promotion_id" form:"promotion_id"`
	ProjectId     string `json:"project_id" form:"project_id"`
	PromotionName string `json:"promotion_name" form:"promotion_name"`
	ProjectName   string `json:"project_name" form:"project_name"`
	AdCallbackId  int64  `json:"ad_callback_id" form:"ad_callback_id"` //  回传表的id
	Mid1          string `json:"mid1" form:"mid1"`
	Mid2          string `json:"mid2" form:"mid2"`
	Mid3          string `json:"mid3" form:"mid3"`
	Mid4          string `json:"mid4" form:"mid4"`
	Mid5          string `json:"mid5" form:"mid5"`
	Mid6          string `json:"mid6" form:"mid6"`
}

func (n Store) CheckParams(context *gin.Context) {
	//1.先按照验证器提供的基本语法，基本可以校验90%以上的不合格参数
	if err := context.ShouldBind(&n); err != nil {
		// 将表单参数验证器出现的错误直接交给错误翻译器统一处理即可
		response.ValidatorError(context, err)
		return
	}

	//  该函数主要是将绑定的数据以 键=>值 形式直接传递给下一步（控制器）
	extraAddBindDataContext := data_transfer.DataAddContext(n, consts.ValidatorPrefix, context)
	if extraAddBindDataContext == nil {
		response.ErrorSystem(context, "HomeNews表单验证器json化失败", "")
	} else {
		// 验证完成，调用控制器,并将验证器成员(字段)递给控制器，保持上下文数据一致性
		(&api.MonitorOcean{}).Store(extraAddBindDataContext)
	}
}
