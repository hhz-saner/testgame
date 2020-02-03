<template>
  <div class="container" v-if="showTask">
    <work :task="task" v-on:submit="submit"></work>
  </div>
  <div class="container" v-else v-on:click="next">
    <div class="content" v-if="showDescription">
      <p class="title">第三种要求：</p>
      <p>从小到大：从小到大对呈现的数字进行排列；</p>
      <p>如屏幕呈现</p>
      <p>7 9 2 0</p>
      <p>正确答案为</p>
      <p class="true-content">0 2 7 9</p>
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
  name: "guide3",

  data() {
    return {
      showTask: false,
      showDescription: true,
      correct: false,
    };
  },
  components: {
    Work
  },
  computed: {},
  methods: {
    next() {
      if (this.showDescription) {
        this.task = this.Utils.getNumberTask(2);
        this.showTask = true;
        this.showDescription = false;
      } else {
        if (this.correct) {
          this.$router.push("/guide4");
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