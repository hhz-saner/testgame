<template>
  <div class="ready-round2 container" v-on:click="next">
    <div class="content">
      <p v-if="correctRate">你的正确率是：{{correctRate}}%</p>
      <p>相信你会在下一轮做得更好的!</p>
      <p>你可以休息一下，也可以</p>
      <p>点击屏幕以继续</p>
    </div>
  </div>
</template>
<script>
export default {
  name: "ready-round2",
  data() {
    return {
      correctRate: sessionStorage.getItem("correct-rate")
    };
  },
  methods: {
    next() {
      this.$router.push("/round2");
    }
  }
};
</script>


<style scoped>
.ready-round2 p {
  margin-bottom: 12vh;
}
</style>