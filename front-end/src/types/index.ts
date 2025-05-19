export function isResponse(obj: any): obj is Api.Response<any> {
    return typeof obj === 'object' && obj.msg !== undefined && obj.code !== undefined;
}
