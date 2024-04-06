package model

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"goskeleton/app/utils/data_bind"
	"time"
)

func CreateMonitorOceanFactory(sqlType string) *MonitorOceanModel {
	return &MonitorOceanModel{BaseModel: BaseModel{DB: UseDbConn(sqlType)}}
}

type MonitorOceanModel struct {
	BaseModel
	SpreadId      string `json:"spread_id" gorm:"spread_id"`
	Acid          string `json:"acid" gorm:"acid"`
	Gid           string `json:"gid" gorm:"gid"`
	Aid           string `json:"aid" gorm:"aid"`
	AidName       string `json:"aid_name" gorm:"aid_name"`
	Cid           string `json:"cid" gorm:"cid"`
	CidName       string `json:"cid_name" gorm:"cid_name"`
	CampaignId    string `json:"campaign_id" gorm:"campaign_id"`
	CampaignName  string `json:"campaign_name" gorm:"campaign_name"`
	Ctype         string `json:"ctype" gorm:"ctype"`
	AdvertiserId  string `json:"advertiser_id" gorm:"advertiser_id"`
	Csite         string `json:"csite" gorm:"csite"`
	ConvertId     string `json:"convert_id" gorm:"convert_id"`
	RequestId     string `json:"request_id" gorm:"request_id"`
	Sl            string `json:"sl" gorm:"sl"`
	Imei          string `json:"imei" gorm:"imei"`
	Idfa          string `json:"idfa" gorm:"idfa"`
	IdfaMd5       string `json:"idfa_md5" gorm:"idfa_md5"`
	Androidid     string `json:"androidid" gorm:"androidid"`
	Oaid          string `json:"oaid" gorm:"oaid"`
	OaidMd5       string `json:"oaid_md5" gorm:"oaid_md5"`
	Os            string `json:"os" gorm:"os"`
	Mac           string `json:"mac" gorm:"mac"`
	Mac1          string `json:"mac1" gorm:"mac1"`
	Ip            string `json:"ip" gorm:"ip"`
	Ipv4          string `json:"ipv4" gorm:"ipv4"`
	Ipv6          string `json:"ipv6" gorm:"ipv6"`
	Ua            string `json:"ua" gorm:"ua"`
	Geo           string `json:"geo" gorm:"geo"`
	Ts            string `json:"ts" gorm:"ts"`
	CallbackParam string `json:"callback_param" gorm:"callback_param"`
	CallbackUrl   string `json:"callback_url" gorm:"callback_url"`
	Model         string `json:"model" gorm:"model"`
	UnionSite     string `json:"union_site" gorm:"union_site"`
	Caid          string `json:"caid" gorm:"caid"`
	CaidMd5       string `json:"caid_md5" gorm:"caid_md5"`
	CreateTime    int32  `json:"create_time" gorm:"create_time"`
	Status        int8   `json:"status" gorm:"status"`
	UploadTime    int32  `json:"upload_time" gorm:"upload_time"`
	Uid           string `json:"uid" gorm:"uid"`
	IsRunning     string `json:"is_running" gorm:"is_running"`
	Type          int8   `json:"type" gorm:"type"`
	Clickid       string `json:"clickid" gorm:"clickid"`
	Adid          string `json:"adid" gorm:"adid"`
	CreativeId    string `json:"creative_id" gorm:"creative_id"`
	CreativeType  string `json:"creative_type" gorm:"creative_type"`
	ProductId     string `json:"product_id" gorm:"product_id"`
	OuterId       string `json:"outer_id" gorm:"outer_id"`
	PromotionId   string `json:"promotion_id" gorm:"promotion_id"`
	ProjectId     string `json:"project_id" gorm:"project_id"`
	PromotionName string `json:"promotion_name" gorm:"promotion_name"`
	ProjectName   string `json:"project_name" gorm:"project_name"`
	Mid1          string `json:"mid1" gorm:"mid1"`
	Mid2          string `json:"mid2" gorm:"mid2"`
	Mid3          string `json:"mid3" gorm:"mid3"`
	Mid4          string `json:"mid4" gorm:"mid4"`
	Mid5          string `json:"mid5" gorm:"mid5"`
	Mid6          string `json:"mid6" gorm:"mid6"`
	TrackId       string `json:"track_id" gorm:"track_id"`
	AdCallbackId  int32  `json:"ad_callback_id" gorm:"ad_callback_id"`
}

func (m *MonitorOceanModel) TableName() string {
	tb := "nv_monitor_ocean"
	tb += "_" + time.Now().Format("20060102")
	return tb
}

func (m *MonitorOceanModel) CreateOneOcean(tx *gorm.DB, ctx *gin.Context) error {
	data := &MonitorOceanModel{}
	if err := data_bind.ShouldBindFormDataToModel(ctx, data); err != nil {
		return err
	} else {
		return m.CreateOne(tx, m.TableName(), data)
	}
}
