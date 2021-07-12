<template>
    <div class="relative bg-gray-50 h-screen">
        <div class="container mx-auto flex flex-row align-center justify-between">
            <div class="leading-10 py-8">
                <logo/>
            </div>
            <div class="py-8">
                <router-link v-if="uid>0" :to="{name:'Url'}" style="color: white"
                             class="text-white text-sm py-2.5 px-4 rounded-md bg-indigo-500 focus:ring focus:ring-indigo-500 ring-offset-2 ring-offset-indigo-100 shadow-xl">
                    URL 管理
                </router-link>
                <router-link v-else :to="{name:'Login'}" style="color: white"
                             class="text-white text-sm py-2.5 px-4 rounded-md bg-indigo-500 focus:ring focus:ring-indigo-500 ring-offset-2 ring-offset-indigo-100 shadow-xl">
                    Log in / Sign up
                </router-link>
            </div>
        </div>
        <div class="container mx-auto flex flex-row align-center justify-center mt-10">
            <div class="form w-2/4">
                <div class="mb-8">
                    <div class="flex rounded-md shadow-lg flex flex-row relative">
                        <input
                            v-model="form.url"
                            class="py-4 px-3 pr-5 bg-white shadow-sm focus:outline-none focus:ring-2 focus:ring-indigo-500 focus:border-transparent mt-1 block w-full sm:text-sm border border-gray-300 rounded-md"
                            type="text" placeholder="Paste long URL here...">
                        <button class="absolute inset-y-0 right-2" type="button" @click="onClick">
                            <svg t="1625814521069" class="icon" viewBox="0 0 1024 1024" version="1.1"
                                 xmlns="http://www.w3.org/2000/svg" p-id="3791" width="24" height="24">
                                <path
                                    d="M922.523841 496.178663c0.11154 0.273223 0.151449 0.550539 0.25685 0.824785 1.806135 4.627391 2.883677 9.622148 2.895957 14.889105 0 0.024559 0.008186 0.050142 0.008186 0.074701 0 0.01228 0.004093 0.022513 0.004093 0.032746 0 5.369288-1.100054 10.464329-2.962471 15.169491-0.060375 0.153496-0.080841 0.308015-0.14224 0.459464-4.202719 10.278087-12.400425 18.460444-22.684652 22.651906-0.016373 0.00614-0.032746 0.008186-0.047072 0.014326-4.803399 1.950422-10.031471 3.074012-15.533788 3.074012L139.680273 553.3692c-22.848381 0-41.3692-18.520819-41.3692-41.3692 0-22.846334 18.520819-41.3692 41.3692-41.3692l644.817981 0L503.31827 189.452863c-16.156982-16.156982-16.156982-42.349527 0-58.505485 16.154935-16.154935 42.349527-16.154935 58.503439 0L913.570935 482.696604c0.080841 0.080841 0.101307 0.190335 0.180102 0.271176C917.470755 486.73764 920.462902 491.209488 922.523841 496.178663zM710.276832 627.590622c16.154935-16.156982 42.349527-16.156982 58.505485 0 16.152888 16.154935 16.152888 42.349527 0 58.503439L561.822732 893.052622c-16.154935 16.154935-42.349527 16.154935-58.503439 0-16.156982-16.154935-16.156982-42.349527 0-58.503439L710.276832 627.590622z"
                                    p-id="3792" fill="#707070"></path>
                            </svg>
                        </button>
                    </div>
                </div>

                <div class="flex items-start mb-3">
                    <div class="flex items-center h-5 ">
                        <input id="comments" v-model="checked" name="comments" type="checkbox"
                               class="focus:ring-indigo-500 h-4 w-4 text-indigo-600 border-gray-300 rounded"/>
                    </div>
                    <div class="ml-3 text-sm">
                        <label for="comments" class="font-medium text-gray-700">Show advanced options
                        </label>
                    </div>
                </div>

                <div class="bg-white p-3 rounded-lg shadow-xl" v-if="checked">
                    <div class="mb-2">
                        <label for="slug" class="block text-sm font-medium text-gray-700">
                            Custom
                        </label>
                        <div class="mt-1 flex rounded-md shadow-sm">
                            <span
                                class="inline-flex items-center px-3 rounded-l-md border border-r-0 border-gray-300 bg-gray-50 text-gray-500 text-sm">
                                    {{ redirectPrefix }}
                            </span>
                            <input type="text" name="slug" id="slug" v-model="form.slug"
                                   class="placeholder-opacity-50 focus:ring-indigo-500 focus:border-indigo-500 flex-1 block w-full rounded-none rounded-r-md sm:text-sm border-gray-300"
                                   placeholder="aabb"/>
                        </div>
                    </div>

                    <div class="mb-2">
                        <label for="expired_at" class="block text-sm font-medium text-gray-700">Until</label>
                        <input type="date" name="expired_at" id="expired_at" v-model="expired_at"
                               autocomplete="given-name"
                               class="mt-1 placeholder-opacity-25 focus:ring-indigo-500 focus:border-indigo-500 block w-full shadow-sm sm:text-sm border-gray-300 rounded-md"/>
                    </div>

                    <div class="mb-2">
                        <label for="name" class="block text-sm font-medium text-gray-700">
                            Description
                        </label>
                        <div class="mb-1">
                            <textarea id="name" name="name" rows="3" v-model="form.name"
                                      class="placeholder-opacity-50 shadow-sm focus:ring-indigo-500 focus:border-indigo-500 mt-1 block w-full sm:text-sm border border-gray-300 rounded-md"/>
                        </div>
                        <!--                        <p class="mt-2 text-sm text-gray-500">-->
                        <!--                            Brief description for your profile. URLs are hyperlinked.-->
                        <!--                        </p>-->
                    </div>

                </div>

            </div>

        </div>
        <div class="container mx-auto flex flex-row align-center justify-center mt-10">

            <div class="p-4 w-2/4 bg-white rounded-md shadow" v-if="hash.length>0">
                <div class="clipboard relative">
                    <div class="text-align bg-gray-100 p-2 py-4 rounded">{{ redirectUrl }}</div>
                    <button type="button" id="copy" class="absolute right-1 top-1" :data-clipboard-text="redirectUrl">
                        <svg t="1625819700101" class="icon" viewBox="0 0 1024 1024" version="1.1"
                             xmlns="http://www.w3.org/2000/svg" p-id="5334" width="24" height="24">
                            <path
                                d="M768 682.666667V170.666667a85.333333 85.333333 0 0 0-85.333333-85.333334H170.666667a85.333333 85.333333 0 0 0-85.333334 85.333334v512a85.333333 85.333333 0 0 0 85.333334 85.333333h512a85.333333 85.333333 0 0 0 85.333333-85.333333zM170.666667 170.666667h512v512H170.666667z m682.666666 85.333333v512a85.333333 85.333333 0 0 1-85.333333 85.333333H256a85.333333 85.333333 0 0 0 85.333333 85.333334h426.666667a170.666667 170.666667 0 0 0 170.666667-170.666667V341.333333a85.333333 85.333333 0 0 0-85.333334-85.333333z"
                                p-id="5335" fill="#8a8a8a"></path>
                        </svg>
                    </button>
                </div>
                <div v-if="svgContent.length > 0" style="width: 200px;text-align: center;margin: 0 auto;"
                     v-html="svgContent"></div>
            </div>

        </div>
    </div>
