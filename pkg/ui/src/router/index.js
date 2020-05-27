import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

/* Layout */
import Layout from '@/layout'

/**
 * Note: sub-menu only appear when route children.length >= 1
 * Detail see: https://panjiachen.github.io/vue-element-admin-site/guide/essentials/router-and-nav.html
 *
 * hidden: true                   if set true, item will not show in the sidebar(default is false)
 * alwaysShow: true               if set true, will always show the root menu
 *                                if not set alwaysShow, when item has more than one children route,
 *                                it will becomes nested mode, otherwise not show the root menu
 * redirect: noRedirect           if set noRedirect will no redirect in the breadcrumb
 * name:'router-name'             the name is used by <keep-alive> (must set!!!)
 * meta : {
    roles: ['admin','editor']    control the page roles (you can set multiple roles)
    title: 'title'               the name show in sidebar and breadcrumb (recommend set)
    icon: 'svg-name'             the icon show in the sidebar
    noCache: true                if set true, the page will no be cached(default is false)
    affix: true                  if set true, the tag will affix in the tags-view
    breadcrumb: false            if set false, the item will hidden in breadcrumb(default is true)
    activeMenu: '/example/list'  if set path, the sidebar will highlight the path you set
  }
 */

/**
 * constantRoutes
 * a base page that does not have permission requirements
 * all roles can be accessed
 */
export const constantRoutes = [
  {
    path: '/redirect',
    component: Layout,
    hidden: true,
    children: [
      {
        path: '/redirect/:path(.*)',
        component: () => import('@/views/redirect/index')
      }
    ]
  },
  {
    path: '/login',
    component: () => import('@/views/login/index'),
    hidden: true
  },
  {
    path: '/auth-redirect',
    component: () => import('@/views/login/auth-redirect'),
    hidden: true
  },
  {
    path: '/404',
    component: () => import('@/views/error-page/404'),
    hidden: true
  },
  {
    path: '/401',
    component: () => import('@/views/error-page/401'),
    hidden: true
  },
  {
    path: '/',
    component: Layout,
    redirect: '/dashboard',
    children: [
      {
        path: 'dashboard',
        component: () => import('@/views/dashboard/index'),
        name: 'Dashboard',
        meta: { title: '首页', icon: 'dashboard', affix: true }
      }
    ]
  }
]

/**
 * asyncRoutes
 * the routes that need to be dynamically loaded based on user roles
 */
