<template>
  <a-layout class="layout" theme="light">
    <a-layout-header>
      <div class="header" >
        <router-link to="/" style="display: block">
          <logo style="color: #555;text-align: center"/>
        </router-link>
      </div>
    </a-layout-header>
    <a-layout-content class="content">
      <router-view/>
    </a-layout-content>
  </a-layout>

</template>

<script>
import {defineComponent, onMounted, reactive, toRefs} from 'vue'
import Logo from "../components/Logo.vue";
import {useRoute} from "vue-router";
import {LinkOutlined,} from '@ant-design/icons-vue'

export default defineComponent({
  components: {
    Logo, LinkOutlined,
  },
  setup() {
    const state = reactive({
      selectedKeys: ['Url']
    })
    const route = useRoute()

    onMounted(() => {
      state.selectedKeys = [route.matched[1].name || "Url"]
    })

    return {
      ...toRefs(state),
    }
  }
})
</script>

<style lang="less">
.ant-layout {
  background-color: #edf2f7;
}
.header{
  width: 1280px;margin: 0 auto;display: flex;
}
.menu {
  .menu-item {
    height: 36px;
    line-height: 36px;
    margin-bottom: 15px;

    a {
      font-weight: 500;
      text-decoration: none;
      color: #999;
      display: block;
      padding: 0 10px;

      &.router-link-active {
        background-color: #39455012;
        color: #555555;
      }
    }
  }
}


.layout-header {
  display: flex;
}

.site-layout-content {
  min-height: 280px;
  padding: 24px;
  background: #fff;
}


.ant-layout-header {
  background-color: transparent;
}

.ant-layout-sider-light {
  background-color: transparent;
}

.content {
  min-height: 700px;
  max-width: 1280px;
  width: 1280px;
  margin: 20px auto;
}
</style>