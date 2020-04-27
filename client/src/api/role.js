/**
 * “角色管理”相关接口
 */
import request from '@/utils/request'

export default {
  // 添加角色
  addRole(data) {
    return request({
      url: '/auth/user/addRole',
      method: 'post',
      data
    })
  },

  // 删除角色
  deleteRole(data) {
    return request({
      url: '/auth/user/role/delete',
      method: 'delete',
      data
    })
  },

  // 查询角色
  queryRole(queryParam, pageParam) {
    return request({
      url: '/auth/user/roleQuery',
      method: 'post',
      data: {
        ...queryParam,
        current: pageParam.current,
        size: pageParam.size
      }
    })
  },

  // 更新角色
  updateRole(data) {
    return request({
      url: '/auth/user/role/updateRole',
      method: 'put',
      data
    })
  },

  // 更新角色的权限
  updateRolePerms(data) {
    return request({
      url: '/auth/user/role/updatePerm',
      method: 'put',
      data
    })
  },

  /**
   * 添加角色的权限
   * @param perm
   */
  addRolePerm(data) {
    return request({
      url: '/sys_role/perm',
      method: 'post',
      data
    })
  },

  /**
   * 删除角色的权限
   * @param perm
   */
  deleteRolePerm(data) {
    return request({
      url: '/sys_role/perm',
      method: 'delete',
      data
    })
  },

  // 查选角色的所有权限值
  findRolePerms(rid) {
    return request({
      url: '/auth/user/rolePerms/' + rid,
      method: 'get'
    })
  }

}

