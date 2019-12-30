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
  </div>
</template>
<script>
export default {
  name: "Work",
  props: ["task"],
  data() {
    return {
      showInput: false,
      showtype: true,
      blackStyle: false,
      index: 0,

      input: "",
      cross: "➕",

      timer: null,

    };
  },
  computed: {},

  created() {
    this.timer = setInterval(this.next, 1000);
  },
  methods: {
    next() {
      this.showtype = false;
      let text = this.task.list[this.index];
      if (text != undefined) {
        this.cross = text;
        this.index++;
      } else {
        clearTimeout(this.timer);
        this.showInput = true;
        if(this.$route.query.auto == "auto"){
          setTimeout(this.autoAnswer, 500);
        }
      }
    },
    submit(value) {
      if (value == "") {
        alert("请输入答案");
        return;
      }
      let answer = value
      value = value
        .trim()
        .toUpperCase();
      this.$emit("submit", {
        question: this.task.typeName + " " + this.task.list.toString(),
        answer: answer,
        correct: this.task.right == value
      });
    },
    autoAnswer() {
      this.submit(this.task.right)
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