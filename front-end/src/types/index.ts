export interface Response<T> {
  msg: string;
  code: number;
  data: T;
}

export function isResponse(obj: any): obj is Response<any> {
  return typeof obj === 'object' && obj.msg !== undefined && obj.code !== undefined;
}
