package overhead

import (
	"life/global"
	"life/model/common/request"
	"life/model/overhead"
	overheadReq "life/model/overhead/request"
)

type OverheadMonthStatisticsService struct {
}

// CreateOverheadMonthStatistics 创建OverheadMonthStatistics记录
// Author [piexlmax](https://github.com/piexlmax)
func (MonthStatisticsService *OverheadMonthStatisticsService) CreateOverheadMonthStatistics(MonthStatistics overhead.OverheadMonthStatistics) (err error) {
	err = global.GVA_DB.Create(&MonthStatistics).Error
	return err
}

// DeleteOverheadMonthStatistics 删除OverheadMonthStatistics记录
// Author [piexlmax](https://github.com/piexlmax)
func (MonthStatisticsService *OverheadMonthStatisticsService) DeleteOverheadMonthStatistics(MonthStatistics overhead.OverheadMonthStatistics) (err error) {
	err = global.GVA_DB.Delete(&MonthStatistics).Error
	return err
}

// DeleteOverheadMonthStatisticsByIds 批量删除OverheadMonthStatistics记录
// Author [piexlmax](https://github.com/piexlmax)
func (MonthStatisticsService *OverheadMonthStatisticsService) DeleteOverheadMonthStatisticsByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]overhead.OverheadMonthStatistics{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOverheadMonthStatistics 更新OverheadMonthStatistics记录
// Author [piexlmax](https://github.com/piexlmax)
func (MonthStatisticsService *OverheadMonthStatisticsService) UpdateOverheadMonthStatistics(MonthStatistics overhead.OverheadMonthStatistics) (err error) {
	err = global.GVA_DB.Save(&MonthStatistics).Error
	return err
}

// GetOverheadMonthStatistics 根据id获取OverheadMonthStatistics记录
// Author [piexlmax](https://github.com/piexlmax)
func (MonthStatisticsService *OverheadMonthStatisticsService) GetOverheadMonthStatistics(id uint) (MonthStatistics overhead.OverheadMonthStatistics, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&MonthStatistics).Error
	return
}

// GetOverheadMonthStatisticsInfoList 分页获取OverheadMonthStatistics记录
// Author [piexlmax](https://github.com/piexlmax)
func (MonthStatisticsService *OverheadMonthStatisticsService) GetOverheadMonthStatisticsInfoList(info overheadReq.OverheadMonthStatisticsSearch) (list []overhead.OverheadMonthStatistics, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&overhead.OverheadMonthStatistics{})
	var MonthStatisticss []overhead.OverheadMonthStatistics
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&MonthStatisticss).Error
	return MonthStatisticss, total, err
}
