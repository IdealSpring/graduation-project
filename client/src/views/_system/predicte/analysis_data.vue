<template>
  <div class="app-container">

    <el-row>
      <!--角色信息-->
      <span class="page-title">企业分析报告</span>
      <router-link to="/predicte_analysis/analysis_manage" style="float:right">
        <el-button type="text" icon="el-icon-back">返回上一页</el-button>
      </router-link>
    </el-row>

    <div style="margin-bottom: 10px;"></div>

    <el-row :gutter="20">
      <el-col :span="6">
        <el-card>
          <div slot="header">
            <span>企业基本信息</span>
          </div>
          <el-form label-position="left" class="demo-table-expand" style="height: 350px;">
            <el-form-item label="纳税人 ID">
              <span>9527</span>
            </el-form-item>
            <el-form-item label="行业代码">
              <span>1522</span>
            </el-form-item>
            <el-form-item label="所在行业">
              <span>IT技术与服务</span>
            </el-form-item>
            <el-form-item label="登记注册代码">
              <span>230</span>
            </el-form-item>
            <el-form-item label="注册类型">
              <span>IT技术</span>
            </el-form-item>
            <el-form-item label="营业时间">
              <span>2020-04-30 18:17:34</span>
            </el-form-item>
            <el-form-item label="创建时间">
              <span>2020-04-30 18:17:34</span>
            </el-form-item>
          </el-form>
        </el-card>
      </el-col>

      <el-col :span="18">
        <el-card>
          <div slot="header">
            <span>企业税额走势图</span>
          </div>
          <div :style="{height:'350px'}" ref="myEchart"></div>
        </el-card>
      </el-col>
    </el-row>

    <el-row :gutter="20">
      <el-col :span="12">
        <el-card>
          <div slot="header">
            <span>经营各项支出占比</span>
          </div>
          <div :style="{height:'350px'}" ref="myEchart2"></div>
        </el-card>
      </el-col>

      <el-col :span="12">
        <el-card>
          <div slot="header">
            <span>进销项发票税额占比</span>
          </div>
          <div :style="{height:'350px'}" ref="myEchart3"></div>
        </el-card>
      </el-col>

    </el-row>

  </div>
</template>

<script>

  import echarts from "echarts"

  export default {
    name: 'AssignPerm',
    data() {
      return {
        option1: {
          xAxis: {
            type: 'category',
            data: ['2015年', '2016年', '2017年', '2018年', '2019年']
          },
          yAxis: {
            type: 'value'
          },
          series: [{
            data: [10, 13, 16, 10, 21],
            type: 'line'
          }]
        },
        option2: {
          tooltip: {
            trigger: 'item',
            formatter: '{a} <br/>{b} : {c} ({d}%)'
          },
          legend: {
            orient: 'vertical',
            left: 'left',
            data: ['水费', '电费', '天然气', '公证费', '餐饮费']
          },
          series: [
            {
              name: '详情:',
              type: 'pie',
              radius: '55%',
              center: ['50%', '60%'],
              data: [
                {value: 335, name: '水费'},
                {value: 310, name: '电费'},
                {value: 234, name: '天然气'},
                {value: 135, name: '公证费'},
                {value: 1548, name: '餐饮费'}
              ],
              emphasis: {
                itemStyle: {
                  shadowBlur: 10,
                  shadowOffsetX: 0,
                  shadowColor: 'rgba(0, 0, 0, 0.5)'
                }
              }
            }
          ]
        },

        option3: {
          tooltip: {
            trigger: 'item',
            formatter: '{a} <br/>{b} : {c} ({d}%)'
          },
          legend: {
            orient: 'vertical',
            left: 'left',
            data: ['进项发票税额', '销项发票税额']
          },
          series: [
            {
              name: '详情',
              type: 'pie',
              radius: '55%',
              center: ['50%', '60%'],
              data: [
                {value: 600, name: '进项发票税额'},
                {value: 310, name: '销项发票税额'}
              ],
              emphasis: {
                itemStyle: {
                  shadowBlur: 10,
                  shadowOffsetX: 0,
                  shadowColor: 'rgba(0, 0, 0, 0.5)'
                }
              }
            }
          ]
        },
      }
    },

    computed: {},
    watch: {},
    created() {
      this.$nextTick(() => {
        this.fetchData()
      })

    },
    methods: {
      // 获取数据
      fetchData() {
        let lineChart = echarts.init(this.$refs.myEchart);
        lineChart.setOption(this.option1)

        echarts.init(this.$refs.myEchart2).setOption(this.option2)
        echarts.init(this.$refs.myEchart3).setOption(this.option3)
      }
    }
  }
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
  .el-row {
    margin-bottom: 20px;
    &:last-child {
      margin-bottom: 0;
    }
  }
  .el-col {
    border-radius: 4px;
  }
  .page-title{
    font-size: 24px;
    font-weight: bold;
    color: #303133;
  }
  .demo-table-expand {
    font-size: 0;
  }
  .demo-table-expand label {
    width: 90px;
    color: #99a9bf;
  }
  .demo-table-expand .el-form-item {
    margin-right: 0;
    margin-bottom: 10px;
  }
</style>
