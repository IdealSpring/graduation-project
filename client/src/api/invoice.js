/**
 * 纳税人档案管理 接口
 */
import request from '@/utils/request'

export default {
  // 删除单条记录
  deleteInvoice(data) {
    return request({
      url: '/auth/data/invoice/delete/' + data,
      method: 'delete'
    })
  },

  // 上传文件
  uploadFile(data) {
    return request({
      url: '/auth/data/invoice/uploadFile',
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
      url: '/auth/data/invoice/queryInvoiceToPage',
      method: 'post',
      data: {
        ...queryParam,
        current: pageParam.current,
        size: pageParam.size
      }
    })
  }

}
