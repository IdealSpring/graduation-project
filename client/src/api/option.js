/**
 * “选项”相关接口
 */
import request from '@/utils/request'

export default {

  listRoleOptions() {
    return request({
      url: '/auth/user/option/role',
      method: 'get',
    })
  }

}





