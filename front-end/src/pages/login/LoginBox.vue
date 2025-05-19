<template>
  <ThemeProvider :color="{ middle: { 'bg-base': '#fff' }, primary: { DEFAULT: '#1896ff' } }">
    <div class="login-box rounded-sm">
      <a-form
        :model="form"
        :wrapperCol="{ span: 24 }"
        @finish="login"
        class="login-form w-[400px] p-lg xl:w-[440px] xl:p-xl h-fit text-text"
      >
        <a-form-item :required="true" name="tenantId">
          <a-select
            v-model:value="form.tenantId"
            :getPopupContainer="triggerNode => {return triggerNode.parentNode || document.body;}"
            placeholder="请选择租户"
          >
            <a-select-option v-for="item in tenantList" :key="item.value" :value="item.value" :disabled="item.disabled">{{ item.label }}</a-select-option>
          </a-select>
        </a-form-item>
        <a-form-item :required="true" name="username">
          <a-input
            v-model:value="form.username"
            autocomplete="new-username"
            placeholder="请输入用户名或邮箱: admin"
            class="login-input h-[40px]"
          />
        </a-form-item>
        <a-form-item :required="true" name="password">
          <a-input
            v-model:value="form.password"
            autocomplete="new-password"
            placeholder="请输入登录密码: 888888"
            class="login-input h-[40px]"
            type="password"
          />
        </a-form-item>
        <a-button htmlType="submit" class="h-[40px] w-full" type="primary" :loading="loading"> 登录 </a-button>
        <a-divider></a-divider>
        <div class="terms">
          登录即代表您同意我们的
          <span class="font-bold">用户条款 </span>、<span class="font-bold"> 数据使用协议 </span>、以及
          <span class="font-bold">Cookie使用协议</span>。
        </div>
      </a-form>
    </div>
  </ThemeProvider>
</template>
<script lang="ts" setup>
  import { reactive, ref, onMounted } from 'vue';
  import { useAccountStore } from '@/store';
  import { ThemeProvider } from 'stepin';
  import {getTenant} from "@/api/account";

  const loading = ref(false);

  const form = reactive<Account.LoginForm>({
    username: 'admin',
    password: 'admin888',
    tenantId: 1,
  });
  const tenantList = ref<Base.SelectOption[]>([]);

  const fetchTenantList = async () => {
   const { data } = await getTenant()
    tenantList.value = data;
  }

  const emit = defineEmits<{
    (e: 'success', fields: Account.LoginForm): void;
    (e: 'failure', reason: string, fields: Account.LoginForm): void;
  }>();

  const accountStore = useAccountStore();
  function login(params: Account.LoginForm) {
    loading.value = true;
    accountStore
      .login(params.username, params.password)
      .then((res) => {
        emit('success', params);
      })
      .catch((e) => {
        emit('failure', e.message, params);
      })
      .finally(() => (loading.value = false));
  }

  onMounted(() => {
    fetchTenantList();
  })
</script>
