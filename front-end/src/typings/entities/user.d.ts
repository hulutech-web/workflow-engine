/// <reference path="../global.d.ts"/>

/** 用户数据库表字段 */
namespace Entity {
  interface User {
    /** 用户id */
    id?: number
    /** 用户名 */
    username?: string
    /** 昵称 */
    nickname?: string
    /** 头像 */
    avatar?: string
    /** 角色id */
    roleId?: string
    /** 岗位id */
    postId?: string
    /** 部门 */
    dept?: string
    /** 租户 */
    tenant?: string
    /** 是否多点登录 */
    isMultipoint?: number
    /** 是否禁用 */
    isDisable?: number
    /** 最后登录ip */
    lastLoginIp?: string
    /** 最后登录时间 */
    lastLoginTime?: string
    /** 超级管理员 */
    softSuper?: boolean
    /** 超级租户 */
    superTenant?: boolean
    /** 创建时间 */
    createTime?: string
    /** 更新时间 */
    updateTime?: string
  }
}
