<template>
  <div class="container" v-if="showTask">
    <work :task="task" v-on:submit="submit"></work>
  </div>
  <div class="container" v-else v-on:click="next">
    <div class="content" v-if="showDescription">
      <p class="title">第五种要求：</p>
      <p>加和：对所有出现的数字进行加和。</p>
      <p>如屏幕呈现</p>
      <p>8 4 6 0 2</p>
      <p>正确答案为</p>
      <p class="true-content">20</p>
      <p class="bottom-tip red">接下来，让我们做一组练习</p>
    </div>
    <div class="content" v-else>
      <div v-if="correct" class="content">
        <p class="title">对了！</p>
        <p class="bottom-tip">点击屏幕以继续</p>
      </div>
      <div v-else>
        <div class="content">
          <p class="title">输入错误。</p>
          <p>请仔细阅读任务要求。集中注意记住出现的字母/数字</p>
          <p class="bottom-tip">点击屏幕以继续</p>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import Work from "@/components/Work.vue";

export default {
  name: "guide5",
    data() {
    return {
      showTask: false,
      showDescription: true,
      correct: false
    };
  },
  components: {
    Work
  },
  computed: {},
  methods: {
    next() {
      if (this.showDescription) {
        this.task = this.Utils.getNumberTask(4);
        this.showTask = true;
        this.showDescription = false;
      } else {
        if (this.correct) {
          this.$router.push("/ready-exercise");
        } else {
          this.showTask = false;
          this.showDescription = true;
        }
      }
    },
    submit(value) {
      this.showTask = false;
      this.correct = value.correct;
    }
  }
};
</script>
