import Vue from 'vue'
import Router from 'vue-router'

const _import = require('./_import_' + process.env.NODE_ENV)
// in development-env not use lazy-loading, because lazy-loading too many pages will cause webpack hot update too slow. so only in production use lazy-loading;
// detail: https://panjiachen.github.io/vue-element-admin-site/#/lazy-loading

Vue.use(Router)

/* Layout */
import Layout from '../views/layout/Layout'

/** note: submenu only apppear when children.length>=1
 *   detail see  https://panjiachen.github.io/vue-element-admin-site/#/router-and-nav?id=sidebar
 **/

/**
 * hidden: true                   if `hidden:true` will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu, whatever its child routes length
 *                                if not set alwaysShow, only more than one route under the children
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noredirect           if `redirect:noredirect` will no redirct in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']     will control the page roles (you can set multiple roles)
    title: 'title'               the name show in submenu and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar,
    noCache: true                if true ,the page will no be cached(default is false)
  }
 **/
export const constantRouterMap = [
  { path: '/login', component: _import('login/index'), hidden: true },
  { path: '/authredirect', component: _import('login/authredirect'), hidden: true },
  { path: '/404', component: _import('errorPage/404'), hidden: true },
  { path: '/401', component: _import('errorPage/401'), hidden: true },
  {
    path: '',
    component: Layout,
    redirect: 'dashboard',
    children: [{
      path: 'dashboard',
      component: _import('dashboard/index'),
      name: 'dashboard',
      meta: { title: '首页', icon: 'dashboard', noCache: true }
    }]
  }
]

export default new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRouterMap
})

