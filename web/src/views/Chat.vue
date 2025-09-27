<template>
  <div class="chat-container">
    <a-card class="chat-card" :bordered="false">
      <template #title>
        <div class="chat-title">
          <h2>聊天功能</h2>
          <p>功能开发中，敬请期待...</p>
        </div>
      </template>

      <div class="chat-content">
        <a-empty
          description="聊天功能开发中"
          :image="Empty.PRESENTED_IMAGE_SIMPLE"
        >
          <template #description>
            <span>聊天功能正在开发中，敬请期待...</span>
          </template>
        </a-empty>

        <div class="chat-actions">
          <a-button type="primary" @click="logout" danger> 退出登录 </a-button>
        </div>
      </div>
    </a-card>
  </div>
</template>

<script setup>
import { useRouter } from "vue-router";
import { message } from "ant-design-vue";
import { Empty } from "ant-design-vue";
import { useUserStore } from "@/stores/user";

const router = useRouter();
const userStore = useUserStore();

const logout = () => {
  userStore.logout();
  message.success("已退出登录");
  router.push("/login");
};
</script>

<style scoped>
.chat-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: #f5f5f5;
  padding: 20px;
}

.chat-card {
  width: 100%;
  max-width: 500px;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.chat-title {
  text-align: center;
  margin-bottom: 0;
}

.chat-title h2 {
  margin: 0 0 8px 0;
  font-size: 24px;
  font-weight: 600;
  color: #1f2937;
}

.chat-title p {
  margin: 0;
  color: #6b7280;
  font-size: 14px;
}

.chat-content {
  text-align: center;
}

.chat-actions {
  margin-top: 24px;
}
</style>
