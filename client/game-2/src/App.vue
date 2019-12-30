<template>
  <div id="app">
    <router-view />
  </div>
</template>

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