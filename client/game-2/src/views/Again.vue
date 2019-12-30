<template>
  <div class="again container">
    <div class="center">
      <div class="content">
        <p class="title">你的正确率是：{{correctRate}}%</p>
        <p>正确率达到 80% 才可以进行正式游戏哦</p>
      </div>
      <a-button class="btn" size="large" type="primary" v-on:click="review">复习规则</a-button>
      <a-button class="btn" size="large" type="primary"  v-on:click="renew">重新练习</a-button>
    </div>
  </div>
</template>

<script>
export default {
  name: "again",
  data() {
    return {
      correctRate:sessionStorage.getItem('correct-rate')
    };
  },
  methods: {
    renew() {
      this.$router.push("/exercise");
    },
    review(){
        this.$router.push("/explain");
    }
  }
};
</script>

<style scoped>
.again .center {
  align-content: center;
  font-size: 33px;
  text-align: center;
  margin-top: 30vh;
}
.btn {
  margin: 10px;
}
</style>