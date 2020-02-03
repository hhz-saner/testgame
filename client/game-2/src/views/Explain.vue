<template>
  <div class="explain container" v-on:click="next">
    <div class="content">
      <div v-if="index==0">
        <p>这个游戏非常简单。</p>
        <p>在每一轮开始之前，首先你需要集中注意力在这个十字上</p>
        <div class="cross">➕</div>
        <p>以及十字下方的文字上（十字下方文字为每一轮任务要求）</p>
      </div>
      <div v-else class="text-left">
        <p>共有5种任务要求。</p>
        <p>当十字消失后，你需要记住接下来连续出现的一连串数字或字母（数量为2-7个不等）。</p>
        <p>之后按要求进行回忆。</p>
        <p>一共有5种不同要求, 接下来我们会一一讲解。</p>
      </div>
    </div>
    <p class="bottom-tip">点击屏幕以继续</p>
  </div>
</template>

<script>
export default {
  name: "explain",
  data() {
    return {
      index: 0,
      loading: false
    };
  },
  created() {},
  components: {},
  methods: {
    next() {
      this.index++;
      if (this.index >= 2) {
        this.$router.push("/guide1");
      }
    }
  }
};
</script>

<style scoped>

</style>