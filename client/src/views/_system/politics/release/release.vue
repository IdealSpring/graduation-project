<template>
  <div class="app-container">
    <!--查询-->
    <el-row>
      <el-input v-model="tableQuery.provinceName" placeholder="政策名称" style="width: 20%;"/>
      <span style="margin-right: 15px;"/>
      <el-tooltip content="搜索" placement="top">
        <el-button icon="el-icon-search" circle @click="fetchData()"></el-button>
      </el-tooltip>
    </el-row>

    <!--占位框-->
    <div style="margin-bottom: 30px;"></div>

    <!--按钮-->
    <el-button type="primary" icon="el-icon-plus" size="mini">添加省份</el-button>

    <!--占位框-->
    <div style="margin-bottom: 30px;"></div>

    <!--表格-->
    <el-table style="width: 100%;"
              border highlight-current-row
              v-loading.body="tableLoading"
              element-loading-text="Loading"
              :data="tableData">
      <el-table-column prop="provinceId" label="政策ID"/>
      <el-table-column prop="provinceName" label="标题"/>
      <el-table-column prop="userSize" label="附件文件"/>

      <el-table-column prop="createTime" label="创建时间">
        <template slot-scope="scope">
          <span v-text="parseTime(scope.row.createTime)"></span>
        </template>
      </el-table-column>
      <el-table-column prop="updateTime" label="更新时间">
        <template slot-scope="scope">
          <span v-text="parseTime(scope.row.updateTime)"></span>
        </template>
      </el-table-column>

      <el-table-column label="操作">
        <template slot-scope="scope">
          <el-tooltip content="编辑" placement="top">
            <el-button size="medium" type="primary" icon="el-icon-edit" circle plain/>
          </el-tooltip>

          <el-tooltip content="详情" placement="top">
            <el-button size="medium" type="info" icon="el-icon-info" circle plain/>
          </el-tooltip>

          <el-tooltip content="删除" placement="top">
            <el-button size="medium" type="danger" icon="el-icon-delete" circle plain/>
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

    <!--弹出窗口：新增/编辑政策-->
    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible" width="30%">
      <el-form :model="temp" label-position="left" label-width="120px">

        <el-form-item label="登录名" prop="uname" v-if="dialogStatus=='create'">
          <el-input v-model="temp.username"></el-input>
        </el-form-item>

        <el-form-item label="昵称" prop="nick">
          <el-input v-model="temp.nick"></el-input>
        </el-form-item>

        <el-form-item label="角色" v-if="dialogStatus=='create'">
          <el-select v-model="temp.roleId" clearable placeholder="请选择">
            <el-option
              :key="item.roleId"
              :label="item.roleName"
              :value="item.roleId">
            </el-option>
          </el-select>
        </el-form-item>

        <el-form-item label="密码" prop="pwd">
          <el-input type="password" v-model="temp.pwd"></el-input>
        </el-form-item>

        <el-form-item label="确认密码" prop="pwd2">
          <el-input type="password" v-model="temp.pwd2"></el-input>
        </el-form-item>

      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button v-if="dialogStatus=='create'" type="primary" @click="">创建</el-button>
        <el-button v-else type="primary" @click="">确定</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script>
  import provinceApi from '@/api/province'
  import { parseTime } from '@/utils'
  import { pageParamNames, confirm, root } from '@/utils/constants'

  export default {
    name: 'PoliticsManage',

    data() {
      return {
        tableQuery: {
          provinceName: null
        },
        tableLoading: false,
        tableData: [],
        tablePage: {
          current: null,
          pages: null,
          size: null,
          total: null
        },
        textMap: {
          update: '编辑政策',
          create: '新增政策'
        },
        dialogStatus: '',
        dialogFormVisible: false,
        temp: {
          idx: null, //tableData中的下标
          userId: null,
          username: null,
          nick: null,
          roleId: null,
          pwd: null,
          pwd2: null,
          roleId: null,
          createTime: null,
          updateTime: null
        },
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
