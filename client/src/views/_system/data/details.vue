<template>
  <div class="app-container">
    <!--查询-->
    <el-row>
      <el-input v-model="tableQuery.detailsId" placeholder="请输入货物明细ID" style="width: 20%;"/>
      <span style="margin-right: 15px;"/>
      <el-tooltip content="搜索" placement="top">
        <el-button icon="el-icon-search" circle @click="fetchData(1)"></el-button>
      </el-tooltip>
    </el-row>

    <!--占位框-->
    <div style="margin-bottom: 30px;"></div>

    <!--按钮-->
    <el-upload
      action="string"
      :http-request="uploadFile"
      :show-file-list="false">
      <el-button slot="trigger" icon="el-icon-upload2"  size="small" type="primary">文件导入</el-button>
    </el-upload>

    <!--占位框-->
    <div style="margin-bottom: 30px;"></div>

    <!--表格-->
    <el-table style="width: 100%;"
              border highlight-current-row
              v-loading.body="tableLoading"
              element-loading-text="Loading"
              :data="tableData">
      <el-table-column prop="id" label="逻辑ID" width="80" align="center"/>
      <el-table-column prop="invoiceId" label="发票ID" width="80" align="center"/>
      <el-table-column prop="goodsName" label="货物名称" align="center"/>
      <el-table-column prop="specifications" label="规格型号" align="center"/>
      <el-table-column prop="unit" label="单位" align="center"/>
      <el-table-column prop="number" label="数量" align="center"/>
      <el-table-column prop="price" label="单价" align="center"/>
      <el-table-column prop="amountMoney" label="金额" align="center"/>
      <el-table-column prop="taxAmount" label="税额" align="center"/>
      <el-table-column prop="productCode" label="商品编码" align="center"/>

      <el-table-column prop="createTime" label="录入系统时间" width="165" align="center">
        <template slot-scope="scope">
          <span v-text="parseTime(scope.row.createTime)"></span>
        </template>
      </el-table-column>

      <el-table-column label="操作" width="70" align="center">
        <template slot-scope="scope">
          <el-tooltip content="删除" placement="top">
            <el-button @click="handleDelete(scope.$index,scope.row)" size="medium" type="danger" icon="el-icon-delete"
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
  import detailsApi from '@/api/details'
  import { parseTime, resetTemp } from '@/utils'
  import debounce from 'lodash/debounce'
  import store from '@/store'

  export default {
    name: 'PoliticsManage',

    data() {
      return {
        tableQuery: {
          detailsId: null
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
          taxpayerId: null,
          industryCode: null,
          registrationType: null,
          openTime: null,
          createTime: null,
          updateTime: null
        },
      }
    },

    created() {
      this.fetchData()
    },

    watch: {
      //延时查询
      'tableQuery.invoiceId': debounce(function() {
        this.fetchData()
      }, 500)
    },

    methods: {
      // 解析时间
      parseTime,

      // 分页
      handleSizeChange(val) {
        this.tablePage.size = val
        this.fetchData()
      },
      handleCurrentChange(val) {
        this.tablePage.current = val
        this.fetchData()
      },

      // 上传文件
      uploadFile(item) {
        let formData = new FormData()
        formData.append('file', item.file)

        detailsApi.uploadFile(formData).then(res => {
          this.$notify({
            title: '成功',
            message: '文件上传成功',
            type: 'success'
          })

          this.fetchData(1)
        }).catch(err => {
          console.log(err)
        })

      },

      //删除
      handleDelete(idx, row) {
        this.$confirm('您确定要永久删除吗？', '提示', confirm).then(() => {
          detailsApi.deleteInvoice(row.id).then(res => {
            this.tableData.splice(idx, 1)
            --this.tablePage.total
            this.dialogFormVisible = false
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
        detailsApi.fetchDataToPage(this.tableQuery, this.tablePage).then(res => {
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
