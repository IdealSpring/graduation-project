<template>
  <div class="app-container">
    <!--查询-->
    <el-row>
      <el-select v-model="standard" clearable placeholder="请选择">
        <el-option
          v-for="item in standardList"
          :key="item.value"
          :label="item.label"
          :value="item.value">
        </el-option>
      </el-select>

      <el-input v-model="tableQuery.taxpayerId" placeholder="企业检索" style="width: 20%;"/>
      <span style="margin-right: 10px;"/>
      <el-tooltip content="搜索" placement="top">
        <el-button icon="el-icon-search" circle @click="fetchData(1)"></el-button>
      </el-tooltip>
    </el-row>

    <!--占位框-->
    <div style="margin-bottom: 15px;"></div>

    <el-row>
      <el-col :span="24">
        <el-card>
          <div slot="header">
            <span>企业纳税额排行榜</span>
            <span style="font-size: smaller; color: #8c939d">-根据配置更新周期自动更新</span>
          </div>
          <div>
            <el-table
              highlight-current-row
              v-loading.body="tableLoading"
              element-loading-text="Loading"
              :data="tableData">
              <el-table-column type="index" :index="indexMethod" align="center"/>
              <el-table-column prop="taxpayerId" label="纳税人ID" width="100" align="center"/>
              <el-table-column prop="industryCode" label="行业代码" align="center"/>
              <el-table-column prop="registrationType" label="登记类型代码" align="center"/>

              <el-table-column prop="openTime" label="开营时间" width="200" align="center">
                <template slot-scope="scope">
                  <span v-text="parseTime(scope.row.createTime)"></span>
                </template>
              </el-table-column>
              <el-table-column prop="registrationType" label="年度累计缴税金额" align="center"/>
              <el-table-column prop="createTime" label="录入系统时间" width="200" align="center">
                <template slot-scope="scope">
                  <span v-text="parseTime(scope.row.createTime)"></span>
                </template>
              </el-table-column>
              <el-table-column label="操作" align="center" width="120">
                <template slot-scope="scope">
                  <el-button
                    @click.native.prevent="handleAnalysisDetails(scope.$index,scope.row)"
                    type="text"
                    size="small">
                    分析报表
                  </el-button>
                </template>
              </el-table-column>

            </el-table>

            <!--占位框-->
            <div style="margin-bottom: 10px;"></div>

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
        </el-card>
      </el-col>
    </el-row>

  </div>
</template>

<script>
  import taxpayerApi from '@/api/taxpayer'
  import { parseTime, resetTemp } from '@/utils'
  import debounce from 'lodash/debounce'
  import store from '@/store'

  export default {
    name: 'PoliticsManage',

    data() {
      return {
        tableQuery: {
          taxpayerId: null
        },
        // 全部，正常，异常
        standard: 1,
        standardList: [
          {
            label: '全部',
            value: 1
          },
          {
            label: '正常企业',
            value: 2
          },
          {
            label: '异常企业',
            value: 3
          }
        ],
        tableLoading: false,
        tableData: [],
        tableDataForModelInd: [{
          name: 'F1-score',
          value: '94.55%'
        }, {
          name: 'Precision',
          value: '94.81%'
        }, {
          name: 'Recall',
          value: '94.57%'
        }, {
          name: 'Accuracy',
          value: '94.56%'
        }, {
          name: 'Kappa',
          value: '88.18%'
        }, {
          name: 'Reliability',
          value: '92.59%'
        }, {
          name: 'RSD',
          value: '5.59%'
        }
        ],
        // 当前用户角色
        currentRole: store.getters.role,
        dialogFormVisible: false,
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
        }
      }
    },

    created() {
      this.fetchData()
    },

    watch: {
      //延时查询
      'tableQuery.taxpayerId': debounce(function() {
        this.fetchData()
      }, 500)
    },

    methods: {
      indexMethod(index) {
        return index + 1
      },

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

      // 预测
      handlePredicte() {
        this.dialogFormVisible = true
      },

      handleAnalysisDetails(idx, row) {
        this.$router.push({ path: '/predicte_analysis/analysis_data' })
      },

      // 上传文件
      uploadFile(item) {
        let formData = new FormData()
        formData.append('file', item.file)

        taxpayerApi.uploadFile(formData).then(res => {
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
          taxpayerApi.deletePolitics(row.taxpayerId).then(res => {
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
        taxpayerApi.fetchDataToPage(this.tableQuery, this.tablePage).then(res => {
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
