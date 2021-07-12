import {createStore} from 'vuex'
import Cookies from 'js-cookie'


export default createStore({
    state() {
        return {
            uid: 0,
        }
    },
    mutations: {
        GET_UID(state, data) {
            state.uid = data
        },
    },
    actions: {
        get_uid({commit}) {
            const uid = Cookies.get('uid')
            if (uid) {
                commit('GET_UID', uid)
            }
        },
    }
})