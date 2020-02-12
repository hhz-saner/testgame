<template>
  <div>
    <!--  提示语部分  -->
    <div class="top">
      <div class="welcomeSpan text-center">欢迎来到当前的小游戏</div>
    </div>
    
    <!--  注意点部分  -->
    <div class="middle">
    </div>

    <!--  选择部分    -->
    <div class="bottom">  
      <a-row>
        <a-col :span="2"></a-col>
        <a-col :span="20">
          <a-input placeholder="请输入学号"   v-model="userNumber"  ref="userNumberInput">
            <a-icon slot="prefix" type="user" />
            <a-icon v-if="userNumber" slot="suffix" type="close-circle" @click="emitEmptyNumber" />
          </a-input>
        </a-col>
        <a-col :span="2"></a-col>
      </a-row>

      <a-row>
        <a-col :span="2"></a-col>
        <a-col :span="20">
          <a-checkbox @change="changeHasKnow">我（及父母）已经同意参与此研究并填写过相应的知情同意书</a-checkbox>
        </a-col>
        <a-col :span="2"></a-col>
      </a-row>

      <a-row>
        <a-col :span="2"></a-col>
        <a-col :span="20">
          <a-button type="primary" class="mgt-30" block  @click="enterPlay">进入游戏</a-button>
        </a-col>
        <a-col :span="2"></a-col>
      </a-row>
    
    </div>
  </div>
</template>

<style >
.top{
  position: fixed;
  top: 10%; 
  width: 100%;
}
.welcomeSpan{
  font-family: 黑体;
  font-weight: bold;
  font-size: 30px;
  padding: 10px;
}
.bottom{
  position: fixed;
  bottom: 10%;
  width: 100%
}
</style>

<script>
export default {
  data () {
    return {
      userNumber: '',
      hasKnow: false,
    }
  },
  methods: {
    // @text 清空操作
    emitEmptyNumber () {
      this.$refs.userNumberInput.focus()
      this.userNumber = ''
    },
    // @text checkbox事件
    changeHasKnow (val){
      this.hasKnow = val.target.checked;
    },
   
    // @text 提交按钮
    enterPlay () {
      if ( this.hasKnow == true && this.userNumber != '' ){
          this.axios({
            method: "post",
            responseType: "json",
            url: "https://game-log.huhhz.me/login",
            data: JSON.stringify({
              user_code: this.userNumber
            }),
            headers: {
              "Content-Type": "application/json"
            }
          }).then(response => {
            sessionStorage.setItem("user_code",response.data.data.UserCode);
            sessionStorage.setItem("game_type",response.data.data.GameOneStatus);

            if (response.data.data.GameOneStatus == 0){
              this.$router.push({ path: "/gameinfo"}) 
            }else if ( response.data.data.GameOneStatus == 1 || response.data.data.GameOneStatus == 2){
              this.$router.push({ path: "/officialplay"}) 
            }else{
              this.$router.push({ path: "/endgame"})
            }        
          });
      }else{
        alert("请填写学号并勾选知情同意书！");
      }
    } 

  },
}
</script>
