package campaign

import "encoding/json"

// UpdateStatusRequest 修改广告计划状态 API Request
type UpdateStatusRequest struct {
	// AdvertiserID 广告主ID
	AdvertiserID int64 `json:"advertiser_id,omitempty"`
	// CampaignID 广告计划ID
	CampaignID int64 `json:"campaign_id,omitempty"`
	// CampaignIDs 与原有的campaign_id字段可以同时填，也可以只填一个;1.传入的计划id不得重复，且至少有一个;传入的campaign_id总数最多20个。2.put_status为3时，会删除广告计划，和计划下的所有广告组，程序化创意，创意
	CampaignIDs []int64 `json:"campaign_ids,omitempty"`
	// PutStatus 操作码; 1-投放、2-暂停、3-删除，传其他数字非法
	PutStatus int `json:"put_status,omitempty"`
}

// Url implement PostRequest interface
func (r UpdateStatusRequest) Url() string {
	return "v1/campaign/update/status"
}

// Encode implement PostRequest interface
func (r UpdateStatusRequest) Encode() []byte {
	ret, _ := json.Marshal(r)
	return ret
}
