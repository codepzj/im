<template>
  <div
    class="flex items-center justify-center h-screen w-full bg-[url('/background.png')] bg-cover bg-center"
  >
    <a-card class="shadow-sm w-96">
      <div class="flex justify-center">
        <img class="my-6 w-32" src="/logo.png" />
      </div>
      <a-form
        :model="RegisterForm"
        class="w-full mx-auto pt-4 px-4"
        @finish="Register"
      >
        <a-form-item
          label="昵称"
          name="nickname"
          :rules="[{ required: true, message: '请输入昵称' }]"
        >
          <a-input v-model:value="RegisterForm.nickname">
            <template #prefix>
              <UserOutlined class="site-form-item-icon" />
            </template>
          </a-input>
        </a-form-item>

        <a-form-item
          label="手机号"
          name="phone"
          :rules="[
            { required: true, message: '请输入手机号' },
            { pattern: /^1[3-9]\d{9}$/, message: '请输入正确的手机号' },
          ]"
        >
          <a-input v-model:value="RegisterForm.phone">
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
            <a-input v-model:value="RegisterForm.phoneCode" class="flex-1">
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

        <a-form-item
          label="密码"
          name="password"
          :rules="[
            { required: true, message: '请输入密码' },
            { min: 6, message: '密码至少6位' },
          ]"
        >
          <a-input-password v-model:value="RegisterForm.password">
            <template #prefix>
              <LockOutlined class="site-form-item-icon" />
            </template>
          </a-input-password>
        </a-form-item>

        <a-form-item
          label="确认密码"
          name="confirmPassword"
          :rules="[
            { required: true, message: '请确认密码' },
            {
              validator: (rule, value, callback) => {
                if (!RegisterForm.password) {
                  callback();
                  return;
                }
                if (value !== RegisterForm.password) {
                  callback('两次输入的密码不一致');
                  return;
                }
                callback();
              },
            },
          ]"
        >
          <a-input-password v-model:value="RegisterForm.confirmPassword">
            <template #prefix>
              <LockOutlined class="site-form-item-icon" />
            </template>
          </a-input-password>
        </a-form-item>

        <a-form-item class="flex justify-center !mt-12">
          <a-button
            :disabled="disabled"
            type="primary"
            html-type="submit"
            class="w-48"
          >
            注册
          </a-button>
        </a-form-item>
      </a-form>

      <div class="text-center mt-4">
        <a-button type="link" @click="goToLogin"> 已有账户？立即登录 </a-button>
      </div>
    </a-card>
  </div>
</template>

<script setup>
import { ref, reactive, computed } from "vue";
import { useRouter } from "vue-router";
import { message } from "ant-design-vue";
import { sendSmsCode, register } from "@/api/user";
import {
  UserOutlined,
  PhoneOutlined,
  SafetyOutlined,
  LockOutlined,
} from "@ant-design/icons-vue";

const router = useRouter();

const smsLoading = ref(false);
const smsCountdown = ref(0);

const RegisterForm = reactive({
  nickname: "",
  phone: "",
  phoneCode: "",
  password: "",
  confirmPassword: "",
});

const sendSmsCodeAction = async () => {
  if (!RegisterForm.phone) {
    message.warning("请先输入手机号");
    return;
  }

  if (!/^1[3-9]\d{9}$/.test(RegisterForm.phone)) {
    message.warning("请输入正确的手机号");
    return;
  }

  smsLoading.value = true;
  try {
    await sendSmsCode(RegisterForm.phone);
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

const Register = async () => {
  try {
    await register(RegisterForm);
    message.success("注册成功，请登录");
    router.push("/login");
  } catch (error) {
    message.error(error.message || "注册失败");
  }
};

const goToLogin = () => {
  router.push("/login");
};

const disabled = computed(() => {
  return !(
    RegisterForm.nickname &&
    RegisterForm.phone &&
    RegisterForm.phoneCode &&
    RegisterForm.password &&
    RegisterForm.confirmPassword
  );
});
</script>
