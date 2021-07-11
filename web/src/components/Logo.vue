<template>
<div class="flex flex-row align-center">
    <em class="t">T</em>iny <em class="u">U</em>RL
</div>
</template>

<script>
import {defineComponent} from 'vue'

export default defineComponent({})
</script>

<style scoped>
em {
  font-style: normal;
}

.t {
  color: #2196f3;
}

.u {
  color: blueviolet;
}

.default {
  color: #333333;
  font-weight: bold;
}

.white {
  color: #eee;
}
</style>