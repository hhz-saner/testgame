import Vue from 'vue'
import Router from 'vue-router'
import Content from './views/Content.vue'
import GameInfo from './views/GameInfo.vue'
import TryPlay from './views/Tryplay.vue'
import OfficialPlay from './views/Officialplay.vue'
import EndGame from './views/EndGame.vue'



Vue.use(Router)

const routers = new Router({
  routes: [
    {
      // @text 正式内容
      path: '/',
      name: 'content',
      component: Content
    },
    {
      //  @text  游戏介绍
      path: '/gameinfo',
      name: 'gameinfo',
      component: GameInfo,
      meta: {
        requireLogin: true,
      }
    },
    {
      // @text 练习
      path: '/tryplay',
      name: 'tryplay',
      component: TryPlay,
      meta: {
        requireLogin: true,
      }
    },
    {
      // @text 正式游戏
      path: '/officialplay',
      name: 'officialplay',
      component: OfficialPlay,
      meta: {
        requireLogin: true,
      }
    },
    {
      // @text 正式游戏
      path: '/endgame',
      name: 'endgame',
      component: EndGame,
      meta: {
        requireLogin: true,
      }
      
    },
  ]
})

routers.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.requireLogin)) {
    if (sessionStorage.getItem('user_code')) {
        next();
    } else {
        next({path:"/"});
    }
  } else {
    next();
  }

})

export default routers
