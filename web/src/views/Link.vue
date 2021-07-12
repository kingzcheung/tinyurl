<template>
    <div class="url">
        <h1>Links</h1>

    </div>
</template>

<script>
import {defineComponent, onMounted, reactive, toRefs} from 'vue'
import {apiListUrl} from "../model/model";
import dayjs from 'dayjs'
import relativeTime from 'dayjs/plugin/relativeTime'

dayjs.extend(relativeTime)

export default defineComponent({
    components: {
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