package overhead

import (
	"life/global"
	"life/model/common/request"
	"life/model/overhead"
	overheadReq "life/model/overhead/request"
)

type OverheadInfoService struct {
}

// CreateOverheadInfo 创建OverheadInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (InfoService *OverheadInfoService) CreateOverheadInfo(Info overhead.OverheadInfo) (err error) {
	err = global.GVA_DB.Create(&Info).Error
	return err
}

// DeleteOverheadInfo 删除OverheadInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (InfoService *OverheadInfoService) DeleteOverheadInfo(Info overhead.OverheadInfo) (err error) {
	err = global.GVA_DB.Delete(&Info).Error
	return err
}

// DeleteOverheadInfoByIds 批量删除OverheadInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (InfoService *OverheadInfoService) DeleteOverheadInfoByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]overhead.OverheadInfo{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOverheadInfo 更新OverheadInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (InfoService *OverheadInfoService) UpdateOverheadInfo(Info overhead.OverheadInfo) (err error) {
	err = global.GVA_DB.Save(&Info).Error
	return err
}

// GetOverheadInfo 根据id获取OverheadInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (InfoService *OverheadInfoService) GetOverheadInfo(id uint) (Info overhead.OverheadInfo, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&Info).Error
	return
}

// GetOverheadInfoInfoList 分页获取OverheadInfo记录
// Author [piexlmax](https://github.com/piexlmax)
func (InfoService *OverheadInfoService) GetOverheadInfoInfoList(info overheadReq.OverheadInfoSearch) (list []overhead.OverheadInfo, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&overhead.OverheadInfo{})
	var Infos []overhead.OverheadInfo
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&Infos).Error
	return Infos, total, err
}
