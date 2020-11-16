export default {
  // Global page headers (https://go.nuxtjs.dev/config-head)
  head: {
    title: '微课教育',
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { hid: 'description', name: 'description', content: '' }
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ]
  },

  // Global CSS (https://go.nuxtjs.dev/config-css)
  css: [
    'ant-design-vue/dist/antd.css',
    "@/assets/css/reset.css"
  ],

  // Plugins to run before rendering page (https://go.nuxtjs.dev/config-plugins)
  plugins: [
    '@/plugins/antd-ui',
    { src: "~/plugins/swiper.js", ssr: false }
  ],

  // Auto import components (https://go.nuxtjs.dev/config-components)
  components: true,

  // Modules for dev and build (recommended) (https://go.nuxtjs.dev/config-modules)
  buildModules: [
    '@nuxtjs/router',
    '@nuxtjs/style-resources'
  ],

  // Modules (https://go.nuxtjs.dev/config-modules)
  modules: [
  ],

  // Build Configuration (https://go.nuxtjs.dev/config-build)
  build: {
  },
  styleResources: {
    scss: [
      "@/assets/css/color.scss"
    ]
  },
}
