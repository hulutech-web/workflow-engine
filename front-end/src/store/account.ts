import { defineStore } from 'pinia';
import http from '../utils/request/http';
import { Response } from '@/types';
import { useMenuStore } from './menu';
import { useAuthStore } from '@/plugins';
import { useLoadingStore } from './loading';
import {login} from "@/api/account";

export interface Profile {
  account: Account;
  permissions: string[];
  role: string;
}
export interface Account {
  username: string;
  avatar: string;
  gender: number;
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
      console.log(res)
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
        .request<Account, Response<Profile>>('/account', 'get')
        .then((response) => {
          if (response.code === 0) {
            const { setAuthorities } = useAuthStore();
            const { account, permissions, role } = response.data;
            this.account = account;
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
