import { createRouter, createWebHistory } from "vue-router";
import { useUserStore } from "@/stores/user";

const routes = [
  {
    path: "/",
    redirect: "/login",
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("@/views/Login.vue"),
  },
  {
    path: "/register",
    name: "Register",
    component: () => import("@/views/Register.vue"),
  },
  {
    path: "/chat",
    name: "Chat",
    component: () => import("@/views/Chat.vue"),
    meta: { requiresAuth: true },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

// 路由守卫
router.beforeEach((to, from, next) => {
  const userStore = useUserStore();

  // 如果访问的是登录或注册页面
  if (to.name === "Login" || to.name === "Register") {
    // 如果用户已登录，重定向到聊天页面
    if (userStore.isLoggedIn) {
      next("/chat");
    } else {
      // 用户未登录，允许访问登录/注册页面
      next();
    }
  } else {
    // 访问其他页面
    // 如果页面需要认证但用户未登录，重定向到登录页
    if (to.meta.requiresAuth && !userStore.isLoggedIn) {
      next("/login");
    } else {
      // 允许访问
      next();
    }
  }
});

export default router;
