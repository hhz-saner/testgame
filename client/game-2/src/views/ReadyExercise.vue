<template>
  <div class="explain container" v-on:click="next">
    <div class="content">
      <div class="title">让我们来做一下练习吧</div>
      <p class="bottom-tip">点击屏幕以继续</p>
    </div>
  </div>
</template>

<script>
export default {
  name: "ready-explain",
  data() {
    return {};
  },
  methods: {
    next() {
      this.$router.push("/exercise");
    }
  }
};
</script>

<style scoped>
</style>