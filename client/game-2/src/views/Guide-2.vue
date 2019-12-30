<template>
  <div class="container" v-if="showTask">
    <work :task="task" v-on:submit="submit"></work>
  </div>
  <div class="container" v-else v-on:click="next">
    <div class="content" v-if="showDescription">
      <p class="title">第二种要求：</p>
      <p>倒序：按照数字/字母出现的顺序从后往前进行回忆；</p>
      <p>如屏幕呈现</p>
      <p>7 9 2 4</p>
      <p>正确答案为</p>
      <p>4 2 9 7</p>
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
        </div>
      </div>
    </div>
    <p class="bottom-tip">点击屏幕以继续</p>
  </div>
</template>

<script>
import Work from "@/components/Work.vue";

export default {
  name: "guide2",
  data() {
    return {
      showTask: false,
      showDescription: true,
      correct: false,
      taskIndex: "number",
    };
  },
  components: {
    Work
  },
  methods: {
    next() {
      if (this.showDescription) {
        this.task = this.Utils.getNumberTask(1);
        this.taskIndex = "number";
        this.showTask = true;
        this.showDescription = false;
      } else {
        if (this.correct) {
          if (this.taskIndex == "number") {
            this.showTask = false;
            this.task = this.Utils.getLetterTask(1);
            this.$nextTick(() => {
              this.showTask = true;
            });
            this.taskIndex = "letter";
          } else {
            this.$router.push("/guide3");
          }
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