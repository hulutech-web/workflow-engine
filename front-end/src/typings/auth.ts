declare namespace Auth {
    interface RoleSimple {
        id: number;
        name: string;
        createdAt: string;
        updatedAt: string;
    }
    interface Role {
        id: number;
        name: string;
        remark: string;
        menus: number[];
        member: number;
        sort: number;
        is_disable: number;
        createdAt: string;
        updatedAt: string;
    }
    interface RoleReq {
        id: number;
        name: string;
        remark: string;
        sort: number;
        is_disable: number;
        menus: number[];
    }

  interface Tenant {
    id: number;
    name: string;
    address: string;
    phone: string;
    email: string;
    domain: string;
    logo: string;
    description: string;
    is_disable: number;
    expired_at: number;
    menus: number[];
    created_at: string;
    updated_at: string;
  }
  
    interface TenantReq {
      id: number;
      name: string;
      address: string;
      phone: string;
      email: string;
      domain: string;
      logo: string;
      description: string;
      is_disable: number;
      expired_at: number;
      menus: number[];
    }

}
