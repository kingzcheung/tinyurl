import {createStore} from 'vuex'
import storejs from 'storejs'


export default createStore({
    state() {
        return {
            user: {},
        }
    },
    mutations: {
        GET_USER(state, data) {
            state.user = data
        },
    },
    actions: {
        get_user({commit}) {
            const user_info = storejs.get('user_info')
            if (user_info) {
                commit('GET_USER', user_info)
            }
        },
    }
})