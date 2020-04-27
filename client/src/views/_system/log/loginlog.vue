<template>
  <div class="app-container">
    <!--查询-->
    <el-row>
      <el-input v-model="tableQuery.provinceName" placeholder="请输入登录地点" style="width: 20%;"/>
      <el-input v-model="tableQuery.provinceName2" placeholder="请输入用户名" style="width: 20%;"/>
      <span style="margin-right: 15px;"/>
      <el-tooltip content="搜索" placement="top">
        <el-button icon="el-icon-search" circle @click="fetchData()"></el-button>
      </el-tooltip>
    </el-row>

    <!--占位框-->
    <div style="margin-bottom: 30px;"></div>

    <!--按钮-->
    <el-button type="danger" icon="el-icon-delete" size="mini">清空</el-button>


    <!--占位框-->
    <div style="margin-bottom: 30px;"></div>

    <!--表格-->
    <el-table style="width: 100%;"
              border highlight-current-row
              v-loading.body="tableLoading"
              element-loading-text="Loading"
              :data="tableData">
      <el-table-column prop="provinceId" label="日志ID" width="65"/>
      <el-table-column prop="provinceName" label="用户名"/>
      <el-table-column prop="userSize" label="昵称"/>
      <el-table-column prop="userSize" label="登录地址"/>
      <el-table-column prop="userSize" label="登录地点"/>
      <el-table-column prop="userSize" label="浏览器"/>
      <el-table-column prop="userSize" label="操作系统"/>

      <el-table-column prop="createTime" label="登录时间">
        <template slot-scope="scope">
          <span v-text="parseTime(scope.row.createTime)"></span>
        </template>
      </el-table-column>

      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-tooltip content="删除" placement="top">
            <el-button @click="" size="medium" type="danger" icon="el-icon-remove"
                       circle plain />
          </el-tooltip>
        </template>
      </el-table-column>
    </el-table>

    <!--占位框-->
    <div style="margin-bottom: 30px;"></div>

    <!--分页-->
    <el-pagination
      @size-change="handleSizeChange"
      @current-change="handleCurrentChange"
      :current-page="tablePage.current"
      :page-sizes="[10, 20, 30, 40, 50]"
      :page-size="tablePage.size"
      layout="total, sizes, prev, pager, next, jumper"
      :total="tablePage.total">
    </el-pagination>

  </div>
</template>

<script>
  import provinceApi from '@/api/province'
  import { parseTime } from '@/utils'
  import { pageParamNames, confirm, root } from '@/utils/constants'

  export default {
    name: 'ReleaseManage',

    data() {
      return {
        tableQuery: {
          provinceName: null,
          provinceName2: null,
        },
        tableLoading: false,
        tableData: [],
        tablePage: {
          current: null,
          pages: null,
          size: null,
          total: null
        }
      }
    },

    created() {
      this.fetchData()
    },

    watch: {},

    methods: {
      // 解析时间
      parseTime,

      // 有操作权限
      hasAdminRole(row) {
        if (this.currentRole.priority > row.Role.priority) {
          return true
        }
        return false
      },

      // 分页
      handleSizeChange(val) {
        this.tablePage.size = val
        this.fetchData()
      },
      handleCurrentChange(val) {
        this.tablePage.current = val
        this.fetchData()
      },

      // 获取表格数据
      fetchData(current) {
        if (current) {
          this.tablePage.current = current
        }

        this.tableLoading = true
        provinceApi.fetchDataToPage(this.tableQuery, this.tablePage).then(res => {
          let page = res.data.page

          this.tableData = page.records
          this.tablePage.current = page.current
          this.tablePage.pages = page.pages
          this.tablePage.size = page.size
          this.tablePage.total = page.total

          this.tableLoading = false
        }).catch(error => {
          this.$message({
            showClose: true,
            message: '数据加载失败: ' + error,
            type: 'error'
          })
        })
      }
    }
  }
</script>

<style rel="stylesheet/scss" lang="scss" scoped>

</style>
