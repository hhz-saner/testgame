<template>
  <div class="exercise container">
    <div v-if="startWork">
      <work :task="task" v-if="showWork" v-on:submit="submit"></work>
      <div class="black-style" v-else></div>
    </div>
    <div v-else class="content">
      <p>正确率{{correctRate}}</p>
      <p>{{right}}/{{reportList.length}}</p>
      <p class="cross">{{number}}/{{count}}</p>
    </div>
  </div>
</template>

<script>
import Work from "@/components/Work.vue";

export default {
  name: "exercise",
  data() {
    return {
      tasks: [],
      reportList: [],
      startWork: false,
      showWork: false,
      right: 0,
      count: 0,
      task: null
    };
  },
  computed: {
    correctRate: function() {
      return Math.round((this.right / this.reportList.length) * 100);
    },
    number: function() {
      return this.reportList.length + 1;
    }
  },
  created() {
    for (let i = 1; i <= 2; i++) {
      this.tasks.push(this.Utils.getNumberTask(0));
      this.tasks.push(this.Utils.getNumberTask(1));
      this.tasks.push(this.Utils.getNumberTask(2));
      this.tasks.push(this.Utils.getNumberTask(4));

      this.tasks.push(this.Utils.getLetterTask(0));
      this.tasks.push(this.Utils.getLetterTask(1));
      this.tasks.push(this.Utils.getLetterTask(3));
    }
    this.count = this.tasks.length;
    this.right = 0;
    this.getTask();
    setTimeout(this.start, 1000);
  },
  components: {
    Work
  },
  methods: {
    getTask() {
      let randomNum = Math.floor(Math.random() * this.tasks.length);
      let task = this.tasks.splice(randomNum, 1)[0];
      this.task = task;
    },
    start() {
      this.$nextTick(() => {
        this.startWork = true;
        this.showWork = true;
      });
    },
    submit(value) {
      this.reportList.push(value);
      if (value.correct) {
        this.right++;
      }
      if (this.tasks.length > 0) {
        this.startWork = false;
        this.showWork = false;
        setTimeout(this.startNextRound, 500);
      } else {
        let correctRate = Math.round((this.right / this.count) * 100);
        if (correctRate < 80) {
          this.$router.push("/again");
        } else {
          this.axios({
            method: "post",
            responseType: "json",
            url: "https://game-log.huhhz.me/report",
            data: JSON.stringify({
              user_code: sessionStorage.getItem("userCode"),
              device_id: localStorage.getItem("device_id"),
              game_id: 2,
              game_type: 1,
              block_id: this.Utils.guid2(),
              data: this.reportList
            }),
            headers: {
              "Content-Type": "application/json"
            }
          }).then(response => {
            if (response.data.code > 0) {
              sessionStorage.setItem("correct-rate", correctRate);
              sessionStorage.setItem("gameStatus", 1);
              this.$router.push("/ready-round1");
            } else {
              alert("请重新点击提交按钮！");
            }
          });
        }
      }
    },
    startNextRound() {
      this.getTask();
      setTimeout(this.start, 1000);
    }
  }
};
</script>

<style scoped>
</style>