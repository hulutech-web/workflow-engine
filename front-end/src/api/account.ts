import { post, get } from '@/utils/request/http';
import { Response } from '@/types';

export const login = (data: Account.LoginForm) => post<Response<Account.LoginResult>>('/account/login', data);

// 租户信息
export const getTenant = () => get<Response<Base.SelectOption[]>>('/account/tenant');

