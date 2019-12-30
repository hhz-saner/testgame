<template>
  <div class="task">
    <div v-if="showInput" class="input-search">
      <a-input-search placeholder="请输入答案" @search="submit" enterButton="提交" size="large" />
    </div>
    <div v-else>
      <div v-if="task!=null">
        <div class="cross">{{cross}}</div>
        <p v-if="showtype" class="cross-type">{{task.typeName}}</p>
      </div>
    </div>
    <div class="blackStyle" v-if="blackStyle"></div>

    <div v-if="type=='exercise'"></div>
  </div>
</template>

<script>
export default {
  name: "Task",
  props: ["tasks", "type"],
  data() {
    return {
      typeTextList: ["顺序", "倒叙", "从大到小", "加和"],
      index: 0,
      showInput: false,
      showtype: true,
      blackStyle: false,
      input: "",
      cross: "➕",
      timer: null,

      tasksCount: 0,
      rightCount: 0,
      task: null,
      reportList: []
    };
  },
  computed: {
    gameType: () => {
      switch (this.type) {
        case "exercise":
          return 1;
        case "round1":
          return 2;
        case "round2":
          return 3;
      }
    }
  },

  created() {
    this.tasksCount = this.tasks.length;
    this.rightCount = 0;
    this.index = 0;
    this.getTask();
    this.timer = setInterval(this.next, 1000);
  },
  methods: {
    startRound() {
      this.getTask();
      this.showtype = true;
      this.blackStyle = false;
      this.timer = setInterval(this.next, 1000);
    },
    next() {
      this.showtype = false;
      let text = this.task.list[this.index];
      if (text != undefined) {
        this.cross = text;
        this.index++;
      } else {
        clearTimeout(this.timer);
        this.showInput = true;
      }
    },
    getTask() {
      let randomNum = Math.floor(Math.random() * this.tasks.length);
      let task = this.tasks.splice(randomNum, 1)[0];
      task.typeName = this.typeTextList[task.type];
      this.task = task;
    },
    submit(value) {
      this.input = value;
      if (this.input == "") {
        alert("请输入答案");
        return;
      }
      let isRight = this.isRight(this.task, this.input);
      if (isRight) {
        this.rightCount++;
      }
      this.reportList.push({
        question: this.task.typeName + " " + this.task.list.toString(),
        answer: this.input,
        correct: isRight
      });
      this.blackStyle = true;
      this.input = "";
      this.index = 0;
      this.showInput = false;
      this.cross = "➕";
      if (this.tasks.length > 0) {
        setTimeout(this.startRound, 1500);
      } else {
        if (
          this.type == "exercise" &&
          Math.round((this.rightCount / this.tasksCount) * 10) < 8
        ) {
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
              game_type: this.gameType,
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
                Math.round((this.rightCount / this.tasksCount) * 100)
              );
              if (this.type == "round1") {
                this.$router.push("/ready-round2");
              } else if (this.type == "round2") {
                this.$router.push("/end");
              } else {
                this.$router.push("/ready-round1");
              }
            } else {
              alert("请重新点击提交按钮！");
            }
          });
        }
      }
    },
    isRight(task, value) {
      value = value.trim().toUpperCase();
      switch (task.type) {
        case 0:
          return task.list.toString() == value.split("").toString();
        case 1:
          return task.list.reverse().toString() == value.split("").toString();
        case 2:
          return (
            task.list
              .sort()
              .reverse()
              .toString() == value.split("").toString()
          );
        case 3:
          return (
            task.list
              .sort()
              .reverse()
              .toString() == value.split("").toString()
          );
        case 4:
          return eval(task.list.join("+")) == value;
      }
    }
  }
};
</script>
<style scoped>
.task {
  text-align: center;
}
.task .input-search {
  margin-top: 30vh;
}

.task .blackStyle {
  background-color: black;
  width: 100%;
  height: 100%;
  position: fixed;
  top: 0;
  bottom: 0;
  left: 0;
  right: 0;
}
</style>