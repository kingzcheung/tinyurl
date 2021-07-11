import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import vitePluginImp from 'vite-plugin-imp'
import {resolve} from 'path'


// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
        vitePluginImp({
            libList: [
                {
                    libName: 'ant-design-vue',
                    style(name) {
                        if (/popconfirm/.test(name)) {
                            // support multiple style file path to import
                            return [
                                'ant-design-vue/es/button/style/index.less',
                                'ant-design-vue/es/popover/style/index.less'
                            ]
                        }
                        if (/col|row/.test(name)) {
                            return 'ant-design-vue/es/grid/style/index.less'
                        }
                        return [
                            `ant-design-vue/es/${name}/style/index.less`
                        ]
                    }
                },
            ]
        })
    ],
    css: {
        preprocessorOptions: {
            less: {
                javascriptEnabled: true,
            },
        },
    },
    build: {
        rollupOptions: {
            input: {
                main: resolve(__dirname, 'index.html'),
                login: resolve(__dirname, 'login.html'),
            },
        }
    }
})
