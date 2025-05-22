declare namespace API {
  type AccountLoginReq = {
    password: string;
    tenantId: number;
    username: string;
  };

  type AccountLoginResp = {
    token?: string;
    userInfo?: UserResp;
  };

  type AccountRegisterReq = {
    confirm: string;
    email: string;
    password: string;
    phone: string;
    tenantId: number;
    username: string;
  };

  type DateTime = {
    'time.Time'?: string;
  };

  type FileResp = {
    cid?: number;
    ext?: string;
    id?: number;
    name?: string;
    path?: string;
    size?: number;
    tenant_id?: number;
    type?: string;
    uri?: string;
    user_id?: number;
  };

  type IdReq = {
    /** 主键ID */
    id: number;
  };

  type MenuAddReq = {
    badge?: string;
    button?: MenuButton[];
    cacheable?: boolean;
    component: string;
    icon?: string;
    menu_type: 'menu' | 'action';
    name: string;
    path: string;
    permission?: string;
    pid: number;
    render_menu?: boolean;
    sort?: number;
    target?: string;
    title: string;
  };

  type MenuButton = {
    id?: number;
    name?: string;
    sort?: number;
    title?: string;
  };

  type MenuButton = {
    id?: number;
    name?: string;
    sort?: number;
    title?: string;
  };

  type MenuEditReq = {
    badge?: string;
    button?: MenuButton[];
    cacheable?: boolean;
    component: string;
    icon?: string;
    id: number;
    menu_type: 'menu' | 'action';
    name: string;
    path: string;
    permission?: string;
    pid: number;
    render_menu?: boolean;
    sort?: number;
    target?: string;
    title: string;
  };

  type MenuResp = {
    badge?: string;
    button?: MenuButton[];
    cacheable?: boolean;
    children?: MenuResp[];
    component?: string;
    icon?: string;
    id?: number;
    menu_type?: string;
    name?: string;
    path?: string;
    permission?: string;
    pid?: number;
    render_menu?: boolean;
    sort?: number;
    target?: string;
    title?: string;
  };

  type PageReq = {
    /** 页码 */
    pageNo?: number;
    /** 每页大小 */
    pageSize?: number;
  };

  type PageResp = {
    /** 总数 */
    count?: number;
    /** 数据 */
    lists?: any;
    /** 页No */
    pageNo?: number;
    /** 每页Size */
    pageSize?: number;
  };

  type Response = {
    code?: number;
    data?: any;
    msg?: string;
  };

  type RoleAddReq = {
    is_admin?: 0 | 1;
    is_disable?: 0 | 1;
    menus?: number[];
    name: string;
    remark?: string;
    sort?: number;
  };

  type RoleEditReq = {
    id: number;
    is_admin?: 0 | 1;
    is_disable?: 0 | 1;
    menus?: number[];
    name: string;
    remark?: string;
    sort?: number;
  };

  type RoleResp = {
    created_at?: DateTime;
    /** 主键 */
    id?: number;
    /** 是否禁用: [0=否, 1=是] */
    is_disable?: number;
    /** 成员数量 */
    member?: number;
    /** 关联菜单 */
    menus?: number[];
    /** 角色名称 */
    name?: string;
    /** 角色备注 */
    remark?: string;
    /** 角色排序 */
    sort?: number;
    updated_at?: DateTime;
  };

  type RoleSimpleResp = {
    created_at?: DateTime;
    /** 主键 */
    id?: number;
    /** 角色名称 */
    name?: string;
    updated_at?: DateTime;
  };

  type SelectOption = {
    disabled?: boolean;
    label?: string;
    value?: any;
  };

  type TenantAddReq = {
    address: string;
    description?: string;
    domain: string;
    email: string;
    expired_at?: number;
    is_disable?: 0 | 1;
    logo?: string;
    menus?: number[];
    name: string;
    phone: string;
  };

  type TenantEditReq = {
    address: string;
    description?: string;
    domain: string;
    email: string;
    expired_at?: number;
    id: number;
    is_disable?: 0 | 1;
    logo?: string;
    menus?: number[];
    name: string;
    phone: string;
  };

  type TenantResp = {
    address?: string;
    created_at?: DateTime;
    description?: string;
    domain?: string;
    email?: string;
    expired_at?: number;
    id?: number;
    is_disable?: number;
    logo?: string;
    menus?: number[];
    name?: string;
    phone?: string;
    updated_at?: DateTime;
  };

  type unnamedApiParams = {
    /** 菜单ID */
    id: number;
  };

  type unnamedApiParams = {
    /** 用户ID */
    id: number;
  };

  type UserAddReq = {
    avatar?: string;
    email?: string;
    is_disable?: 0 | 1;
    is_multipoint?: 0 | 1;
    nickname?: string;
    password: string;
    phone: string;
    role_id?: number;
    tenant_id?: number;
    username: string;
  };

  type UserEditReq = {
    avatar?: string;
    email?: string;
    id: number;
    is_disable?: 0 | 1;
    is_multipoint?: 0 | 1;
    nickname?: string;
    phone: string;
    role_id?: number;
    tenant_id?: number;
    username: string;
  };

  type UserQueryReq = {
    email?: string;
    is_disable?: 0 | 1 | -1;
    phone?: string;
    role_id?: number;
    tenant_id?: number;
    username?: string;
  };

  type UserResp = {
    avatar?: string;
    created_at?: DateTime;
    email?: string;
    id?: number;
    is_disable?: number;
    is_multipoint?: number;
    nickname?: string;
    phone?: string;
    role?: { id?: number; is_admin?: number; name?: string };
    tenant?: { id?: number; name?: string };
    updated_at?: DateTime;
    username?: string;
  };

  type UserSelfResp = {
    permissions?: string[];
    user?: UserResp;
  };

  type UserUpdateReq = {
    avatar?: string;
    confirm_password?: string;
    id: number;
    password?: string;
  };

  type ValidateServiceParams = {
    /** 验证码角度 */
    angle: string;
    /** 验证码key */
    captcha_key: string;
  };
}
