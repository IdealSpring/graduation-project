/**
 * 发行管理 API 数据接口
 */
import request from '@/utils/request'

export default {
  // 添加省份
  addProvince(param) {
    return request({
      url: '/auth/release/addProvince',
      method: 'post',
      data: {
        ...param
      }
    })
  },

  // 获取分页数据
  fetchDataToPage(queryParam, pageParam) {
    return request({
      url: '/auth/release/queryProvince',
      method: 'post',
      data: {
        ...queryParam,
        current: pageParam.current,
        size: pageParam.size
      }
    })
  }
}
