/**
 * “选项”相关接口
 */
import request from '@/utils/request'

export default {

  // 其获取系统所有角色
  listRoleOptions() {
    return request({
      url: '/auth/user/option/role',
      method: 'get',
    })
  },

  // 其获取系统所有角色
  listOprovinceOptions() {
    return request({
      url: '/auth/release/option/province',
      method: 'get',
    })
  }

}





