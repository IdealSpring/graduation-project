/**
 * 税收政策下发 接口
 */
import request from '@/utils/request'

export default {
  // 用户获取未读Politics
  getPoliticsNotify(data) {
    return request({
      url: '/auth/politics/release/politicsNotify',
      method: 'post',
      data
    })
  },

  // 删除单条记录
  deletePolitics(data) {
    return request({
      url: '/auth/politics/release/politics/delete/' + data,
      method: 'delete',
    })
  },

  // 添加政策
  addPolitics(data) {
    return request({
      url: '/auth/politics/release/addPolitics',
      method: 'post',
      data
    })
  },

  // 获取分页数据
  fetchDataToPage(queryParam, pageParam) {
    return request({
      url: '/auth/politics/release/queryPoliticsToPage',
      method: 'post',
      data: {
        ...queryParam,
        current: pageParam.current,
        size: pageParam.size
      }
    })
  }

}
