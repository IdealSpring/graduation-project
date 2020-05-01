/**
 * 纳税人档案管理 接口
 */
import request from '@/utils/request'

export default {

  // 下载
  downloadFile() {
    return request({
      url: '/auth/predicte/fileDownload',
      method: 'post',
      responseType: 'blob'
    })
  }

}
