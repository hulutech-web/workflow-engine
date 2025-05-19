declare namespace Account {
    interface LoginForm {
        tenantId: number;
        username: string;
        password: string;
    }
    interface LoginResult {
        token: string;
        userInfo: UserInfo;
    }
    interface UserInfo {
        id: number;
        username: string;
        phone: string;
        email: string;
        avatar: string;
        nickname: string;
        is_multipoint: number;
        is_disable: number;
        role: {
            id: number;
            name: string;
            is_admin: number;
        };
        tenant: {
            id: number;
            name: string;
        };
        created_at: string;
        updated_at: string;
    }
}