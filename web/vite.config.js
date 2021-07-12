import {defineConfig} from 'vite'
import vue from '@vitejs/plugin-vue'
import {resolve} from 'path'


// https://vitejs.dev/config/
export default defineConfig({
    plugins: [
        vue(),
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
