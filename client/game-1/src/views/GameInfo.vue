<template>
  <div class="heigh100 maxheight100">

    <div  class="heigh100"  v-if="index==1 || index==2 || index==-1 || index==10 || index==11" @click="next">
        <!-- 说明1  -->
        <div v-if="index==-1">
          <div class="game_prompt">
            <div>这个游戏非常简单,在每一轮开始之前,首先你需要集中注意力在这个+上</div>
          </div>
          <div class="plus">
            <div class="font-size-50">+</div>
          </div>
          <div class="bottom text-center">点击任意区域继续</div>
        </div> 

         <!--   选择正确  -->
        <div v-if="index==1">
          <div class="arrow">
            <div class="font-size-50" > 对啦！</div>
          </div>
        </div>

        <!--   选择错误  -->
        <div v-if="index==2">
          <div class="arrow">
            <div class="font-size-30" > 不对呀， 再试一次？......</div>
          </div>
        </div>

        <!-- 演示结束 -->
        <div v-if="index==10" >
          <div class="game_prompt">
            <div>在正式游戏中，你需要又快又准确的连续作出反应。每一次操作中间有几秒钟的休息时间。如果反应不及时，游戏会自动跳转至下一轮。请记住：在每一轮开始时集中注意力在十字上（+）</div>
            <div class="bottom text-center">下面先做一下练习吧。</div>
          </div>
        </div>

         <!-- 进入练习 -->
        <div v-if="index==11" >
          <div class="game_prompt">
            <div class="top text-center">准备好了吗？</div>
            <div class="bottom text-center">点击屏幕已开始</div>
          </div>
        </div>

    </div>

  
    <div class="heigh100" v-if="index==0">
        <!--   箭头选择页  -->
        <div>
        <div class="game_prompt">
          <div v-if="count == 0">接下来你会看到排成一排的五个 > ,比如：</div>
          <div v-if="count == 1">但有的时候，箭头方向不是一致的 ,比如：</div>
        </div>
        <div class="arrow">
          <div class="font-size-50" >>>{{arrowStr}}>></div>
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
          <a-row>
            <a-col :span="2"></a-col>
            <a-col :span="20">
              <div class="font-size-20 mgt-10">你需要指出中间箭头的方向（点击上方按钮进行选择），你的答案是？</div>
            </a-col>
            <a-col :span="2"></a-col>
          </a-row>
        </div>
      </div>
    </div>

  </div>
</template>

<style>
.game_prompt {
  font-family: 黑体;
  font-weight: bold;
  font-size: 20px;
  padding: 10px;
}
.arrow {
  position: fixed;
  top: 30%;
  left: 30%;
}
.plus {
  position: fixed;
  top: 30%;
  left: 45%;
}
</style>

<script>
export default {
  data() {
    return {
      arrowStr: ">",
      arrowLeft: "<",
      arrowRight: ">",
      chooseResult: 0,
      index: -1,
      count:0 ,
    };
  },
  methods: {
    // @text 跳转下一组
    next() {
        if(this.chooseResult == 0 && this.count == 0){
          // 第一次点击
          this.index = 0;
        }else if (this.chooseResult == 1 && this.count<2){
          // @text 选择正确，跳转下一组数据（两组以内）
          this.index = 0;
          this.arrowStr = "<";
        }else if(this.chooseResult == 2){
          // @text 选择错误，返回上一组数据
          this.index = 0;
        }else if(this.chooseResult == 1 && this.count==2 && this.index == 1){
          this.index = 10;
        }else if(this.index == 10){
            this.index = 11;
        }else if(this.index == 11){
            this.$router.push({ path: "/tryplay"}) 
        }
    },

    // @text 按钮的点击事件--左
    chooseArrowLeft() {
      if (this.arrowStr == this.arrowLeft){
        this.chooseResult = 1;
        this.count++;
        this.index = 1;
      }else{
        this.chooseResult = 2;
        this.index = 2;
      }
    },
    // @text 按钮的点击事件--右
    chooseArrowRight() {
      if (this.arrowStr == this.arrowRight){
        this.chooseResult = 1;
        this.count++;
        this.index = 1;
      }else{
        this.chooseResult = 2;
        this.index = 2;
      }
    }
    
  }
};
</script>
