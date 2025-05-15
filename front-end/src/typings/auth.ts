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
        menu_ids: string
    }
}
