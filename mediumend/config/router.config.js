export default [
  // user
  {
    path: '/user',
    component: '../layouts/UserLayout',
    routes: [
      { path: '/user', redirect: '/user/login' },
      { path: '/user/login', component: './User/Login' },
    //   { path: '/user/register', component: './User/Register' },
    //   { path: '/user/register-result', component: './User/RegisterResult' },
    ],
  },
  // app
  {
    path: '/',
    component: '../layouts/BasicLayout',
    Routes: ['src/pages/Authorized'],
    // authority: ['admin', 'user'],
    routes: [
      // dashboard
      { path: '/', redirect: '/user/login' },
      
      // profile
      {
        path: "/profile",
        name: "profile",
        icon: "audit",
        component: "./Profile/Form"
        // routes: [
        //   {
        //     path: "/profile/basic",
        //     name: "basic",
        //     component: "./Profile/BasicProfile"
        //   }
        // ]
      },
      {
        path: "/rank",
        name: "rank",
        icon: "audit",
        component: "./ranking/Form"
        // routes: [
        //   {
        //     path: "/profile/basic",
        //     name: "basic",
        //     component: "./Profile/BasicProfile"
        //   }
        // ]
      },

      // list
      // {
      //   path: '/list',
      //   icon: 'table',
      //   name: 'list',
      //   routes: [
      //     {
      //       path: '/list/table-list',
      //       name: 'searchtable',
      //       component: './List/TableList',
      //     },
      //   ],
      // },
      
      // {
      //   name: 'exception',
      //   icon: 'warning',
      //   path: '/exception',
      //   routes: [
      //     // exception
      //     {
      //       path: '/exception/403',
      //       name: 'not-permission',
      //       component: './Exception/403',
      //     },
      //     {
      //       path: '/exception/404',
      //       name: 'not-find',
      //       component: './Exception/404',
      //     },
      //     {
      //       path: '/exception/500',
      //       name: 'server-error',
      //       component: './Exception/500',
      //     },
      //     {
      //       path: '/exception/trigger',
      //       name: 'trigger',
      //       hideInMenu: true,
      //       component: './Exception/TriggerException',
      //     },
      //   ],
      // },
      
      {
        component: '404',
      },
    ],
  },
];
