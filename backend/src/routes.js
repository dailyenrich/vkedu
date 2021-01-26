const index = {
    path: '/',
    name: '控制面板',
    component: resolve=>require(['@/components/common/Layout'], resolve),
    children: [
        {
            path: '/',
            component: resolve=>require(['@/components/index/Index'], resolve),
            name: "数据中心"
        }
    ]
}

const auth = {
    path: '/user',
    component: resolve=>require(['@/components/common/UserLayout'], resolve),
    children: [
        {
            path: 'login',
            component: resolve=>require(['@/components/auth/Login'], resolve),
            name: "登录"
        }
    ]
}

const cource = {
    path: '/cource',
    name: "课程管理",
    component: resolve=>require(['@/components/common/Layout'], resolve),
    children: [
        {
            path: '/course/list',
            component: resolve=>require(['@/components/course/List'], resolve),
            name: "课程列表"
        },
        {
            path: '/course/create',
            component: resolve=>require(['@/components/course/Store'], resolve),
            name: "课程添加"
        },
        {
            path: '/course/update/:id',
            component: resolve=>require(['@/components/course/Edit'], resolve),
            name: "课程修改"
        }
    ]
}

const chapter = {
    path: '/cource/:cid/chapter/list',
    name: "章节管理",
    component: resolve=>require(['@/components/common/Layout'], resolve),
    children: [
        {
            path: '/cource/:cid/chapter/list',
            component: resolve=>require(['@/components/chapter/List'], resolve),
            name: "章节列表"
        },
        {
            path: '/cource/:cid/chapter/create',
            component: resolve=>require(['@/components/chapter/Store'], resolve),
            name: "章节添加"
        },
        {
            path: '/cource/:cid/chapter/update/:id',
            component: resolve=>require(['@/components/chapter/Edit'], resolve),
            name: "章节修改"
        }
    ]
}

const section = {
    path: '/chapter/:cid/section/list',
    name: "小节管理",
    component: resolve=>require(['@/components/common/Layout'], resolve),
    children: [
        {
            path: '/chapter/:cid/section/list',
            component: resolve=>require(['@/components/section/List'], resolve),
            name: "小节列表"
        },
        {
            path: '/chapter/:cid/section/create',
            component: resolve=>require(['@/components/section/Store'], resolve),
            name: "小节添加"
        },
        {
            path: '/chapter/:cid/section/update/:id',
            component: resolve=>require(['@/components/chapter/Edit'], resolve),
            name: "小节修改"
        }
    ]
}
export default [
    index,
    auth,
    cource,
    chapter,
    section
];
