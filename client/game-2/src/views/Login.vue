<template>
  <div class="login container">
    <div class="title">欢迎来到当前的小游戏！</div>
    <div class="login-form">
      <a-row type="flex" justify="center" align="middle" style>
        <a-input
          class="item"
          size="large"
          v-model="userCode"
          placeholder="请输入学号"
          ref="userCodeInput"
        >
          <a-icon slot="prefix" type="user" />
          <a-icon v-if="userCode" slot="suffix" type="close-circle" @click="emitEmpty" />
        </a-input>
      </a-row>
      <a-row type="flex" justify="center" align="middle" style>
        <a-checkbox class="item agree" @change="onChange">我（及父母）已经同意参与此研究并填写过相应的知情同意书</a-checkbox>
      </a-row>
      <a-row type="flex" justify="center" align="middle">
        <a-button
          class="item"
          size="large"
          type="primary"
          :block="true"
          v-on:click="submitHandle"
        >开始游戏</a-button>
      </a-row>
    </div>
  </div>
</template>

<script>
export default {
  name: "login",
  data() {
    return {
      userCode: "",
      check: false
    };
  },
  components: {},
  methods: {
    emitEmpty() {
      this.$refs.userCodeInput.focus();
      this.userCode = "";
    },
    onChange(e) {
      this.check = e.target.checked;
    },
    submitHandle() {
      if (this.userCode == "" || this.check == false) {
        alert("请输入学号并同意和填写知情同意书");
      } else {
        this.getUserInfo();
      }
    },
    getUserInfo() {
      this.axios({
        method: "post",
        responseType: "json",
        url: "https://game-log.huhhz.me/login",
        data: JSON.stringify({
          user_code: this.userCode
        }),
        headers: {
          "Content-Type": "application/json"
        }
      }).then(response => {
        if (response.data.code > 0) {
          sessionStorage.setItem("userCode", response.data.data.UserCode);
          sessionStorage.setItem("gameStatus", response.data.data.GameTwoStatus);
           if (response.data.data.GameTwoStatus == 1) {
            this.$router.push("/ready-round1");
          } else if (response.data.data.GameTwoStatus == 2) {
            this.$router.push("/ready-round2");
          } else if (response.data.data.GameTwoStatus == 3) {
            this.$router.push("/finished");
          } else {
            this.$router.push("/explain");
          }
        } else {
          alert("请重新点击登录按钮");
        }
      });
    }
  }
};
</script>


<style scoped>
.login .login-form {
  margin-top: 55vh;
}
.login .login-form .agree {
  font-size: 2.6vw;
}
.login .login-form .item {
  margin-top: 1vh;
}
</style>