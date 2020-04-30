/**
 * 纳税人档案管理 接口
 */
import request from '@/utils/request'

export default {
  // 删除单条记录
  deletePolitics(data) {
    return request({
      url: '/auth/data/taxpayer/delete/' + data,
      method: 'delete'
    })
  },

  // 上传文件
  uploadFile(data) {
    return request({
      url: '/auth/data/taxpayer/uploadFile',
      method: 'post',
      data: data,
      headers: {
        'Content-Type': 'multipart/form-data;charset=utf-8'
      }
    })

  },

  // 获取分页数据
  fetchDataToPage(queryParam, pageParam) {
    return request({
      url: '/auth/data/taxpayer/queryTaxpayerToPage',
      method: 'post',
      data: {
        ...queryParam,
        current: pageParam.current,
        size: pageParam.size
      }
    })
  }

}
