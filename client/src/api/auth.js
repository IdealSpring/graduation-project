/**
 * 登录相关接口
 */
import request from '@/utils/request'

export default {

  // 通过用户名和密码登录
  loginByUsername(username, password) {
    const data = {
      username: username,
      password: password
    }
    return request({
      url: '/login',
      method: 'post',
      data
    })
  },

  // 用户退出
  logout() {
    return request({
      url: '/auth/user/logout',
      method: 'post'
    })
  },

  // 获取用户信息
  getUserInfo(token) {
    return request({
      url: '/auth/user/info',
      method: 'get'
    })
  },

}

