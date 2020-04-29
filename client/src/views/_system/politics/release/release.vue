<template>
  <div class="app-container">
    <!--查询-->
    <el-row>
      <el-input v-model="tableQuery.title" placeholder="请输入标题" style="width: 20%;"/>
      <span style="margin-right: 15px;"/>
      <el-tooltip content="搜索" placement="top">
        <el-button icon="el-icon-search" circle @click="fetchData(1)"></el-button>
      </el-tooltip>
    </el-row>

    <!--占位框-->
    <div style="margin-bottom: 30px;"></div>

    <!--按钮-->
    <el-button type="primary" icon="el-icon-plus" size="mini" v-if="hasFilePerm()" @click="handleCreate()">下发文件
    </el-button>

    <!--占位框-->
    <div style="margin-bottom: 30px;"></div>

    <!--表格-->
    <el-table style="width: 100%;"
              border highlight-current-row
              v-loading.body="tableLoading"
              element-loading-text="Loading"
              :data="tableData">
      <el-table-column prop="politicsId" label="政策ID" width="65" align="center"/>
      <el-table-column prop="title" label="标题" align="center"/>
      <el-table-column prop="abstract" label="正文摘要" align="center"/>
      <el-table-column prop="enclosurePath" label="附件" align="center"/>

      <el-table-column prop="createTime" label="创建时间" width="160" align="center">
        <template slot-scope="scope">
          <span v-text="parseTime(scope.row.createTime)"></span>
        </template>
      </el-table-column>

      <el-table-column label="操作" width="120" align="center">
        <template slot-scope="scope">
          <el-tooltip content="详情" placement="top">
            <el-button @click="handleShow(scope.$index,scope.row)" size="medium" type="primary" icon="el-icon-view" circle plain/>
          </el-tooltip>

          <el-tooltip content="删除" placement="top" v-if="hasFilePerm()">
            <el-button @click="handleDelete(scope.$index,scope.row)" size="medium" type="danger" icon="el-icon-delete" circle plain/>
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

    <el-dialog title="下发政策文件" :visible.sync="dialogFormVisible" width="60%">
      <el-form :model="temp" label-position="left" label-width="120px">

        <el-form-item label="标题">
          <el-input v-model="temp.title" placeholder="请输入标题"></el-input>
        </el-form-item>

        <el-form-item label="正文摘要">
          <el-input type="textarea" placeholder="请输入正文摘要"
                    :autosize="{ minRows: 2, maxRows: 4}"
                    v-model="temp.abstract" maxlength="100" show-word-limit>
          </el-input>
        </el-form-item>

        <el-form-item label="正文">
          <el-input type="textarea" placeholder="请输入正文"
                    :autosize="{ minRows: 9, maxRows: 12}"
                    v-model="temp.content">
          </el-input>
        </el-form-item>

        <el-form-item label="附件">
          <el-upload
            action="https://jsonplaceholder.typicode.com/posts/"
            multiple>
            <el-button size="small" type="primary">点击上传</el-button>
            <div slot="tip" class="el-upload__tip">只能上传不超过5M文件</div>
          </el-upload>
        </el-form-item>
      </el-form>

      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button type="primary" @click="createData()">创建</el-button>
      </div>
    </el-dialog>

    <el-dialog title="政策详情" :visible.sync="dialogDetialVisible" width="60%">
      <el-form :model="temp" label-position="left" label-width="120px">

        <el-form-item label="标题">
          <el-input v-model="temp.title" placeholder="请输入标题" :disabled="true"></el-input>
        </el-form-item>

        <el-form-item label="正文摘要">
          <el-input type="textarea" placeholder="请输入正文摘要"
                    :autosize="{ minRows: 2, maxRows: 4}"
                    v-model="temp.abstract" maxlength="100" show-word-limit :disabled="true">
          </el-input>
        </el-form-item>

        <el-form-item label="正文">
          <el-input type="textarea" placeholder="请输入正文"
                    :autosize="{ minRows: 9, maxRows: 12}"
                    v-model="temp.content" :disabled="true">
          </el-input>
        </el-form-item>
      </el-form>
    </el-dialog>

  </div>
</template>

<script>
  import releaseApi from '@/api/release'
  import { parseTime, resetTemp } from '@/utils'
  import debounce from 'lodash/debounce'
  import store from '@/store'

  export default {
    name: 'PoliticsManage',

    data() {
      return {
        tableQuery: {
          // 重要
          politicsType: null,
          // 其它
          title: null
        },
        tableLoading: false,
        tableData: [],
        // 当前用户角色
        currentRole: store.getters.role,
        dialogFormVisible: false,
        dialogDetialVisible: false,
        tablePage: {
          current: null,
          pages: null,
          size: null,
          total: null
        },
        temp: {
          idx: null, //tableData中的下标
          politicsId: null,
          title: null,
          abstract: null,
          content: null,
          enclosurePath: null,
          createTime: null,
          politicsType: null,
          readingSigns: null
        }
      }
    },

    created() {
      this.initData()
      this.fetchData()
    },

    watch: {
      //延时查询
      'tableQuery.title': debounce(function() {
        this.fetchData()
      }, 500)
    },

    methods: {
      // 解析时间
      parseTime,

      // 有操作权限
      hasFilePerm() {
        return this.currentRole.perms.some(p => p.val == 'm:politics:release')
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

      // 初始化数据
      initData() {
        this.tableQuery.politicsType = 'guide'
      },

      // 展示
      handleShow(idx, row) {
        this.temp = Object.assign({}, row) // copy obj
        this.temp.idx = idx

        this.dialogDetialVisible = true
      },

      //删除
      handleDelete(idx, row) {
        this.$confirm('您确定要永久删除吗？', '提示', confirm).then(() => {
          releaseApi.deletePolitics(row.politicsId).then(res => {
            this.tableData.splice(idx, 1)
            --this.tablePage.total
            this.dialogFormVisible = false
            this.$message.success('删除成功')
          })
        }).catch(() => {
          this.$message.info('已取消删除')
        })

      },

      // 创建
      handleCreate() {
        resetTemp(this.temp)
        this.dialogFormVisible = true
      },
      createData() {
        this.temp.politicsType = this.tableQuery.politicsType
        releaseApi.addPolitics(this.temp).then((res) => {
          this.temp.politicsId = res.data.politicsId//后台传回来新增记录的id
          this.temp.createTime = res.data.createTime//后台传回来新增记录的时间

          this.tableData.unshift(Object.assign({}, this.temp))
          ++this.tablePage.total
          this.dialogFormVisible = false
          this.$message.success('添加角色成功')
        })
      },

      // 获取表格数据
      fetchData(current) {
        if (current) {
          this.tablePage.current = current
        }

        this.tableLoading = true
        releaseApi.fetchDataToPage(this.tableQuery, this.tablePage).then(res => {
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
