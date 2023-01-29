package overhead

import (
	"life/global"
	"life/model/common/request"
	"life/model/overhead"
	overheadReq "life/model/overhead/request"
)

type OverheadCategoryService struct {
}

// CreateOverheadCategory 创建OverheadCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *OverheadCategoryService) CreateOverheadCategory(category overhead.OverheadCategory) (err error) {
	err = global.GVA_DB.Create(&category).Error
	return err
}

// DeleteOverheadCategory 删除OverheadCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *OverheadCategoryService) DeleteOverheadCategory(category overhead.OverheadCategory) (err error) {
	err = global.GVA_DB.Delete(&category).Error
	return err
}

// DeleteOverheadCategoryByIds 批量删除OverheadCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *OverheadCategoryService) DeleteOverheadCategoryByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]overhead.OverheadCategory{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateOverheadCategory 更新OverheadCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *OverheadCategoryService) UpdateOverheadCategory(category overhead.OverheadCategory) (err error) {
	err = global.GVA_DB.Save(&category).Error
	return err
}

// GetOverheadCategory 根据id获取OverheadCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *OverheadCategoryService) GetOverheadCategory(id uint) (category overhead.OverheadCategory, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&category).Error
	return
}

// GetOverheadCategoryInfoList 分页获取OverheadCategory记录
// Author [piexlmax](https://github.com/piexlmax)
func (categoryService *OverheadCategoryService) GetOverheadCategoryInfoList(info overheadReq.OverheadCategorySearch) (list []overhead.OverheadCategory, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&overhead.OverheadCategory{})
	var categorys []overhead.OverheadCategory
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&categorys).Error
	return categorys, total, err
}
