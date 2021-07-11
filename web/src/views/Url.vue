<template>
    <div class="url">
        <h1>URLs</h1>
        <a-card>
            <a-table :data-source="data" row-key="slug">
                <a-table-column key="slug" title="Slug" data-index="slug">
                    <template #default="{record}">
                        <a target="_blank" :href="`/r/${record.slug}`">{{ record.slug }}</a>
                    </template>
                </a-table-column>
                <a-table-column key="url" title="URL" data-index="url">
                    <template #default="{record}">
                        <a target="_blank" :href="record.url">{{ record.url }}</a>
                    </template>
                </a-table-column>
                <a-table-column key="name" title="说明" data-index="name"/>
                <a-table-column key="type" title="类型" data-index="type">
                    <template #default="{record}">
                        <a-tag color="default" v-if="record.type===0">默认</a-tag>
                        <a-tag color="orange" v-if="record.type===1">自定义</a-tag>
                    </template>
                </a-table-column>
                <a-table-column key="created_at" title="添加时间" align="right" data-index="created_at">
                    <template #default="{record}">
                        {{ dayjs.unix(record.created_at).fromNow() }}
                    </template>
                </a-table-column>
                <a-table-column align="right" title="操作">
                    <template #default="{record}">
                        <router-link style="color:#1890ff" :to="{name:'UrlAnalytics',params:{slug:record.slug}}">
                            <bar-chart-outlined/>
                            数据分析
                        </router-link>
                    </template>
                </a-table-column>
            </a-table>
        </a-card>
    </div>
</template>

<script>
import {defineComponent, onMounted, reactive, toRefs} from 'vue'
import {apiListUrl} from "../model/model";
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'
import {BarChartOutlined} from '@ant-design/icons-vue'

dayjs.extend(relativeTime)

export default defineComponent({
    components: {
        BarChartOutlined,
    },
    setup() {
        const state = reactive({
            data: []
        })

        onMounted(() => {
            apiListUrl().then(({data}) => {
                state.data = data
            })
        })

        return {
            ...toRefs(state),
            dayjs,
        }
    }
})
</script>

<style lang="less" scoped>

h1 {
    margin: 0 0 10px;
    font-weight: 700;
    font-size: 30px;
}
</style>