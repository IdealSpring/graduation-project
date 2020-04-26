<template>
  <div class="app-container">
    <!--中国地图-->
    <div class="echarts">
      <div :style="{height:'540px',width:'1140px'}" ref="myEchart"></div>
    </div>

  </div>
</template>

<script>
  import echarts from "echarts";
  import 'echarts/map/js/china.js' // 引入中国地图数据

  export default {
    name: 'ShowManage',
    props: ["renderData"],
    data() {
      return {
        chart: null,
      };
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
      initEchartMap() {
        //console.log(this.renderData)
        //这里我们用固定的数据，真正使用时，用父组件传递来的数据替换series即可
        let myChart = echarts.init(this.$refs.myEchart);
        window.onresize = myChart.resize;
        myChart.setOption({
          backgroundColor: "#fff",
          title: {
            text: '每日货盘运行图',
            left: 'center',
            textStyle: {
              color: '#000000'
            }
          },
          tooltip : {
            trigger: 'item'
          },
          dataRange: {
            show: false,
            min: 0,
            max: 1000,
            text: ['High', 'Low'],
            realtime: true,
            calculable: true,
            color: ['orangered', 'yellow', 'lightskyblue']
          },
          visualMap : {
            show : true,
            min : 0,
            max : 255,
            calculable : true,
            inRange : {
              color : ['aqua', 'lime', 'yellow', 'orange', '#ff3333']
            },
            textStyle : {
              color : '#12223b'
            }
          },
          geo: { // 这个是重点配置区
            map: 'china', // 表示中国地图
            roam: true,
            layoutCenter : ['50%', '50%'],
            layoutSize : "130%",
            label: {
              normal: {
                show: true, // 是否显示对应地名
                textStyle: {
                  color: 'rgba(0,0,0,0.4)'
                }
              },
            },
            itemStyle: {
              normal: {
                areaColor: '#37376e',
                borderColor: '#fff'
              },
              emphasis: {
                areaColor: null,
                shadowOffsetX: 0,
                shadowOffsetY: 0,
                shadowBlur: 20,
                borderWidth: 0,
                shadowColor: 'rgba(0, 0, 0, 0.5)'
              }
            }
          },
          series: [
            {
              "type": "lines",
              "zlevel": 2,
              "effect": {
                "show": true,
                "period": 4,
                "trailLength": 0.02,
                "symbol": "arrow",
                "symbolSize": 5
              },
              "lineStyle": {
                "normal": {
                  "width": 1,
                  "opacity": 0.6,
                  "curveness": 0.2
                }
              },
              "tooltip": {
                "trigger": "item"
              },
              "data": [
                {
                  "fromName": "北京",
                  "toName": "吉林",
                  "value": 3,
                  "coords": [
                    [
                      116.46,
                      39.92
                    ],
                    [
                      "125.815",
                      "44.258"
                    ],
                  ]
                },
                {
                  "fromName": "北京",
                  "toName": "江西",
                  "value": 3,
                  "coords": [
                    [
                      116.46,
                      39.92
                    ],
                    [
                      "116.005",
                      "28.663"
                    ],
                  ]
                },
              ]
            },

            {
              "name": "总局位置",
              "type": "scatter",
              "coordinateSystem": "geo",
              "zlevel": 2,
              "tooltip": {
                "trigger": "item"
              },
              "label": {
                "normal": {
                  "show": true,
                  "position": "right",
                  "color": "#00ffff",
                  "formatter": "{b}",
                  "textStyle": {
                    "color": "#EE0000"
                  }
                },
                "emphasis": {
                  "show": true
                }
              },
              "symbol": "circle",
              "symbolSize": 20,
              "itemStyle": {
                "normal": {
                  "show": true,
                  "color": "#EE0000"
                }
              },
              "data": [
                {
                  "name": "北京",
                  "value": [
                    116.46,
                    39.92
                  ]
                }
              ]
            },

            {
              "name": "发行省份",
              "type": "effectScatter",
              "coordinateSystem": "geo",
              "zlevel": 2,
              "rippleEffect": {
                "period": 4,
                "brushType": "stroke",
                "scale": 4
              },
              "tooltip": {
                "trigger": "item"
              },
              "label": {
                "normal": {
                  "show": true,
                  "position": "left",
                  "offset": [
                    -5,
                    5
                  ],
                  "formatter": "{b}"
                },
                "emphasis": {
                  "show": true
                }
              },
              "hoverAnimation": true,
              "symbol": "circle",
              "symbolSize": 5,
              "itemStyle": {
                "normal": {
                  "show": false,
                  "color": "#f00"
                }
              },
              "data": [
                {
                  "name": "吉林",
                  "value": [
                    "125.815",
                    "44.258",
                    3
                  ]
                },
                {
                  "name": "江西",
                  "value": [
                    "116.005",
                    "28.663",
                    3
                  ]
                },
              ]
            }
          ]
        })
      }
    }
  }
</script>

<style rel="stylesheet/scss" lang="scss" scoped></style>
