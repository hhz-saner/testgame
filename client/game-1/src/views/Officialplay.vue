<template>
  <div class="heigh100 maxheight100">

    <div class="heigh100" v-if="index==0" >
        <!--  焦点屏幕  -->
        <div>
          <div class="plus">
            <div class="font-size-50">+</div>
          </div>
        </div>
    </div>

    <div class="heigh100" v-if="index==1">
        <!--   箭头选择页  -->
        <div v-if="index==1">
        <div class="arrow">
          <div class="font-size-50" >{{arrowStr}}</div>
        </div>
        <div class="bottom">
          <a-row>
            <a-col :span="2"></a-col>
            <a-col :span="8">
              <a-button type="primary" class="mgt-30 font-size-20" block @click="chooseArrowLeft"> {{arrowLeft}}
              </a-button>
            </a-col>
            <a-col :span="4"></a-col>
            <a-col :span="8">
              <a-button type="primary" class="mgt-30 font-size-20" block @click="chooseArrowRight" > {{arrowRight}}
              </a-button>
            </a-col>
            <a-col :span="2"></a-col>
          </a-row>
        </div>
      </div>
    </div>

    <div class="heigh100" v-if="index==11" @click="next()">
       <div class="game_prompt">
            <div class="text-center">您的正确率：{{this.getAveragePoint()}}</div>
            <div class="font-size-20 text-center">您的平均反应时长：{{this.getAverageTime()}}</div>
            <div class="text-center">相信你在下一轮会做的更好！</div>
            <div class="bottom text-center">
                  <p>可以休息一下， 也可以点击屏幕继续！</p>
                  <p>点击屏幕继续</p>
            </div>
          </div>
    </div>

    <div class="heigh100" v-if="index==12" @click="next(true)">
       <div class="game_prompt">
            <div class="text-center">您的正确率：{{this.getAveragePoint()}}</div>
            <div class="font-size-20 text-center">您的平均反应时长：{{this.getAverageTime()}}</div>
        </div>
        <div class="bottom text-center">
          <p>{{this.endContent}}</p>
        </div>
    </div>

  </div>
</template>

<script>
export default {
  data() {
    return {
      arrowStr: ">>>>>",
      arrowLeft: "<",
      arrowRight: ">",
      rightCount: 0,
      mistakenCount: 0,
      index: 0,
      count: 0,
      timerA: null,
      timerB: null,
      timerC: null,
      timerD: null,
      maxCount: 59,
      trailCount: 0,
      firstPoint: 0,
      secondPoint: 0,
      averagePoint: 0,
      choice:'',
      exampleArrow: [">>>>>",">><>>","<<><<","<<<<<"],
      notice:'',
      data:[],
      averagePoint: '',
      correct: false,
      showTime: 0,
      clicktime: 0,
      clickTimes: 0,
      use_time: 0,
      endContent:'请稍后， 数据正在上传',
    };
  },
  created(){
    if (sessionStorage.getItem('game_type') > 3){
      this.$router.push({ path: "/endgame"}) 
    }
    if (sessionStorage.getItem('game_type') == 0 ){
       this.$router.push({ path: "/gameinfo"}) 
    }
    if (sessionStorage.getItem('game_type') == 1){
      this.$router.push({ path: "/tryplay"})
    }
    if (sessionStorage.getItem('game_type') == 2 || sessionStorage.getItem('game_type') == 3 ){
      this.$router.push({ path: "/officialplay"})
    }
    this.clickTimes = 0;
    this.use_time = 0;
    this.ArrowArr = this.getArrowArr();
    this.startPlayTrail();
  },

  methods: {
    next(reportOK){
      if(sessionStorage.getItem('game_type') == 2 || sessionStorage.getItem('game_type') == 3){
        location.reload()
      }else{
        this.$router.push({ path: "/endgame"}) 
      }
    },
    getArrowArr(){
       function randomsort(a, b) { 
          return Math.random()>.5 ? -1 : 1; 
      } 
      var arr = [
          0,0,0,0,0,0,0,0,0,0,0,0,0,0,0,
          1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,
          2,2,2,2,2,2,2,2,2,2,2,2,2,2,2,
          3,3,3,3,3,3,3,3,3,3,3,3,3,3,3,
          ]; 
      arr.sort(randomsort);
      return arr;     
    },

    startPlayTrail(){
     // @text 超过最大数 ，出结果
      if ( this.count > this.maxCount){
        if (sessionStorage.getItem('game_type') == 2){
          this.index = 11;
        }else{
          this.index = 12;
        }
        // @text ajax上传数据
        var game_type = parseInt(sessionStorage.getItem('game_type'));
      
        this.reportOK = false;
        this.axios({
            method: "post",
            responseType: "json",
            url: "https://game-log.huhhz.me/report",
            data: JSON.stringify({
               user_code: sessionStorage.getItem('user_code'),
               device_id: localStorage.getItem('device_id'),
               game_id: 1,
               game_type: game_type,
               block_id: this.Utils.guid2(),
               data: this.data
            }),
            headers: {
              "Content-Type": "application/json"
            }
          }).then(response => {
            sessionStorage.setItem('game_type',game_type+1);
            this.endContent = '点击屏幕继续';
            this.reportOK = true;
          });
        return;
      }
      this.index = 0;
      this.choice = '';
      this.notice = '';
      this.correct = false;
      this.showTime = 0;
      this.clicktime = 0;

      // 两秒中跳转至目标选择  
      var key1 = this.count;
      var key = this.ArrowArr.pop();
      this.arrowStr = this.exampleArrow[key];
      this.timerA = setTimeout(this.showArrow,200);
    },

    showArrow() {
      this.showTime = new Date().getTime();
      this.index = 1;
      this.timerB = setTimeout(this.showNotice,1250);
    },

    showNotice() {
      if (this.showTime == 0 || this.clicktime == 0){
        this.use_time = this.use_time;
      }else{
        this.use_time += (this.clicktime - this.showTime);
        this.clickTimes++;
      }
      var useTime = (this.showTime == 0 || this.clicktime == 0)?0:this.clicktime-this.showTime;

      // @text 需要存储的数据
      var dataVal = {question:this.arrowStr,answer:this.choice,correct:this.correct,use_time:useTime};
      this.data.push(dataVal)

      this.count++;
      this.index = 2;
      this.mistakenCount++;
      this.timerC = setTimeout(this.startPlayTrail,800);
    },

    // @text 算平均正确率
    getAveragePoint(){
      this.averagePoint = Math.floor((this.rightCount / this.count ) * 100); 
      return this.averagePoint + '%';
    },

    // @text 算平均时长
    getAverageTime(){
      if(this.clickTimes != 0){
        return Math.floor((this.use_time / this.clickTimes)) + 'ms'
      }
      return 0 + 'ms';
    },

    // @text 按钮的点击事件--左
    chooseArrowLeft() {
      this.clicktime = new Date().getTime();
      clearTimeout(this.timerB);
      this.choice = this.arrowLeft;
      if (this.arrowStr.substr(2,1) == this.arrowLeft){
        this.rightCount++;
        this.correct = true;
      }else{
        this.mistakenCount++;
        this.correct = false;
      }
      this.showNotice();
    },
  
    // @text 按钮的点击事件--右
    chooseArrowRight() {
      this.clicktime = new Date().getTime();
      clearTimeout(this.timerB);
      this.choice = this.arrowRight;
      if (this.arrowStr.substr(2,1) == this.arrowRight){
        this.rightCount++;
        this.correct = true;
      }else{
        this.mistakenCount++;
        this.correct = false;
      }
      this.showNotice();
    },


  }
};
</script>
