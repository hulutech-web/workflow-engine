declare namespace Api {
  interface Response<T> {
    msg: string;
    code: number;
    data: T;
  }
}


