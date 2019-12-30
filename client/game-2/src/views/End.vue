<template>
  <div class="end container" v-on:click="next">
    <div class="center">
      <div>你的正确率是：{{correctRate}}%</div>
      <p class="bottom-tip">点击屏幕以继续</p>
    </div>
  </div>
</template>
<script>
export default {
  name: "end",
  data() {
    return {
      index: 0,
      correctRate:sessionStorage.getItem("correct-rate")
    };
  },
  methods: {
    next() {
        this.$router.push("/finished");
    }
  }
};
</script>


<style scoped>
.end .center {
  align-content: center;
  font-size: 33px;
  text-align: center;
  margin-top: 30vh;
}
</style>