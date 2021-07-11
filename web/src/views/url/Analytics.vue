<template>
    <div class="analytics">
        <a-page-header
            style="margin-bottom: 20px;"
            :title="slug"
            @back="() => $router.back()"
        />
        <a-card title="趋势" hoverable style="border-radius: 5px;height: 100%;margin-bottom: 20px;"
                :bodyStyle="{padding:'0'}">
            <div id="main" style="width: 100%;height: 300px"></div>
        </a-card>
        <a-row type="flex" :gutter="16" style="margin-bottom: 20px;" align="stretch">
            <a-col :span="24">
                <a-card hoverable
                        :tab-list="tabList"
                        :active-tab-key="key"
                        @tabChange="key => onTabChange(key, 'key')"
                        style="border-radius: 5px;height: 100%;" :bodyStyle="{padding:'0'}">
                    <div id="item" style="width: 100%;height: 300px"></div>
                    <a-table :data-source="data" row-key="item" :pagination="false">
                        <a-table-column key="item" title="ITEM" data-index="item"/>
                        <a-table-column key="count" title="COUNT" data-index="count" align="right"/>
                    </a-table>
                </a-card>
            </a-col>
            <a-col :span="0" hoverable>
                <a-card title="地域分布" hoverable style="border-radius: 5px;height: 100%;" :bodyStyle="{padding:'0'}">
                    <div id="world" style="width: 100%;height:300px"></div>
                </a-card>
            </a-col>
        </a-row>
    </div>
</template>

<script>
import {defineComponent, onMounted, reactive, ref, toRefs} from 'vue'
import {apiAnalytics} from "../../model/model";
import * as echarts from 'echarts/core';
import {GridComponent, LegendComponent, TitleComponent, TooltipComponent} from 'echarts/components';
import {BarChart, LineChart} from 'echarts/charts';
import {CanvasRenderer, SVGRenderer} from 'echarts/renderers';

echarts.use(
    [GridComponent, LineChart, CanvasRenderer, SVGRenderer, BarChart, LegendComponent, TooltipComponent, TitleComponent]
);

export default defineComponent({
    props: {
        slug: {
            type: String,
            required: true
        },
    },
    setup(props) {
        const key = ref("source")
        const tabList = reactive([
            {key: 'source', tab: 'Sources'},
            {key: 'region', tab: 'Regions'},
            {key: 'os', tab: 'OS'},
            {key: 'device', tab: 'Device'},
            {key: 'browser', tab: 'Browser'},
        ])
        const state = reactive({
            data: [],
            days: []
        })


        const _initItemChart = () => {
            const chartDom = document.getElementById('item');
            const myChart = echarts.init(chartDom);
            let option;

            option = {

                tooltip: {
                    trigger: 'axis',
                    axisPointer: {
                        type: 'shadow'
                    }
                },

                grid: {
                    left: '3%',
                    right: '4%',
                    bottom: '3%',
                    containLabel: true
                },
                xAxis: {
                    type: 'value',
                    boundaryGap: [0, 0.01]
                },
                yAxis: {
                    type: 'category',
                    data: state.data||[].map(v => v.item)
                },
                series: [
                    {
                        name: '2011年',
                        type: 'bar',
                        data: state.data||[].map(v => v.count)
                    },

                ]
            };

            option && myChart.setOption(option);
        }

        const _initChart = (days) => {
            const chartDom = document.getElementById('main');
            const myChart = echarts.init(chartDom);
            let option;

            option = {
                xAxis: {
                    type: 'category',
                    boundaryGap: false,
                    data: days.map(v => v.item)
                },
                yAxis: {
                    type: 'value'
                },
                series: [{
                    data: days.map(v => v.count),
                    type: 'line',
                    areaStyle: {}
                }]
            };

            option && myChart.setOption(option);
        }

        const fetchData = () => {
            apiAnalytics(props.slug, {type: key.value}).then(({data}) => {
                state.data = data;
                _initItemChart()
            })
            apiAnalytics(props.slug, {type: "day"}).then(({data}) => {
                state.days = data;
                _initChart(state.days||[])
            })
        }
        onMounted(() => {
            fetchData()
        })

        const onTabChange = (k) => {
            key.value = k
            fetchData()
        }

        return {
            ...toRefs(state),
            onTabChange,
            key,
            tabList
        }
    }
})
</script>

<style scoped>

</style>