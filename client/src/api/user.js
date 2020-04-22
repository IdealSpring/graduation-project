/**
 * “用户管理”相关接口
 */
import request from '@/utils/request'

export default {

  // 查询用户
  queryUser(queryParam,pageParam) {
    return request({
      // url: '/sys_user/query',
      url: '/auth/user/query',
      method: 'post',
      data: {
        ...queryParam,
        current: pageParam.current,
        size: pageParam.size
      }
    })
  },
  // 更新用户信息
  updateUser(data) {
    return request({
      url: '/auth/user/update',
      method: 'put',
      data
    })
  },

  // 修改密码
  updatePwd(data) {
    return request({
      url: '/auth/user/pwd',
      method: 'put',
      data
    })
  },

  // 添加用户
  addUser(data) {
    return request({
      url: '/auth/user/addUser',
      method: 'post',
      data
    })
  },

  // 删除用户
  deleteUser(data) {
    return request({
      // url: '/auth/user/delete',
      url: '/auth/user/delete/' + data,
      method: 'delete',
      // data
    })
  },

  /**
   * 更新用户的角色
   * @param perm
   */
  updateUserRoles(data) {
    return request({
      url: '/sys_user/role',
      method: 'patch',
      data
    })
  }

}

