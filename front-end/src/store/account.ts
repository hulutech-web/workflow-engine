import { defineStore } from 'pinia';
import http from '../utils/request/http';
import { useMenuStore } from './menu';
import { useAuthStore } from '@/plugins';
import { useLoadingStore } from './loading';
import {login} from "@/api/account";

export interface Profile {
  user: Account;
  permissions: string[];
  role: string;
}
export interface Account {
  username: string;
  avatar: string;
}

export type TokenResult = {
  token: string;
  expires: number;
};
export const useAccountStore = defineStore('account', {
  state() {
    return {
      account: {} as Account,
      permissions: [] as string[],
      role: '',
      logged: true,
    };
  },
  actions: {
    async login(username: string, password: string) {
      // return http
      //   .request<TokenResult, Response<TokenResult>>('/login', 'post_json', { username, password })
      //   .then(async (response) => {
      //     if (response.code === 0) {
      //       this.logged = true;
      //       http.setAuthorization(`${response.data.token}`, new Date(response.data.expires));
      //       await useMenuStore().getMenuList();
      //       return response.data;
      //     } else {
      //       return Promise.reject(response);
      //     }
      //   });
      const data = {
        tenantId: 1,
        username: username,
        password: password,
      } as Account.LoginForm;
      const res = await login(data)
      if (res) {
        this.logged = true;
        http.setAuthorization(`${res.data.token}`, 7200 * 1000, 'Authorization');
        await useMenuStore().getMenuList();
        return res.data;
      }
    },
    async logout() {
      return new Promise<boolean>((resolve) => {
        localStorage.removeItem('stepin-menu');
        http.removeAuthorization();
        this.logged = false;
        resolve(true);
      });
    },
    async profile() {
      const { setAuthLoading } = useLoadingStore();
      setAuthLoading(true);
      return http
        .request('/user/self', 'get')
        .then((response) => {
          if (response) {
            const { setAuthorities } = useAuthStore();
            const { user, permissions, role } = response.data;
            console.log(user, permissions, role)
            this.account = user;
            this.permissions = permissions;
            this.role = role;
            setAuthorities(permissions);
            return response.data;
          } else {
            return Promise.reject(response);
          }
        })
        .finally(() => setAuthLoading(false));
    },
    setLogged(logged: boolean) {
      this.logged = logged;
    },
  },
});
