import { defineStore } from "pinia";
import { ref, computed } from "vue";

export const useUserStore = defineStore(
  "user",
  () => {
    // 状态
    const userInfo = ref(null);
    const isLoggedIn = ref(false);

    // 设置用户信息
    const setUserInfo = (info) => {
      userInfo.value = info;
    };

    // 设置登录状态
    const setLoggedIn = (loggedIn) => {
      isLoggedIn.value = loggedIn;
    };

    // 退出登录
    const logout = () => {
      userInfo.value = null;
      isLoggedIn.value = false;
    };

    return {
      // 状态
      userInfo,
      isLoggedIn,
      // 方法
      setUserInfo,
      setLoggedIn,
      logout,
    };
  },
  {
    persist: true,
    session: false,
  }
);
