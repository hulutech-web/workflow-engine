import { post, get } from '@/utils/request/http';

// menu list
export const getMenuList = () => get('/auth/menu/list');

// menu add
export const addMenu = (data: any) => post('/auth/menu/add', data);

// menu edit
export const editMenu = (data: any) => post('/auth/menu/edit', data);

// menu delete
export const deleteMenu = (id: number) => get('/auth/menu/delete', { id });