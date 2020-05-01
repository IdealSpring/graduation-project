<template>
  <div class="app-container">
    <!--查询-->
    <el-row>
      <el-button type="primary" @click="handlePredicte()">配置预测参数</el-button>
      <span style="margin-right: 15px;"/>
      <el-input v-model="tableQuery.taxpayerId" placeholder="企业检索" style="width: 20%;"/>
      <span style="margin-right: 10px;"/>
      <el-tooltip content="搜索" placement="top">
        <el-button icon="el-icon-search" circle @click="fetchData(1)"></el-button>
      </el-tooltip>
    </el-row>

    <!--占位框-->
    <div style="margin-bottom: 15px;"></div>

    <el-row :gutter="20">
      <el-col :span="19">
        <el-card>
          <div slot="header">
            <span>异常企业名单</span>
            <span style="font-size: smaller; color: #8c939d">-根据配置更新周期自动更新</span>
            <el-button style="float: right; margin-right: 10px; padding: 7px 5px" size="mini" type="primary"
                       icon="el-icon-download" @click="downloadFile()">企业名单
            </el-button>
            <el-button style="float: right; margin-right: 10px; padding: 7px 5px" size="mini" type="success"
                       icon="el-icon-download"  @click="downloadFile()">算法模型
            </el-button>
          </div>
          <div>
            <el-table
              border highlight-current-row
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
              <el-table-column prop="createTime" label="录入系统时间" width="200" align="center">
                <template slot-scope="scope">
                  <span v-text="parseTime(scope.row.createTime)"></span>
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
              :page-sizes="[7, 20, 30, 40, 50]"
              :page-size="tablePage.size"
              layout="total, sizes, prev, pager, next, jumper"
              :total="tablePage.total">
            </el-pagination>

          </div>
        </el-card>
      </el-col>

      <el-col :span="5">
        <el-card>
          <div slot="header">
            <span>算法模型指标</span>
          </div>
          <div>
            <el-table
              v-loading.body="tableLoading"
              element-loading-text="Loading"
              :data="tableDataForModelInd">
              <el-table-column prop="name" label="Name"/>
              <el-table-column prop="value" label="Value"/>
            </el-table>

          </div>
        </el-card>
      </el-col>
    </el-row>

    <!--弹出窗口：新增/编辑用户-->
    <el-dialog title="预测配置参数" :visible.sync="dialogFormVisible" width="40%">
      <el-form :model="temp" label-position="left" label-width="120px">

        <el-form-item label="执行频率" prop="">
          <el-select v-model="execRate" clearable placeholder="请选择">
            <el-option
              v-for="item in execRateList"
              :key="item.value"
              :label="item.label"
              :value="item.value">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="算法模型" prop="industryCode">
          <el-upload
            action="string"
            :auto-upload="false">
            <el-button slot="trigger" icon="el-icon-upload2" size="mini"  type="primary">本地模型</el-button>
          </el-upload>
        </el-form-item>
        <el-form-item label="模型属性" prop="industryCode">
          <el-upload
            action="string"
            :auto-upload="false">
            <el-button slot="trigger" icon="el-icon-upload2" size="mini"  type="primary">算法使用属性</el-button>
          </el-upload>
        </el-form-item>

      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取消</el-button>
        <el-button @click="dialogFormVisible = false" type="primary">确定</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script>
  import taxpayerApi from '@/api/taxpayer'
  import predicteApi from '@/api/predicte'
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
        // 执行周期
        execRate: null,
        execRateList: [
          {
            label: '每天',
            value: 1
          },
          {
            label: '每周',
            value: 2
          },
          {
            label: '每月',
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
          size: 7,
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

      downloadFile() {
        predicteApi.downloadFile().then(data => {
          if (!data) {
            this.$message.error('下载内容为空')
            return
          }
          let url = window.URL.createObjectURL(new Blob([data]))
          let link = document.createElement('a')
          link.style.display = 'none'
          link.href = url
          link.setAttribute('download', 'data' + '.zip')

          document.body.appendChild(link)
          link.click()
          //释放URL对象所占资源
          window.URL.revokeObjectURL(url)
          //用完即删
          document.body.removeChild(link)
        }).catch(err => {
          console.log('err: ', err)
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
