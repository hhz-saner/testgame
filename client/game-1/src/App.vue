<template>
  <div id="app">
    <router-view/>
  </div>
</template>

<style>
*{
  margin: 0px;
  padding: 0px;
}
.mgt-10 {
  margin-top: 10px;
}
.mgt-30 {
  margin-top: 30px;
}
.text-center {
  text-align: center;
}
.vhidden {
  visibility: hidden;
}
.font-size-20 {
  font-size: 20px;
}
.font-size-30 {
  font-size: 30px;
}
.font-size-50 {
    font-size: 50px;
}
.font-size-100 {
  font-size: 100px;
}
.heigh100{
  height: 100%;
}
.maxheight100{
  max-height:100%;
}

#app {
  background: rgb(248, 246, 246);
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  position: absolute;
  width: 100%;
  height: 100%;
}
.ant-btn-primary{
  background-color:black;
  border-color:black;
}
.ant-checkbox-checked .ant-checkbox-inner {
    background-color: black;
    border-color: black;
}
.ant-input-affix-wrapper:hover .ant-input:not(.ant-input-disabled) {
    border-color: black;
    border-right-width: 1px !important;
}
.ant-btn-primary:hover, .ant-btn-primary:focus {
    color: #fff;
    background-color: black;
    border-color: black;
}
.ant-btn:hover, .ant-btn:focus, .ant-btn:active, .ant-btn.active {
    background: #fff;
    text-decoration: none;
}
.ant-btn:hover, .ant-btn:focus {
    color: #fff;
    background-color: black;
    border-color: black;
}
</style>

<script>
export default {
  name: "app",
  created() {
    this.setDeviceId();
  },
  methods: {
    setDeviceId() {
      if (!localStorage.getItem("device_id")) {
       localStorage.setItem("device_id", this.Utils.guid2());
      }
    }
  }
};
</script>
