<template>
  <router-view />
</template>

<script setup>

</script>

<style lang="less">
html,body {
    margin: 0;
    padding: 0;

}
html,body,#app {
    height: 100%;
}
a {
    text-decoration: none;
  &:active,&:focus,&:link,&:visited {
    color: #555555;
  }
}
#app {
  position: relative;
}
</style>