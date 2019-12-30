<template>
  <div class="round2 container">
    <work :task="task" v-if="showWork" v-on:submit="submit"></work>
    <div class="blackStyle" v-else></div>
  </div>
</template>

<script>
import Work from "@/components/Work.vue";

export default {
  name: "round2",
  data() {
    return {
      tasks: [],
      reportList: [],
      showWork: true,
      right: 0,
      count: 0,
      task: null
    };
  },
  created() {
    for (let j = 2; j <= 7; j++) {
      this.tasks.push(this.Utils.getNumberTask(0, j));
      this.tasks.push(this.Utils.getNumberTask(1, j));
      this.tasks.push(this.Utils.getNumberTask(2, j));
      this.tasks.push(this.Utils.getNumberTask(4, j));

      this.tasks.push(this.Utils.getLetterTask(0, j));
      this.tasks.push(this.Utils.getLetterTask(1, j));
      this.tasks.push(this.Utils.getLetterTask(3, j));
    }
    this.count = this.tasks.length;
    this.right = 0;
    this.getTask();
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
    submit(value) {
      this.reportList.push(value);
      if (value.correct) {
        this.right++;
      }
      if (this.tasks.length > 0) {
        this.showWork = false;
        setTimeout(this.startNextRound, 1500);
      } else {
        this.axios({
          method: "post",
          responseType: "json",
          url: "https://game-log.huhhz.me/report",
          data: JSON.stringify({
            user_code: sessionStorage.getItem("userCode"),
            device_id: localStorage.getItem("device_id"),
            game_id: 2,
            game_type: 3,
            block_id: this.Utils.guid2(),
            data: this.reportList
          }),
          headers: {
            "Content-Type": "application/json"
          }
        }).then(response => {
          if (response.data.code > 0) {
            sessionStorage.setItem(
              "correct-rate",
              Math.round((this.right / this.count) * 100)
            );
            sessionStorage.setItem("gameStatus", 3);

            this.$router.push("/end");
          } else {
            alert("请重新点击提交按钮！");
          }
        });
      }
    },
    startNextRound() {
      this.getTask();
      this.$nextTick(() => {
        this.showWork = true;
      });
    }
  }
};
</script>