<template>
  <div class="app-container">
    <div v-if="hasAdminRole()">
      <el-row  :gutter="20">
        <el-col :span="10">
          <el-card>
            <div :style="{height:'480px'}"  style="background: #F5F5F5; border-radius: 4px;" ref="myEchart"></div>
          </el-card>
        </el-col>

        <el-col  :span="14">
          <el-card>
            <div :style="{height:'480px'}"  style="background: #F5F5F5; border-radius: 4px;" ref="myEchart2"></div>
          </el-card>
        </el-col>
      </el-row>
    </div>
    <div v-else>
      <el-row>
        <el-col :span="24">
          <el-card>
            <div :style="{height:'480px'}"  style="background: #F5F5F5; border-radius: 4px;" ref="myEchart3"></div>
          </el-card>
        </el-col>
      </el-row>

    </div>



  </div>
</template>

<script>
  import echarts from "echarts"
  import store from '@/store'

  export default {
    name: 'StatisticsManage',
    data() {
      return {
        chart: null,
        // 当前用户角色
        currentRole: store.getters.role,

        pieOption: {
          title: {
            text: '本年度税收占比',
            subtext: '—起始时间为去年今天',
            left: 'center'
          },
          tooltip: {
            trigger: 'item',
            formatter: '{a} <br/>{b} : {c} ({d}%)'
          },
          legend: {
            orient: 'vertical',
            left: 'left',
            data: null
          },
          series: [
            {
              name: '详情:',
              type: 'pie',
              radius: '55%',
              center: ['50%', '60%'],
              data: null,
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

        lineOption: {
          title: {
            text: '历年税收金额走势图'
          },
          tooltip: {
            trigger: 'axis'
          },
          legend: {
            data: null
          },
          grid: {
            left: '3%',
            right: '4%',
            bottom: '3%',
            containLabel: true
          },
          toolbox: {
            feature: {
              saveAsImage: {}
            }
          },
          xAxis: {
            type: 'category',
            boundaryGap: false,
            data: null
          },
          yAxis: {
            type: 'value'
          },
          series: null
        },

        provinceLineOption: {
          title: {
            text: '吉林省各地区税收走势图'
          },
          tooltip: {
            trigger: 'axis'
          },
          legend: {
            data: null
          },
          grid: {
            left: '3%',
            right: '4%',
            bottom: '3%',
            containLabel: true
          },
          toolbox: {
            feature: {
              saveAsImage: {}
            }
          },
          xAxis: {
            type: 'category',
            boundaryGap: false,
            data: null
          },
          yAxis: {
            type: 'value'
          },
          series: null
        },

      };
    },
    created() {
      this.fetchData()
    },
    mounted() {
      this.initEchartMap();
    },
    beforeDestroy() {
      if (!this.chart) {
        return;
      }
      this.chart.dispose();
      this.chart = null;
    },
    methods: {
      // 判断是总局管理员还是省级管理员
      hasAdminRole() {
        if (this.currentRole.roleId == 1) {
          return true
        }

        return false
      },
      // 获取数据
      fetchData() {
        // 总局管理员数据
        if (this.hasAdminRole()) {
          // 饼状图
          this.pieOption.legend.data = ['北京', '吉林', '江西']
          this.pieOption.series[0].data = [
            {value: 335, name: '北京'},
            {value: 310, name: '吉林'},
            {value: 234, name: '江西'}
          ]

          // 折线图
          this.lineOption.legend.data = ['北京', '吉林', '江西']
          this.lineOption.xAxis.data = ['2015年', '2016年', '2017年', '2018年', '2019年']
          this.lineOption.series = [
            {
              name: '北京',
              type: 'line',
              stack: '总量',
              data: [120, 132, 101, 134, 90, 230]
            },
            {
              name: '吉林',
              type: 'line',
              stack: '总量',
              data: [220, 182, 191, 234, 290, 330]
            },
            {
              name: '江西',
              type: 'line',
              stack: '总量',
              data: [150, 232, 201, 154, 190, 330]
            }
          ]

        } else { // 省级管理员数据

          // 折线图
          this.provinceLineOption.legend.data = ['长春', '吉林', '白山', '白城']
          this.provinceLineOption.xAxis.data = ['2015年', '2016年', '2017年', '2018年', '2019年']
          this.provinceLineOption.series = [
            {
              name: '长春',
              type: 'line',
              stack: '总量',
              data: [120, 132, 101, 134, 90, 230]
            },
            {
              name: '吉林',
              type: 'line',
              stack: '总量',
              data: [220, 182, 191, 234, 290, 330]
            },
            {
              name: '白山',
              type: 'line',
              stack: '总量',
              data: [112, 145, 178, 190, 220, 255]
            },
            {
              name: '白城',
              type: 'line',
              stack: '总量',
              data: [224, 232, 332, 112, 143, 123]
            },
          ]
        }



      },
      // 初始化图形
      initEchartMap() {
        if (this.hasAdminRole()) {
          let pieChart = echarts.init(this.$refs.myEchart);
          let lineChart = echarts.init(this.$refs.myEchart2);

          pieChart.setOption(this.pieOption)
          lineChart.setOption(this.lineOption)
        } else {
          let lineChart = echarts.init(this.$refs.myEchart3);
          lineChart.setOption(this.provinceLineOption)
        }

      }
    }
  }
</script>

<style rel="stylesheet/scss" lang="scss" scoped>
  .el-col {
    border-radius: 4px;
  }
  .bg-purple-dark {
    background: #99a9bf;
  }
  .bg-purple {
    background: #d3dce6;
  }
  .bg-purple-light {
    background: #e5e9f2;
  }
  .grid-content {
    border-radius: 4px;
    min-height: 36px;
  }
</style>