export const asyncRoutes = [
  {
    path: '/operation',
    component: Layout,
    alwaysShow: true,
    name: 'Operation',
    meta: {
      title: '运营管理',
      icon: 'list'
    },
    children: [
      {
        path: '/users',
        component: () => import('@/views/users/list'),
        name: 'UserList',
        hidden: false,
        meta: { title: '用户列表', icon: 'list' }
      },
    ]
  },
  {
    path: '/novels',
    component: Layout,
    alwaysShow: true,
    name: 'Novels',
    meta: {
      title: '小说管理',
      icon: 'list'
    },
    children: [
      {
        path: '/categorys',
        component: () => import('@/views/category/list'),
        name: 'CategoryList',
        hidden: false,
        meta: { title: '分类列表', icon: 'list' }
      },
      {
        path: '/novel/books',
        component: () => import('@/views/books/list'),
        name: 'BooksList',
        hidden: false,
        meta: { title: '小说列表', icon: 'list' }
      },
      {
        path: '/novel/create',
        component: () => import('@/views/books/create'),
        name: 'CreateBook',
        hidden: true,
        meta: { title: '新建小说', icon: 'create' }
      },
      {
        path: '/novel/update',
        component: () => import('@/views/books/update'),
        name: 'UpdateBook',
        hidden: true,
        meta: { title: '修改小说', icon: 'create' }
      },
      {
        path: '/novel/view',
        component: () => import('@/views/books/view'),
        name: 'getBook',
        hidden: true,
        meta: { title: '查看小说', icon: 'create' }
      },
      {
        path: '/novel/chapters',
        component: () => import('@/views/chapters/list'),
        name: 'ChaptersList',
        hidden: true,
        meta: { title: '章节列表', icon: 'list' }
      },
      {
        path: '/novel/chapters/create',
        component: () => import('@/views/chapters/create'),
        name: 'CreateChapters',
        hidden: true,
        meta: { title: '新建章节', icon: 'create' }
      },
      {
        path: '/novel/chapters/update',
        component: () => import('@/views/chapters/update'),
        name: 'UpdateChapters',
        hidden: true,
        meta: { title: '修改章节', icon: 'create' }
      },
      {
        path: '/novel/uploadnovel/create',
        component: () => import('@/views/books/upload'),
        name: 'uploadNovel',
        hidden: false,
        meta: { title: '上传小说', icon: 'list' }
      },
      {
        path: '/novel/comments/lists',
        component: () => import('@/views/comments/list'),
        name: 'commentsList',
        hidden: false,
        meta: { title: '评论列表', icon: 'list' }
      },
    ]
  },
  {
    path: '/comics',
    component: Layout,
    alwaysShow: true,
    name: 'Comics',
    meta: {
      title: '漫画管理',
      icon: 'list'
    },
    children: [
      {
        path: '/comics/categorys',
        component: () => import('@/views/comicCategory/list'),
        name: '/comics/CategoryList',
        hidden: false,
        meta: { title: '分类列表', icon: 'list' }
      },
      {
        path: '/comics/comic',
        component: () => import('@/views/comics/list'),
        name: 'BooksList',
        hidden: false,
        meta: { title: '漫画列表', icon: 'list' }
      },
      {
        path: '/comics/create',
        component: () => import('@/views/comics/create'),
        name: 'CreateComics',
        hidden: true,
        meta: { title: '新建漫画', icon: 'comics' }
      },
      {
        path: '/comics/update',
        component: () => import('@/views/comics/update'),
        name: 'UpdateComics',
        hidden: true,
        meta: { title: '修改漫画', icon: 'create' }
      },
       {
         path: '/comics/view',
         component: () => import('@/views/comics/view'),
         name: 'getComics',
         hidden: true,
         meta: { title: '查看漫画', icon: 'create' }
       },
      // {
      //   path: '/novel/chapters',
      //   component: () => import('@/views/chapters/list'),
      //   name: 'ChaptersList',
      //   hidden: true,
      //   meta: { title: '章节列表', icon: 'list' }
      // },
      // {
      //   path: '/novel/chapters/create',
      //   component: () => import('@/views/chapters/create'),
      //   name: 'CreateChapters',
      //   hidden: true,
      //   meta: { title: '新建章节', icon: 'create' }
      // },
      // {
      //   path: '/novel/chapters/update',
      //   component: () => import('@/views/chapters/update'),
      //   name: 'UpdateChapters',
      //   hidden: true,
      //   meta: { title: '修改章节', icon: 'create' }
      // },
      // {
      //   path: '/novel/uploadnovel/create',
      //   component: () => import('@/views/books/upload'),
      //   name: 'uploadNovel',
      //   hidden: false,
      //   meta: { title: '上传小说', icon: 'list' }
      // },
      {
        path: '/comics/comments/lists',
        component: () => import('@/views/comicComments/list'),
        name: 'ComicCommentsList',
        hidden: false,
        meta: { title: '评论列表', icon: 'list' }
      },
    ]
  },

  // 404 page must be placed at the end !!!
  { path: '*', redirect: '/404', hidden: true }
]

const createRouter = () => new Router({
  // mode: 'history', // require service support
  scrollBehavior: () => ({ y: 0 }),
  routes: constantRoutes
})

const router = createRouter()

// Detail see: https://github.com/vuejs/vue-router/issues/1234#issuecomment-357941465
export function resetRouter() {
  const newRouter = createRouter()
  router.matcher = newRouter.matcher // reset router
}

export default router
