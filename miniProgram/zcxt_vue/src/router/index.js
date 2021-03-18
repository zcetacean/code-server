import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);


import Layout from "@/views/Layout/index.vue";

const routes = [
  {
    path: "/",
    redirect: "login",
    hidden: true,
  },
  {
    path: "/login",
    name: "Login",
    hidden: true,
    meta: {
      name: '登录'
    },
    component: () => import("../views/Login/index.vue")
  },
  // 学生管理
  {
    path: "/student",
    name: "Student",
    redirect: "index",
    meta: {
      name: '学生管理'
    },
    component: Layout,
    children: [
      {
        path: "/index",
        name: "Index",
        component: () => import("../views/Student/index.vue"),
      }
    ]
  },
  // 综测管理
  {
    path: "/verify",
    name: "Verify",
    meta: {
      name: '综测管理'
    },
    component: Layout,
    children: [
      {
        path: "/verify",
        name: "Item",
        component: () => import("../views/Verify/index.vue"),
      }
    ]
  },
  // 组织管理
  {
    path: "/organization",
    name: "Organization",
    meta: {
      name: '组织管理'
    },
    component: Layout,
    children: [
      {
        path: "/tree",
        name: "Tree",
        component: () => import("../views/Organization/Tree.vue"),
      },
      {
        path: "/item",
        name: "Item",
        component: () => import("../views/Organization/Item.vue"),
      }
    ]
  },
  // 管理员
  {
    path: "/admin",
    name: "Admin",
    meta: {
      name: '管理员'
    },
    component: Layout,
    children: [
      {
        path: "/admin",
        name: "Admin",
        component: () => import("../views/Admin/index.vue"),
      }
    ]
  }
];

const router = new VueRouter({
  routes
});

export default router;
