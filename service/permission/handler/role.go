package handler

import (
	"context"
	"edas/service/permission/db"
	proto "edas/service/permission/proto"
	"edas/share/errors"
	"edas/share/util"
	"go.uber.org/zap"
)

func (p *PermissionServiceHandler) CreateRole(ctx context.Context, req *proto.CreateRoleRequest, res *proto.CreateRoleResponse) error {
	// 1. 判断资源是否存在
	oldItem, err := db.SelectRoleByName(req.Name)
	if err != nil {
		p.logger.Error("error", zap.Error(err))
		return errors.ErrorRoleFailed
	}
	if oldItem != nil {
		return errors.ErrorRoleAlreadyExists
	}
	// 2. 创建角色
	record := util.MustUUID()
	role := db.Role{
		Record:   record,
		Name:     req.Name,
		Sequence: req.Sequence,
		Memo:     req.Memo,
		Creator:  req.Creator,
	}
	err = db.InsertRole(&role)
	if err != nil {
		p.logger.Error("error", zap.Any("创建角色失败", err))
		return errors.ErrorCreateRoleFailed
	}
	return nil
}

func (p *PermissionServiceHandler) DeleteRole(ctx context.Context, req *proto.DeleteRoleRequest, res *proto.DeleteRoleResponse) error {
	// 1. 判断资源是否存在
	oldItem, err := db.SelectRoleByRecord(req.Record)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrorNotFoundRole
	}
	// 2. 删除角色
	err = db.DeleteRole(req.Record)
	if err != nil {
		p.logger.Error("error", zap.Any("删除角色失败", err))
		return errors.ErrorDeleteRoleFailed
	}
	return nil
}

func (p *PermissionServiceHandler) UpdateRole(ctx context.Context, req *proto.UpdateRoleRequest, res *proto.UpdateRoleResponse) error {
	// 1. 判断资源是否存在
	oldItem, err := db.SelectRoleByRecord(req.Record)
	if err != nil {
		return err
	} else if oldItem == nil {
		return errors.ErrorNotFoundRole
	} else if req.Name != "" && req.Name != oldItem.Name {
		// 2. 根据name判断是否已经注册过
		item, err := db.SelectRoleByName(req.Name)
		if err != nil {
			p.logger.Error("error", zap.Error(err))
			return errors.ErrorRoleFailed
		}
		if item != nil {
			return errors.ErrorRoleAlreadyExists
		}
	}
	// 3. 更新角色信息
	role := db.Role{
		Record:   req.Record,
		Name:     req.Name,
		Sequence: req.Sequence,
		Memo:     req.Memo,
		Creator:  req.Creator,
	}
	err = db.UpdateRole(&role)
	if err != nil {
		p.logger.Error("error", zap.Any("更新角色失败", err))
		return errors.ErrorUpdateRoleFailed
	}
	return nil
}

func (p *PermissionServiceHandler) QueryRole(ctx context.Context, req *proto.QueryRoleRequest, res *proto.QueryRoleResponse) error {
	return nil
}

func (p *PermissionServiceHandler) GetRole(ctx context.Context, req *proto.GetRoleRequest, res *proto.GetRoleResponse) error {
	if req.Record != "" {
		// 1. 根据RecordId判断资源是否存在
		role, err := db.SelectRoleByRecord(req.Record)
		if err != nil {
			return err
		} else if role == nil {
			return errors.ErrorNotFoundRole
		}
		// 2. 构造返回结构
		res.Record = role.Record
		res.Name = role.Name
		res.Sequence = role.Sequence
		res.Memo = role.Memo
		res.Creator = role.Creator
		return nil
	} else if req.Name != "" {
		// 1. 根name判断资源是否存在
		role, err := db.SelectRoleByRecord(req.Name)
		if err != nil {
			return err
		} else if role == nil {
			return errors.ErrorNotFoundRole
		}
		// 2. 构造返回结构
		res.Record = role.Record
		res.Name = role.Name
		res.Sequence = role.Sequence
		res.Memo = role.Memo
		res.Creator = role.Creator
		return nil
	}
	return nil
}
