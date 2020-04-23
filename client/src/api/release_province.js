/**
 * 发行管理 API 数据接口
 */
import request from '@/utils/request'

export default {
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
