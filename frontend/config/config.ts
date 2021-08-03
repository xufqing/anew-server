// https://umijs.org/config/
import { defineConfig } from 'umi';
import defaultSettings from './defaultSettings';
import proxy from './proxy';
const { REACT_APP_ENV } = process.env;

export default defineConfig({
  hash: true,
  antd: {},
  dva: {
    hmr: true,
  },
  layout: {
    // https://umijs.org/zh-CN/plugins/plugin-layout
    locale: true,
    siderWidth: 208,
    ...defaultSettings,
  },
  dynamicImport: {
    loading: '@ant-design/pro-layout/es/PageLoading',
  },
  targets: {
    ie: 11,
  },
  // umi routes: https://umijs.org/docs/routing
  routes: [
    {
      path: '/user',
      layout: false,
      routes: [
        {
          path: '/user/login',
          layout: false,
          name: '登录',
          component: './auth/login',
        },
        {
          path: '/user',
          redirect: '/user/login',
        },
      ],
    },
    {
      path: '/account',
      name: 'account',
      routes: [
        {
          path: '/account/settings',
          name: '账户设置',
          component: './account/settings',
        },
        {
          path: '/account',
          redirect: '/account/settings',
        },
      ],
    },
    // system
    {
      path: '/system',
      name: 'system',
      routes: [
        {
          path: '/system/user',
          name: '用户管理',
          component: './system/user',
        },
        {
          path: '/system/dept',
          name: '部门管理',
          component: './system/dept',
        },
        {
          path: '/system/menu',
          name: '菜单管理',
          component: './system/menu',
        },
        {
          path: '/system/role',
          name: '角色管理',
          component: './system/role',
        },
        {
          path: '/system/api',
          name: '接口管理',
          component: './system/api',
        },
        {
          path: '/system/dict',
          name: '字典管理',
          component: './system/dict',
        },
        {
          path: '/system/operlog',
          name: '日志管理',
          component: './system/operlog',
        },
        {
          path: '/index',
          redirect: '/index',
        },
      ],
    },
    // asset
    {
      path: '/asset',
      name: 'asset',
      routes: [
        {
          path: '/asset/host',
          name: '主机管理',
          component: './asset/host',
        },
        {
          path: '/asset/console',
          name: '终端管理',
          layout: false,
          component: './asset/console',
        },
        {
          path: '/index',
          redirect: '/index',
        },
      ],
    },
    {
      component: '404',
    },
  ],
  // Theme for antd: https://ant.design/docs/react/customize-theme-cn
  theme: {
    'primary-color': defaultSettings.primaryColor,
  },
  // esbuild is father build tools
  // https://umijs.org/plugins/plugin-esbuild
  esbuild: {},
  title: false,
  ignoreMomentLocale: true,
  proxy: proxy[REACT_APP_ENV || 'dev'],
  manifest: {
    basePath: '/',
  },
  // Fast Refresh 热更新
  fastRefresh: {},
  // openAPI: [
  //   {
  //     requestLibPath: "import { request } from 'umi'",
  //     // 或者使用在线的版本
  //     // schemaPath: "https://gw.alipayobjects.com/os/antfincdn/M%24jrzTTYJN/oneapi.json"
  //     schemaPath: join(__dirname, 'oneapi.json'),
  //     mock: false,
  //   },
  //   {
  //     requestLibPath: "import { request } from 'umi'",
  //     schemaPath: 'https://gw.alipayobjects.com/os/antfincdn/CA1dOm%2631B/openapi.json',
  //     projectName: 'swagger',
  //   },
  // ],
  nodeModulesTransform: {
    type: 'none',
  },
  // mfsu启用后有部分样式丢失，据说bug，等umi修复
  // mfsu: {},
  webpack5: {},
  exportStatic: {},
});
