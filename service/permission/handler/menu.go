package handler

import (
	"context"
	"edas/service/permission/db"
	proto "edas/service/permission/proto"
	"edas/share/errors"
	"edas/share/util"
	"go.uber.org/zap"
)

// PermissionServiceHandler 结构体
type PermissionServiceHandler struct {
	logger *zap.Logger
}

func (p *PermissionServiceHandler) CreateMenu(ctx context.Context, req *proto.CreateMenuRequest, res *proto.CreateMenuResponse) error {
	// 1. 获取请求参数
	name := req.Name
	sequence := req.Sequence
	icon := req.Icon
	router := req.Router
	hidden := req.Hidden
	parentId := req.ParentId
	parentPath := req.ParentPath
	creator := req.Creator
	// 2. 根据name判断是否已经注册过
	menu, err := db.SelectMenuByName(name)
	if err != nil {
		p.logger.Error("error", zap.Error(err))
		return errors.ErrorMenuFailed
	}
	if menu != nil {
		return errors.ErrorMenuAlreadyExists
	}
	// 3. 未查询到记录就进行创建
	record := util.MustUUID()
	err = db.InsertMenu(sequence, hidden, icon, router, parentId, parentPath, creator, record, name)
	if err != nil {
		p.logger.Error("error", zap.Error(err))
		return errors.ErrorMenuFailed
	}
	return nil
}

func (p *PermissionServiceHandler) DeleteMenu(ctx context.Context, req *proto.DeleteMenuRequest, res *proto.DeleteMenuResponse) error {
	record := req.Record
	// 1. 判断资源是否存在
	oldItem, err := db.SelectMenuByRecord(record)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrorNotFoundMenuItem
	}

	// 2. 若该菜单含有子级，不能删除
	result, err := db.SelectMenuByParentID(record)
	if err != nil {
		p.logger.Error("error", zap.Error(err))
		return errors.ErrorMenuFailed
	} else if len(result) > 0 {
		return errors.ErrorNotAllowDeleteWithChild
	}

	// 3. 正常删除菜单
	err = db.DeleteMenu(record)
	if err != nil {
		p.logger.Error("error", zap.Any("删除菜单失败", err))
		return errors.ErrorMenuFailed
	}
	return nil
}

func (p *PermissionServiceHandler) UpdateMenu(ctx context.Context, req *proto.UpdateMenuRequest, res *proto.UpdateMenuResponse) error {
	// 1. 判断资源是否存在
	oldItem, err := db.SelectMenuByRecord(req.Record)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrorNotFoundMenuItem
	} else if req.Name != "" && oldItem.Name != req.Name {
		// 2. 根据name判断是否已经注册过
		name, err := db.SelectMenuByName(req.Name)
		if err != nil {
			p.logger.Error("error", zap.Error(err))
			return errors.ErrorMenuFailed
		}
		if name != nil {
			return errors.ErrorMenuAlreadyExists
		}
	}
	// 3. 更新菜单
	menu := db.Menu{
		RecordId:   req.Record,
		Name:       req.Name,
		Sequence:   req.Sequence,
		Icon:       req.Icon,
		Router:     req.Router,
		Hidden:     req.Hidden,
		ParentId:   req.ParentId,
		ParentPath: req.ParentPath,
		Creator:    req.Creator,
	}
	err = db.UpdateMenu(&menu)
	if err != nil {
		p.logger.Error("error", zap.Any("更新菜单失败", err))
		return errors.ErrorUpdateMenuFailed
	}
	return nil
}

func (p *PermissionServiceHandler) QueryMenu(ctx context.Context, req *proto.QueryMenuRequest, res *proto.QueryMenuResponse) error {

	return nil
}

func (p *PermissionServiceHandler) GetMenu(ctx context.Context, req *proto.GetMenuRequest, res *proto.GetMenuResponse) error {
	menu, err := db.SelectMenuByRecord(req.Record)
	if err != nil {
		return err
	} else if menu == nil {
		return errors.ErrorNotFoundMenuItem
	}
	res.Record = menu.RecordId
	res.Name = menu.Name
	res.Sequence = menu.Sequence
	res.Icon = menu.Icon
	res.Router = menu.Router
	res.Hidden = menu.Hidden
	res.ParentId = menu.ParentId
	res.ParentPath = menu.ParentPath
	res.Creator = menu.Creator
	return nil
}
