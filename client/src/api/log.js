/**
 * 日志 接口
 */
import request from '@/utils/request'

export default {
  // 删除所有操作日志
  deleteAllOperationLog() {
    return request({
      url: '/auth/log/operation/deleteAll',
      method: 'delete',
    })
  },

  // 删除所有登录日志
  deleteAllLoginLog() {
    return request({
      url: '/auth/log/loggin/deleteAll',
      method: 'delete',
    })
  },

  // 删除单条记录
  deleteLog(data) {
    return request({
      url: '/auth/log/operation/delete/' + data,
      method: 'delete',
    })
  },

  // 获取分页数据
  fetchDataToPage(queryParam, pageParam) {
    return request({
      url: '/auth/log/operation/queryDataToPage',
      method: 'post',
      data: {
        ...queryParam,
        current: pageParam.current,
        size: pageParam.size
      }
    })
  }

}
