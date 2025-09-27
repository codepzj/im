<template>
  <div
    class="flex items-center justify-center h-screen w-full bg-[url('/background.png')] bg-cover bg-center"
  >
    <a-card class="shadow-sm w-96">
      <div class="flex justify-center">
        <img class="my-6 w-32" src="/logo.png" />
      </div>

      <!-- 登录方式切换 -->
      <div class="flex justify-center mb-6">
        <a-segmented
          class="!my-2"
          v-model:value="loginType"
          :options="[
            { label: '密码登录', value: 'password' },
            { label: '验证码登录', value: 'sms' },
          ]"
        />
      </div>

      <!-- 密码登录表单 -->
      <a-form
        v-if="loginType === 'password'"
        :model="passwordForm"
        class="w-full mx-auto pt-4 px-4"
        @finish="handlePasswordLogin"
      >
        <a-form-item
          label="手机号"
          name="phone"
          :rules="[
            { required: true, message: '请输入手机号' },
            { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号' },
          ]"
        >
          <a-input v-model:value="passwordForm.phone">
            <template #prefix>
              <PhoneOutlined class="site-form-item-icon" />
            </template>
          </a-input>
        </a-form-item>

        <a-form-item
          label="密码"
          name="password"
          :rules="[{ required: true, message: '请输入密码' }]"
        >
          <a-input-password v-model:value="passwordForm.password">
            <template #prefix>
              <LockOutlined class="site-form-item-icon" />
            </template>
          </a-input-password>
        </a-form-item>

        <a-form-item class="flex justify-center !mt-12">
          <a-button
            :disabled="!passwordForm.phone || !passwordForm.password"
            type="primary"
            html-type="submit"
            class="w-48"
            :loading="loading"
          >
            登录
          </a-button>
        </a-form-item>
      </a-form>

      <!-- 验证码登录表单 -->
      <a-form
        v-if="loginType === 'sms'"
        :model="smsForm"
        class="w-full mx-auto pt-4 px-4"
        @finish="handleSmsLogin"
      >
        <a-form-item
          label="手机号"
          name="phone"
          :rules="[
            { required: true, message: '请输入手机号' },
            { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号' },
          ]"
        >
          <a-input v-model:value="smsForm.phone">
            <template #prefix>
              <PhoneOutlined class="site-form-item-icon" />
            </template>
          </a-input>
        </a-form-item>

        <a-form-item
          label="验证码"
          name="phoneCode"
          :rules="[{ required: true, message: '请输入验证码' }]"
        >
          <div class="flex gap-2">
            <a-input v-model:value="smsForm.phoneCode" class="flex-1">
              <template #prefix>
                <SafetyOutlined class="site-form-item-icon" />
              </template>
            </a-input>
            <a-button
              @click="sendSmsCodeAction"
              :disabled="smsCountdown > 0"
              :loading="smsLoading"
            >
              {{ smsCountdown > 0 ? `${smsCountdown}s` : "获取验证码" }}
            </a-button>
          </div>
        </a-form-item>

        <a-form-item class="flex justify-center !mt-12">
          <a-button
            :disabled="!smsForm.phone || !smsForm.phoneCode"
            type="primary"
            html-type="submit"
            class="w-48"
            :loading="loading"
          >
            登录
          </a-button>
        </a-form-item>
      </a-form>

      <div class="text-center mt-4">
        <a-button type="link" @click="goToRegister">
          还没有账户？立即注册
        </a-button>
      </div>
    </a-card>
  </div>
</template>
<script setup>
import { reactive, ref } from "vue";
import { useRouter } from "vue-router";
import { message } from "ant-design-vue";
import { useUserStore } from "@/stores/user";
import { sendSmsCode, login, loginWithSms } from "@/api/user";
import {
  PhoneOutlined,
  LockOutlined,
  SafetyOutlined,
} from "@ant-design/icons-vue";

const router = useRouter();
const userStore = useUserStore();

// 登录方式
const loginType = ref("password");
const loading = ref(false);
const smsLoading = ref(false);
const smsCountdown = ref(0);

// 密码登录表单
const passwordForm = reactive({
  phone: "",
  password: "",
});

// 验证码登录表单
const smsForm = reactive({
  phone: "",
  phoneCode: "",
});

// 发送验证码
const sendSmsCodeAction = async () => {
  if (!smsForm.phone) {
    message.warning("请先输入手机号");
    return;
  }

  if (!/^1[3-9]\d{9}$/.test(smsForm.phone)) {
    message.warning("请输入正确的手机号");
    return;
  }

  smsLoading.value = true;
  try {
    await sendSmsCode(smsForm.phone);
    message.success("验证码已发送");
    startCountdown();
  } catch (error) {
    message.error(error.message || "发送验证码失败");
  } finally {
    smsLoading.value = false;
  }
};

const startCountdown = () => {
  smsCountdown.value = 60;
  const timer = setInterval(() => {
    smsCountdown.value--;
    if (smsCountdown.value <= 0) {
      clearInterval(timer);
    }
  }, 1000);
};

// 密码登录
const handlePasswordLogin = async () => {
  loading.value = true;
  try {
    const response = await login(passwordForm);
    userStore.setLoggedIn(true);
    message.success("登录成功");
    router.push("/chat");
  } catch (error) {
    message.error(error.message || "登录失败");
  } finally {
    loading.value = false;
  }
};

// 验证码登录
const handleSmsLogin = async () => {
  loading.value = true;
  try {
    const response = await loginWithSms(smsForm);
    userStore.setLoggedIn(true);
    message.success("登录成功");
    router.push("/chat");
  } catch (error) {
    message.error(error.message || "登录失败");
  } finally {
    loading.value = false;
  }
};

const goToRegister = () => {
  router.push("/register");
};
</script>