</template>

<script>

import {computed, defineComponent, onMounted, reactive, toRefs} from 'vue'
import {apiShorten} from "../model/model";
import ClipboardJS from "clipboard"
import Logo from "../components/Logo.vue";
import QRCode from 'qrcode-svg'
import {useStore} from "vuex";


const reg = /^((ht|f)tps?):\/\/[\w\-]+(\.[\w\-]+)+([\w\-.,@?^=%&:\/~+#]*[\w\-@?^=%&\/~+#])?$/i

export default defineComponent({
    components: {
        Logo,
    },
    setup() {
        const state = reactive({
            hash: '',
            checked: false,
            svgContent: '',
            expired_at: '',
            form: {
                url: '',
                slug: '',
                expired_at: 0,
                type: 0,
                name: ''
            }
        })
        const store = useStore()

        store.dispatch('get_uid')

        const uid = computed(() => store.state.uid)

        const redirectPrefix = computed(() => {
            return `${location.origin}/r/`
        })

        const redirectUrl = computed(() => {
            return redirectPrefix.value + state.hash
        })


        const onClick = () => {
            if (state.form.url.trim().length === 0) return;
            if (!reg.test(state.form.url)) {
                alert("Invalid URL.")
                return;
            }

            if (state.expired_at.length > 0) {
                state.form.expired_at = (new Date(state.expired_at)).getTime() / 1000;
            }


            if (state.form.slug.trim().length > 0) {
                state.form.type = 1
            }

            apiShorten(state.form).then(({data}) => {
                state.hash = data.slug || '';
                const qrcode = new QRCode({
                    content: `${location.origin}/r/${state.hash}`,
                    container: "svg-viewbox", //Responsive use
                    join: true //Crisp rendering and 4-5x reduced file size
                });
                state.svgContent = qrcode.svg();
                // clean
                state.form.slug = ''
                state.form.expired_at = 0
                state.form.name = ''
            })
        }

        onMounted(() => {
            const clip = new ClipboardJS("#copy")
            clip.on("success", e => {
                alert('Copy success!')
                e.clearSelection();
            })

        })

        return {
            onClick,
            ...toRefs(state),
            redirectPrefix,
            redirectUrl,
            uid,
        }
    }
})
</script>
