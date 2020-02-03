<template>
  <div class="ready-round1 container" v-on:click="next">
    <div class="center" v-if="index==0" >
      <div>你的正确率是：{{correctRate}}% 不错！</div>
      <p class="bottom-tip">点击屏幕以继续</p>
    </div>
    <div class="content" v-if="index==1">
      <p>看来你已经很快掌握了规则,</p>
      <p>接下来就是正式游戏啦。</p>
      <p>一共有2个回合，每个回合42道题目。</p>
      <p>请记住：</p>
      <p>在每一轮开始时集中注意力在十字上</p>
      <p>（<span class="plus">+</span>）</p>
      <p>并记清楚每轮的任务要求</p>
      <p>开始第一回合！</p>
      <p class="bottom-tip">点击屏幕以继续</p>
    </div>
  </div>
</template>
<script>
export default {
  name: "ready-round1",
  data() {
    return {
      index:0,
      correctRate: sessionStorage.getItem("correct-rate")
    };
  },
  created() {
    this.index = this.correctRate?0:1;
  },
  methods: {
    next() {
      this.index++;
      if (this.index >= 2) {
        this.$router.push("/round1");
      }
    }
  }
};
</script>


<style scoped>
.ready-round1 .center {
  align-content: center;
  font-size: 33px;
  text-align: center;
  margin-top: 30vh;
}
.ready-round1 .plus {
  font-size: 60px;
}
</style>