export const asyncRouterMap = [

  {
    path: '/user',
    component: Layout,
    meta: { perm: 'm:user', title: '用户管理', icon: 'peoples' },
    children: [
      {
        path: 'user_manage',
        name: 'user_manage',
        component: _import('_system/user/index'),
        meta: { perm: 'm:user:info', title: '信息管理', noCache: true }
      },
      {
        path: 'role_manage',
        name: 'role_manage',
        component: _import('_system/role/index'),
        meta: { perm: 'm:user:role', title: '角色管理', noCache: true }
      },
      {
        hidden: true,
        path: 'role_manage/:roleId/assign_perm',
        name: 'role_manage_assign_perm',
        component: _import('_system/role/assign_perm'),
        meta: { hiddenTag: true, title: '角色授权' }
      }
    ]
  },
  {
    path: '/release',
    component: Layout,
    meta: { perm: 'm:release', title: '发行管理', icon: 'international' },
    children: [
      {
        path: 'release_manage',
        name: 'release_manage',
        component: _import('_system/release/sys/index'),
        meta: { perm: 'm:release:sys', title: '系统发行', noCache: true }
      },
      {
        path: 'show_manage',
        name: 'show_manage',
        component: _import('_system/release/show/index'),
        meta: { perm: 'm:release:show', title: '发行省份', noCache: true }
      }
    ]
  },
  {
    path: '/predicte_analysis',
    component: Layout,
    meta: { perm: 'm:predicte_analysis', title: '预测分析', icon: 'documentation' },
    children: [
      {
        path: 'predicte_manage',
        name: 'predicte_manage',
        component: _import('errorPage/404'),
        meta: { perm: 'm:predicte_analysis:predicte', title: '异常企业预测', noCache: true }
      },
      {
        path: 'analysis_manage',
        name: 'analysis_manage',
        component: _import('errorPage/404'),
        meta: { perm: 'm:predicte_analysis:analysis', title: '异常企业分析', noCache: true }
      }
    ]
  },
  {
    path: '/statistics',
    component: Layout,
    alwaysShow: true,
    meta: { perm: 'm:statistics', title: '数据统计', icon: 'chart' },
    children: [
      {
        path: 'statistics_provincialtax_manage',
        name: 'statistics_provincialtax_manage',
        component: _import('_system/statistics/statistics'),
        meta: { perm: 'm:statistics:provincialtax', title: '各省税收数据', noCache: true }
      },
      {
        path: 'statistics_areataxdata_manage',
        name: 'statistics_areataxdata_manage',
        component: _import('_system/statistics/statistics'),
        meta: { perm: 'm:statistics:areataxdata', title: '各地区税收数据', noCache: true }
      },
      {
        path: 'statistics_enterprisedata_manage',
        name: 'statistics_enterprisedata_manage',
        component: _import('_system/statistics/statistics'),
        meta: { perm: 'm:statistics:enterprisedata', title: '企业数据', noCache: true }
      }
    ]
  },
  {
    path: '/politics',
    component: Layout,
    meta: { perm: 'm:politics', title: '政务管理', icon: 'message' },
    children: [
      {
        path: 'politics_release_manage',
        name: 'politics_release_manage',
        component: _import('_system/politics/release/release'),
        meta: { perm: 'm:politics:release', title: '税收政策下发', noCache: true }
      },
      {
        path: 'politics_guide_manage',
        name: 'politics_guide_manage',
        component: _import('_system/politics/release/report-guide'),
        meta: { perm: 'm:politics:guide', title: '业务指导协调', noCache: true }
      },
      {
        path: 'politics_upRelease_manage',
        name: 'politics_upRelease_manage',
        component: _import('_system/politics/release/release'),
        meta: { perm: 'm:politics:upRelease', title: '上级政策文件', noCache: true }
      },
      {
        path: 'politics_report_manage',
        name: 'politics_report_manage',
        component: _import('_system/politics/release/report-guide'),
        meta: { perm: 'm:politics:report', title: '工作汇报', noCache: true }
      },
      {
        path: 'politics_notify_manage',
        name: 'politics_notify_manage',
        component: _import('_system/politics/release/notify'),
        meta: { perm: 'm:politics:notify', title: '发布通知', noCache: true }
      }
    ]
  },
  {
    path: '/data',
    component: Layout,
    meta: { perm: 'm:data', title: '数据管理', icon: 'chart' },
    children: [
      {
        path: 'data_archives_manage',
        name: 'data_archives_manage',
        component: _import('_system/data/taxpayer'),
        meta: { perm: 'm:data:archives', title: '纳税人档案', noCache: true }
      },
      {
        path: 'data_invoice_manage',
        name: 'data_invoice_manage',
        component: _import('_system/data/invoice'),
        meta: { perm: 'm:data:invoice', title: '发票数据', noCache: true }
      },
      {
        path: 'data_details_manage',
        name: 'data_details_manage',
        component: _import('_system/data/details'),
        meta: { perm: 'm:data:details', title: '发票明细', noCache: true }
      }
    ]
  },
  {
    path: '/log',
    component: Layout,
    meta: { perm: 'm:log', title: '日志管理', icon: 'documentation' },
    children: [
      {
        path: 'log_loginlog_manage',
        name: 'log_loginlog_manage',
        component: _import('_system/log/loginlog'),
        meta: { perm: 'm:log:loginlog', title: '登录日志', noCache: true }
      },
      {
        path: 'log_operlog_manage',
        name: 'log_operlog_manage',
        component: _import('_system/log/operlog'),
        meta: { perm: 'm:log:operlog', title: '操作日志', noCache: true }
      }
    ]
  },

  /*
  {
    path: '/menu1',
    component: Layout,
    children: [{
      path: 'index',
      name: 'menu1',
      component: _import('menu/menu1'),
      meta: { perm:'m:menu1', title: '菜单1', icon: 'icon' }
    }]
  },


  {
    path: '/menu2',
    component: Layout,
    children: [{
      path: 'index',
      name: 'menu2',
      component: _import('menu/menu2'),
      meta: { perm:'m:menu2', title: '菜单2', icon: 'icon' }
    }]
  },

  {
    path: '/menu3',
    component: Layout,
    meta: {
      perm:'m:menu3',
      title: '菜单3',
      icon: 'chart'
    },
    children: [
      { path: 'menu3_1', component: _import('menu/menu3_1'), name: 'menu3_1', meta: { perm:'m:menu3:1', title: '菜单3-1', icon: 'chart', noCache: true }},
      { path: 'menu3_2', component: _import('menu/menu3_2'), name: 'menu3_2', meta: { perm:'m:menu3:2', title: '菜单3-2', icon: 'chart', noCache: true }},
      { path: 'menu3_3', component: _import('menu/menu3_3'), name: 'menu3_3', meta: { perm:'m:menu3:3', title: '菜单3-3', icon: 'chart', noCache: true }}
    ]
  },


  {
    path: '/menu4',
    name: 'menu4',
    component: Layout,
    redirect: '/menu4/menu4_1/a',
    meta: {
      perm:'m:menu4',
      title: '菜单4',
      icon: 'example'
    },
    children: [
      {
        path: '/menu4/menu4_1',
        name: 'menu4_1',
        component: _import('menu/menu4_1/index'),
        redirect: '/menu4/menu4_1/a',
        meta: {
          perm:'m:menu4:1',
          title: '菜单4-1',
          icon: 'table'
        },
        children: [
          { path: 'a', name: 'menu4_1_a', component: _import('menu/menu4_1/a'), meta: { perm:'m:menu4:1:a', title: '菜单4-1-a' }},
          { path: 'b', name: 'menu4_1_b', component: _import('menu/menu4_1/b'), meta: { perm:'m:menu4:1:b', title: '菜单4-1-b' }},
          { path: 'c', name: 'menu4_1_c', component: _import('menu/menu4_1/c'), meta: { perm:'m:menu4:1:c', title: '菜单4-1-c' }}
        ]
      },
      { path: 'menu4/menu4_2', name: 'menu4_2', icon: 'tab', component: _import('menu/menu4_2/index'), meta: {perm:'m:menu4:2', title: '菜单4-2' }}
    ]
  },

  { path: '*', redirect: '/404', hidden: true }*/
]
