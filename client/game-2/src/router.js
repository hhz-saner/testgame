import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)


import Login from './views/Login.vue'
import Explain from './views/Explain.vue'
import Guide1 from './views/Guide-1.vue'
import Guide2 from './views/Guide-2.vue'
import Guide3 from './views/Guide-3.vue'
import Guide4 from './views/Guide-4.vue'
import Guide5 from './views/Guide-5.vue'
import ReadyExercise from './views/ReadyExercise.vue'

import Exercise from './views/Exercise.vue'
import Again from './views/Again.vue'
import ReadyRound1 from './views/ReadyRound1.vue'
import Round1 from './views/Round1.vue'
import ReadyRound2 from './views/ReadyRound2.vue'
import Round2 from './views/Round2.vue'
import End from './views/End.vue'
import Finished from './views/Finished.vue'



const routers = new Router({
    routes: [
        {
            path: '/',
            name: 'login',
            component: Login
        },

        {
            path: '/explain',
            name: 'explain',
            component: Explain,
            meta: {
                requireLogin: true,
                status:0
            }
        },
        {
            path: '/guide1',
            name: 'guide1',
            component: Guide1,
            meta: {
                requireLogin: true,
                status:0
            }
        },
        {
            path: '/guide2',
            name: 'guide2',
            component: Guide2,
            meta: {
                requireLogin: true,
                status:0
            }
        },
        {
            path: '/guide3',
            name: 'guide3',
            component: Guide3,
            meta: {
                requireLogin: true,
                status:0
            }
        },
        {
            path: '/guide4',
            name: 'guide4',
            component: Guide4,
            meta: {
                requireLogin: true,
                status:0
            }
        },
        {
            path: '/guide5',
            name: 'guide5',
            component: Guide5,
            meta: {
                requireLogin: true,
                status:0
            }
        },
        {
            path: '/ready-exercise',
            name: 'ready-exercise',
            component: ReadyExercise,
            meta: {
                requireLogin: true
            }
        },
        {
            path: '/exercise',
            name: 'exercise',
            component: Exercise,
            meta: {
                requireLogin: true
            }
        },
        {
            path: '/again',
            name: 'again',
            component: Again,
            meta: {
                requireLogin: true
            }
        },
        {
            path: '/ready-round1',
            name: 'ready-round1',
            component: ReadyRound1,
            meta: {
                requireLogin: true
            }
        },
        {
            path: '/round1',
            name: 'round1',
            component: Round1,
            meta: {
                requireLogin: true
            }
        },
        {
            path: '/ready-round2',
            name: 'ready-round2',
            component: ReadyRound2,
            meta: {
                requireLogin: true
            }
        },
        {
            path: '/round2',
            name: 'round2',
            component: Round2,
            meta: {
                requireLogin: true
            }
        },
        {
            path: '/end',
            name: 'end',
            component: End,
            meta: {
                requireLogin: true
            }
        },
        {
            path: '/finished',
            name: 'finished',
            component: Finished,
            meta: {
                requireLogin: true
            }
        }

    ]
})


routers.beforeEach((to, from, next) => {
    if (to.matched.some(record => record.meta.requireLogin)) {  // 判断该路由是否需要登录权限
        if (sessionStorage.getItem('userCode')) {  // 判断当前用户的登录信息loginInfo是否存在
            next();
        } else {
            next({
                path: '/'
            })
        }
    } else {
        next();
    }

})

export default routers
