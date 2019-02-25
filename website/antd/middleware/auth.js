import axios from 'axios'

export default function({ store, error, req, route, redirect }) {
  if (!store.state.user) {
    if (process.server) {
      if (req.headers.cookie) {
        const cookieparser = require('cookieparser')
        const parsed = cookieparser.parse(req.headers.cookie)
        if (parsed.token) {
          store.commit('SET_TOKEN', parsed.token)
          axios.defaults.headers.common.Cookie = req.headers.cookie
          axios.get(`/api/user/get`).then(res => {
            if (res.data.msg === '登录成功') {
              const user = res.data.data
              store.commit('SET_USER', user)
            } else {
              redirect({
                path: '/user/login?callbackUrl=' + route.currentRoute.path
              })
            }
          })
        }
      } else {
        redirect({ path: '/user/login?callbackUrl=' + route.currentRoute.path })
      }
    } else {
      axios.get(`/api/user/get`).then(res => {
        if (res.data.msg === '登录成功') {
          const user = res.data.data
          store.commit('SET_USER', user)
        } else {
          route.push({
            path: '/user/login?callbackUrl=' + route.currentRoute.path
          })
        }
      })
    }
  }
}
