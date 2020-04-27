<template>
  <div class="app-container">
    <!--查询-->
    <el-row>
      <el-input v-model="tableQuery.moduleName" placeholder="请输入模块" style="width: 20%;"/>
      <el-input v-model="tableQuery.username" placeholder="请输入操作人员" style="width: 20%;"/>
      <span style="margin-right: 15px;"/>
      <el-tooltip content="搜索" placement="top">
        <el-button icon="el-icon-search" circle @click="fetchData()"></el-button>
      </el-tooltip>
    </el-row>

    <!--占位框-->
    <div style="margin-bottom: 30px;"></div>

    <!--按钮-->
    <el-button type="danger" icon="el-icon-delete" size="mini" @click="deleteAllOperationLog()">清空</el-button>


    <!--占位框-->
    <div style="margin-bottom: 30px;"></div>

    <!--表格-->
    <el-table style="width: 100%;"
              border highlight-current-row
              v-loading.body="tableLoading"
              element-loading-text="Loading"
              :data="tableData">
      <el-table-column prop="logId" label="日志ID" width="65" align="center"/>
      <el-table-column prop="moduleName" label="系统模块" align="center"/>
      <el-table-column prop="operationType" label="操作类型" align="center"/>
      <el-table-column prop="method" label="请求方式" align="center"/>
      <el-table-column prop="ip" label="主机" align="center"/>
      <el-table-column prop="uri" label="操作接口" align="center"/>
      <el-table-column prop="status" label="操作状态" align="center"/>
      <el-table-column prop="username" label="操作人员" align="center"/>

      <el-table-column prop="createTime" label="操作时间" width="160" align="center">
        <template slot-scope="scope">
          <span v-text="parseTime(scope.row.createTime)"></span>
        </template>
      </el-table-column>

      <el-table-column label="操作" width="100" align="center">
        <template slot-scope="scope">
          <el-tooltip content="删除" placement="top">
            <el-button @click="handleDelete(scope.$index,scope.row)" size="medium" type="danger" icon="el-icon-remove"
                       circle plain/>
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
  import logApi from '@/api/log'
  import { parseTime } from '@/utils'
  import debounce from 'lodash/debounce'
  import { pageParamNames, confirm, root } from '@/utils/constants'

  export default {
    name: 'ReleaseManage',

    data() {
      return {
        tableQuery: {
          isLoginOrLogout: false,
          moduleName: null,
          username: null,
          nick: null
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

    watch: {
      //延时查询
      'tableQuery.moduleName': debounce(function() {
        this.fetchData()
      }, 500),
      'tableQuery.username': debounce(function() {
        this.fetchData()
      }, 500)
    },

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

      // 删除所有操作日志
      deleteAllOperationLog() {
        this.$confirm('您确定要清除所有操作日志？', '提示', confirm).then(() => {
          logApi.deleteAllOperationLog().then(res => {
            this.fetchData()

            this.$message.success('删除成功')
          })
        }).catch(() => {
          this.$message.info('已取消删除')
        })
      },

      // 删除
      handleDelete(idx, row) {
        this.$confirm('您确定要永久删除该条日志？', '提示', confirm).then(() => {
          logApi.deleteLog(row.logId).then(res => {
            this.tableData.splice(idx, 1)
            --this.tablePage.total
            this.$message.success('删除成功')
          })
        }).catch(() => {
          this.$message.info('已取消删除')
        })

      },

      // 获取表格数据
      fetchData(current) {
        if (current) {
          this.tablePage.current = current
        }

        this.tableLoading = true
        logApi.fetchDataToPage(this.tableQuery, this.tablePage).then(res => {
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
