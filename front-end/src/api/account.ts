import { post, get } from '@/utils/request/http';

export const login = (data: Account.LoginForm) => post<Account.LoginResult>('/account/login', data);

// 租户信息
export const getTenant = () => get<Base.SelectOption[]>('/account/tenant');

