package overhead

import (
	"life/global"
	"life/model/common/request"
	"life/model/overhead"
	overheadReq "life/model/overhead/request"
)

type OverheadDailyStatisticsService struct {
}

// CreateOverheadDailyStatistics 创建OverheadDailyStatistics记录
// Author [piexlmax](https://github.com/piexlmax)
func (DailyStatisticsService *OverheadDailyStatisticsService) CreateOverheadDailyStatistics(DailyStatistics overhead.OverheadDailyStatistics) (err error) {
	err = global.GVA_DB.Create(&DailyStatistics).Error
	return err
}

// DeleteOverheadDailyStatistics 删除OverheadDailyStatistics记录
// Author [piexlmax](https://github.com/piexlmax)
func (DailyStatisticsService *OverheadDailyStatisticsService) DeleteOverheadDailyStatistics(DailyStatistics overhead.OverheadDailyStatistics) (err error) {
	err = global.GVA_DB.Delete(&DailyStatistics).Error
	return err
}

// DeleteOverheadDailyStatisticsByIds 批量删除OverheadDailyStatistics记录
// Author [piexlmax](https://github.com/piexlmax)
func (DailyStatisticsService *OverheadDailyStatisticsService) DeleteOverheadDailyStatisticsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]overhead.OverheadDailyStatistics{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOverheadDailyStatistics 更新OverheadDailyStatistics记录
// Author [piexlmax](https://github.com/piexlmax)
func (DailyStatisticsService *OverheadDailyStatisticsService) UpdateOverheadDailyStatistics(DailyStatistics overhead.OverheadDailyStatistics) (err error) {
	err = global.GVA_DB.Save(&DailyStatistics).Error
	return err
}

// GetOverheadDailyStatistics 根据id获取OverheadDailyStatistics记录
// Author [piexlmax](https://github.com/piexlmax)
func (DailyStatisticsService *OverheadDailyStatisticsService) GetOverheadDailyStatistics(id uint) (DailyStatistics overhead.OverheadDailyStatistics, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&DailyStatistics).Error
	return
}

// GetOverheadDailyStatisticsInfoList 分页获取OverheadDailyStatistics记录
// Author [piexlmax](https://github.com/piexlmax)
func (DailyStatisticsService *OverheadDailyStatisticsService) GetOverheadDailyStatisticsInfoList(info overheadReq.OverheadDailyStatisticsSearch) (list []overhead.OverheadDailyStatistics, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&overhead.OverheadDailyStatistics{})
	var DailyStatisticss []overhead.OverheadDailyStatistics
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&DailyStatisticss).Error
	return DailyStatisticss, total, err
}
