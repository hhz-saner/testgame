<template>
  <div class="explain container" v-on:click="next">
    <div class="content">
      <div class="title">
        <p>让我们来做一下练习吧！</p>
        <p>接下来的14道题目将帮助你熟悉规则。</p>
      </div>
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